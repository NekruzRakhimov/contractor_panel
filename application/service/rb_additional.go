package service

import (
	"bytes"
	"contractor_panel/domain/model"
	"encoding/json"
	"fmt"
	log "github.com/sirupsen/logrus"
	"io/ioutil"
	"net/http"
	"time"
)

const (
	TempDateCompleter = " 0:00:00"
	TempDateEnd       = " 23:59:59"
)

func GetBrands() (model.Brand, error) {
	brand := model.Brand{}
	client := &http.Client{
		Timeout: 120 * time.Second,
	}
	uri := "http://89.218.153.38:8081/AQG_ULAN/hs/integration/brandlist"
	req, err := http.NewRequest("GET", uri, nil)
	req.Header.Set("Content-Type", "application/json") // This makes it work
	req.SetBasicAuth("http_client", "123456")

	if err != nil {
		log.Println(err)
	}
	resp, err := client.Do(req)
	if err != nil {
		log.Println(err)
		return brand, err
	}

	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		log.Println(err)
		return brand, err
	}

	defer resp.Body.Close()
	if err != nil {
		log.Println(err)
		return brand, err
	}
	body = bytes.TrimPrefix(body, []byte("\xef\xbb\xbf")) // Or []byte{239, 187, 191}

	err = json.Unmarshal(body, &brand)
	if err != nil {
		log.Println(err)
		return brand, err
	}

	return brand, nil

}

func GetSales(reqBrand model.ReqBrand) (model.Sales, error) {
	var sales model.Sales

	date := model.ReqBrand{
		ClientCode:     reqBrand.ClientCode,
		DateStart:      reqBrand.DateStart + TempDateCompleter,
		DateEnd:        reqBrand.DateEnd + TempDateEnd,
		Type:           "sales_brand_only",
		TypeValue:      "",
		TypeParameters: nil,
		//SchemeType:     reqBrand.SchemeType,
		SchemeType: "LS",
	}
	//for _, value := range brandInfo {
	//	date.TypeParameters = append(date.TypeParameters, value.Brand)
	//}

	reqBodyBytes := new(bytes.Buffer)
	json.NewEncoder(reqBodyBytes).Encode(&date)
	fmt.Println(">>> ", reqBodyBytes)

	//parm.Add("datestart", "01.01.2022 0:02:09")
	//parm.Add("dateend", "01.01.2022 0:02:09")
	client := &http.Client{}
	log.Println("запрос", reqBodyBytes)
	uri := "http://89.218.153.38:8081/AQG_ULAN/hs/integration/getdata"
	req, err := http.NewRequest("POST", uri, reqBodyBytes)
	req.Header.Set("Content-Type", "application/json") // This makes it work
	req.SetBasicAuth("http_client", "123456")

	if err != nil {
		log.Println(err)
	}
	resp, err := client.Do(req)
	if err != nil {
		log.Println(err)
		return sales, err
	}

	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		log.Println(err)
		return sales, err
	}
	//log.Println("BODYYYYYYYYYYYYYYYYYYYYYYYYYYYYYYYYYYY", string(body))

	defer resp.Body.Close()
	if err != nil {
		log.Println(err)
		return sales, err
	}
	body = bytes.TrimPrefix(body, []byte("\xef\xbb\xbf")) // Or []byte{239, 187, 191}

	err = json.Unmarshal(body, &sales)
	if err != nil {
		log.Println(err)
		return sales, err
	}

	return sales, nil
}

