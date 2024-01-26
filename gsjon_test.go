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
func TestTobool(t *testing.T) {
	jsonstr := `
	[{"questionId":"12057","questionName":"全新机(包装盒无破损,配件齐全且原装,可无原 机膜和卡针)","classId":"1","type":1,"single":"1"},{"questionId":"12097","questionName":"机身弯曲情况","classId":"3","type":3,"single":"1"},{"questionId":"12066","questionName":"屏幕外观","classId":"3","type":3,"single":"2"},{"questionId":"12077","questionName":"屏幕显示","classId":"3","type":2,"single":"2"},{"questionId":"12088","questionName":"电池健康效率","classId":"3","type":2,"single":"2"},{"questionId":"12100","questionName":"维修情况","classId":"3","type":2,"single":"2"},{"questionId":"12106","questionName":"零件维修情况","classId":"3","type":3},{"questionId":"12115","questionName":"受潮状况","classId":"3","type":2},{"questionId":"12119","questionName":"开机状态","classId":"3","type":2},{"questionId":"9368","questionName":"是否全新","classId":"3","type":4}]
	`
	path := "@this.#.single.@tobool"
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
		path := "{output:{items:{updateTime:PaginateOut.#.Fupdate_time.@tostring,id:PaginateOut.#.Fid.@tostring,identify:PaginateOut.#.Fidentify.@tostring,status:PaginateOut.#.Fstatus.@tonum,createTime:PaginateOut.#.Fcreate_time.@tostring,storeName:PaginateOut.#.Fstore_name.@tostring,merchantId:PaginateOut.#.Fmerchant_id.@tostring,merchantName:PaginateOut.#.Fmerchant_name.@tostring,operateName:PaginateOut.#.Foperate_name.@tostring,storeId:PaginateOut.#.Fstore_id.@tostring}|@group,pageInfo:{pageIndex:input.pageIndex.@tostring,pageSize:input.pageSize.@tostring,total:PaginateTotalOut.@tostring}}}"
		result := gjson.Get(value, path).String()
		fmt.Println(result)

	})

	t.Run("complex2", func(t *testing.T) {
		value := `{"code":"0","message":"ok","items":[{"id":9,"advertiserID":"123","name":"新年豪礼","position":"123","did":120,"landingPage":"http://domain.com/active.html","beginAt":"2023-01-12 00:00:00","endAt":"2023-01-30 00:00:00","createdAt":"2023-05-31 18:09:02","updatedAt":"2023-05-31 18:09:02"},{"id":8,"advertiserID":"123","name":"新年豪礼","position":"123","did":120,"landingPage":"http://domain.com/active.html","beginAt":"2023-01-12 00:00:00","endAt":"2023-01-30 00:00:00","createdAt":"2023-05-31 18:07:10","updatedAt":"2023-05-31 18:07:10"},{"id":7,"advertiserID":"123","name":"新年豪礼","position":"123","did":120,"landingPage":"http://domain.com/active.html","beginAt":"2023-01-12 00:00:00","endAt":"2023-01-30 00:00:00","createdAt":"2023-05-31 18:06:34","updatedAt":"2023-05-31 18:06:34"},{"id":6,"advertiserID":"123","name":"新年豪礼","position":"123","did":120,"landingPage":"http://domain.com/active.html","beginAt":"2023-01-12 00:00:00","endAt":"2023-01-30 00:00:00","createdAt":"2023-05-31 16:53:12","updatedAt":"2023-05-31 16:53:12"},{"id":5,"advertiserID":"123","name":"新年豪礼","position":"123","did":120,"landingPage":"http://domain.com/active.html","beginAt":"2023-01-12 00:00:00","endAt":"2023-01-30 00:00:00","createdAt":"2023-05-31 14:26:20","updatedAt":"2023-05-31 14:26:20"},{"id":4,"advertiserID":"123","name":"新年豪礼","position":"123","did":120,"landingPage":"http://domain.com/active.html","beginAt":"2023-01-12 00:00:00","endAt":"2023-01-30 00:00:00","createdAt":"2023-05-31 14:18:50","updatedAt":"2023-05-31 14:18:50"},{"id":3,"advertiserID":"123","name":"新年豪礼","position":"123","did":120,"landingPage":"http://domain.com/active.html","beginAt":"2023-01-12 00:00:00","endAt":"2023-01-30 00:00:00","createdAt":"2023-05-31 14:18:08","updatedAt":"2023-05-31 14:18:08"},{"id":2,"advertiserID":"123","name":"新年豪礼","position":"123","did":120,"landingPage":"http://domain.com/active.html","beginAt":"2023-01-12 00:00:00","endAt":"2023-01-30 00:00:00","createdAt":"2023-05-31 14:16:08","updatedAt":"2023-05-31 14:16:08"}],"pagination":{"index":0,"size":10,"total":8}}`
		path := "{out:{code:code.@tostring,message:message.@tostring,items:{id:items.#.id.@tostring,name:items.#.name.@tostring,position:items.#.position.@tostring,landingPage:items.#.landingPage.@tostring,updatedAt:items.#.updatedAt.@tostring,advertiserId:items.#.advertiserId.@tostring,did:items.#.did.@tostring,beginAt:items.#.beginAt.@tostring,endAt:items.#.endAt.@tostring,createdAt:items.#.createdAt.@tostring}|@group,pagination:{index:pagination.index.@tostring,size:pagination.size.@tostring,total:pagination.total.@tostring}}}"
		result := gjson.Get(value, path).String()
		fmt.Println(result)
	})

}

