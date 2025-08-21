package tools

import "time"

type RespFeeRecords struct {
	FeeRecords []struct {
		BillDate                  string    `json:"bill_date"`
		BillType                  int       `json:"bill_type"`
		CustomerID                string    `json:"customer_id"`
		Region                    string    `json:"region"`
		RegionName                string    `json:"region_name"`
		CloudServiceType          string    `json:"cloud_service_type"`
		ResourceType              string    `json:"resource_type"`
		CloudServiceTypeName      string    `json:"cloud_service_type_name"`
		ResourceTypeName          string    `json:"resource_type_name"`
		EffectiveTime             time.Time `json:"effective_time"`
		ExpireTime                time.Time `json:"expire_time"`
		ResourceID                string    `json:"resource_id"`
		ResourceName              string    `json:"resource_name"`
		ResourceTag               any       `json:"resource_tag"`
		ProductID                 string    `json:"product_id"`
		ProductName               string    `json:"product_name"`
		ProductSpecDesc           string    `json:"product_spec_desc"`
		SkuCode                   string    `json:"sku_code"`
		SpecSize                  any       `json:"spec_size"`
		SpecSizeMeasureID         any       `json:"spec_size_measure_id"`
		TradeID                   string    `json:"trade_id"`
		ID                        string    `json:"id"`
		TradeTime                 time.Time `json:"trade_time"`
		EnterpriseProjectID       string    `json:"enterprise_project_id"`
		EnterpriseProjectName     string    `json:"enterprise_project_name"`
		ChargeMode                string    `json:"charge_mode"`
		OrderID                   string    `json:"order_id"`
		PeriodType                string    `json:"period_type"`
		UsageType                 any       `json:"usage_type"`
		Usage                     any       `json:"usage"`
		UsageMeasureID            any       `json:"usage_measure_id"`
		FreeResourceUsage         any       `json:"free_resource_usage"`
		FreeResourceMeasureID     any       `json:"free_resource_measure_id"`
		RiUsage                   any       `json:"ri_usage"`
		RiUsageMeasureID          any       `json:"ri_usage_measure_id"`
		UnitPrice                 float64   `json:"unit_price"`
		Unit                      string    `json:"unit"`
		OfficialAmount            float64   `json:"official_amount"` //官网价
		DiscountAmount            float64   `json:"discount_amount"` //优惠金额
		Amount                    float64   `json:"amount"`          //应付金额
		CashAmount                float64   `json:"cash_amount"`
		CreditAmount              float64   `json:"credit_amount"`
		CouponAmount              float64   `json:"coupon_amount"`
		FlexipurchaseCouponAmount float64   `json:"flexipurchase_coupon_amount"`
		StoredCardAmount          float64   `json:"stored_card_amount"`
		BonusAmount               float64   `json:"bonus_amount"`
		DebtAmount                float64   `json:"debt_amount"`
		AdjustmentAmount          int       `json:"adjustment_amount"`
		MeasureID                 int       `json:"measure_id"`
		Formula                   string    `json:"formula"`
		SubServiceTypeCode        any       `json:"sub_service_type_code"`
		SubServiceTypeName        any       `json:"sub_service_type_name"`
		SubResourceTypeCode       any       `json:"sub_resource_type_code"`
		SubResourceTypeName       any       `json:"sub_resource_type_name"`
		SubResourceID             any       `json:"sub_resource_id"`
		SubResourceName           any       `json:"sub_resource_name"`
		ConsumeTime               time.Time `json:"consume_time"`
		RelativeOrderID           string    `json:"relative_order_id"`
	} `json:"fee_records"`
	TotalCount int    `json:"total_count"`
	Currency   string `json:"currency"`
}
