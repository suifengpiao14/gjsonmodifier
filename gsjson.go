package gjsonmodifier

import (
	"context"
	"encoding/json"
	"fmt"
	"strings"
	"unsafe"

	"github.com/d5/tengo/v2"
	"github.com/d5/tengo/v2/stdlib"
	"github.com/pkg/errors"
	"github.com/spf13/cast"
	"github.com/tidwall/gjson"
	"github.com/tidwall/sjson"
)

func bytesString(b []byte) string {
	return *(*string)(unsafe.Pointer(&b))
}

// unwrap removes the '[]' or '{}' characters around json
func unwrap(json string) string {
	json = trim(json)
	if len(json) >= 2 && (json[0] == '[' || json[0] == '{') {
		json = json[1 : len(json)-1]
	}
	return json
}

func trim(s string) string {
left:
	if len(s) > 0 && s[0] <= ' ' {
		s = s[1:]
		goto left
	}
right:
	if len(s) > 0 && s[len(s)-1] <= ' ' {
		s = s[:len(s)-1]
		goto right
	}
	return s
}

func init() {
	//支持大小写转换
	gjson.AddModifier("case", func(json, arg string) string {
		if arg == "upper" {
			return strings.ToUpper(json)
		}
		if arg == "lower" {
			return strings.ToLower(json)
		}
		return json
	})
	gjson.AddModifier("tonum", tonum)       // 列转换为数字,可用于字符串数字排序
	gjson.AddModifier("tostring", tostring) // 列转换为数字,可用于字符串数字排序
	gjson.AddModifier("combine", combine)

	gjson.AddModifier("leftJoin", leftJoin)
	gjson.AddModifier("index", index)
	gjson.AddModifier("concat", concat)     //将二维数组按行合并成一维数组,可用于多列索引
	gjson.AddModifier("unique", unique)     //数组去重
	gjson.AddModifier("multi", multi)       //两数想成
	gjson.AddModifier("in", in)             //值在范围内
	gjson.AddModifier("replace", replace)   //替换
	gjson.AddModifier("basePath", basePath) //获取基本路径
	gjson.AddModifier("eval", eval)         //eval 脚本
}

func combine(jsonStr, arg string) string {
	res := gjson.Parse(jsonStr).Array()
	var out []byte
	out = append(out, '{')
	var keys []gjson.Result
	var values []gjson.Result
	for i, value := range res {
		switch i {
		case 0:
			keys = value.Array()
		case 1:
			values = value.Array()
		}
	}
	vlen := len(values)
	var kvMap = map[string]struct{}{}
	for k, v := range keys {
		key := v.String()
		if _, ok := kvMap[key]; ok {
			continue
		}
		out = append(out, v.Raw...)
		out = append(out, ':')
		if k < vlen {
			value := values[k]
			out = append(out, value.Raw...)
		} else {
			out = append(out, '"', '"')
		}
		kvMap[key] = struct{}{}
	}
	out = append(out, '}')
	return bytesString(out)
}

func leftJoin(jsonStr, arg string) string {
	if arg == "" {
		return jsonStr
	}
	if arg[0] != '[' && arg[0] != '{' {
		arg = fmt.Sprintf("[%s]", arg)
	}
	var container = jsonStr
	var sels []subSelector
	var selsLen = 0
	var err error
	ok := false
	sels, _, ok = parseSubSelectors(arg)
	if !ok {
		return container
	}
	selsLen = len(sels)
	if selsLen < 2 {
		return jsonStr
	}

	if selsLen%2 != 0 {
		err = errors.Errorf("leftJoin:The path contained in the parameter must be an even number,got:%s", arg)
		panic(err)
	}
	for i := len(sels) - 1; i >= 0; i = i - 2 {
		left := sels[i-1]
		right := sels[i]
		sub := leftJoin2Path(container, left.path, right.path)
		leftRowGetPath := getParentPath(left.path)
		leftRowSetPath := nameOfLast(leftRowGetPath) // 获取数组下标
		container, err = sjson.SetRaw(container, leftRowSetPath, sub)
		if err != nil {
			err = errors.WithMessage(err, "leftJoin:")
			panic(err)
		}
	}
	out := gjson.Get(container, "@this.0").String() // 返回数组的第一个
	return out

}