func TestIn(t *testing.T) {
	t.Run("simple", func(t *testing.T) {
		data := `
	{
	  "users": [
	    {"name": "Alice", "age": 25},
	    {"name": "Bob", "age": 30},
	    {"name": "Charlie", "age": 35}
	  ]
	}
	`
		result := gjson.Get(data, "users.#(age.@in:25,30)#")
		fmt.Println(result)
	})
}

func TestTestQuery(t *testing.T) {

	t.Run("complex", func(t *testing.T) {
		data := "{\"Description\":\"获取一个或者多个橱窗下的所有上线广告\",\"Path\":\"/api/v1/banner/list\",\"Method\":\"POST\",\"Header\":[{\"fullname\":\"content-type\",\"type\":\"string\",\"required\":\"是\",\"default\":\"application/json\",\"description\":\"文件格式\",\"example\":\"\"},{\"fullname\":\"appid\",\"type\":\"string\",\"required\":\"是\",\"default\":\"\",\"description\":\"访问服务的备案id\",\"example\":\"er45e7r\"},{\"fullname\":\"signature\",\"type\":\"string\",\"required\":\"是\",\"default\":\"\",\"description\":\"签名,外网访问需开启签名\",\"example\":\"erefdsf154\"}],\"common\":{\"header\":\"\\n\\n|参数名|类型|必选|默认值|说明|示例|\\n|:----    |:---|:----- |:-----   |:-----   |:-----   |\\n|content-type| string|是|application/json|文件格式||\\n|appid|string|是||访问服务的备案id|er45e7r|\\n|signature|string|是||签名,外网访问需开启签名|erefdsf154|\\n\",\"pagination\":{\"RequestBody\":[{\"fullname\":\"index\",\"type\":\"string\",\"format\":\"int\",\"required\":\"是\",\"description\":\"页索引,0开始\",\"default\":\"0\",\"example\":\"\"},{\"fullname\":\"size\",\"type\":\"string\",\"format\":\"int\",\"required\":\"是\",\"description\":\"每页数量\",\"default\":\"10\",\"example\":\"\"}],\"ResponseBody\":[{\"fullname\":\"pagination\",\"type\":\"object\",\"description\":\"对象\",\"example\":\"\"},{\"fullname\":\"pagination.index\",\"type\":\"string\",\"description\":\"{{Get . `RequestBody.#(fullname=\\\"index\\\")#\",\"example\":\"0.description`}}\"},{\"fullname\":\"pagination.size\",\"type\":\"string\",\"description\":\"{{Get . `RequestBody.#(fullname=\\\"size\\\")#\",\"example\":\"0.description`}}\"},{\"fullname\":\"pagination.total\",\"type\":\"string\",\"description\":\"总数\",\"example\":\"60\"}]},\"code\":{\"ResponseBody\":[{\"fullname\":\"code\",\"type\":\"string\",\"description\":\"业务状态码\",\"example\":\"0\"},{\"fullname\":\"message\",\"type\":\"string\",\"description\":\"业务提示\",\"example\":\"ok\"}]},\"response\":{\"ok\":\"\\n```json \\n{\\n  \\\"code\\\": \\\"0\\\",\\n  \\\"message\\\": \\\"ok\\\"\\n}\\n``` \\n\",\"error\":\"\\n```json\\n{\\n    \\\"code\\\":\\\"xxx\\\",\\n    \\\"message\\\":\\\"xxx提示\\\"\\n}\\n```\\n\"},\"curlTpl\":\"\\n\\n```bash\\nRegestBody='[[.RequestBody]]'\\n[[.Bash]]\\ncurl -X[[.Method]] [[range $k,$v:=.Headers]] [[if (eq $k \\\"signature\\\") ]] [[$v = \\\"$signature\\\"]] [[end]]  -H '[[$k]]:\\\"[[$v]]\\\"' [[end]] -d '$RequestBody' '[[.URL]]'\\n```\\n\"},\"Server\":[{\"name\":\"dev\",\"host\":\"http://localhost:801\",\"proxy\":\"\",\"description\":\"开发环境\"},{\"name\":\"test\",\"host\":\"test.domain.com\",\"proxy\":\"\",\"description\":\"测试环境\"},{\"name\":\"prod\",\"host\":\"http://domain.com\",\"proxy\":\"\",\"description\":\"正式环境\"}],\"Contact\":[{\"name\":\"彭政\",\"phone\":\"15999646794\"}],\"Service\":{\"preScript\":[{\"language\":\"bash\",\"script\":\"signature=\\\"baqsh\\\"\"},{\"language\":\"javascript\",\"script\":\"var signature=\\\"javascript\\\"\"}]},\"Variable\":[{\"name\":\"serviceId\",\"value\":\"1234,xyuientg,74ere\",\"description\":\"服务ID\"}],\"adminWindowUpdateRequestBody\":[{\"fullname\":\"id\",\"type\":\"string\",\"format\":\"int\",\"required\":\"是\",\"description\":\"主键\",\"default\":\"\",\"example\":\"1\"},{\"fullname\":\"endpoint\",\"type\":\"string\",\"format\":\"\",\"required\":\"是\",\"description\":\"广告位终端(app名称、具体web网站等)\",\"default\":\"\",\"example\":\"kele_youpin_app_ios\"},{\"fullname\":\"code\",\"type\":\"string\",\"format\":\"\",\"required\":\"是\",\"description\":\"位置编码\",\"default\":\"\",\"example\":\"index_head_1\"},{\"fullname\":\"title\",\"type\":\"string\",\"format\":\"\",\"required\":\"是\",\"description\":\"位置名称\",\"default\":\"\",\"example\":\"可口可乐\"},{\"fullname\":\"remark\",\"type\":\"string\",\"format\":\"\",\"required\":\"是\",\"description\":\"位置描述\",\"default\":\"\",\"example\":\"赞助广告\"},{\"fullname\":\"contentTypes\",\"type\":\"string\",\"format\":\"\",\"required\":\"是\",\"description\":\"广告素材(类型),text-文字,image-图片,vido-视频,多个逗号分隔\",\"default\":\"\",\"example\":\"text\"},{\"fullname\":\"width\",\"type\":\"string\",\"format\":\"int\",\"required\":\"是\",\"description\":\"橱窗宽度,单位px\",\"default\":\"985\",\"example\":\"\"},{\"fullname\":\"high\",\"type\":\"string\",\"format\":\"int\",\"required\":\"是\",\"description\":\"橱窗高度,单位px\",\"default\":\"211\",\"example\":\"\"}],\"RequestBody\":[{\"fullname\":\"{{Table . adminWindowUpdateRequestBody.#(fullname#(\\\"endpoint\\\",\\\"code\\\")) \\\"\",\"type\":\"{#.fullname.@basePath}\",\"format\":\"string\",\"required\":\"{#.format}\",\"description\":\"是\",\"default\":\"{#.description}\",\"example\":\"\"},{\"fullname\":\"conditionValue\",\"type\":\"string\",\"format\":\"\",\"required\":\"是\",\"description\":\"条件数据\",\"default\":\"json字符串,用来收集条件数据,比如地区、用户标签值等,具体字段和后台支持的条件相匹配\",\"example\":\"bannar\"}],\"ResponseBody\":[{\"fullname\":\"code.@basePath\",\"type\":\"string\",\"description\":\"业务状态码\",\"example\":\"0\"},{\"fullname\":\"items[].id\",\"type\":\"string\",\"description\":\"id\",\"example\":\"0\"},{\"fullname\":\"items[].title\",\"type\":\"string\",\"description\":\"banner标题\",\"example\":\"\"},{\"fullname\":\"items[].image\",\"type\":\"string\",\"description\":\"图片地址\",\"example\":\"http://image.service.cn/banner_1.jpg\"},{\"fullname\":\"items[].link\",\"type\":\"string\",\"description\":\"图片地址\",\"example\":\"http://image.service.cn/banner_1.jpg\"},{\"fullname\":\"items[].link\",\"type\":\"string\",\"description\":\"图片地址\",\"example\":\"http://image.service.cn/banner_1.jpg\"}]}"
		path := `adminWindowUpdateRequestBody.#(fullname.@in:"endpoint,code")#`
		out := TestQuery(data, path)
		fmt.Println(out)
	})
}