func GetSalesNEw(reqBrand model.ReqBrand) (model.Sales, error) {
	var sales model.Sales

	date := model.ReqBrand{
		ClientCode:     reqBrand.ClientCode,
		DateStart:      reqBrand.DateStart + TempDateCompleter,
		DateEnd:        reqBrand.DateEnd + TempDateEnd,
		Type:           "sales_total",
		TypeValue:      "",
		TypeParameters: nil,
		//SchemeType:     reqBrand.SchemeType,
		SchemeType: "LS",
	}
	//for _, value := range brandInfo {
	//	date.TypeParameters = append(date.TypeParameters, value.Brand)
	//}

	reqBodyBytes := new(bytes.Buffer)
	json.NewEncoder(reqBodyBytes).Encode(&date)
	fmt.Println(">>> ", reqBodyBytes)

	//parm.Add("datestart", "01.01.2022 0:02:09")
	//parm.Add("dateend", "01.01.2022 0:02:09")
	client := &http.Client{
		Timeout: 120 * time.Second,
	}
	log.Println("запрос", reqBodyBytes)
	uri := "http://89.218.153.38:8081/AQG_ULAN/hs/integration/getdata"
	req, err := http.NewRequest("POST", uri, reqBodyBytes)
	req.Header.Set("Content-Type", "application/json") // This makes it work
	req.SetBasicAuth("http_client", "123456")

	if err != nil {
		log.Println(err)
	}
	resp, err := client.Do(req)
	if err != nil {
		log.Println(err)
		return sales, err
	}

	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		log.Println(err)
		return sales, err
	}
	//log.Println("BODYYYYYYYYYYYYYYYYYYYYYYYYYYYYYYYYYYY", string(body))

	defer resp.Body.Close()
	if err != nil {
		log.Println(err)
		return sales, err
	}
	body = bytes.TrimPrefix(body, []byte("\xef\xbb\xbf")) // Or []byte{239, 187, 191}

	err = json.Unmarshal(body, &sales)
	if err != nil {
		log.Println(err)
		return sales, err
	}

	return sales, nil
}

func AddBrand(brandName string) (model.AddBrand, error) {
	brand := model.AddBrand{BrandName: brandName}
	reqBodyBytes := new(bytes.Buffer)
	json.NewEncoder(reqBodyBytes).Encode(&brand)

	client := &http.Client{
		Timeout: 120 * time.Second,
	}
	uri := "http://89.218.153.38:8081/AQG_ULAN/hs/integration/createbrand"
	req, err := http.NewRequest("POST", uri, reqBodyBytes)
	req.Header.Set("Content-Type", "application/json") // This makes it work
	req.SetBasicAuth("http_client", "123456")

	if err != nil {
		log.Println(err)
	}
	resp, err := client.Do(req)
	if err != nil {
		log.Println(err)
		return brand, err
	}

	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		log.Println(err)
		return brand, err
	}

	defer resp.Body.Close()
	if err != nil {
		log.Println(err)
		return brand, err
	}
	body = bytes.TrimPrefix(body, []byte("\xef\xbb\xbf")) // Or []byte{239, 187, 191}
	log.Println(string(body))
	err = json.Unmarshal(body, &brand)
	if err != nil {
		log.Println(err)
		return brand, err
	}

	return brand, nil

}

func GetSalesBrand(reqBrand model.ReqBrand, brandInfo []model.BrandInfo) (model.Sales, error) {
	var sales model.Sales

	date := model.ReqBrand{
		ClientCode:     reqBrand.ClientCode,
		DateStart:      reqBrand.DateStart + TempDateCompleter,
		DateEnd:        reqBrand.DateEnd + TempDateEnd,
		Type:           "sales_brand_only",
		TypeValue:      "",
		TypeParameters: nil,
	}
	//for _, value := range brandInfo {
	//	date.TypeParameters = append(date.TypeParameters, value.Brand)
	//}

	reqBodyBytes := new(bytes.Buffer)
	json.NewEncoder(reqBodyBytes).Encode(&date)
	fmt.Println(">>> ", reqBodyBytes)

	//parm.Add("datestart", "01.01.2022 0:02:09")
	//parm.Add("dateend", "01.01.2022 0:02:09")
	client := &http.Client{
		Timeout: 120 * time.Second,
	}
	log.Println(reqBodyBytes)
	uri := "http://89.218.153.38:8081/AQG_ULAN/hs/integration/getdata"
	req, err := http.NewRequest("POST", uri, reqBodyBytes)
	req.Header.Set("Content-Type", "application/json") // This makes it work
	req.SetBasicAuth("http_client", "123456")

	if err != nil {
		log.Println(err)
	}
	resp, err := client.Do(req)
	if err != nil {
		log.Println(err)
		return sales, err
	}

	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		log.Println(err)
		return sales, err
	}
	log.Println("BODYYYYYYYYYYYYYYYYYYYYYYYYYYYYYYYYYYY", string(body))

	defer resp.Body.Close()
	if err != nil {
		log.Println(err)
		return sales, err
	}
	body = bytes.TrimPrefix(body, []byte("\xef\xbb\xbf")) // Or []byte{239, 187, 191}

	err = json.Unmarshal(body, &sales)
	if err != nil {
		log.Println(err)
		return sales, err
	}

	return sales, nil

}