func leftJoin2Path(jsonStr, leftPath, rightPath string) string { //合并2个元素
	res := gjson.Parse(jsonStr)
	firstRowPath := getParentPath(leftPath)
	secondRowPath := getParentPath(rightPath)
	firstRef := fmt.Sprintf("[[%s]|@concat,%s]", unwrap(leftPath), firstRowPath)
	firstRefArr := res.Get(firstRef).Array()
	indexArr, rowArr := firstRefArr[0].Array(), firstRefArr[1].Array()
	secondMapPath := fmt.Sprintf("[[%s]|@concat,%s]|@combine", unwrap(rightPath), secondRowPath)
	secondMap := res.Get(secondMapPath).Map()
	secondDefault := map[string]gjson.Result{}
	for _, v := range secondMap {
		for key, value := range v.Map() {
			raw := `""`
			if value.Type == gjson.Number {
				raw = `0`
			}
			secondDefault[key] = gjson.Result{Type: value.Type, Str: `""`, Num: 0, Raw: raw}
		}
		break
	}

	var out []byte
	out = append(out, '[')

	for i, index := range indexArr {
		if i > 0 {
			out = append(out, ',')
		}
		row := rowArr[i].Map()
		secondV, ok := secondMap[index.String()]
		secondVMap := secondDefault
		if ok {
			secondVMap = secondV.Map()
		}
		for k, v := range secondVMap {
			if _, ok := row[k]; ok {
				k = fmt.Sprintf("%s1", k)
			}
			row[k] = v
		}
		out = append(out, '{')
		j := 0
		for k, v := range row {
			if j > 0 {
				out = append(out, ',')
			}
			j++
			out = append(out, fmt.Sprintf(`"%s"`, k)...)
			out = append(out, ':')
			out = append(out, v.Raw...)

		}
		out = append(out, '}')
	}
	out = append(out, ']')
	outStr := bytesString(out)
	return outStr
}

func index(jsonStr, arg string) string {
	if arg == "" {
		return jsonStr
	}
	res := gjson.Parse(jsonStr)
	if arg[0] != '[' && arg[0] != '{' {
		arg = fmt.Sprintf(`[%s]`, arg) // 统一使用复合索引方式处理
	}
	rowPath := getParentPath(arg)
	refPath := fmt.Sprintf("[%s|@values,%s]", arg, rowPath)
	refArr := res.Get(refPath).Array()
	keyArr, rowArr := refArr[0].Array(), refArr[1].Array()
	indexMapArr := make(map[string][]gjson.Result)
	indexArr := concatColumn("-", keyArr...)
	for i, indexKey := range indexArr {
		if _, ok := indexMapArr[indexKey]; !ok {
			indexMapArr[indexKey] = make([]gjson.Result, 0)
		}
		indexMapArr[indexKey] = append(indexMapArr[indexKey], rowArr[i])
	}

	var out []byte
	out = append(out, '{')
	i := 0
	for k, arr := range indexMapArr {
		if i > 0 {
			out = append(out, ',')
		}
		i++
		out = append(out, fmt.Sprintf(`"%s"`, k)...)
		out = append(out, ':')
		out = append(out, '[')
		for j, row := range arr {
			if j > 0 {
				out = append(out, ',')
			}
			out = append(out, row.Raw...)
		}
		out = append(out, ']')
	}
	out = append(out, '}')
	outStr := bytesString(out)
	return outStr
}

func tonum(value string, arg string) (num string) {
	num = strings.Trim(value, `'"`)
	if num == "" {
		num = "0"
	}
	return num
}