func TestEval(t *testing.T) {

	data := `[{"fullname":"items[].id"},{"fullname":"items[].name"}]`
	t.Run("simple", func(t *testing.T) {
		path := `#.fullname.@basePath.@eval:value == "id" ? "是":"否"`
		out := TestQuery(data, path)
		fmt.Println(out)
	})
	t.Run("filter", func(t *testing.T) {
		path := `#(fullname.@basePath.@eval:(value!="name"&&value!="updated"))#`
		out := TestQuery(data, path)
		fmt.Println(out)
	})

	t.Run("inArray", func(t *testing.T) {
		path := `#(fullname.@basePath.@eval:(inArray(["name"])))#`
		out := TestQuery(data, path)
		fmt.Println(out)
	})
	t.Run("notInArray", func(t *testing.T) {
		path := `#(fullname.@basePath.@eval:(notInArray(["id"])))#`
		out := TestQuery(data, path)
		fmt.Println(out)
	})
	t.Run("kvMap", func(t *testing.T) {
		path := `#.fullname.@basePath.@eval:(kvMap({"id":"id1"},false))`
		out := TestQuery(data, path)
		fmt.Println(out)
	})

	t.Run("collection", func(t *testing.T) {
		path := `#(fullname.@basePath.@eval:(collection.any(["id","name1"],func(index,item){return value==item})))#`
		out := TestQuery(data, path)
		fmt.Println(out)
	})

	t.Run("basePathAddPrefix", func(t *testing.T) {
		path := `#.fullname.@basePath.@eval:inArray(["id"])?"user"+firstUpper(value):""`
		out := TestQuery(data, path)
		fmt.Println(out)
	})

}