func GetPurchase(reqBrand model.ReqBrand) (model.Purchase, error) {
	var purchase model.Purchase

	date := model.ReqBrand{
		ClientCode:     reqBrand.ClientCode,
		DateStart:      reqBrand.DateStart + TempDateCompleter,
		DateEnd:        reqBrand.DateEnd + TempDateEnd,
		Type:           "purchase_total",
		TypeValue:      "",
		TypeParameters: nil,
		Contracts:      nil,
		SchemeType:     reqBrand.SchemeType,
	}
	//for _, value := range brandInfo {
	//	date.TypeParameters = append(date.TypeParameters, value.Brand)
	//}

	reqBodyBytes := new(bytes.Buffer)
	json.NewEncoder(reqBodyBytes).Encode(&date)
	fmt.Println(">>> ", reqBodyBytes)

	//parm.Add("datestart", "01.01.2022 0:02:09")
	//parm.Add("dateend", "01.01.2022 0:02:09")
	client := &http.Client{
		Timeout: 120 * time.Second,
	}
	log.Println(reqBodyBytes)
	uri := "http://89.218.153.38:8081/AQG_ULAN/hs/integration/getdata"
	req, err := http.NewRequest("POST", uri, reqBodyBytes)
	req.Header.Set("Content-Type", "application/json") // This makes it work
	req.SetBasicAuth("http_client", "123456")

	if err != nil {
		log.Println(err)
	}
	resp, err := client.Do(req)
	if err != nil {
		log.Println(err)
		return purchase, err
	}

	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		log.Println(err)
		return purchase, err
	}
	log.Println("BODYYYYYYYYYYYYYYYYYYYYYYYYYYYYYYYYYYY", string(body))

	defer resp.Body.Close()
	if err != nil {
		log.Println(err)
		return purchase, err
	}
	body = bytes.TrimPrefix(body, []byte("\xef\xbb\xbf")) // Or []byte{239, 187, 191}

	err = json.Unmarshal(body, &purchase)
	if err != nil {
		log.Println(err)
		return purchase, err
	}

	return purchase, nil

}

func GetPurchaseBrandOnly(reqBrand model.ReqBrand) (model.Purchase, error) {
	var purchase model.Purchase

	date := model.ReqBrand{
		ClientCode:     reqBrand.ClientCode,
		DateStart:      reqBrand.DateStart + TempDateCompleter,
		DateEnd:        reqBrand.DateEnd + TempDateEnd,
		Type:           "purchase_brand_only",
		TypeValue:      "",
		TypeParameters: nil,
		Contracts:      reqBrand.Contracts,
	}
	//for _, value := range brandInfo {
	//	date.TypeParameters = append(date.TypeParameters, value.Brand)
	//}

	reqBodyBytes := new(bytes.Buffer)
	json.NewEncoder(reqBodyBytes).Encode(&date)
	fmt.Println(">>> ", reqBodyBytes)

	//parm.Add("datestart", "01.01.2022 0:02:09")
	//parm.Add("dateend", "01.01.2022 0:02:09")
	client := &http.Client{
		Timeout: 120 * time.Second,
	}
	log.Println(reqBodyBytes)
	uri := "http://89.218.153.38:8081/AQG_ULAN/hs/integration/getdata"
	req, err := http.NewRequest("POST", uri, reqBodyBytes)
	req.Header.Set("Content-Type", "application/json") // This makes it work
	req.SetBasicAuth("http_client", "123456")

	if err != nil {
		log.Println(err)
	}
	resp, err := client.Do(req)
	if err != nil {
		log.Println(err)
		return purchase, err
	}

	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		log.Println(err)
		return purchase, err
	}
	log.Println("BODYYYYYYYYYYYYYYYYYYYYYYYYYYYYYYYYYYY", string(body))

	defer resp.Body.Close()
	if err != nil {
		log.Println(err)
		return purchase, err
	}
	body = bytes.TrimPrefix(body, []byte("\xef\xbb\xbf")) // Or []byte{239, 187, 191}

	err = json.Unmarshal(body, &purchase)
	if err != nil {
		log.Println(err)
		return purchase, err
	}

	return purchase, nil

}

