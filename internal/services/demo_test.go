package services

//import (
//	. "github.com/smartystreets/goconvey/convey"
//)
//
//func MockRiskCtlService(t *testing.T) (riskctl.RiskCtlService, func()) {
//	//redis, teardownRedis := tests.MockRedis(t)
//	riskctlSvc := NewRiskCtlService(nil, nil)
//	teardown := func() {
//		//teardownRedis()
//	}
//	return riskctlSvc, teardown
//}
//
//func TestRiskCtlService_CheckPay(t *testing.T) {
//	riskctlSvc, teardown := MockRiskCtlService(t)
//	defer teardown()
//
//	// 金额不超过限制允许支付
//	SkipConvey("under-limit-checkpay-result", t, func() {
//		//So(,ShouldEqual,riskctl.SuccessCheckPayResult)
//	})
//	// 同一个城市允许支付
//	SkipConvey("same-city-checkpay-result", t, func() {
//		//So(,ShouldEqual,riskctl.SuccessCheckPayResult)
//	})
//	// 不同城市 && 风控门店 && 超过限制金额
//	SkipConvey("diff-city-risk-meid", t, func() {
//		//So(,ShouldEqual,riskctl.OffsiteCheckPayResult)
//	})
//	// 不同城市 && 非风控门店, 允许支付
//	SkipConvey("diff-city-no-risk-meid", t, func() {
//		//So(,ShouldEqual,riskctl.SuccessCheckPayResult)
//	})
//	// 不同城市 && 风控门店 && 不超过限制金额
//	Convey("normal: diff-city-risk-meid-under-limit", t, func() {
//		r := &riskctl.CheckPay{
//			ClientIP: "1.12.7.255",
//		}
//		_, err := riskctlSvc.CheckPay(r)
//		fmt.Println(1111, err)
//		So(err, ShouldEqual, nil)
//	})
//}