func TestBasePathAddPrefix(t *testing.T) {

	data := `[{"fullname":"items[].id"},{"fullname":"items[].name"}]`
	t.Run("string", func(t *testing.T) {
		path := `#.fullname.@basePathAddPrefix:"user"`
		out := TestQuery(data, path)
		fmt.Println(out)
	})
	t.Run("eval", func(t *testing.T) {
		path := `#.fullname.@basePathAddPrefix:kvMap({"id":"user"},"info")`
		out := TestQuery(data, path)
		fmt.Println(out)
	})
	t.Run("evalNoDefault", func(t *testing.T) {
		path := `#.fullname.@basePathAddPrefix:kvMap({"id":"user"},"")`
		out := TestQuery(data, path)
		fmt.Println(out)
	})

}

func TestParseSubSelectors(t *testing.T) {
	t.Run("simple path", func(t *testing.T) {
		path := `id.name`
		sels, out, ok := ParseSubSelectors(path)
		fmt.Println(sels, out, ok)
	})
	t.Run("complete path", func(t *testing.T) {
		path := `{name:{id:id.name}|@group}|a`
		sels, out, ok := ParseSubSelectors(path)
		fmt.Println(sels, out, ok)
	})

}

func TestGetAllPath(t *testing.T) {
	data := `{"code":0,"services":[{"id":6,"servers":[]},{"id":1,"servers":[{"name":"dev","title":"开发环境"},{"name":"prod","title":"开发环境"}]}]}`
	paths := GetAllPath(data)
	fmt.Println(paths)
}