func GetBrandSales(reqBrand model.ReqBrand) (model.Sales, error) {
	var sales model.Sales

	date := model.ReqBrand{
		ClientCode:     reqBrand.ClientCode,
		DateStart:      reqBrand.DateStart + TempDateCompleter,
		DateEnd:        reqBrand.DateEnd + TempDateEnd,
		Type:           "sales_brand_only",
		TypeValue:      reqBrand.TypeValue,
		TypeParameters: reqBrand.TypeParameters,
	}
	//for _, value := range brandInfo {
	//	date.TypeParameters = append(date.TypeParameters, value.Brand)
	//}

	reqBodyBytes := new(bytes.Buffer)
	json.NewEncoder(reqBodyBytes).Encode(&date)
	fmt.Println(">>> ", reqBodyBytes)

	//parm.Add("datestart", "01.01.2022 0:02:09")
	//parm.Add("dateend", "01.01.2022 0:02:09")
	client := &http.Client{
		Timeout: 120 * time.Second,
	}
	log.Println(reqBodyBytes)
	uri := "http://89.218.153.38:8081/AQG_ULAN/hs/integration/getdata"
	req, err := http.NewRequest("POST", uri, reqBodyBytes)
	req.Header.Set("Content-Type", "application/json") // This makes it work
	req.SetBasicAuth("http_client", "123456")

	if err != nil {
		log.Println(err)
	}
	resp, err := client.Do(req)
	if err != nil {
		log.Println(err)
		return sales, err
	}

	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		log.Println(err)
		return sales, err
	}

	defer resp.Body.Close()
	body = bytes.TrimPrefix(body, []byte("\xef\xbb\xbf")) // Or []byte{239, 187, 191}

	log.Println("BODY: ", string(body))
	err = json.Unmarshal(body, &sales)
	if err != nil {
		log.Println(err)
		return sales, err
	}

	return sales, nil

}

