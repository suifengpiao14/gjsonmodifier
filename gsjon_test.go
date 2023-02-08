package gjsonmodifier

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/assert"
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
func TestLeftJoin2(t *testing.T) {

	PaginateOut := `[{"Fauto_create_time":"2023-01-17 17:33:53","Fauto_update_time":"2023-01-17 17:34:08","Fdetect_sn_id":"15454","Fend_node":"201","Fend_node_time":"2023-01-17 10:35:10","Fid":"2","Fop_id":"12","Fop_name":"张三","Forder_id":"16688102","Fsound_dur_time":"10","Fsound_record_id":"10001055","Fsound_url":"http://s1-1251010403.file.myqcloud.com/xian_yu_x_z/order/order_record/18170456-01171722.mp3","Fstart_node":"0","Fstart_node_time":"2023-01-17 10:35:00","Fvalid":"1"},{"Fauto_create_time":"2023-01-17 14:05:56","Fauto_update_time":"2023-01-1717:31:06","Fdetect_sn_id":"15454","Fend_node":"201","Fend_node_time":"2023-01-17 10:35:10","Fid":"1","Fop_id":"4234","Fop_name":"13528768996","Forder_id":"16688102","Fsound_dur_time":"10","Fsound_record_id":"10001054","Fsound_url":"http://s1-1251010403.file.myqcloud.com/xian_yu_x_z/order/order_record/18170456-01171722.mp3","Fstart_node":"0","Fstart_node_time":"2023-01-17 10:35:00","Fvalid":"1"}]`

	startNodeTitle := `[{"Fsound_node_id":"0","startNodeTitle":"无"},{"Fsound_node_id":"101","startNodeTitle":"手动建单"},{"Fsound_node_id":"102","startNodeTitle":"到店验机"},{"Fsound_node_id":"103","startNodeTitle":"开始质检"},{"Fsound_node_id":"104","startNodeTitle":"重新估价"},{"Fsound_node_id":"105","startNodeTitle":"继续付款"},{"Fsound_node_id":"201","startNodeTitle":"订单完成（确定）"},{"Fsound_node_id":"202","startNodeTitle":"订单关闭"},{"Fsound_node_id":"203","startNodeTitle":"自动关闭"},{"Fsound_node_id":"204","startNodeTitle":"返回首页"}]`
	endNodeTitle := `[{"Fsound_node_id":"0","endNodeTitle":"无"},{"Fsound_node_id":"101","endNodeTitle":"手动建单"},{"Fsound_node_id":"102","endNodeTitle":"到店验机"},{"Fsound_node_id":"103","endNodeTitle":"开始质检"},{"Fsound_node_id":"104","endNodeTitle":"重新估价"},{"Fsound_node_id":"105","endNodeTitle":"继续付款"},{"Fsound_node_id":"201","endNodeTitle":"订单完成（确定）"},{"Fsound_node_id":"202","endNodeTitle":"订单关闭"},{"Fsound_node_id":"203","endNodeTitle":"自动关闭"},{"Fsound_node_id":"204","endNodeTitle":"返回首页"}]`

	jsonstr := fmt.Sprintf("[%s,%s,%s]", PaginateOut, startNodeTitle, endNodeTitle)
	path := "@leftJoin:[@this.0.#.Fstart_node,@this.1.#.Fsound_node_id,@this.0.#.Fend_node,@this.2.#.Fsound_node_id]"
	out := gjson.Get(jsonstr, path).String()
	fmt.Println(out)

}