func TestInnerGroup(t *testing.T) {
	t.Run("simple", func(t *testing.T) {
		data := `{"name":[[],[[1,2],[3,4]]],"title":[[],["开发环境","正式环境"]]}`
		expected := `[[],[{"name":"dev","title":"开发环境"},{"name":"prod","title":"正式环境"}]]`
		_ = expected
		out := groupPlus(data, "0")
		fmt.Println(out)
	})

	t.Run("complexe 3", func(t *testing.T) {
		path := `{code:code.@tostring,message:message.@tostring,services:{navs:{name:services.#.navs.#.name.@tostring,title:services.#.navs.#.title.@tostring,route:services.#.navs.#.route.@tostring,sort:services.#.navs.#.sort.@tostring}|@groupPlus:1,id:services.#.id.@tostring,name:services.#.name.@tostring,title:services.#.title.@tostring,document:services.#.document.@tostring,createdAt:services.#.createdAt.@tostring,updatedAt:services.#.updatedAt.@tostring,servers:{name:services.#.servers.#.name.@tostring,title:services.#.servers.#.title.@tostring}|@groupPlus:1}|@groupPlus:0,pagination:{index:pagination.index.@tostring,size:pagination.size.@tostring,total:pagination.total.@tostring}}`
		data := `{"code":0,"message":"","services":[{"id":6,"name":"advertise1","title":"广告服务","document":"","createdAt":"2023-12-02 23:01:04","updatedAt":"2023-12-02 23:01:04","servers":[],"navs":[]},{"id":1,"name":"advertise","title":"广告服务","document":"","createdAt":"2023-11-25 22:32:16","updatedAt":"2023-11-25 22:32:16","servers":[{"name":"dev","title":"开发环境"},{"name":"prod","title":"开发环境"}],"navs":[{"name":"creative","title":"广告创意","route":"/advertise/creativeList","sort":99},{"name":"plan","title":"广告计划","route":"/advertise/planList","sort":98},{"name":"window","title":"橱窗","route":"/advertise/windowList","sort":97},{"name":"crativeList","title":"广告服务","route":"/creativeList","sort":4}]}],"pagination":{"index":0,"size":10,"total":2}}`
		newData := gjson.Get(data, path).String()
		fmt.Println(newData)
	})
}

