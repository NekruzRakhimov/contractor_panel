package model

type ContractMiniInfo struct {
	ID                        int     `json:"id"`
	PrevContractId            int     `json:"-" gorm:"-"`
	ContractorName            string  `json:"contractor_name"`
	Beneficiary               string  `json:"beneficiary,omitempty"`
	ContractNumber            string  `json:"contract_number"`
	ContractName              string  `json:"contract_name"`
	ContractType              string  `json:"contract_type"`
	Status                    string  `json:"status"`
	Author                    string  `json:"author"`
	Amount                    float32 `json:"amount"`
	CreatedAt                 string  `json:"created_at,omitempty"`
	UpdatedAt                 string  `json:"updated_at,omitempty"`
	IsExtendContract          bool    `json:"is_extend_contract"`
	ExtendDate                string  `json:"extend_date"`
	AdditionalAgreementNumber int     `json:"additional_agreement_number"`
	EndDate                   string  `json:"end_date"`
	StartDate                 string  `json:"start_date"`
}

type ContractWithJsonB struct {
	ID                        int             `json:"id"`
	Type                      string          `json:"type"`
	PrevContractId            int             `json:"-"`
	Status                    string          `json:"status"` //вынести статус в отдельную таблицу
	Requisites                string          `json:"requisites"`
	Manager                   string          `json:"manager"`
	KAM                       string          `json:"kam"`
	SupplierCompanyManager    string          `json:"supplier_company_manager"`
	ContractParameters        string          `json:"contract_parameters"`
	WithTemperatureConditions bool            `json:"with_temperature_conditions"`
	Products                  string          `json:"products"`
	Discounts                 string          `json:"discounts"`
	Comment                   string          `json:"comment"`
	CreatedAt                 string          `json:"created_at,omitempty"`
	UpdatedAt                 string          `json:"updated_at,omitempty"`
	IsIndivid                 bool            `json:"is_individ"`
	IsExtendContract          bool            `json:"is_extend_contract"`
	ExtendDate                string          `json:"extend_date"`
	DiscountBrand             []DiscountBrand `json:"discount_brand"`
	AdditionalAgreementNumber int             `json:"additional_agreement_number"`
	ExtContractCode           string          `json:"ext_contract_code"`
	View                      string          `json:"view"`
	Regions                   string          `json:"regions"`
}

type Contract struct {
	ID                        int                    `json:"id"`
	Type                      string                 `json:"type"`
	PrevContractId            int                    `json:"-" gorm:"-"`
	Status                    string                 `json:"status"`
	Requisites                Requisites             `json:"requisites"`
	Manager                   string                 `json:"manager,omitempty"`
	KAM                       string                 `json:"kam,omitempty"`
	SupplierCompanyManager    SupplierCompanyManager `json:"supplier_company_manager"`
	ContractParameters        ContractParameters     `json:"contract_parameters"`
	WithTemperatureConditions bool                   `json:"with_temperature_conditions"`
	Products                  []Product              `json:"products"`
	Discounts                 []Discount             `json:"discounts"`
	Comment                   string                 `json:"comment"`
	CreatedAt                 string                 `json:"created_at,omitempty"`
	UpdatedAt                 string                 `json:"updated_at,omitempty"`
	IsExtendContract          bool                   `json:"is_extend_contract"`
	ExtendDate                string                 `json:"extend_date"`
	AdditionalAgreementNumber int                    `json:"additional_agreement_number"`
	IsIndivid                 bool                   `json:"is_individ"`
	//	Brand           string `json:"brand"`
	//	DiscountPercent string `json:"discount_percent"`
	DiscountBrand   []DiscountBrand `json:"discount_brand"`
	ExtContractCode string          `json:"ext_contract_code"`
	View            string          `json:"view"`

	Regions    []Regions `json:"regions"`
	ClientCode string    `json:"client_code"`
}
type Regions struct {
	RegionName string `json:"region_name"`
	RegionCode string `json:"region_code"`
}

type PriceType struct {
	PriceTypeName     string `json:"pricetype_name"`
	PriceTypeCode     string `json:"pricetype_code"`
	PriceTypeCurrency string `json:"pricetype_currency"`
	ClientBin         string `json:"client_bin"`
}

type Requisites struct {
	ContractorName         string `json:"contractor_name"`
	Beneficiary            string `json:"beneficiary,omitempty"`
	BankOfBeneficiary      string `json:"bank_of_beneficiary,omitempty"`
	BankBeneficiaryAddress string `json:"bank_beneficiary_address"`
	SwiftCode              string `json:"swift_code"`
	BIN                    string `json:"bin,omitempty"`
	IIC                    string `json:"iic,omitempty"`
	Phone                  string `json:"phone,omitempty"`
	AccountNumber          string `json:"account_number,omitempty"`
}

type SupplierCompanyManager struct {
	Country   string `json:"country"`
	WorkPhone string `json:"work_phone,omitempty"`
	Email     string `json:"email,omitempty"`
	Skype     string `json:"skype,omitempty"`
	Phone     string `json:"phone,omitempty"`
	Position  string `json:"position,omitempty"`
	// Основание - надо понять как правильно назвать это
	Base     string `json:"base,omitempty"`
	FullName string `json:"full_name"`
}