//func FoundBrandDiscount(rbReq model.RBRequest) []model.RbDTO {
//	var rbBrands []model.RbDTO
//	var rbBrand model.RbDTO
//	var totalBrandsDiscount []model.TotalBrandDiscount
//	var BrandTotal model.TotalBrandDiscount
//
//	// f, err := excelize.OpenFile("files/reports/rb/rb_report.xlsx")
//	// if err != nil {
//	// 	fmt.Println(err)
//	// 	return rbBrands
//	// }
//
//	// берем бренды по Бину
//
//	// тут возвращаю
//	//brand AS brand_name, discount_percent, contract_id - 3 fields
//
//	// имя брендов и 1 номер договора
//	dataBrand, contractNumber := repository.GetIDByBIN(rbReq.BIN)
//	log.Println("ДАННЫЕ БРЕНДОВ", dataBrand)
//
//	// тут что ищем?
//	// dataBrand - ID, discount, brand
//
//	for _, value := range dataBrand {
//		var floatPercent float32
//		valuePercent, err := strconv.ParseFloat(value.DiscountPercent, 32)
//		if err != nil {
//			// do something sensible
//		}
//		floatPercent = float32(valuePercent)
//
//		// тут добавляю бренды в TypeParameters
//		//reqBrand.TypeParameters = append(reqBrand.TypeParameters, value.BrandName)
//		//ID              int     `json:"id"` -- отдадим сами
//		//ContractNumber  string  `json:"contract_number"`  -- взял
//		//StartDate       string  `json:"start_date"`  -- возьмем из запроса
//		//EndDate         string  `json:"end_date"` -- возьмем из запроса
//		//BrandName       string  `json:"brand_name,omitempty"` -- отдадим сами
//		//DiscountPercent float32 `json:"discount_percent"` -- отдадим сами
//		//DiscountAmount  float32 `json:"discount_amount"` --  отдадим тоже сами
//		BrandTotal.ContractNumber = contractNumber
//		BrandTotal.BrandName = value.BrandName
//		BrandTotal.DiscountPercent = floatPercent
//		BrandTotal.Id, _ = strconv.Atoi(value.ContractParam)
//
//		totalBrandsDiscount = append(totalBrandsDiscount, BrandTotal)
//	}
//
//	log.Println("reqBrand.TypeParameters", reqBrand.TypeParameters)
//
//	// BrandName string  `json:"brand_name"`
//	//	Amount    float32 `json:"amount"` - сумма чего тогда
//
//	// reqBrand -> он дает массив брендов
//
//	sales, _ := GetBrandSales(reqBrand)
//
//	// Берет определенные бренды из 1С:
//	counter := 1
//	for _, sale := range sales.SalesArr {
//
//		//   {
//		//            "product_name": "7Stick жевательная резинка Арбуз 14,5 г ",
//		//            "product_code": "00000074577",
//		//            "total": 3600,
//		//            "qnt_total": 36,
//		//            "date": "2022-01-03T00:00:00",
//		//            "brand_code": "000000137",
//		//            "brand_name": "7Stick"
//		//        },
//
//		count := float32(0)
//		//	тут будет список брендов
//
//		//TODO: тут ты можешь сразу записать в экселе наименование,кол-во,код, сумму, бренд,
//		//sale.BrandName
//
//		//TODO: после того как мы записали все бренды, мы должны посчитать от него общую сумму
//		for _, brand := range reqBrand.TypeParameters {
//			log.Println("BRAND", brand)
//
//			// мы нашли схожие бренды, что мы должны сделать?
//			if sale.BrandName == brand {
//				count += sale.Total // - это сколько было продано по данному товару
//				//TODO: ты итог должен записать и после чего какой процент и только потом сумму скидки
//
//				//TOTAL
//				// Percent
//
//				counter += 1
//				strCount := fmt.Sprint(counter)
//
//				//log.Println("Наименование: ", sale.ProductName)
//
//				//log.Println("Количество:", sale.QntTotal)
//				//log.Println("Номер продукта:", sale.ProductCode)
//				//log.Println("Стоимость: ", sale.Total)
//				//log.Println("БРЕНД:", sale.BrandName)
//				f.SetCellValue("Sheet1", "A"+strCount, sale.ProductName)
//				f.SetCellValue("Sheet1", "B"+strCount, sale.QntTotal)
//				f.SetCellValue("Sheet1", "C"+strCount, sale.ProductCode)
//				f.SetCellValue("Sheet1", "D"+strCount, sale.Total)
//				f.SetCellValue("Sheet1", "E"+strCount, sale.BrandName)
//				if err := f.SaveAs("reportRB_brand1.xlsx"); err != nil {
//					fmt.Println("ERRRRRRRRRRRRRRRRRRRRRRRRRRRRROOOOOOOOOOORRRRRRRRRRR", err)
//
//				}
//
//				// Тут он пустой
//				fmt.Println("totalBrandsDiscount", totalBrandsDiscount)
//				fmt.Println("LEN:", len(totalBrandsDiscount))
//				if len(totalBrandsDiscount) == 0 {
//
//					//TODO: тут записываем каждый бренд и его общую сумму, это и есть TOTAL но по скидкам
//
//					log.Println("сРАБАОТЛО")
//					BrandTotal.BrandName = brand
//					BrandTotal.Amount = count
//
//					totalBrandsDiscount = append(totalBrandsDiscount, BrandTotal)
//
//				}
//
//				//  в начале так как массив длина его 0 - поэтому мы заранее добавили 1 бренд
//				// после чего делаем проверку, если нашли схожий бренд - просто делаем + суммы к этому бренду
//				for i, check := range totalBrandsDiscount {
//					log.Println("___________________________________________________________________")
//
//					if brand != check.BrandName {
//						BrandTotal.BrandName = brand
//						BrandTotal.Amount = count
//						fmt.Println("записался в массив")
//						log.Println("записался в массив")
//						totalBrandsDiscount = append(totalBrandsDiscount, BrandTotal)
//					} else if brand == check.BrandName {
//						fmt.Println("обновили  массив")
//						log.Println("обновили  массив")
//						totalBrandsDiscount[i].Amount += count
//					}
//
//				}
//
//				// когда мы посчитали общую сумму данного бренда мы должны найти его скидку
//
//			}
//
//		}
//
//	}
//
//	for _, value := range totalBrandsDiscount {
//		fmt.Println("ИтогСуммы: ", value.Amount)
//		TotalPercent := (value.Amount * value.DiscountPercent) / 100
//		log.Println("Сумма скидки: ", TotalPercent)
//		log.Println("Скидка: ", value.DiscountPercent)
//		//ID              int     `json:"id"`
//		//ContractNumber  string  `json:"contract_number"`
//		//StartDate       string  `json:"start_date"`
//		//EndDate         string  `json:"end_date"`
//		//BrandName       string  `json:"brand_name,omitempty"`
//		//DiscountPercent float32 `json:"discount_percent"`
//		//DiscountAmount  float32 `json:"discount_amount"`
//
//		rbBrand.ID = value.Id
//		rbBrand.ContractNumber = value.ContractNumber
//		rbBrand.StartDate = reqBrand.DateStart
//		rbBrand.EndDate = reqBrand.DateEnd
//		rbBrand.DiscountPercent = value.DiscountPercent
//		rbBrand.DiscountAmount = TotalPercent
//
//	}
//
//	//ID              int     `json:"id"`
//	//ContractNumber  string  `json:"contract_number"`
//	//StartDate       string  `json:"start_date"`
//	//EndDate         string  `json:"end_date"`
//	//BrandName       string  `json:"brand_name,omitempty"`
//	//DiscountPercent float32 `json:"discount_percent"`
//	//DiscountAmount  float32 `json:"discount_amount"`
//	//TODO: вернуть даннные:
//	// 1. процент
//	// 2. Сумму скидки
//	// 3. Имя Бренда
//	// 4.  номер договора??
//	// 5. ID Договора
//
//	log.Println(rbBrands, "ОТВЕТ ТВОЕЙ МОДЕЛИ")
//
//	return rbBrands
//
//}