func TestToArray(t *testing.T) {

	data := `{"db":{"export":{"export_template":{"id":{"database":"export","table":"export_template","name":"id","goType":"int","dbType":"bigint(11) unsigned","comment":"自增ID","nullable":"false","enums":"","autoIncrement":"true","default":"","onUpdate":"false","unsigned":"true","size":"11"},"tenant_id":{"database":"export","table":"export_template","name":"tenant_id","goType":"string","dbType":"varchar(128)","comment":"租户标识","nullable":"false","enums":"","autoIncrement":"false","default":"","onUpdate":"false","unsigned":"false","size":"128"},"template_name":{"database":"export","table":"export_template","name":"template_name","goType":"string","dbType":"varchar(64)","comment":"模板名称","nullable":"false","enums":"","autoIncrement":"false","default":"","onUpdate":"false","unsigned":"false","size":"64"},"title":{"database":"export","table":"export_template","name":"title","goType":"string","dbType":"varchar(64)","comment":"导出任务标题","nullable":"false","enums":"","autoIncrement":"false","default":"","onUpdate":"false","unsigned":"false","size":"64"},"filed_meta":{"database":"export","table":"export_template","name":"filed_meta","goType":"string","dbType":"varchar(512)","comment":"导出文件标题和数据字段映射关系[{\\\\\\\"name\\\\\\\":\\\\\\\"data_key\\\\\\\",\\\\\\\"title\\\\\\\":\\\\\\\"数据标题\\\\\\\"}]","nullable":"false","enums":"","autoIncrement":"false","default":"","onUpdate":"false","unsigned":"false","size":"512"},"http_tpl":{"database":"export","table":"export_template","name":"http_tpl","goType":"string","dbType":"text","comment":"代理发起http请求模板","nullable":"false","enums":"","autoIncrement":"false","default":"","onUpdate":"false","unsigned":"false","size":"0"},"http_script":{"database":"export","table":"export_template","name":"http_script","goType":"string","dbType":"text","comment":"请求前后执行的脚本","nullable":"false","enums":"","autoIncrement":"false","default":"","onUpdate":"false","unsigned":"false","size":"0"},"callback_tpl":{"database":"export","table":"export_template","name":"callback_tpl","goType":"string","dbType":"text","comment":"导出结束后回调请求模板","nullable":"false","enums":"","autoIncrement":"false","default":"","onUpdate":"false","unsigned":"false","size":"0"},"callback_script":{"database":"export","table":"export_template","name":"callback_script","goType":"string","dbType":"text","comment":"回调前后执行的脚本","nullable":"false","enums":"","autoIncrement":"false","default":"","onUpdate":"false","unsigned":"false","size":"0"},"request_interval":{"database":"export","table":"export_template","name":"request_interval","goType":"string","dbType":"varchar(10)","comment":"循环请求获取数据的间隔时间,单位毫秒-ms,秒-s,小时h","nullable":"false","enums":"","autoIncrement":"false","default":"1s","onUpdate":"false","unsigned":"false","size":"10"},"max_exec_time":{"database":"export","table":"export_template","name":"max_exec_time","goType":"string","dbType":"varchar(15)","comment":"任务处理最长时间,单位秒-s","nullable":"false","enums":"","autoIncrement":"false","default":"","onUpdate":"false","unsigned":"false","size":"15"},"expired":{"database":"export","table":"export_template","name":"expired","goType":"string","dbType":"varchar(10)","comment":"文件过期时间,单位小时-h,月-m","nullable":"false","enums":"","autoIncrement":"false","default":"1d","onUpdate":"false","unsigned":"false","size":"10"},"async":{"database":"export","table":"export_template","name":"async","goType":"string","dbType":"enum('true','false')","comment":"是否异步执行(true-是,false-否)","nullable":"false","enums":"true,false","autoIncrement":"false","default":"true","onUpdate":"false","unsigned":"false","size":"0"},"remark":{"database":"export","table":"export_template","name":"remark","goType":"string","dbType":"varchar(256)","comment":"备注","nullable":"false","enums":"","autoIncrement":"false","default":"","onUpdate":"false","unsigned":"false","size":"256"},"creator_id":{"database":"export","table":"export_template","name":"creator_id","goType":"string","dbType":"varchar(64)","comment":"创建者ID","nullable":"false","enums":"","autoIncrement":"false","default":"","onUpdate":"false","unsigned":"false","size":"64"},"creator_name":{"database":"export","table":"export_template","name":"creator_name","goType":"string","dbType":"varchar(64)","comment":"创建者名称","nullable":"false","enums":"","autoIncrement":"false","default":"","onUpdate":"false","unsigned":"false","size":"64"},"deleted_at":{"database":"export","table":"export_template","name":"deleted_at","goType":"string","dbType":"datetime","comment":"删除时间","nullable":"true","enums":"","autoIncrement":"false","default":"NULL","onUpdate":"false","unsigned":"false","size":"0"},"created_at":{"database":"export","table":"export_template","name":"created_at","goType":"string","dbType":"datetime","comment":"创建时间","nullable":"false","enums":"","autoIncrement":"false","default":"current_timestamp()","onUpdate":"false","unsigned":"false","size":"0"},"updated_at":{"database":"export","table":"export_template","name":"updated_at","goType":"string","dbType":"datetime","comment":"更新时间","nullable":"false","enums":"","autoIncrement":"false","default":"current_timestamp()","onUpdate":"true","unsigned":"false","size":"0"}},"export_task":{"id":{"database":"export","table":"export_task","name":"id","goType":"int","dbType":"bigint(11) unsigned","comment":"自增ID","nullable":"false","enums":"","autoIncrement":"true","default":"","onUpdate":"false","unsigned":"true","size":"11"},"template_id":{"database":"export","table":"export_task","name":"template_id","goType":"int","dbType":"bigint(11) unsigned","comment":"export_template 表ID","nullable":"false","enums":"","autoIncrement":"false","default":"0","onUpdate":"false","unsigned":"true","size":"11"},"tenant_id":{"database":"export","table":"export_task","name":"tenant_id","goType":"string","dbType":"varchar(128)","comment":"租户标识","nullable":"false","enums":"","autoIncrement":"false","default":"","onUpdate":"false","unsigned":"false","size":"128"},"template_name":{"database":"export","table":"export_task","name":"template_name","goType":"string","dbType":"varchar(64)","comment":"模板名称","nullable":"false","enums":"","autoIncrement":"false","default":"","onUpdate":"false","unsigned":"false","size":"64"},"creator_id":{"database":"export","table":"export_task","name":"creator_id","goType":"string","dbType":"varchar(64)","comment":"创建者ID","nullable":"false","enums":"","autoIncrement":"false","default":"","onUpdate":"false","unsigned":"false","size":"64"},"creator_name":{"database":"export","table":"export_task","name":"creator_name","goType":"string","dbType":"varchar(64)","comment":"创建者名称","nullable":"false","enums":"","autoIncrement":"false","default":"","onUpdate":"false","unsigned":"false","size":"64"},"filename":{"database":"export","table":"export_task","name":"filename","goType":"string","dbType":"varchar(256)","comment":"文件名","nullable":"false","enums":"","autoIncrement":"false","default":"","onUpdate":"false","unsigned":"false","size":"256"},"title":{"database":"export","table":"export_task","name":"title","goType":"string","dbType":"varchar(64)","comment":"任务标题","nullable":"false","enums":"","autoIncrement":"false","default":"","onUpdate":"false","unsigned":"false","size":"64"},"md5":{"database":"export","table":"export_task","name":"md5","goType":"string","dbType":"varchar(64)","comment":"指纹","nullable":"false","enums":"","autoIncrement":"false","default":"","onUpdate":"false","unsigned":"false","size":"64"},"status":{"database":"export","table":"export_task","name":"status","goType":"string","dbType":"enum('queuing','exporting','success','fail')","comment":"任务状态(queuing-排队中,exporting-正在导出,success-成功,fail-失败)","nullable":"false","enums":"queuing,exporting,success,fail","autoIncrement":"false","default":"queuing","onUpdate":"false","unsigned":"false","size":"0"},"callback_status":{"database":"export","table":"export_task","name":"callback_status","goType":"string","dbType":"enum('init','doing','success','fail')","comment":"回调状态(init-初始化,doing-回调中,success-成功,fail-失败)","nullable":"false","enums":"init,doing,success,fail","autoIncrement":"false","default":"init","onUpdate":"false","unsigned":"false","size":"0"},"timeout":{"database":"export","table":"export_task","name":"timeout","goType":"string","dbType":"varchar(15)","comment":"任务处理超时时间","nullable":"false","enums":"","autoIncrement":"false","default":"","onUpdate":"false","unsigned":"false","size":"15"},"size":{"database":"export","table":"export_task","name":"size","goType":"int","dbType":"int(11) unsigned","comment":"文件大小,单位B","nullable":"false","enums":"","autoIncrement":"false","default":"0","onUpdate":"false","unsigned":"true","size":"11"},"url":{"database":"export","table":"export_task","name":"url","goType":"string","dbType":"varchar(256)","comment":"下载地址","nullable":"false","enums":"","autoIncrement":"false","default":"","onUpdate":"false","unsigned":"false","size":"256"},"remark":{"database":"export","table":"export_task","name":"remark","goType":"string","dbType":"varchar(256)","comment":"备注","nullable":"false","enums":"","autoIncrement":"false","default":"","onUpdate":"false","unsigned":"false","size":"256"},"expired_at":{"database":"export","table":"export_task","name":"expired_at","goType":"string","dbType":"datetime","comment":"文件过期时间","nullable":"true","enums":"","autoIncrement":"false","default":"NULL","onUpdate":"false","unsigned":"false","size":"0"},"created_at":{"database":"export","table":"export_task","name":"created_at","goType":"string","dbType":"datetime","comment":"创建时间","nullable":"false","enums":"","autoIncrement":"false","default":"current_timestamp()","onUpdate":"false","unsigned":"false","size":"0"},"updated_at":{"database":"export","table":"export_task","name":"updated_at","goType":"string","dbType":"datetime","comment":"更新时间","nullable":"false","enums":"","autoIncrement":"false","default":"current_timestamp()","onUpdate":"true","unsigned":"false","size":"0"}}}}}`
	t.Run("no arg", func(t *testing.T) {
		path := "db.export.export_template.@toArray"
		out := gjson.Get(data, path).String()
		fmt.Println(out)
	})

	t.Run("with arg", func(t *testing.T) {
		path := "db.export.export_template.@toArray:identity"
		out := gjson.Get(data, path).String()
		fmt.Println(out)
	})
}

