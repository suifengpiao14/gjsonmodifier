package gsjsonmodifier

import (
	"fmt"
	"testing"

	"github.com/tidwall/gjson"
)

func TestCombine(t *testing.T) {
	jsonstr := `
	{"_errCode":"0","_errStr":"SUCCESS","_data":{"items":[{"id":"1","qid":"12057","qname":"全新机(包装盒无破损,配件齐全且原装,可无原 机膜和卡针)","classId":"1","className":"手机","step":"3","stepName":"维修情况","sort":"1","isMultiple":"0","createTime":"2022-08-09 11:50:22","updateTime":"2022-11-28 14:33:59"},{"id":"2","qid":"12097","qname":"机身弯曲情况","classId":"3","className":"平板","step":"2","stepName":"成色情况","sort":"0","isMultiple":"0","createTime":"2022-08-09 11:50:22","updateTime":"2022-08-09 11:50:22"},{"id":"3","qid":"12066","qname":"屏幕外观","classId":"3","className":"平板","step":"2","stepName":"成色情况","sort":"0","isMultiple":"0","createTime":"2022-08-09 11:50:22","updateTime":"2022-08-09 11:50:22"},{"id":"4","qid":"12077","qname":"屏幕显示","classId":"3","className":"平板","step":"2","stepName":"成色情况","sort":"0","isMultiple":"0","createTime":"2022-08-09 11:50:22","updateTime":"2022-08-09 11:50:22"},{"id":"5","qid":"12088","qname":"电池健康效率","classId":"3","className":"平板","step":"3","stepName":"维修情况","sort":"0","isMultiple":"0","createTime":"2022-08-09 11:50:22","updateTime":"2022-08-09 11:50:22"},{"id":"6","qid":"12100","qname":"维修情况","classId":"3","className":"平板","step":"3","stepName":"维修情况","sort":"0","isMultiple":"0","createTime":"2022-08-09 11:50:22","updateTime":"2022-08-09 11:50:22"},{"id":"7","qid":"12106","qname":"零件维修情况","classId":"3","className":"平板","step":"3","stepName":"维修情况","sort":"0","isMultiple":"0","createTime":"2022-08-09 11:50:22","updateTime":"2022-08-09 11:50:22"},{"id":"8","qid":"12115","qname":"受潮状况","classId":"3","className":"平板","step":"3","stepName":"维修情况","sort":"0","isMultiple":"0","createTime":"2022-08-09 11:50:22","updateTime":"2022-08-09 11:50:22"},{"id":"9","qid":"12119","qname":"开机状态","classId":"3","className":"平板","step":"4","stepName":"功能情况","sort":"0","isMultiple":"0","createTime":"2022-08-09 11:50:22","updateTime":"2022-08-09 11:50:22"},{"id":"10","qid":"9368","qname":"是否全新","classId":"3","className":"平板","step":"4","stepName":"功能情况","sort":"0","isMultiple":"0","createTime":"2022-08-09 11:50:22","updateTime":"2022-08-09 11:50:22"}],"pageInfo":{"pageIndex":"0","pageSize":"10","total":"164"}}}
	`

	path := `["classId":_data.items.#.classId,"className":_data.items.#.className]|@combine`
	v := gjson.Get(jsonstr, path).String()
	fmt.Println(v)

}

func TestLeftJoin(t *testing.T) {
	jsonstr := `
	[{"questionId":"12057","questionName":"全新机(包装盒无破损,配件齐全且原装,可无原 机膜和卡针)","classId":"1","type":1},{"questionId":"12097","questionName":"机身弯曲情况","classId":"3","type":3},{"questionId":"12066","questionName":"屏幕外观","classId":"3","type":3},{"questionId":"12077","questionName":"屏幕显示","classId":"3","type":2},{"questionId":"12088","questionName":"电池健康效率","classId":"3","type":2},{"questionId":"12100","questionName":"维修情况","classId":"3","type":2},{"questionId":"12106","questionName":"零件维修情况","classId":"3","type":3},{"questionId":"12115","questionName":"受潮状况","classId":"3","type":2},{"questionId":"12119","questionName":"开机状态","classId":"3","type":2},{"questionId":"9368","questionName":"是否全新","classId":"3","type":4}]
	`
	jsonstr2 := `
	[{"classId":"1","className":"手机","type":1},{"classId":"3","className":"平板","type":2},{"classId":"3","className":"平板","type":3}]
	`
	jsonstr3 := `
	[{"typeName":"自营","type":1},{"typeName":"合作","type":2}]
	`
	jstr := fmt.Sprintf("[%s,%s,%s]", jsonstr, jsonstr2, jsonstr3)
	path1 := "@leftJoin:[@this.0.#.classId,@this.1.#.classId]"
	path := "@leftJoin:[[@this.0.#.classId,@this.0.#.type],[@this.1.#.classId,@this.1.#.type],@this.0.#.type,@this.2.#.type]"
	out1 := gjson.Get(jstr, path1).String()
	out := gjson.Get(jstr, path).String()
	fmt.Println(out1)
	fmt.Println(out)

}