func GetSalesSKU(reqBrand model.ReqBrand) (model.Sales, error) {
	var sales model.Sales

	date := model.ReqBrand{
		ClientCode:     reqBrand.ClientCode,
		DateStart:      reqBrand.DateStart + TempDateCompleter,
		DateEnd:        reqBrand.DateEnd + TempDateEnd,
		Type:           "sales_brand_only",
		TypeValue:      reqBrand.TypeValue,
		TypeParameters: reqBrand.TypeParameters,
	}
	//for _, value := range brandInfo {
	//	date.TypeParameters = append(date.TypeParameters, value.Brand)
	//}

	reqBodyBytes := new(bytes.Buffer)
	json.NewEncoder(reqBodyBytes).Encode(&date)
	fmt.Println(">>> ", reqBodyBytes)

	//parm.Add("datestart", "01.01.2022 0:02:09")
	//parm.Add("dateend", "01.01.2022 0:02:09")
	client := &http.Client{
		Timeout: 120 * time.Second,
	}
	log.Println(reqBodyBytes)
	uri := "http://89.218.153.38:8081/AQG_ULAN/hs/integration/getdata"
	req, err := http.NewRequest("POST", uri, reqBodyBytes)
	req.Header.Set("Content-Type", "application/json") // This makes it work
	req.SetBasicAuth("http_client", "123456")

	if err != nil {
		log.Println(err)
	}
	resp, err := client.Do(req)
	if err != nil {
		log.Println(err)
		return sales, err
	}

	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		log.Println(err)
		return sales, err
	}
	log.Println("BODYYYYYYYYYYYYYYYYYYYYYYYYYYYYYYYYYYY", string(body))

	defer resp.Body.Close()
	if err != nil {
		log.Println(err)
		return sales, err
	}
	body = bytes.TrimPrefix(body, []byte("\xef\xbb\xbf")) // Or []byte{239, 187, 191}

	err = json.Unmarshal(body, &sales)
	if err != nil {
		log.Println(err)
		return sales, err
	}

	return sales, nil

}

func PresentationDiscount(rbReq model.RBRequest) (model.Purchase, error) {
	var purchase model.Purchase

	date := model.ReqBrand{
		ClientCode:     rbReq.BIN,
		DateStart:      rbReq.PeriodFrom + TempDateCompleter,
		DateEnd:        rbReq.PeriodTo + TempDateEnd,
		Type:           "purchase",
		TypeValue:      "",
		TypeParameters: nil,
	}
	//for _, value := range brandInfo {
	//	date.TypeParameters = append(date.TypeParameters, value.Brand)
	//}

	reqBodyBytes := new(bytes.Buffer)
	json.NewEncoder(reqBodyBytes).Encode(&date)
	fmt.Println(">>> ", reqBodyBytes)

	//parm.Add("datestart", "01.01.2022 0:02:09")
	//parm.Add("dateend", "01.01.2022 0:02:09")
	client := &http.Client{
		Timeout: 120 * time.Second,
	}
	log.Println(reqBodyBytes)
	uri := "http://89.218.153.38:8081/AQG_ULAN/hs/integration/getdata"
	req, err := http.NewRequest("POST", uri, reqBodyBytes)
	req.Header.Set("Content-Type", "application/json") // This makes it work
	req.SetBasicAuth("http_client", "123456")

	if err != nil {
		log.Println(err)
	}
	resp, err := client.Do(req)
	if err != nil {
		log.Println(err)
		return purchase, err
	}

	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		log.Println(err)
		return purchase, err
	}
	log.Println("BODYYYYYYYYYYYYYYYYYYYYYYYYYYYYYYYYYYY", string(body))

	defer resp.Body.Close()
	if err != nil {
		log.Println(err)
		return purchase, err
	}
	body = bytes.TrimPrefix(body, []byte("\xef\xbb\xbf")) // Or []byte{239, 187, 191}

	err = json.Unmarshal(body, &purchase)
	if err != nil {
		log.Println(err)
		return purchase, err
	}

	return purchase, nil

}

