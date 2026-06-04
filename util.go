package gjsonmodifier

import (
	"strings"

	"github.com/tidwall/gjson"
)

func TrimSpaces(s string) string {
	return strings.Trim(s, "\r\n\t\v\f ")
}

// FindPathByValue 递归遍历 JSON，找到值等于 targetValue 的 key 的 gjson 路径
// 例如 body=`{"_param":{"pageIndex":"{.PageNumber}"}}`, targetValue=`{.PageNumber}` → `_param.pageIndex`
func FindPathByValue(body string, targetValue string) (path string) {
	parsed := gjson.Parse(body)
	if !parsed.IsObject() {
		return ""
	}
	var search func(gjson.Result, string)
	search = func(obj gjson.Result, prefix string) {
		obj.ForEach(func(key, value gjson.Result) bool {
			currentPath := key.String()
			if prefix != "" {
				currentPath = prefix + "." + currentPath
			}
			if value.String() == targetValue {
				path = currentPath
				return false
			}
			if value.IsObject() {
				search(value, currentPath)
				if path != "" {
					return false
				}
			}
			return true
		})
	}
	search(parsed, "")
	return path
}
