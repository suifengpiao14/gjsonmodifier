package gjsonmodifier

import (
	"bytes"
	"encoding/json"
	"fmt"
	"sort"
	"strconv"
	"strings"
	"unsafe"

	"github.com/pkg/errors"
	"github.com/spf13/cast"
	"github.com/suifengpiao14/funcs"
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
	gjson.AddModifier("tostring", tostring) // 列转换为字符
	gjson.AddModifier("tobool", tobool)     // 列转换为boolean 值
	gjson.AddModifier("combine", combine)
	gjson.AddModifier("groupPlus", groupPlus) // 官方版group 加强版,支持平铺到第几层

	gjson.AddModifier("leftJoin", leftJoin)
	gjson.AddModifier("index", index)
	gjson.AddModifier("concat", concat)       //将二维数组按行合并成一维数组,可用于多列索引
	gjson.AddModifier("unique", unique)       //数组去重
	gjson.AddModifier("multi", multi)         //两数想成
	gjson.AddModifier("in", in)               //值在范围内
	gjson.AddModifier("replace", replace)     //替换
	gjson.AddModifier("basePath", basePath)   //获取基本路径
	gjson.AddModifier("camelCase", camelCase) //转驼峰 arg=lower时，首字母小写
	gjson.AddModifier("snakeCase", snakeCase) //转蛇型
	gjson.AddModifier("toArray", toArray)     //map转数组
	gjson.AddModifier("toMap", toMap)         //数组转对象
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
	sels, _, ok = ParseSubSelectors(arg)
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
		sub := leftJoin2Path(container, left.Path, right.Path)
		leftRowGetPath := getParentPath(left.Path)
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
	num = _trimQuotation(value)
	if num == "" {
		num = "0"
	}
	return num
}
func tobool(value string, arg string) (out string) {
	bolStr := _trimQuotation(value)
	bol := cast.ToBool(bolStr)
	out = _formatOut(bol)
	return out
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

func groupPlus(json, arg string) string {
	res := gjson.Parse(json)
	if !res.IsObject() {
		return ""
	}

	var all [][]byte
	res.ForEach(func(key, value gjson.Result) bool {
		if !value.IsArray() {
			return true
		}
		var idx int
		value.ForEach(func(_, value gjson.Result) bool {
			if idx == len(all) {
				all = append(all, []byte{})
			}
			all[idx] = append(all[idx], ("," + key.Raw + ":" + value.Raw)...)
			idx++
			return true
		})
		return true
	})
	level, _ := strconv.Atoi(arg)
	var data []byte
	data = append(data, '[')
	for i, item := range all {

		if i > 0 {
			data = append(data, ',')
		}
		var raw []byte

		raw = append(raw, '{')
		raw = append(raw, item[1:]...)
		raw = append(raw, '}')
		subRaw := raw
		rawLevel := level
		if rawLevel > 0 {
			rawLevel--
			rawS := groupPlus(string(raw), strconv.Itoa(rawLevel))
			subRaw = []byte(rawS)
		}
		data = append(data, subRaw...)
	}
	data = append(data, ']')
	return string(data)
}

// _trimBracket 删除开始结尾的(),涉及脚本函数有用
func _trimBracket(s string) (out string) {
	s = strings.TrimSpace(s)
	l := len(s)
	if l > 1 && s[0] == '(' && s[l-1] == ')' {
		s = s[1 : l-1]
	}
	return s
}

// _trimBracket 删除开始结尾的""",gjson modifer 入参 jsonStr 字符串时，会增加"",所以入参需要剔除
func _trimQuotation(s string) (out string) {
	s = strings.TrimSpace(s)
	l := len(s)
	if l > 1 && s[0] == '"' && s[l-1] == '"' {
		s = s[1 : l-1]
	}
	return s
}

func _formatOut(i interface{}) (out string) {
	out = cast.ToString(i)
	if _, ok := i.(string); ok {
		out = fmt.Sprintf(`"%s"`, out)
	}
	return out
}

func camelCase(jsonStr string, arg string) (out string) {
	arg = _trimQuotation(arg)
	if arg == "upper" {
		out = funcs.ToCamel(jsonStr)
	} else {
		out = funcs.ToLowerCamel(jsonStr)
	}

	out = _formatOut(out)
	return out
}

func snakeCase(jsonStr string, arg string) (out string) {
	out = funcs.ToSnakeCase(jsonStr)
	out = _formatOut(out)
	return out
}

type sortImp struct {
	Key  string `json:"key"`
	Text string `json:"text"`
}

type sortImps []sortImp

func (a sortImps) Len() int           { return len(a) }
func (a sortImps) Swap(i, j int)      { a[i], a[j] = a[j], a[i] }
func (a sortImps) Less(i, j int) bool { return a[i].Key < a[j].Key }

func (sm sortImps) Array() (texts []string) {
	texts = make([]string, 0)
	for _, s := range sm {
		texts = append(texts, s.Text)
	}
	return texts
}

func toArray(jsonStr string, arg string) (out string) {
	result := gjson.Parse(jsonStr)

	level, key := 0, ""
	argArr := strings.SplitN(arg, ",", 2)
	var err error
	switch len(argArr) {
	case 2:
		level, err = strconv.Atoi(argArr[0])
		key = argArr[1]
		if err != nil {
			key = argArr[0]
			level, _ = strconv.Atoi(argArr[1])
		}
	case 1:
		level, err = strconv.Atoi(argArr[0])
		if err != nil {
			key = argArr[0]
		}
	}
	m := result.Map()
	for i := 0; i < level; i++ {
		tmp := make(map[string]gjson.Result)
		for identify, row := range m {
			for subIdentify, subRow := range row.Map() {
				tmp[fmt.Sprintf("%s.%s", identify, subIdentify)] = subRow
			}
		}
		m = tmp
	}
	sm := make(sortImps, 0)
	for identify, row := range m {
		raw := strings.TrimSpace(row.Raw)
		if key != "" { // 增加key
			raw = strings.TrimLeft(raw, "{")
			raw = strings.TrimRight(raw, "}")
			if raw == "" {
				raw = fmt.Sprintf(`{"%s":"%s"}`, key, identify)
			} else {
				raw = fmt.Sprintf(`{"%s":"%s",%s}`, key, identify, raw)
			}
		}
		sm = append(sm, sortImp{
			Key:  identify,
			Text: raw,
		})
	}
	sort.Sort(sm)

	out = fmt.Sprintf(`[%s]`, strings.Join(sm.Array(), ","))
	return out
}

func toMap(jsonStr string, arg string) (out string) {
	result := gjson.Parse(jsonStr)
	var w bytes.Buffer
	w.WriteString("{")
	for i, row := range result.Array() {
		if i > 0 {
			w.WriteString(",")
		}
		key := row.Get(arg).String()
		pair := fmt.Sprintf(`"%s":%s`, key, row.Raw)
		w.WriteString(pair)
	}
	w.WriteString("}")
	out = w.String()
	return out
}

func basePath(jsonStr string, arg string) (out string) {
	out = _trimQuotation(jsonStr)
	lastDotIndex := strings.LastIndex(out, ".")
	if lastDotIndex > -1 {
		out = out[lastDotIndex+1:]
	}
	out = _formatOut(out)
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
		subs, newPath, ok := ParseSubSelectors(path)
		if !ok {
			return path
		}
		if len(subs) == 0 {
			return newPath // todo 验证返回内容
		}
		path = subs[0].Path // 取第一个路径计算父路径
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
	Name string
	Path string
}

// copy from gjson
// ParseSubSelectors returns the subselectors belonging to a '[path1,path2]' or
// '{"field1":path1,"field2":path2}' type subSelection. It's expected that the
// first character in path is either '[' or '{', and has already been checked
// prior to calling this function.
func ParseSubSelectors(path string) (sels []subSelector, out string, ok bool) {
	modifier := 0
	depth := 1
	colon := 0
	start := 1
	i := 1
	pushSel := func() {
		var sel subSelector
		if colon == 0 {
			sel.Path = path[start:i]
		} else {
			sel.Name = path[start:colon]
			sel.Path = path[colon+1 : i]
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

// GetAllPath 获取json中的所有路径
func GetAllPath(jsonStr string) (paths []string) {
	paths = make([]string, 0)
	result := gjson.Parse(jsonStr)
	allResult := getAllJsonResult(result)
	for _, result := range allResult {
		subPath := result.Path(jsonStr)
		paths = append(paths, subPath)
	}
	return paths
}

func getAllJsonResult(result gjson.Result) (allResult []gjson.Result) {
	allResult = make([]gjson.Result, 0)
	result.ForEach(func(key, value gjson.Result) bool {
		if !value.IsArray() && !value.IsObject() {
			allResult = append(allResult, value)
		} else {
			subAllResult := getAllJsonResult(value)
			allResult = append(allResult, subAllResult...)
		}
		return true
	})
	return
}
