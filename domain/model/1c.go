package model

type Brand struct {
	BrandArr []struct {
		BrandName string `json:"brand_name"`
		BrandCode string `json:"brand_code"`
	} `json:"brand_arr"`
}

type TotalAmountBrand struct {
	Total     float32 `json:"total"`
	QntTotal  float32 `json:"qnt_total"`
	Date      string  `json:"date"`
	BrandCode string  `json:"brand_code"`
	BrandName string  `json:"brand_name"`
}

type Sales struct {
	SalesArr []SalesArr `json:"sales_arr"`
}

type SalesNew struct {
	SalesArrNEW []SalesArrNEW `json:"sales_arr"`
}

type SalesArr struct {
	ProductName string  `json:"product_name"`
	ProductCode string  `json:"product_code"`
	Total       float64 `json:"total"`
	QntTotal    float64 `json:"qnt_total"`
	Date        string  `json:"date"`
	BrandCode   string  `json:"brand_code"`
	BrandName   string  `json:"brand_name"`
	RegionCode  string  `json:"region_code"`
	RegionName  string  `json:"region_name"`
}

type SalesArrNEW struct {
	ProductName string  `json:"product_name"`
	ProductCode string  `json:"product_code"`
	Total       float64 `json:"total"`
	QntTotal    float32 `json:"qnt_total"`
	Date        string  `json:"date"`
	BrandCode   string  `json:"brand_code"`
	BrandName   string  `json:"brand_name"`
	RegionCode  string  `json:"region_code"`
	RegionName  string  `json:"region_name"`
}

type Purchase struct {
	PurchaseArr []PurchaseArr `json:"purchase_arr"`
}

type PurchaseArr struct {
	ProductName  string  `json:"product_name"`
	ProductCode  string  `json:"product_code"`
	Total        float64 `json:"total"`
	QntTotal     float64 `json:"qnt_total"`
	Date         string  `json:"date"`
	BrandCode    string  `json:"brand_code"`
	BrandName    string  `json:"brand_name"`
	ContractCode string  `json:"contract_code"`
}

type DateSales struct {
	Datestart      string   `json:"datestart"`
	Dateend        string   `json:"dateend"`
	ClientBin      string   `json:"client_bin"`
	Type           string   `json:"type"`
	TypeValue      string   `json:"typeValue"`
	TypeParameters []string `json:"type_parameters"`
}

type AddBrand struct {
	BrandName string `json:"brand_name"`
	BrandCode string `json:"brand_code,omitempty"`
}

type DiscountBrand struct {
	Id              int     `json:"id"`
	BrandName       string  `json:"brand_name"`
	BrandCode       string  `json:"brand_code"`
	DiscountPercent float64 `json:"discount_percent"`
	ContractId      int     `json:"contract_id,omitempty"`
}

type BrandInfo struct {
	Id              int     `json:"id"`
	ContractInfo    int     `json:"contract_info"`
	Brand           string  `json:"brand"`
	DiscountPercent float32 `json:"discount_percent"`
	ContractId      int     `json:"contract_id"`
	Total           float32 `json:"total"`
	DiscountSum     float32 `json:"discount_sum"`
}

type ReqBrand struct {
	ClientCode     string   `json:"client_bin"`
	Beneficiary    string   `json:"beneficiary"`
	DateStart      string   `json:"datestart"`
	DateEnd        string   `json:"dateend"`
	Type           string   `json:"type"`
	TypeValue      string   `json:"typeValue"`
	TypeParameters []string `json:"typeParameters"`
	Contracts      []string `json:"contracts"`
	SchemeType     string   `json:"scheme_type"`
}

type GetData1CRequest struct {
	ClientBin      string   `json:"client_bin"`
	Beneficiary    string   `json:"beneficiary"`
	DateStart      string   `json:"datestart"`
	DateEnd        string   `json:"dateend"`
	Type           string   `json:"type"`
	TypeValue      string   `json:"typeValue"`
	TypeParameters []string `json:"typeParameters"`
	Contracts      []string `json:"contracts"`
}

type GetData1CResponse struct {
	SalesArr    []GetData1CProducts `json:"sales_arr"`
	PurchaseArr []GetData1CProducts `json:"purchase_arr"`
	PaymentsArr []GetData1CPayments `json:"payments_arr"`
}

type GetData1CPayments struct {
	Organization string                     `json:"organization"`
	DataArr      []GetData1CPaymentsDataArr `json:"data_arr"`
}

type GetData1CPaymentsDataArr struct {
	Expanse             string `json:"expanse"`
	Income              string `json:"income"`
	PaymentOrderExpanse string `json:"paymentOrderExpanse"`
	PaymentOrderIncome  string `json:"paymentOrderIncome"`
	Date                string `json:"date"`
}

type GetData1CProducts struct {
	ProductName string  `json:"product_name"`
	ProductCode string  `json:"product_code"`
	Total       float32 `json:"total"`
	QntTotal    float32 `json:"qnt_total"`
	Date        string  `json:"date"`
	BrandCode   string  `json:"brand_code"`
	BrandName   string  `json:"brand_name"`
}

type TotalBrandDiscount struct {
	BrandName       string  `json:"brand_name"`
	Amount          float32 `json:"amount"`
	DiscountPercent float32 `json:"discount_percent"`
	Id              int     `json:"id"`
	ContractNumber  string  `json:"contract_number"`
}

type ContractParam struct {
	Id             int    `json:"id"`
	Bin            string `json:"bin"`
	ContractNumber string `json:"contract_number"`
}

type BrandAndPercent struct {
	ContractID      string `json:"contract_id"`
	Brand           string `json:"brand"`
	ContractNumber  string `json:"contract_number"`
	BrandCode       string `json:"brand_code"`
	DiscountPercent string `json:"discount_percent"`
}

type RespPriceType struct {
	PricetypeArr []struct {
		PricetypeName         string `json:"pricetype_name"`
		PricetypeCode         string `json:"pricetype_code"`
		PricetypeCurrency     string `json:"pricetype_currency"`
		ClientBin             string `json:"client_bin"`
		PriceTypeCurrencyName string `json:"pricetype_currency_name"`
	} `json:"pricetype_arr"`
}

type PriceTypeAndCode struct {
	PricetypeName string `json:"pricetype_name"`
	PricetypeCode string `json:"pricetype_code"`
	ClientBin     string `json:"client_bin,omitempty"`
}

type PriceTypeResponse struct {
	PricetypeCode string `json:"pricetype_code"`
	PricetypeName string `json:"pricetype_name"`
	ClientBin     string `json:"client_bin"`
}

type CurrencyArr struct {
	CurrencyArr []struct {
		CurrencyName string `json:"currency_name"`
		CurrencyCode string `json:"currency_code"`
	} `json:"currency_arr"`
}

type ConvertCurrency struct {
	CurrencyName string `json:"currency_name"`
	CurrencyCode string `json:"currency_code"`
}