func TestToMap(t *testing.T) {
	data := `[{"identity":"callback_tpl","database":"export","table":"export_template","name":"callback_tpl","goType":"string","dbType":"text","comment":"导出结束后回调请求模板","nullable":"false","enums":"","autoIncrement":"false","default":"","onUpdate":"false","unsigned":"false","size":"0"},{"identity":"max_exec_time","database":"export","table":"export_template","name":"max_exec_time","goType":"string","dbType":"varchar(15)","comment":"任务处理最长时间,单位秒-s","nullable":"false","enums":"","autoIncrement":"false","default":"","onUpdate":"false","unsigned":"false","size":"15"},{"identity":"created_at","database":"export","table":"export_template","name":"created_at","goType":"string","dbType":"datetime","comment":"创建时间","nullable":"false","enums":"","autoIncrement":"false","default":"current_timestamp()","onUpdate":"false","unsigned":"false","size":"0"},{"identity":"title","database":"export","table":"export_template","name":"title","goType":"string","dbType":"varchar(64)","comment":"导出任务标题","nullable":"false","enums":"","autoIncrement":"false","default":"","onUpdate":"false","unsigned":"false","size":"64"},{"identity":"http_tpl","database":"export","table":"export_template","name":"http_tpl","goType":"string","dbType":"text","comment":"代理发起http请求模板","nullable":"false","enums":"","autoIncrement":"false","default":"","onUpdate":"false","unsigned":"false","size":"0"},{"identity":"http_script","database":"export","table":"export_template","name":"http_script","goType":"string","dbType":"text","comment":"请求前后执行的脚本","nullable":"false","enums":"","autoIncrement":"false","default":"","onUpdate":"false","unsigned":"false","size":"0"},{"identity":"callback_script","database":"export","table":"export_template","name":"callback_script","goType":"string","dbType":"text","comment":"回调前后执行的脚本","nullable":"false","enums":"","autoIncrement":"false","default":"","onUpdate":"false","unsigned":"false","size":"0"},{"identity":"request_interval","database":"export","table":"export_template","name":"request_interval","goType":"string","dbType":"varchar(10)","comment":"循环请求获取数据的间隔时间,单位毫秒-ms,秒-s,小时h","nullable":"false","enums":"","autoIncrement":"false","default":"1s","onUpdate":"false","unsigned":"false","size":"10"},{"identity":"expired","database":"export","table":"export_template","name":"expired","goType":"string","dbType":"varchar(10)","comment":"文件过期时间,单位小时-h,月-m","nullable":"false","enums":"","autoIncrement":"false","default":"1d","onUpdate":"false","unsigned":"false","size":"10"},{"identity":"remark","database":"export","table":"export_template","name":"remark","goType":"string","dbType":"varchar(256)","comment":"备注","nullable":"false","enums":"","autoIncrement":"false","default":"","onUpdate":"false","unsigned":"false","size":"256"},{"identity":"creator_id","database":"export","table":"export_template","name":"creator_id","goType":"string","dbType":"varchar(64)","comment":"创建者ID","nullable":"false","enums":"","autoIncrement":"false","default":"","onUpdate":"false","unsigned":"false","size":"64"},{"identity":"id","database":"export","table":"export_template","name":"id","goType":"int","dbType":"bigint(11) unsigned","comment":"自增ID","nullable":"false","enums":"","autoIncrement":"true","default":"","onUpdate":"false","unsigned":"true","size":"11"},{"identity":"template_name","database":"export","table":"export_template","name":"template_name","goType":"string","dbType":"varchar(64)","comment":"模板名称","nullable":"false","enums":"","autoIncrement":"false","default":"","onUpdate":"false","unsigned":"false","size":"64"},{"identity":"deleted_at","database":"export","table":"export_template","name":"deleted_at","goType":"string","dbType":"datetime","comment":"删除时间","nullable":"true","enums":"","autoIncrement":"false","default":"NULL","onUpdate":"false","unsigned":"false","size":"0"},{"identity":"updated_at","database":"export","table":"export_template","name":"updated_at","goType":"string","dbType":"datetime","comment":"更新时间","nullable":"false","enums":"","autoIncrement":"false","default":"current_timestamp()","onUpdate":"true","unsigned":"false","size":"0"},{"identity":"async","database":"export","table":"export_template","name":"async","goType":"string","dbType":"enum('true','false')","comment":"是否异步执行(true-是,false-否)","nullable":"false","enums":"true,false","autoIncrement":"false","default":"true","onUpdate":"false","unsigned":"false","size":"0"},{"identity":"creator_name","database":"export","table":"export_template","name":"creator_name","goType":"string","dbType":"varchar(64)","comment":"创建者名称","nullable":"false","enums":"","autoIncrement":"false","default":"","onUpdate":"false","unsigned":"false","size":"64"},{"identity":"tenant_id","database":"export","table":"export_template","name":"tenant_id","goType":"string","dbType":"varchar(128)","comment":"租户标识","nullable":"false","enums":"","autoIncrement":"false","default":"","onUpdate":"false","unsigned":"false","size":"128"},{"identity":"filed_meta","database":"export","table":"export_template","name":"filed_meta","goType":"string","dbType":"varchar(512)","comment":"导出文件标题和数据字段映射关系[{\\\\\\\\\\\\\\\"name\\\\\\\\\\\\\\\":\\\\\\\\\\\\\\\"data_key\\\\\\\\\\\\\\\",\\\\\\\\\\\\\\\"title\\\\\\\\\\\\\\\":\\\\\\\\\\\\\\\"数据标题\\\\\\\\\\\\\\\"}]","nullable":"false","enums":"","autoIncrement":"false","default":"","onUpdate":"false","unsigned":"false","size":"512"}]`
	path := `@this.@toMap:name|callback_tpl`
	out := gjson.Get(data, path).String()
	fmt.Println(out)
}