func tostring(value string, arg string) (str string) {
	if value == "" {
		return ""
	}
	c := value[0]
	if c == byte('\'') || c == byte('"') {
		return value
	}
	str = fmt.Sprintf("\"%s\"", value) // todo 优化，字符串内包含'\"情况
	return str
}

func unique(json string, arg string) (outStr string) {
	json = TrimSpaces(json)
	m := map[string]struct{}{}
	arr := gjson.Parse(json).Array()
	for _, v := range arr {
		if _, ok := m[v.Raw]; !ok {
			m[v.Raw] = struct{}{}
		}
	}
	isComma := false
	var out []byte
	out = append(out, '[')
	for raw := range m {
		if isComma {
			out = append(out, ',')
		}
		isComma = true
		out = append(out, raw...)

	}
	out = append(out, ']')
	outStr = bytesString(out)
	return outStr
}

func multi(json string, arg string) (v string) {
	res := gjson.Parse(json)
	var f float64 = 1.0
	if res.IsArray() {
		for _, raw := range res.Array() {
			for _, item := range raw.Array() {
				f *= item.Float()
			}
		}
	} else if res.IsObject() {
		for _, item := range res.Map() {
			f *= item.Float()
		}
	}

	if arg == "" {
		arg = "integer"
	}
	switch arg {
	case "int", "integer", "number":
		v = fmt.Sprintf("%d", int64(f))
	case "float":
		v = fmt.Sprintf("%f", f)
	}
	return v
}

func in(jsonStr string, arg string) (out string) {
	parsed := gjson.Parse(jsonStr).String()
	values := strings.Split(arg, ",")
	for _, value := range values {
		if parsed == value {
			return jsonStr
		}
	}
	return ""
}

func orderBy(json string, arg string) (num string) {
	num = strings.Trim(json, `'"`)
	return num
}

func concat(jsonStr, arg string) string {
	resArr := gjson.Parse(jsonStr).Array()
	sep := "-"
	if arg != "" {
		sep = arg
	}
	arr := concatColumn(sep, resArr...)
	b, err := json.Marshal(arr)
	if err != nil {
		err = errors.WithMessage(err, "gsjson.concat")
		panic(err)
	}
	out := string(b)
	return out

}

func replace(jsonStr string, arg string) (out string) {
	arr := strings.Split(arg, "-")
	replacer := strings.NewReplacer(arr...)
	out = replacer.Replace(jsonStr)
	return out
}

func eval(jsonStr string, arg string) (out string) {
	parameters := map[string]interface{}{
		"value": strings.Trim(jsonStr, `'"`),
	}
	l := len(arg)
	if l > 1 {
		if arg[0] == '(' && arg[l-1] == ')' {
			arg = arg[1 : l-1]
		}
	}
	res, err := tengoEval(context.Background(),
		arg,
		parameters,
	)
	if err != nil {
		panic(err)
	}
	out = cast.ToString(res)
	if _, ok := res.(string); ok {
		out = fmt.Sprintf(`"%s"`, out) // 增加引号
	}
	return out
}

func tengoEval(ctx context.Context, expr string, params map[string]interface{}) (interface{}, error) {

	expr = strings.TrimSpace(expr)
	if expr == "" {
		return nil, fmt.Errorf("empty expression")
	}
	var code = `
		fmt := import("fmt")
	    times := import("times")
		collection := import("enum")
		text := import("text")
		inArray := func(val,seq) {
			for x in seq { 
				if val==x{
					return true
				}
			}
			return false
		}
		__res__ := (%s)
	`
	code = fmt.Sprintf(code, expr)

	script := tengo.NewScript([]byte(code))
	script.EnableFileImport(true)
	mods := stdlib.GetModuleMap("times", "enum", "text", "fmt")
	script.SetImports(mods)
	for pk, pv := range params {
		err := script.Add(pk, pv)
		if err != nil {
			return nil, fmt.Errorf("script add: %w", err)
		}
	}
	compiled, err := script.RunContext(ctx)
	if err != nil {
		return nil, fmt.Errorf("script run: %w", err)
	}
	return compiled.Get("__res__").Value(), nil

}

