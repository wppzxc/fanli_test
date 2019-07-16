package utils

import (
	"fmt"
	"github.com/wpp/fanli_test/pkg/types"
	"testing"
)

func Test01(t *testing.T) {
	item := types.Item{
		Id: 2581,
		ExtendDocument: "###拼多多免单来袭请注意看清要求\n###1.领取【1】元券，拍【零痕补水保湿面膜2片】选项，券后【2.8】元拍下\n###2.实发【抽纸一包】\n###3.禁止使用平台券，禁止联系商家咨询\n###4.有任何问题找群主\n###5.券无代表活动结束\n###6.拍完付款后请重新进入活动网址，点订单页面查询是否正常",
		GoodsImageUrl: "",
		GoodsName: "零痕天丝面膜补水保湿收缩毛孔美白淡斑玻尿酸精华淡化痘印男女",
		RefundAmount: "2.80",
		MinGroupPrice: "3.80",
		CouponDiscount: "1.00",
		StartTime: 1563197400,
		StopTime: 1563283800,
	}
	str := GetMsg(item)
	fmt.Printf(str)
}

func Test02(t *testing.T) {
	oldItems := []types.Item{{
		Id: 2581,
		ExtendDocument: "###拼多多免单来袭请注意看清要求\n###1.领取【1】元券，拍【零痕补水保湿面膜2片】选项，券后【2.8】元拍下\n###2.实发【抽纸一包】\n###3.禁止使用平台券，禁止联系商家咨询\n###4.有任何问题找群主\n###5.券无代表活动结束\n###6.拍完付款后请重新进入活动网址，点订单页面查询是否正常",
		GoodsImageUrl: "",
		GoodsName: "零痕天丝面膜补水保湿收缩毛孔美白淡斑玻尿酸精华淡化痘印男女",
		RefundAmount: "2.80",
		MinGroupPrice: "3.80",
		CouponDiscount: "1.00",
		StartTime: 1563197400,
		StopTime: 1563283800,
	}, {
		Id: 2582,
		ExtendDocument: "###拼多多免单来袭请注意看清要求\n###1.领取【1】元券，拍【零痕补水保湿面膜2片】选项，券后【2.8】元拍下\n###2.实发【抽纸一包】\n###3.禁止使用平台券，禁止联系商家咨询\n###4.有任何问题找群主\n###5.券无代表活动结束\n###6.拍完付款后请重新进入活动网址，点订单页面查询是否正常",
		GoodsImageUrl: "",
		GoodsName: "零痕天丝面膜补水保湿收缩毛孔美白淡斑玻尿酸精华淡化痘印男女",
		RefundAmount: "2.80",
		MinGroupPrice: "3.80",
		CouponDiscount: "1.00",
		StartTime: 1563197400,
		StopTime: 1563283800,
	}, {
		Id: 2584,
		ExtendDocument: "###拼多多免单来袭请注意看清要求\n###1.领取【1】元券，拍【零痕补水保湿面膜2片】选项，券后【2.8】元拍下\n###2.实发【抽纸一包】\n###3.禁止使用平台券，禁止联系商家咨询\n###4.有任何问题找群主\n###5.券无代表活动结束\n###6.拍完付款后请重新进入活动网址，点订单页面查询是否正常",
		GoodsImageUrl: "",
		GoodsName: "零痕天丝面膜补水保湿收缩毛孔美白淡斑玻尿酸精华淡化痘印男女",
		RefundAmount: "2.80",
		MinGroupPrice: "3.80",
		CouponDiscount: "1.00",
		StartTime: 1563197400,
		StopTime: 1563283800,
	}}
	 newItems := []types.Item{{
		 Id: 2581,
		 ExtendDocument: "###拼多多免单来袭请注意看清要求\n###1.领取【1】元券，拍【零痕补水保湿面膜2片】选项，券后【2.8】元拍下\n###2.实发【抽纸一包】\n###3.禁止使用平台券，禁止联系商家咨询\n###4.有任何问题找群主\n###5.券无代表活动结束\n###6.拍完付款后请重新进入活动网址，点订单页面查询是否正常",
		 GoodsImageUrl: "",
		 GoodsName: "零痕天丝面膜补水保湿收缩毛孔美白淡斑玻尿酸精华淡化痘印男女",
		 RefundAmount: "2.80",
		 MinGroupPrice: "3.80",
		 CouponDiscount: "1.00",
		 StartTime: 1563197400,
		 StopTime: 1563283800,
	 }, {
		 Id: 2583,
		 ExtendDocument: "###拼多多免单来袭请注意看清要求\n###1.领取【1】元券，拍【零痕补水保湿面膜2片】选项，券后【2.8】元拍下\n###2.实发【抽纸一包】\n###3.禁止使用平台券，禁止联系商家咨询\n###4.有任何问题找群主\n###5.券无代表活动结束\n###6.拍完付款后请重新进入活动网址，点订单页面查询是否正常",
		 GoodsImageUrl: "",
		 GoodsName: "零痕天丝面膜补水保湿收缩毛孔美白淡斑玻尿酸精华淡化痘印男女",
		 RefundAmount: "2.80",
		 MinGroupPrice: "3.80",
		 CouponDiscount: "1.00",
		 StartTime: 1563197400,
		 StopTime: 1563283800,
	 }, {
		 Id: 2582,
		 ExtendDocument: "###拼多多免单来袭请注意看清要求\n###1.领取【1】元券，拍【零痕补水保湿面膜2片】选项，券后【2.8】元拍下\n###2.实发【抽纸一包】\n###3.禁止使用平台券，禁止联系商家咨询\n###4.有任何问题找群主\n###5.券无代表活动结束\n###6.拍完付款后请重新进入活动网址，点订单页面查询是否正常",
		 GoodsImageUrl: "",
		 GoodsName: "零痕天丝面膜补水保湿收缩毛孔美白淡斑玻尿酸精华淡化痘印男女",
		 RefundAmount: "2.80",
		 MinGroupPrice: "3.80",
		 CouponDiscount: "1.00",
		 StartTime: 1563197400,
		 StopTime: 1563283800,
	 }, {
		 Id: 2585,
		 ExtendDocument: "###拼多多免单来袭请注意看清要求\n###1.领取【1】元券，拍【零痕补水保湿面膜2片】选项，券后【2.8】元拍下\n###2.实发【抽纸一包】\n###3.禁止使用平台券，禁止联系商家咨询\n###4.有任何问题找群主\n###5.券无代表活动结束\n###6.拍完付款后请重新进入活动网址，点订单页面查询是否正常",
		 GoodsImageUrl: "",
		 GoodsName: "零痕天丝面膜补水保湿收缩毛孔美白淡斑玻尿酸精华淡化痘印男女",
		 RefundAmount: "2.80",
		 MinGroupPrice: "3.80",
		 CouponDiscount: "1.00",
		 StartTime: 1563197400,
		 StopTime: 1563283800,
	 }}
	 
	 result := GetDiffItems(oldItems, newItems)
	 fmt.Println(result)
}