func TestRename(t *testing.T) {
	jsonstr := `[{"Fsound_node_id":"0","startNodeTitle":"无"},{"Fsound_node_id":"101","startNodeTitle":"手动建单"},{"Fsound_node_id":"102","startNodeTitle":"到店验机"},{"Fsound_node_id":"103","startNodeTitle":"开始质检"},{"Fsound_node_id":"104","startNodeTitle":"重新估价"},{"Fsound_node_id":"105","startNodeTitle":"继续付款"},{"Fsound_node_id":"201","startNodeTitle":"订单完成（确定）"},{"Fsound_node_id":"202","startNodeTitle":"订单关闭"},{"Fsound_node_id":"203","startNodeTitle":"自动关闭"},{"Fsound_node_id":"204","startNodeTitle":"返回首页"}]`
	path := "{Fsound_node_id:@this.#.Fsound_node_id,endNodeTitle:@this.#.startNodeTitle}|@group"
	out := gjson.Get(jsonstr, path).String()
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

func TestTostring(t *testing.T) {
	t.Run("value empty", func(t *testing.T) {
		value := ""
		out := tostring(value, "")
		assert.Equal(t, "", out)
	})
	t.Run("value string", func(t *testing.T) {
		value := "ok"
		out := tostring(value, "")
		assert.Equal(t, "", out)
	})
	t.Run("value number", func(t *testing.T) {
		value := "1.2"
		out := tostring(value, "")
		assert.Equal(t, "\"1.2\"", out)
	})

	t.Run("complex ", func(t *testing.T) {
		value := "{\"input\":{\"pageIndex\":0,\"pageSize\":10},\"PaginateTotalOut\":11,\"PaginateOut\":[{\"Fcreate_time\":\"2022-07-21 18:34:35\",\"Fid\":\"11\",\"Fidentify\":\"fghfg\",\"Fmerchant_id\":\"44\",\"Fmerchant_name\":\"商户名称\",\"Foperate_name\":\"彭政\",\"Fstatus\":\"1\",\"Fstore_id\":\"12\",\"Fstore_name\":\"店铺名称\",\"Fupdate_time\":\"2022-07-21 18:34:35\"},{\"Fcreate_time\":\"2022-07-20 18:51:58\",\"Fid\":\"10\",\"Fidentify\":\"100000\",\"Fmerchant_id\":\"100000\",\"Fmerchant_name\":\"商户名称100000\",\"Foperate_name\":\"彭政100000\",\"Fstatus\":\"1\",\"Fstore_id\":\"1211\",\"Fstore_name\":\"店铺名称100000\",\"Fupdate_time\":\"2022-07-21 18:34:11\"},{\"Fcreate_time\":\"2022-07-20 10:22:30\",\"Fid\":\"9\",\"Fidentify\":\"fghfg\",\"Fmerchant_id\":\"44\",\"Fmerchant_name\":\"商户名称\",\"Foperate_name\":\"彭政\",\"Fstatus\":\"1\",\"Fstore_id\":\"12\",\"Fstore_name\":\"店铺名称\",\"Fupdate_time\":\"2022-07-20 10:22:30\"},{\"Fcreate_time\":\"2022-06-06 10:48:15\",\"Fid\":\"8\",\"Fidentify\":\"222\",\"Fmerchant_id\":\"222\",\"Fmerchant_name\":\"222\",\"Foperate_name\":\"pengkuan\",\"Fstatus\":\"1\",\"Fstore_id\":\"222\",\"Fstore_name\":\"222\",\"Fupdate_time\":\"2022-06-06 10:48:45\"},{\"Fcreate_time\":\"2022-06-06 10:47:48\",\"Fid\":\"7\",\"Fidentify\":\"1111\",\"Fmerchant_id\":\"1111\",\"Fmerchant_name\":\"1111\",\"Foperate_name\":\"pengkuan\",\"Fstatus\":\"2\",\"Fstore_id\":\"1111\",\"Fstore_name\":\"1111\",\"Fupdate_time\":\"2022-06-06 10:47:48\"},{\"Fcreate_time\":\"2022-06-06 10:44:29\",\"Fid\":\"6\",\"Fidentify\":\"312321\",\"Fmerchant_id\":\"321321\",\"Fmerchant_name\":\"321321\",\"Foperate_name\":\"pengkuan\",\"Fstatus\":\"2\",\"Fstore_id\":\"3213213\",\"Fstore_name\":\"12321321\",\"Fupdate_time\":\"2022-06-06 10:48:40\"},{\"Fcreate_time\":\"2022-06-02 14:41:31\",\"Fid\":\"5\",\"Fidentify\":\"abced123f\",\"Fmerchant_id\":\"44\",\"Fmerchant_name\":\"商户名称\",\"Foperate_name\":\"彭政\",\"Fstatus\":\"1\",\"Fstore_id\":\"12\",\"Fstore_name\":\"店铺名称\",\"Fupdate_time\":\"2022-06-06 10:48:27\"},{\"Fcreate_time\":\"2022-06-01 18:14:20\",\"Fid\":\"4\",\"Fidentify\":\"fghfg\",\"Fmerchant_id\":\"44\",\"Fmerchant_name\":\"商户名称\",\"Foperate_name\":\"彭政\",\"Fstatus\":\"2\",\"Fstore_id\":\"12\",\"Fstore_name\":\"店铺名称\",\"Fupdate_time\":\"2022-06-06 10:48:51\"},{\"Fcreate_time\":\"2022-06-01 18:13:50\",\"Fid\":\"3\",\"Fidentify\":\"abced123f\",\"Fmerchant_id\":\"44\",\"Fmerchant_name\":\"商户名称\",\"Foperate_name\":\"彭政\",\"Fstatus\":\"1\",\"Fstore_id\":\"12\",\"Fstore_name\":\"店铺名称\",\"Fupdate_time\":\"2022-06-01 18:13:50\"},{\"Fcreate_time\":\"2022-06-01 17:32:17\",\"Fid\":\"2\",\"Fidentify\":\"15963\",\"Fmerchant_id\":\"1\",\"Fmerchant_name\":\"测试2\",\"Foperate_name\":\"彭政\",\"Fstatus\":\"2\",\"Fstore_id\":\"3\",\"Fstore_name\":\"测试门店2\",\"Fupdate_time\":\"2022-06-01 18:06:01\"}]}"
		path := "{output:{items:{updateTime:PaginateOut.#.Fupdate_time.@tostring,id:PaginateOut.#.Fid.@tostring,identify:PaginateOut.#.Fidentify.@tostring,status:PaginateOut.#.Fstatus|tonum,createTime:PaginateOut.#.Fcreate_time.@tostring,storeName:PaginateOut.#.Fstore_name.@tostring,merchantId:PaginateOut.#.Fmerchant_id.@tostring,merchantName:PaginateOut.#.Fmerchant_name.@tostring,operateName:PaginateOut.#.Foperate_name.@tostring,storeId:PaginateOut.#.Fstore_id.@tostring}|@group,pageInfo:{pageIndex:input.pageIndex.@tostring,pageSize:input.pageSize.@tostring,total:PaginateTotalOut.@tostring}}}"
		result := gjson.Get(value, path).String()
		fmt.Println(result)

	})

}