func basePath(jsonStr string, arg string) (out string) {
	out = jsonStr
	lastDotIndex := strings.LastIndex(jsonStr, ".")
	if lastDotIndex > -1 {
		out = jsonStr[lastDotIndex+1:]
		if out[len(out)-1] == '"' { // jsonStr 带有"",需要补齐开头的"
			out = fmt.Sprintf(`"%s`, out)
		}
	}
	return out
}

// concatColumn 合并一行中的所有数据，复合索引使用
func concatColumn(sep string, columns ...gjson.Result) (out []string) {
	out = make([]string, 0)
	clen := len(columns)
	if clen == 0 {
		return out
	}
	rlen := len(columns[0].Array())
	for i := 0; i < rlen; i++ {
		row := make([]string, 0)
		for j := 0; j < clen; j++ {
			column := columns[j].Array()
			row = append(row, column[i].String())
		}
		out = append(out, strings.Join(row, sep))
	}
	return out
}

func getParentPath(path string) string {
	path = TrimSpaces(path)
	if path[0] == '[' || path[0] == '{' {
		subs, newPath, ok := parseSubSelectors(path)
		if !ok {
			return path
		}
		if len(subs) == 0 {
			return newPath // todo 验证返回内容
		}
		path = subs[0].path // 取第一个路径计算父路径
	}
	path = nameOfPrefix(path)
	path = strings.Trim(path, ".#")
	if path == "" {
		path = "@this"
	}
	return path
}

// nameOfLast returns the name of the last component
func nameOfLast(path string) string {
	for i := len(path) - 1; i >= 0; i-- {
		if path[i] == '|' || path[i] == '.' {
			if i > 0 {
				if path[i-1] == '\\' {
					continue
				}
			}
			return path[i+1:]
		}
	}
	return path
}

// nameOfPrefix returns the name of the path except the last component
func nameOfPrefix(path string) string {
	for i := len(path) - 1; i >= 0; i-- {
		if path[i] == '|' || path[i] == '.' {
			if i > 0 {
				if path[i-1] == '\\' {
					continue
				}
			}
			return path[:i+1]
		}
	}
	return path
}

type subSelector struct {
	name string
	path string
}

// copy from gjson
// parseSubSelectors returns the subselectors belonging to a '[path1,path2]' or
// '{"field1":path1,"field2":path2}' type subSelection. It's expected that the
// first character in path is either '[' or '{', and has already been checked
// prior to calling this function.
func parseSubSelectors(path string) (sels []subSelector, out string, ok bool) {
	modifier := 0
	depth := 1
	colon := 0
	start := 1
	i := 1
	pushSel := func() {
		var sel subSelector
		if colon == 0 {
			sel.path = path[start:i]
		} else {
			sel.name = path[start:colon]
			sel.path = path[colon+1 : i]
		}
		sels = append(sels, sel)
		colon = 0
		modifier = 0
		start = i + 1
	}
	for ; i < len(path); i++ {
		switch path[i] {
		case '\\':
			i++
		case '@':
			if modifier == 0 && i > 0 && (path[i-1] == '.' || path[i-1] == '|') {
				modifier = i
			}
		case ':':
			if modifier == 0 && colon == 0 && depth == 1 {
				colon = i
			}
		case ',':
			if depth == 1 {
				pushSel()
			}
		case '"':
			i++
		loop:
			for ; i < len(path); i++ {
				switch path[i] {
				case '\\':
					i++
				case '"':
					break loop
				}
			}
		case '[', '(', '{':
			depth++
		case ']', ')', '}':
			depth--
			if depth == 0 {
				pushSel()
				path = path[i+1:]
				return sels, path, true
			}
		}
	}
	return
}

// TestQuery 提供测试语法函数
func TestQuery(data string, query string) (out string) {
	out = gjson.Get(data, query).String()
	return out
}
