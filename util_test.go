package gjsonmodifier

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestFindPathByValue(t *testing.T) {
	t.Run("simple top level", func(t *testing.T) {
		body := `{"pageIndex":"{.PageNumber}","pageSize":"100"}`
		path := FindPathByValue(body, "{.PageNumber}")
		assert.Equal(t, "pageIndex", path)
	})

	t.Run("nested", func(t *testing.T) {
		body := `{"data":{"page":"{.PageNumber}"}}`
		path := FindPathByValue(body, "{.PageNumber}")
		assert.Equal(t, "data.page", path)
	})

	t.Run("real business data", func(t *testing.T) {
		body := `{"_head":{"_callerServiceId":"211001","_groupNo":"1","_interface":"MerchantWallet.Lianlian.Trade.List","_invokeId":"xyxzaccount-1780490980","_msgType":"request","_remark":"","_timestamps":"1780490980","_version":"0.01"},"_param":{"login_token":"32795311326224755fa996320f4571b5::327953::1::1326224::75","merchantId":"220390584","memberCode":"prod-HSB-2026010502220390584","tradeType":"","status":"","tradeNo":"","subSysOrdNum":"","finTradeNo":"","outTradeNo":"","searchKeyword":"","isExport":"1","startAt":"2026-06-01","endAt":"2026-06-03 23:59:59","pageIndex":"{.PageNumber}","pageSize":"","md5Time":"29674849"}}`
		path := FindPathByValue(body, "{.PageNumber}")
		assert.Equal(t, "_param.pageIndex", path)
	})

	t.Run("not found", func(t *testing.T) {
		body := `{"pageIndex":"0","pageSize":"100"}`
		path := FindPathByValue(body, "{.PageNumber}")
		assert.Equal(t, "", path)
	})

	t.Run("empty body", func(t *testing.T) {
		path := FindPathByValue("", "{.PageNumber}")
		assert.Equal(t, "", path)
	})

	t.Run("invalid json", func(t *testing.T) {
		path := FindPathByValue("not json", "{.PageNumber}")
		assert.Equal(t, "", path)
	})
}