func GetSales1C(rb model.ReqBrand, typeData string) (model.Sales, error) {
	var sales model.Sales

	date := model.ReqBrand{
		ClientCode:     rb.ClientCode,
		DateStart:      rb.DateStart + TempDateCompleter,
		DateEnd:        rb.DateEnd + TempDateEnd,
		Type:           typeData,
		TypeValue:      "",
		TypeParameters: nil,
	}
	//for _, value := range brandInfo {
	//	date.TypeParameters = append(date.TypeParameters, value.Brand)
	//}

	reqBodyBytes := new(bytes.Buffer)
	json.NewEncoder(reqBodyBytes).Encode(&date)
	//fmt.Println(">>> ", reqBodyBytes)

	//parm.Add("datestart", "01.01.2022 0:02:09")
	//parm.Add("dateend", "01.01.2022 0:02:09")
	client := &http.Client{
		Timeout: 120 * time.Second,
	}
	log.Println("request--->", reqBodyBytes)
	uri := "http://89.218.153.38:8081/AQG_ULAN/hs/integration/getdata"
	req, err := http.NewRequest("POST", uri, reqBodyBytes)
	req.Header.Set("Content-Type", "application/json") // This makes it work
	req.SetBasicAuth("http_client", "123456")

	if err != nil {
		log.Println(err)
	}
	resp, err := client.Do(req)
	if err != nil {
		log.Println(err)
		return sales, err
	}

	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		log.Println(err)
		return sales, err
	}
	//	log.Println("BODYYYYYYYYYYYYYYYYYYYYYYYYYYYYYYYYYYY", body)

	defer resp.Body.Close()
	if err != nil {
		log.Println(err)
		return sales, err
	}
	body = bytes.TrimPrefix(body, []byte("\xef\xbb\xbf")) // Or []byte{239, 187, 191}

	err = json.Unmarshal(body, &sales)
	if err != nil {
		log.Println(err)
		return sales, err
	}

	return sales, nil

}

func GetDataFrom1C(request model.GetData1CRequest) (response model.GetData1CResponse, err error) {
	request.DateStart += TempDateCompleter
	request.DateEnd += TempDateEnd

	reqBodyBytes := new(bytes.Buffer)
	if err = json.NewEncoder(reqBodyBytes).Encode(&request); err != nil {
		return model.GetData1CResponse{}, err
	}

	uri := "http://89.218.153.38:8081/AQG_ULAN/hs/integration/getdata"
	req, err := http.NewRequest("POST", uri, reqBodyBytes)
	if err != nil {
		log.Println(err)
	}

	req.Header.Set("Content-Type", "application/json") // This makes it work
	req.SetBasicAuth("http_client", "123456")

	client := &http.Client{
		Timeout: 120 * time.Second,
	}
	resp, err := client.Do(req)
	if err != nil {
		log.Println(err)
		return model.GetData1CResponse{}, err
	}

	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		log.Println(err)
		return model.GetData1CResponse{}, err
	}

	defer resp.Body.Close()
	log.Println("[1C_get_data]|[body] = ", string(body))

	body = bytes.TrimPrefix(body, []byte("\xef\xbb\xbf")) // Or []byte{239, 187, 191}

	err = json.Unmarshal(body, &response)
	if err != nil {

		log.Println(err)
		return model.GetData1CResponse{}, err
	}
	return response, nil

}