type ContractParameters struct {
	ContractNumber string  `json:"contract_number"`
	ContractName   string  `json:"contract_name"`
	ContractAmount float32 `json:"contract_amount"`

	// update fields
	CurrencyName  string `json:"currency_name,omitempty"`
	CurrencyCode  string `json:"currency_code,omitempty"`
	PricetypeName string `json:"pricetype_name,omitempty"`
	PricetypeCode string `json:"pricetype_code,omitempty"`

	Prepayment                float32  `json:"prepayment,omitempty"`
	DateOfDelivery            string   `json:"date_of_delivery,omitempty"`
	FrequencyDeferredDiscount string   `json:"frequency_deferred_discount,omitempty"` //Кратность расчета отложенной скидки TODO: возможно нужно поменять
	DeliveryAddress           []string `json:"delivery_address,omitempty"`
	DeliveryTimeInterval      int      `json:"delivery_time_interval,omitempty"` //интервал времени поставки после поступления денежых средств
	ReturnTimeDelivery        int      `json:"return_time_delivery,omitempty"`   //время возврата при условии не поставки
	ContractDate              string   `json:"contract_date,omitempty"`
	StartDate                 string   `json:"start_date"`
	EndDate                   string   `json:"end_date"`
	IsExtendContract          bool     `json:"is_extend_contract"`
	ExtendDate                string   `json:"extend_date"`
}

type Product struct {
	ProductNumber    string     `json:"product_number,omitempty"`
	ProductName      string     `json:"product_name"`
	Price            float64    `json:"price"`
	Currency         string     `json:"currency"`
	Substance        string     `json:"substance,omitempty"`
	StorageCondition string     `json:"storage_condition,omitempty"`
	Producer         string     `json:"producer,omitempty"`
	Sku              string     `json:"sku,omitempty"`
	SkuName          string     `json:"sku_name,omitempty"`
	Plan             float64    `json:"plan,omitempty"`
	DiscountPercent  float64    `json:"discount_percent,omitempty"`
	PriceType        *PriceType `json:"price_type,omitempty"`
}

type Discount struct {
	Name            string           `json:"name,omitempty"`
	Code            string           `json:"code"`
	DiscountAmount  int              `json:"discount_amount,omitempty"`
	IsSelected      bool             `json:"is_selected"`
	PeriodFrom      string           `json:"period_from"`
	IsSale          bool             `json:"is_sale"`
	PeriodTo        string           `json:"period_to"`
	DiscountPercent float64          `json:"discount_percent"`
	GrowthPercent   float64          `json:"growth_percent"`
	Periods         []DiscountPeriod `json:"periods,omitempty"`
	DiscountBrands  []DiscountBrands `json:"discount_brands"`
	Products        []Product        `json:"products"`
}
type DiscountPeriod struct {
	PeriodFrom      string  `json:"period_from"`
	PeriodTo        string  `json:"period_to"`
	TotalAmount     float64 `json:"total_amount"`
	RewardAmount    int     `json:"reward_amount"`
	DiscountPercent float64 `json:"discount_percent"`
	Type            string  `json:"type,omitempty"`
	Name            string  `json:"name,omitempty"`
	PurchaseAmount  float64 `json:"purchase_amount,omitempty"`
	SalesAmount     float64 `json:"sales_amount,omitempty"`
	GrowthPercent   float64 `json:"growth_percent,omitempty"`
	//DiscountAmount      float32 `json:"discount_amount,omitempty"`
	//GraceDays           string  `json:"grace_days,omitempty"`
	//PaymentMultiplicity string  `json:"payment_multiplicity,omitempty"`
	//Amount              float32 `json:"amount,omitempty"`
	//Site                string  `json:"site,omitempty"`
	//Other               string  `json:"other"`
	//Comments            string  `json:"comments,omitempty"`
}

type DiscountBrands struct {
	PeriodFrom string     `json:"period_from"`
	PeriodTo   string     `json:"period_to"`
	Brands     []BrandDTO `json:"brands"`
}

type BrandDTO struct {
	DiscountPercent float64 `json:"discount_percent"`
	PurchaseAmount  float64 `json:"purchase_amount"`
	SalesAmount     float64 `json:"sales_amount"`
	BrandName       string  `json:"brand_name"`
	BrandCode       string  `json:"brand_code"`
}

//type Brand struct {
//	ID              int    `json:"id"`
//	Brand           string `json:"brand"`
//	BrandCode       string `json:"brand_code"`
//	DiscountPercent string `json:"discount_percent"`
//}

type RBRequest struct {
	BIN            string `json:"bin"`
	Type           string `json:"type"`
	ContractorName string `json:"contractor_name"`
	PeriodFrom     string `json:"period_from"`
	PeriodTo       string `json:"period_to"`
	ClientCode     string `json:"client_code"`
	//DoubtedDiscounts []DoubtedDiscount `json:"doubted_discounts"`
}

type RbDTO struct {
	ID                   int     `json:"id"`
	ContractNumber       string  `json:"contract_number"`
	StartDate            string  `json:"start_date"`
	EndDate              string  `json:"end_date"`
	TypePeriod           string  `json:"type_period"`
	BrandName            string  `json:"brand_name,omitempty"`
	ProductCode          string  `json:"product_code,omitempty"`
	DiscountPercent      float64 `json:"discount_percent"`
	DiscountAmount       float64 `json:"discount_amount"`
	TotalWithoutDiscount float64 `json:"TotalWithoutDiscount"`
	LeasePlan            float64 `json:"lease_plan"`
	RewardAmount         float64 `json:"reward_amount"`
	DiscountType         string  `json:"discount_type"`
	Status               string  `json:"status"`
	RegionName           string  `json:"region_name"`
	RegionCode           string  `json:"region_code"`
}