func TestIndex(t *testing.T) {
	jsonstr := `
	[{"questionId":"12057","questionName":"全新机(包装盒无破损,配件齐全且原装,可无原 机膜和卡针)","classId":"1"},{"questionId":"12097","questionName":"机身弯曲情况","classId":"3"},{"questionId":"12066","questionName":"屏幕外观","classId":"3"},{"questionId":"12077","questionName":"屏幕显示","classId":"3"},{"questionId":"12088","questionName":"电池健康效率","classId":"3"},{"questionId":"12100","questionName":"维修情况","classId":"3"},{"questionId":"12106","questionName":"零件维修情况","classId":"3"},{"questionId":"12115","questionName":"受潮状况","classId":"3"},{"questionId":"12119","questionName":"开机状态","classId":"3"},{"questionId":"9368","questionName":"是否全新","classId":"3"}]
	`
	path := "@this|@index:#.classId"
	out := gjson.Get(jsonstr, path).String()
	fmt.Println(out)
}

func TestTonum(t *testing.T) {
	jsonstr := `
	[{"questionId":"12057","questionName":"全新机(包装盒无破损,配件齐全且原装,可无原 机膜和卡针)","classId":"1","type":1},{"questionId":"12097","questionName":"机身弯曲情况","classId":"3","type":3},{"questionId":"12066","questionName":"屏幕外观","classId":"3","type":3},{"questionId":"12077","questionName":"屏幕显示","classId":"3","type":2},{"questionId":"12088","questionName":"电池健康效率","classId":"3","type":2},{"questionId":"12100","questionName":"维修情况","classId":"3","type":2},{"questionId":"12106","questionName":"零件维修情况","classId":"3","type":3},{"questionId":"12115","questionName":"受潮状况","classId":"3","type":2},{"questionId":"12119","questionName":"开机状态","classId":"3","type":2},{"questionId":"9368","questionName":"是否全新","classId":"3","type":4}]
	`
	path := "@this.#.questionId.@tonum"
	out := gjson.Get(jsonstr, path).String()

	fmt.Println(out)

}

func TestUnique(t *testing.T) {
	jsonstr := `
	[{"questionId":"12057","questionName":"全新机(包装盒无破损,配件齐全且原装,可无原 机膜和卡针)","classId":"1","type":1},{"questionId":"12097","questionName":"机身弯曲情况","classId":"3","type":3},{"questionId":"12066","questionName":"屏幕外观","classId":"3","type":3},{"questionId":"12077","questionName":"屏幕显示","classId":"3","type":2},{"questionId":"12088","questionName":"电池健康效率","classId":"3","type":2},{"questionId":"12100","questionName":"维修情况","classId":"3","type":2},{"questionId":"12106","questionName":"零件维修情况","classId":"3","type":3},{"questionId":"12115","questionName":"受潮状况","classId":"3","type":2},{"questionId":"12119","questionName":"开机状态","classId":"3","type":2},{"questionId":"9368","questionName":"是否全新","classId":"3","type":4}]
	`
	path := "@this.#.classId|@unique"
	v := gjson.Get(jsonstr, path)
	out := v.String()
	fmt.Println(out)

}
func TestMulti(t *testing.T) {
	jsonstr := `
	{"price":3,"num":2}
	`
	jsonstr2 := `
	["2",3]
	`
	path := "@multi"
	v := gjson.Get(jsonstr, path).String()
	v2 := gjson.Get(jsonstr2, path).String()
	fmt.Println(v)
	fmt.Println(v2)

}
