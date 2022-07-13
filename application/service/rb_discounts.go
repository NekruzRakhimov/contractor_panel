package service

import (
	"contractor_panel/domain/model"
	"errors"
	"fmt"
	log "github.com/sirupsen/logrus"
	"strconv"
	"strings"
	"time"
)

const (
	sheet    = "Итог"
	RB1Name  = "Скидка за объем закупа"
	RB2Name  = "Скидка  на группы товаров"
	RB3Name  = "Скидка за выполнение плана закупа по препаратам"
	RB4Name  = "Скидка за представленность"
	RB5Name  = "Скидка за фиксированную сумму рб за выполнение плана закупа по брендам"
	RB6Name  = "Скидка на МТЗ"
	RB7Name  = "Скидка за выполнение плана продаж по бренду"
	RB8Name  = "Скидка на РБ от закупа продукции %"
	RB9Name  = "Скидка на РБ по филиалам"
	RB10Name = "Скидка РБ по логистике"
	RB11Name = "Скидка  РБ за поддержание ассортимента"
	RB12Name = "Скидка РБ объем за закупку в промежутке времени"
	RB13Name = "Скидка  РБ за прирост продаж"
	RB14Name = "Скидка за выполнение объема продаж по препаратам"
	RB15Name = "Скидка  за выполнение плана продаж"
	RB16Name = "Скидка  за выполнение плана продаж %"
	RB17Name = "Скидка за фактический объем продаж"

	RB18Name = "Скидка РБ за получение отчетности"
)

const (
	RB1Code = "TOTAL_AMOUNT_OF_SELLING"

	RB2Code = "DISCOUNT_BRAND"
	RB3Code = "DISCOUNT_PLAN_LEASE"

	RB4Code  = "DISCOUNT_FOR_REPRESENTATION"
	RB5Code  = "DISCOUNT_FOR_FIX_SUM_MTZ"
	RB6Code  = "DISCOUNT_FOR_PLAN_MTZ"
	RB7Code  = "DISCOUNT_FOR_FULFILLING_SALES"
	RB8Code  = "DISCOUNT_FOR_LEASE_GENERAL"
	RB9Code  = "DISCOUNT_FOR_FILIAL"
	RB10Code = "DISCOUNT_FOR_LOGISTIC"
	RB11Code = "DISCOUNT_FOR_ASSORTMENT"
	RB12Code = "RB_DISCOUNT_FOR_PURCHASE_PERIOD"
	RB13Code = "RB_DISCOUNT_FOR_SALES_GROWTH"
	RB14Code = "DISCOUNT_PLAN_SALE"

	RB15Code = "DISCOUNT_PLAN_SALE_REWARD"
	RB16Code = "DISCOUNT_PLAN_SALE_PERCENT"
	RB17Code = "RB_DISCOUNT_ACTUAL_SALES"

	RB18Code = "DISCOUNT_FOR_REPORTING"
)

func GetRB1stType(request model.RBRequest, contracts []model.Contract) ([]model.RbDTO, error) {
	//TODO: посмотри потом
	//testBin := "060840003599"
	//req := model.ReqBrand{
	//	ClientCode:   request.BIN,
	//	Beneficiary: request.ContractorName,
	//	DateStart:   request.PeriodFrom,
	//	DateEnd:     request.PeriodTo,
	//	Type:        "sales",
	//}

	//externalCodes := GetExternalCode(request.BIN)
	//contractsCode := JoinContractCode(externalCodes)

	reqBrand := model.ReqBrand{
		ClientCode:     request.ClientCode,
		DateStart:      request.PeriodFrom,
		DateEnd:        request.PeriodTo,
		TypeValue:      "",
		TypeParameters: nil,
		//Contracts:      contractsCode, // необходимо получить коды контрактов
	}
	purchase, _ := GetPurchase(reqBrand)
	totalAmount := GetPurchaseTotalAmount(purchase)

	//totalPurchaseCode := CountPurchaseByCode(purchase)
	//
	//present := model.ReqBrand{
	//	ClientCode:      request.BIN,
	//	Beneficiary:    "",
	//	DateStart:      request.PeriodFrom,
	//	DateEnd:        request.PeriodTo,
	//	Type:           "",
	//	TypeValue:      "",
	//	TypeParameters: nil,
	//	Contracts:      nil,
	//}
	//
	//sales, err := GetSales1C(present, "sales_brand_only")
	//sales, err := GetSales(req)
	//if err != nil {
	//	return nil, err
	//}

	fmt.Printf("###%+v\n", contracts)
	log.Printf("[PURCHASE] %f ", totalAmount)

	contractRB := DefiningRBReport(contracts, totalAmount, request)

	return contractRB, nil
}
func DefiningRBReport(contracts []model.Contract, totalAmount float64, request model.RBRequest) (contractsRB []model.RbDTO) {
	for _, contract := range contracts {
		var contractRB []model.RbDTO
		for _, discount := range contract.Discounts {
			if discount.Code == "TOTAL_AMOUNT_OF_SELLING" && discount.IsSelected {
				log.Printf("\n[CONTRACT_DISCOUNT][%s] %+v\n", contract.ContractParameters.ContractNumber, contract.Discounts)
				contractRB = DiscountToReportRB(discount, contract, totalAmount, request)
			}
		}
		contractsRB = append(contractsRB, contractRB...)
	}

	return contractsRB
}
func DiscountToReportRB(discount model.Discount, contract model.Contract, totalAmount float64, request model.RBRequest) (contractsRB []model.RbDTO) {
	fmt.Println("<begin>")
	var contractRB model.RbDTO

	if len(discount.Periods) > 0 {
		contractRB = model.RbDTO{
			ID:             contract.ID,
			ContractNumber: contract.ContractParameters.ContractNumber,
			DiscountType:   RB1Name,
			StartDate:      discount.Periods[0].PeriodFrom,
			EndDate:        discount.Periods[0].PeriodTo,
		}

		var (
			maxDiscountAmount float64 // сумма закупа
			maxRewardAmount   int     // Сумма вознаграждения
			maxLeasePlan      float64 // план закупа
			isCompleted       bool
		)

		for _, period := range discount.Periods {
			if period.PeriodFrom >= request.PeriodFrom && period.PeriodTo <= request.PeriodTo {
				if float64(period.TotalAmount) <= totalAmount {
					if period.TotalAmount >= maxLeasePlan {
						log.Printf("\n[CONTRACT_PERIODS][%s] %+v\n", contract.ContractParameters.ContractNumber, discount.Periods)
						maxDiscountAmount = float64(period.RewardAmount)
						maxRewardAmount = period.RewardAmount
						maxLeasePlan = period.TotalAmount
						isCompleted = true
					}
				} /*else {
					maxRewardAmount = period.RewardAmount
					maxLeasePlan = period.TotalAmount
				}*/
			}
		}

		if !isCompleted && len(discount.Periods) > 0 {
			maxRewardAmount = discount.Periods[0].RewardAmount
			maxLeasePlan = discount.Periods[0].TotalAmount
		}

		// Сумма скидки	| Сумма вознаграждения	| План закупа

		contractRB.RewardAmount = float64(maxRewardAmount)
		contractRB.LeasePlan = maxLeasePlan
		contractRB.DiscountAmount = maxDiscountAmount

		//if len(discount.Periods) > 1 && totalAmount >= discount.Periods[1].TotalAmount && discount.Periods[1].RewardAmount > discount.Periods[0].RewardAmount {
		//	fmt.Printf("worked [totalAmount = %d AND discount.Periods[0].TotalAmount = %d]\n", totalAmount, discount.Periods[0].TotalAmount)
		//	contractRB.DiscountAmount = float32(discount.Periods[0].RewardAmount)
		//} else if totalAmount >= discount.Periods[0].TotalAmount {
		//	fmt.Printf("worked [totalAmount = %d AND discount.Periods[0].TotalAmount = %d]\n", totalAmount, discount.Periods[0].TotalAmount)
		//	contractRB.DiscountAmount = float32(discount.Periods[0].RewardAmount)
		//}
	}
	contractsRB = append(contractsRB, contractRB)

	fmt.Println("<end>")
	return contractsRB
}

func GetRB15ThType(req model.RBRequest, contracts []model.Contract) ([]model.RbDTO, error) {
	var rbDTOsl []model.RbDTO

	for _, contract := range contracts {
		for _, discount := range contract.Discounts {
			fmt.Println(contract.ContractParameters.ContractNumber, "номер договора")
			fmt.Println(discount.Code == RB15Code && discount.IsSelected == true, "Условия")
			if discount.Code == RB15Code && discount.IsSelected == true { // здесь сравниваешь тип скидки и берешь тот тип который тебе нужен
				fmt.Println("Условия ТРУ")
				for _, period := range discount.Periods {
					RbDTO := model.RbDTO{
						ContractNumber:  contract.ContractParameters.ContractNumber,
						StartDate:       period.PeriodFrom,
						EndDate:         period.PeriodTo,
						TypePeriod:      period.Name,
						DiscountPercent: period.DiscountPercent,
						//DiscountAmount:  discountAmount,
						//RewardAmount:         reward,
						//TotalWithoutDiscount: amount,
						LeasePlan:    period.TotalAmount,
						DiscountType: RB15Name,
					}

					//sales, _ := GetSales(reqBrand)
					//amount := CountSales(sales)
					periodFrom, _ := ConvertStringTime(period.PeriodFrom)
					periodTo, _ := ConvertStringTime(period.PeriodTo)
					reqperiodFrom, _ := ConvertStringTime(req.PeriodFrom)
					reqperiodTo, _ := ConvertStringTime(req.PeriodTo)
					if (reqperiodFrom.After(periodFrom) || reqperiodFrom.Equal(periodFrom)) && (reqperiodTo.After(periodTo) || reqperiodTo.Equal(periodTo)) {
						amount := GetSalesByPeriods(req.PeriodFrom, period.PeriodTo, req.ClientCode, contract.View)
						if amount > period.TotalAmount {
							//discountAmount := amount * period.DiscountPercent / 100
							//RbDTO.DiscountAmount = discountAmount
							RbDTO.RewardAmount = float64(period.RewardAmount)
							RbDTO.StartDate = req.PeriodFrom
							RbDTO.EndDate = period.PeriodTo
							RbDTO.TotalWithoutDiscount = amount
							rbDTOsl = append(rbDTOsl, RbDTO)

						} else {
							RbDTO.DiscountAmount = 0
							RbDTO.StartDate = req.PeriodFrom
							RbDTO.EndDate = period.PeriodTo

							rbDTOsl = append(rbDTOsl, RbDTO)
						}

						//// брать продажи ДНО и ДКП
						//
						//RbDTO.StartDate = period.PeriodTo
						//RbDTO.EndDate = req.PeriodTo
						//RbDTO.DiscountAmount = 0
						////RbDTO.DiscountType = "Нет скидок"
						//RbDTO.DiscountType = "Кейс #1"
						//rbDTOsl = append(rbDTOsl, RbDTO)

					} else if (periodFrom.After(reqperiodFrom) || reqperiodFrom.Equal(periodFrom)) && (periodTo.After(reqperiodTo) || reqperiodTo.Equal(periodTo)) {
						amount := GetSalesByPeriods(period.PeriodFrom, req.PeriodTo, req.ClientCode, contract.View)
						if amount > period.TotalAmount {
							//discountAmount := amount * period.DiscountPercent / 100
							RbDTO.RewardAmount = float64(period.RewardAmount)
							RbDTO.StartDate = period.PeriodFrom
							RbDTO.EndDate = req.PeriodTo
							RbDTO.TotalWithoutDiscount = amount

							rbDTOsl = append(rbDTOsl, RbDTO)

						} else {
							RbDTO.DiscountAmount = 0
							RbDTO.StartDate = period.PeriodFrom
							RbDTO.EndDate = req.PeriodTo
							rbDTOsl = append(rbDTOsl, RbDTO)
						}

						//RbDTO.StartDate = req.PeriodFrom
						//RbDTO.EndDate = period.PeriodFrom
						//RbDTO.DiscountAmount = 0
						////RbDTO.DiscountType = "Нет скидок"
						//RbDTO.DiscountType = "Кейс #2"
						//rbDTOsl = append(rbDTOsl, RbDTO)

					} else if (reqperiodFrom.After(periodFrom) || reqperiodFrom.Equal(periodFrom)) && (periodTo.After(reqperiodTo) || reqperiodTo.Equal(periodTo)) {
						//}else if (reqperiodFrom.After(periodFrom) &&  periodTo.After(reqperiodTo)) || (reqperiodFrom.Equal(periodFrom) && reqperiodTo.Equal(periodTo)){
						amount := GetSalesByPeriods(req.PeriodFrom, req.PeriodTo, req.ClientCode, contract.View)
						if amount > period.TotalAmount {
							//discountAmount := amount * period.DiscountPercent / 100
							//RbDTO.DiscountAmount = discountAmount
							RbDTO.RewardAmount = float64(period.RewardAmount)
							RbDTO.StartDate = req.PeriodFrom
							RbDTO.EndDate = req.PeriodTo
							RbDTO.TotalWithoutDiscount = amount

							rbDTOsl = append(rbDTOsl, RbDTO)

						} else {
							RbDTO.DiscountAmount = 0
							RbDTO.StartDate = req.PeriodFrom
							RbDTO.EndDate = req.PeriodTo
							rbDTOsl = append(rbDTOsl, RbDTO)
						}

						//4 кейс
					} else if (periodFrom.After(reqperiodFrom) || reqperiodFrom.Equal(periodFrom)) && (reqperiodTo.After(periodTo) || reqperiodTo.Equal(periodTo)) {
						//}else if ( periodFrom.After(reqperiodFrom) && reqperiodTo.After(periodTo)) || (reqperiodFrom.Equal(periodFrom) && reqperiodTo.Equal(periodTo)){
						amount := GetSalesByPeriods(period.PeriodFrom, period.PeriodTo, req.ClientCode, contract.View)

						if amount > period.TotalAmount {
							discountAmount := amount * period.DiscountPercent / 100
							RbDTO.DiscountAmount = discountAmount
							RbDTO.StartDate = period.PeriodFrom
							RbDTO.EndDate = period.PeriodTo
							RbDTO.TotalWithoutDiscount = amount
							if discountAmount > 0 {
								rbDTOsl = append(rbDTOsl, RbDTO)
							}
						} else {
							RbDTO.StartDate = period.PeriodFrom
							RbDTO.EndDate = period.PeriodTo
							RbDTO.DiscountAmount = 0
							rbDTOsl = append(rbDTOsl, RbDTO)
						}

						//RbDTO.StartDate = req.PeriodFrom
						//RbDTO.EndDate = period.PeriodFrom
						//RbDTO.DiscountAmount = 0
						//RbDTO.DiscountType = "Кейс #4 (начало)"
						//rbDTOsl = append(rbDTOsl, RbDTO)
						//RbDTO.DiscountType = "Кейс #4 (конец)"
						//RbDTO.StartDate = period.PeriodTo
						//RbDTO.EndDate = req.PeriodTo
						//
						//rbDTOsl = append(rbDTOsl, RbDTO)

					}

				}

			}

		}

	}

	return rbDTOsl, nil
}

func GetRB16ThType(req model.RBRequest, contracts []model.Contract) ([]model.RbDTO, error) {
	var rbDTOsl []model.RbDTO

	for _, contract := range contracts {
		for _, discount := range contract.Discounts {
			//if amount >= period.TotalAmount

			fmt.Println(contract.ContractParameters.ContractNumber, "номер договора")
			if discount.Code == RB16Code && discount.IsSelected == true {
				// здесь сравниваешь тип скидки и берешь тот тип который тебе нужен
				fmt.Println("Условия ТРУ")
				for _, period := range discount.Periods {
					rbDTOsl = FillRbwithPeriod(period, req, contract, RB16Name)
				}
			}
		}
	}

	return rbDTOsl, nil
}

func FillRbwithPeriod(period model.DiscountPeriod, req model.RBRequest, contract model.Contract, RBNAME string) []model.RbDTO {
	var rbDTOsl []model.RbDTO
	periodFrom, _ := ConvertStringTime(period.PeriodFrom)
	periodTo, _ := ConvertStringTime(period.PeriodTo)
	reqperiodFrom, _ := ConvertStringTime(req.PeriodFrom)
	reqperiodTo, _ := ConvertStringTime(req.PeriodTo)

	RbDTO := model.RbDTO{
		ContractNumber:  contract.ContractParameters.ContractNumber,
		StartDate:       period.PeriodFrom,
		EndDate:         period.PeriodTo,
		TypePeriod:      period.Name,
		DiscountPercent: period.DiscountPercent,
		//DiscountAmount:  discountAmount,
		//RewardAmount:         reward,
		//TotalWithoutDiscount: amount,
		LeasePlan:    period.TotalAmount,
		DiscountType: RBNAME,
	}
	if (reqperiodFrom.After(periodFrom) || reqperiodFrom.Equal(periodFrom)) && (reqperiodTo.After(periodTo) || reqperiodTo.Equal(periodTo)) {
		amount := GetSalesByPeriods(req.PeriodFrom, period.PeriodTo, req.ClientCode, contract.View)
		if amount > period.TotalAmount {
			discountAmount := amount * period.DiscountPercent / 100
			RbDTO.DiscountAmount = discountAmount
			RbDTO.StartDate = req.PeriodFrom
			RbDTO.EndDate = period.PeriodTo
			RbDTO.TotalWithoutDiscount = amount
			if discountAmount > 0 {
				rbDTOsl = append(rbDTOsl, RbDTO)
			}
		} else {
			RbDTO.DiscountAmount = 0
			RbDTO.StartDate = req.PeriodFrom
			RbDTO.EndDate = period.PeriodTo

			rbDTOsl = append(rbDTOsl, RbDTO)
		}

		//// брать продажи ДНО и ДКП
		//
		//RbDTO.StartDate = period.PeriodTo
		//RbDTO.EndDate = req.PeriodTo
		//RbDTO.DiscountAmount = 0
		////RbDTO.DiscountType = "Нет скидок"
		//RbDTO.DiscountType = "Кейс #1"
		//rbDTOsl = append(rbDTOsl, RbDTO)

	} else if (periodFrom.After(reqperiodFrom) || reqperiodFrom.Equal(periodFrom)) && (periodTo.After(reqperiodTo) || reqperiodTo.Equal(periodTo)) {
		amount := GetSalesByPeriods(period.PeriodFrom, req.PeriodTo, req.ClientCode, contract.View)
		if amount > period.TotalAmount {
			discountAmount := amount * period.DiscountPercent / 100
			RbDTO.DiscountAmount = discountAmount
			RbDTO.StartDate = period.PeriodFrom
			RbDTO.EndDate = req.PeriodTo
			RbDTO.TotalWithoutDiscount = amount
			if discountAmount > 0 {
				rbDTOsl = append(rbDTOsl, RbDTO)
			}
		} else {
			RbDTO.DiscountAmount = 0
			RbDTO.StartDate = period.PeriodFrom
			RbDTO.EndDate = req.PeriodTo
			rbDTOsl = append(rbDTOsl, RbDTO)
		}

		//RbDTO.StartDate = req.PeriodFrom
		//RbDTO.EndDate = period.PeriodFrom
		//RbDTO.DiscountAmount = 0
		////RbDTO.DiscountType = "Нет скидок"
		//RbDTO.DiscountType = "Кейс #2"
		//rbDTOsl = append(rbDTOsl, RbDTO)

	} else if (reqperiodFrom.After(periodFrom) || reqperiodFrom.Equal(periodFrom)) && (periodTo.After(reqperiodTo) || reqperiodTo.Equal(periodTo)) {
		//}else if (reqperiodFrom.After(periodFrom) &&  periodTo.After(reqperiodTo)) || (reqperiodFrom.Equal(periodFrom) && reqperiodTo.Equal(periodTo)){
		amount := GetSalesByPeriods(req.PeriodFrom, req.PeriodTo, req.ClientCode, contract.View)
		if amount > period.TotalAmount {
			discountAmount := amount * period.DiscountPercent / 100
			RbDTO.DiscountAmount = discountAmount
			RbDTO.StartDate = req.PeriodFrom
			RbDTO.EndDate = req.PeriodTo
			RbDTO.TotalWithoutDiscount = amount
			if discountAmount > 0 {
				rbDTOsl = append(rbDTOsl, RbDTO)
			}
		} else {
			RbDTO.DiscountAmount = 0
			RbDTO.StartDate = req.PeriodFrom
			RbDTO.EndDate = req.PeriodTo
			rbDTOsl = append(rbDTOsl, RbDTO)
		}

		//4 кейс
	} else if (periodFrom.After(reqperiodFrom) || reqperiodFrom.Equal(periodFrom)) && (reqperiodTo.After(periodTo) || reqperiodTo.Equal(periodTo)) {
		//}else if ( periodFrom.After(reqperiodFrom) && reqperiodTo.After(periodTo)) || (reqperiodFrom.Equal(periodFrom) && reqperiodTo.Equal(periodTo)){
		amount := GetSalesByPeriods(period.PeriodFrom, period.PeriodTo, req.ClientCode, contract.View)

		if amount > period.TotalAmount {
			discountAmount := amount * period.DiscountPercent / 100
			RbDTO.DiscountAmount = discountAmount
			RbDTO.StartDate = period.PeriodFrom
			RbDTO.EndDate = period.PeriodTo
			RbDTO.TotalWithoutDiscount = amount
			if discountAmount > 0 {
				rbDTOsl = append(rbDTOsl, RbDTO)
			}
		} else {
			RbDTO.StartDate = period.PeriodFrom
			RbDTO.EndDate = period.PeriodTo
			RbDTO.DiscountAmount = 0
			rbDTOsl = append(rbDTOsl, RbDTO)
		}

		//RbDTO.StartDate = req.PeriodFrom
		//RbDTO.EndDate = period.PeriodFrom
		//RbDTO.DiscountAmount = 0
		//RbDTO.DiscountType = "Кейс #4 (начало)"
		//rbDTOsl = append(rbDTOsl, RbDTO)
		//RbDTO.DiscountType = "Кейс #4 (конец)"
		//RbDTO.StartDate = period.PeriodTo
		//RbDTO.EndDate = req.PeriodTo
		//
		//rbDTOsl = append(rbDTOsl, RbDTO)

	}

	fmt.Println("УСЛОВИЯ ПРОШЛИ")
	fmt.Println("ПЕРИОДЫ которые прошли", period)

	//fmt.Println("AMOUNT", amount, period.SalesAmount)
	//if amount >= period.TotalAmount {
	//discountAmount := amount * period.DiscountPercent / 100

	//} else {
	//rbDTOsl, _ = GetNil12Rb(rbDTOsl, contract, period, RB17Name)

	return rbDTOsl

}

func ConvertStringTime(timeStr string) (time.Time, error) {
	layoutISO := "02.1.2006"
	periodFrom, err := time.Parse(layoutISO, timeStr)
	if err != nil {
		fmt.Println(err)
		return time.Time{}, err

	}

	return periodFrom, nil

}

func GetRB17ThType(req model.RBRequest, contracts []model.Contract) ([]model.RbDTO, error) {
	var rbDTOsl []model.RbDTO

	for _, contract := range contracts {
		fmt.Println(contract.ID, "ID договора")

		//fmt.Println("данные продаж")
		//fmt.Println("Сумма", sales)

		for _, discount := range contract.Discounts {
			fmt.Println(contract.ContractParameters.ContractNumber, "номер договора")

			if discount.Code == RB17Code && discount.IsSelected == true { // здесь сравниваешь тип скидки и берешь тот тип который тебе нужен
				fmt.Println("Условия ТРУ")
				for _, period := range discount.Periods {
					periodFrom, _ := ConvertStringTime(period.PeriodFrom)
					periodTo, _ := ConvertStringTime(period.PeriodTo)
					reqperiodFrom, _ := ConvertStringTime(req.PeriodFrom)
					reqperiodTo, _ := ConvertStringTime(req.PeriodTo)

					RbDTO := model.RbDTO{
						ContractNumber:  contract.ContractParameters.ContractNumber,
						StartDate:       period.PeriodFrom,
						EndDate:         period.PeriodTo,
						TypePeriod:      period.Name,
						DiscountPercent: period.DiscountPercent,
						//DiscountAmount:  discountAmount,
						//RewardAmount:         reward,
						//TotalWithoutDiscount: amount,
						LeasePlan:    period.TotalAmount,
						DiscountType: RB17Name,
					}
					if (reqperiodFrom.After(periodFrom) || reqperiodFrom.Equal(periodFrom)) && (reqperiodTo.After(periodTo) || reqperiodTo.Equal(periodTo)) {
						amount := GetSalesByPeriods(req.PeriodFrom, period.PeriodTo, req.ClientCode, contract.View)
						discountAmount := amount * period.DiscountPercent / 100
						RbDTO.DiscountAmount = discountAmount
						RbDTO.StartDate = req.PeriodFrom
						RbDTO.EndDate = period.PeriodTo
						RbDTO.TotalWithoutDiscount = amount
						if discountAmount > 0 {
							rbDTOsl = append(rbDTOsl, RbDTO)
						}

						//// брать продажи ДНО и ДКП
						//
						//RbDTO.StartDate = period.PeriodTo
						//RbDTO.EndDate = req.PeriodTo
						//RbDTO.DiscountAmount = 0
						////RbDTO.DiscountType = "Нет скидок"
						//RbDTO.DiscountType = "Кейс #1"
						//rbDTOsl = append(rbDTOsl, RbDTO)

					} else if (periodFrom.After(reqperiodFrom) || reqperiodFrom.Equal(periodFrom)) && (periodTo.After(reqperiodTo) || reqperiodTo.Equal(periodTo)) {
						amount := GetSalesByPeriods(period.PeriodFrom, req.PeriodTo, req.ClientCode, contract.View)
						discountAmount := amount * period.DiscountPercent / 100
						RbDTO.DiscountAmount = discountAmount
						RbDTO.StartDate = period.PeriodFrom
						RbDTO.EndDate = req.PeriodTo
						RbDTO.TotalWithoutDiscount = amount
						if discountAmount > 0 {
							rbDTOsl = append(rbDTOsl, RbDTO)
						}

						//RbDTO.StartDate = req.PeriodFrom
						//RbDTO.EndDate = period.PeriodFrom
						//RbDTO.DiscountAmount = 0
						////RbDTO.DiscountType = "Нет скидок"
						//RbDTO.DiscountType = "Кейс #2"
						//rbDTOsl = append(rbDTOsl, RbDTO)

					} else if (reqperiodFrom.After(periodFrom) || reqperiodFrom.Equal(periodFrom)) && (periodTo.After(reqperiodTo) || reqperiodTo.Equal(periodTo)) {
						//}else if (reqperiodFrom.After(periodFrom) &&  periodTo.After(reqperiodTo)) || (reqperiodFrom.Equal(periodFrom) && reqperiodTo.Equal(periodTo)){
						amount := GetSalesByPeriods(req.PeriodFrom, req.PeriodTo, req.ClientCode, contract.View)
						discountAmount := amount * period.DiscountPercent / 100
						RbDTO.DiscountAmount = discountAmount
						RbDTO.StartDate = req.PeriodFrom
						RbDTO.EndDate = req.PeriodTo
						RbDTO.TotalWithoutDiscount = amount
						if discountAmount > 0 {
							rbDTOsl = append(rbDTOsl, RbDTO)
						}

						//4 кейс
					} else if (periodFrom.After(reqperiodFrom) || reqperiodFrom.Equal(periodFrom)) && (reqperiodTo.After(periodTo) || reqperiodTo.Equal(periodTo)) {
						//}else if ( periodFrom.After(reqperiodFrom) && reqperiodTo.After(periodTo)) || (reqperiodFrom.Equal(periodFrom) && reqperiodTo.Equal(periodTo)){
						amount := GetSalesByPeriods(period.PeriodFrom, period.PeriodTo, req.ClientCode, contract.View)
						discountAmount := amount * period.DiscountPercent / 100
						RbDTO.DiscountAmount = discountAmount
						RbDTO.StartDate = period.PeriodFrom
						RbDTO.EndDate = period.PeriodTo
						RbDTO.TotalWithoutDiscount = amount
						if discountAmount > 0 {
							rbDTOsl = append(rbDTOsl, RbDTO)
						}

						//RbDTO.StartDate = req.PeriodFrom
						//RbDTO.EndDate = period.PeriodFrom
						//RbDTO.DiscountAmount = 0
						//RbDTO.DiscountType = "Кейс #4 (начало)"
						//rbDTOsl = append(rbDTOsl, RbDTO)
						//RbDTO.DiscountType = "Кейс #4 (конец)"
						//RbDTO.StartDate = period.PeriodTo
						//RbDTO.EndDate = req.PeriodTo
						//
						//rbDTOsl = append(rbDTOsl, RbDTO)

					}

					fmt.Println("УСЛОВИЯ ПРОШЛИ")
					fmt.Println("ПЕРИОДЫ которые прошли", period)

					//fmt.Println("AMOUNT", amount, period.SalesAmount)
					//if amount >= period.TotalAmount {
					//discountAmount := amount * period.DiscountPercent / 100

					//} else {
					//rbDTOsl, _ = GetNil12Rb(rbDTOsl, contract, period, RB17Name)
				}

			}

		}

	}

	return rbDTOsl, nil
}

func GetSalesByPeriods(start, end string, code string, contractType string) float64 {
	reqBrand := model.ReqBrand{
		ClientCode:     code,
		DateStart:      start,
		DateEnd:        end,
		TypeValue:      "",
		TypeParameters: nil,
		SchemeType:     contractType,
		//
	}
	sales, _ := GetSalesNEw(reqBrand)
	amount := CountSales(sales)

	return amount

}

func GetSalesRegionsTotalAmount(SalesArr model.Sales, regions []model.Regions) (totalAmount float64) {
	for _, region := range regions {
		for _, sale := range SalesArr.SalesArr {
			if sale.RegionCode == region.RegionCode {
				totalAmount += float64(sale.Total)
			}
		}
	}

	return totalAmount
}

func GetPurchaseTotalAmount(purchases model.Purchase) (totalAmount float64) {
	for _, purchase := range purchases.PurchaseArr {
		totalAmount += purchase.Total
	}

	return totalAmount
}

//func GetRB2ndType(rbReq model.RBRequest) []model.RbDTO {
//	brandTotal := map[string]float32{}
//	var rbDtoSl []model.RbDTO
//
//	rbBrand := model.ReqBrand{
//		ClientCode: rbReq.BIN,
//		DateStart: rbReq.PeriodFrom,
//		DateEnd:   rbReq.PeriodTo,
//	}
//
//	// берем бренды и их Total // общую сумму не зависимо от договора
//	sales, _ := GetSales(rbBrand)
//
//	// тут считаем общую сумму каждого бренда
//	for _, sale := range sales.SalesArr {
//		// считаем общую сумму по брендам, и чтобы они не дублировались
//		brandTotal[sale.BrandName] += sale.Total
//	}
//
//	// берем скидки по брендам и название брендов
//	dataBrands, err := repository.GetIDByBIN(rbReq.BIN)
//	if err != nil {
//		return nil
//	}
//	fmt.Println("dataBrand", dataBrands)
//
//	for brandName, total := range brandTotal {
//
//		for _, brand := range dataBrands {
//			// сравниваем бренды, то есть если бин - 160140011654- то у него всего 2 бренда
//			//[Sante:1579 Silver Care:19410]
//			if brand.Brand == brandName {
//				value, _ := strconv.ParseFloat(brand.DiscountPercent, 32)
//				dicsount := float32(value)
//				TotalPercent := (total * dicsount) / 100
//				rbdro := model.RbDTO{
//					ID:                   0,
//					ContractNumber:       brand.ContractNumber,
//					StartDate:            rbReq.PeriodFrom,
//					EndDate:              rbReq.PeriodTo,
//					TypePeriod:           "",
//					BrandName:            brandName,
//					ProductCode:          "",
//					DiscountPercent:      dicsount,
//					DiscountAmount:       TotalPercent,
//					TotalWithoutDiscount: 0,
//					LeasePlan:            0,
//					RewardAmount:         0,
//					DiscountType:         RB2Name,
//				}
//				rbDtoSl = append(rbDtoSl, rbdro)
//
//			}
//		}
//	}
//	//}
//
//	return rbDtoSl
//
//}

//func GetRB2ndType(rbReq model.RBRequest) []model.RbDTO {
//	brandTotal := map[string]float32{}
//	var rbDtoSl []model.RbDTO

func GetRegionsSales(salesIn model.Sales, regions []model.Regions) (salesOut model.Sales) {
	var sales []model.SalesArr
	for _, region := range regions {
		for _, sale := range salesIn.SalesArr {
			if region.RegionCode == sale.RegionCode {
				sales = append(sales, sale)
			}
		}
	}

	salesOut.SalesArr = sales
	return salesOut
}

func GetRB2ndType(rb model.RBRequest, contracts []model.Contract) (rbDTO []model.RbDTO) {
	var schema []string

	for _, contract := range contracts {
		for _, regionCode := range contract.Regions {
			schema = append(schema, regionCode.RegionCode)
		}

		req := model.ReqBrand{
			ClientCode: rb.ClientCode,
			DateStart:  rb.PeriodFrom,
			DateEnd:    rb.PeriodTo,
			SchemeType: contract.View,
		}
		sales, _ := GetSales(req)

		mapBrands := CountSalesByBrand(sales)

		innerMapBrands := mapBrands
		if contract.View == "PF" {
			regionSales := GetRegionsSales(sales, contract.Regions)
			innerMapBrands = CountSalesByBrand(regionSales)
		}

		for _, discount := range contract.Discounts {
			if discount.Code == RB2Code && discount.IsSelected == true {

				for _, discountBrand := range discount.DiscountBrands {
					//   01.01.2022                01.01.2022       31.03.2022                 <= 31.03.2022
					//if discountBrand.PeriodFrom >= rb.PeriodFrom && discountBrand.PeriodTo <= rb.PeriodTo {

					for _, dataBrand := range discountBrand.Brands {
						for brand, total := range innerMapBrands {
							if brand == dataBrand.BrandName {
								var discountAmount float64
								//if total >= dataBrand.PurchaseAmount {
								discountAmount = total * dataBrand.DiscountPercent / 100

								rbDTO = append(rbDTO, model.RbDTO{
									ContractNumber:       contract.ContractParameters.ContractNumber,
									StartDate:            contract.ContractParameters.StartDate,
									EndDate:              contract.ContractParameters.EndDate,
									BrandName:            dataBrand.BrandName,
									ProductCode:          dataBrand.BrandCode,
									DiscountPercent:      dataBrand.DiscountPercent,
									TotalWithoutDiscount: total,
									DiscountAmount:       discountAmount,
									DiscountType:         RB2Name,
								})
							}

						}

					}
				}
			}
		}
	}
	//}

	return rbDTO
}

//func GetRB3rdType(request model.RBRequest, contracts []model.Contract) ([]model.RbDTO, error) {
//
//	//req := model.ReqBrand{
//	//	ClientCode:      request.BIN,
//	//	Beneficiary:    request.ContractorName,
//	//	DateStart:      request.PeriodFrom,
//	//	DateEnd:        request.PeriodTo,
//	//	Type:           "sales",
//	//	TypeValue:      "sku",
//	//	TypeParameters: GetAllProductsSku(contracts),
//	//}
//	//
//	//sales, err := GetBrandSales(req)
//	//if err != nil {
//	//	return nil, err
//	//}
//
//	externalCodes := GetExternalCode(request.BIN)
//	contractsCode := JoinContractCode(externalCodes)
//
//	req := model.ReqBrand{
//		ClientCode:      request.BIN,
//		DateStart:      request.PeriodFrom,
//		DateEnd:        request.PeriodTo,
//		TypeValue:      "",
//		TypeParameters: nil,
//		Contracts:      contractsCode, // необходимо получить коды контрактов
//	}
//	purchase, _ := GetPurchase(req)
//	totalAmount := GetPurchaseTotalAmount(purchase)
//
//	//fmt.Printf("req \n\n%+v\n\n", req)
//	//fmt.Printf("SALES \n\n%+v\n\n", sales)
//
//	var RBs []model.RbDTO
//	fmt.Println("*********************************************")
//	for _, contract := range contracts {
//		for _, product := range contract.Products {
//			//total := GetTotalSalesForSku(sales, product.Sku)
//			rb := model.RbDTO{
//				ID:              contract.ID,
//				ContractNumber:  contract.ContractParameters.ContractNumber,
//				StartDate:       contract.ContractParameters.StartDate,
//				EndDate:         contract.ContractParameters.EndDate,
//				ProductCode:     product.Sku,
//				DiscountPercent: product.DiscountPercent,
//				LeasePlan:       product.Plan,
//				DiscountType:    RB3Name,
//			}
//			if totalAmount >= float64(product.Plan) {
//				rb.DiscountAmount = float32(totalAmount) * rb.DiscountPercent / 100
//			} else {
//				rb.DiscountAmount = 0
//			}
//
//			RBs = append(RBs, rb)
//
//		}
//	}
//	fmt.Println("*********************************************")
//	return RBs, nil
//}\

func GetRB3rdType(request model.RBRequest, contracts []model.Contract) ([]model.RbDTO, error) {

	//externalCodes := GetExternalCode(request.BIN)
	//contractsCode := JoinContractCode(externalCodes)
	req := model.ReqBrand{
		ClientCode:     request.ClientCode,
		DateStart:      request.PeriodFrom,
		DateEnd:        request.PeriodTo,
		TypeValue:      "",
		TypeParameters: nil,
		//Contracts:      contractsCode, // необходимо получить коды контрактов
	}
	purchase, _ := GetPurchase(req)
	totalAmount := GetPurchaseTotalAmount(purchase)

	var RBs []model.RbDTO

	for _, contract := range contracts {
		for _, discount := range contract.Discounts {
			if discount.Code == RB3Code && discount.IsSelected == true {
				for _, product := range discount.Products {
					//total := GetTotalSalesForSku(sales, product.Sku)
					rb := model.RbDTO{
						ID:                   contract.ID,
						ContractNumber:       contract.ContractParameters.ContractNumber,
						StartDate:            contract.ContractParameters.StartDate,
						EndDate:              contract.ContractParameters.EndDate,
						ProductCode:          product.Sku,
						DiscountPercent:      product.DiscountPercent,
						TotalWithoutDiscount: totalAmount,
						LeasePlan:            product.Plan,
						DiscountType:         RB3Name,
					}
					if totalAmount >= product.Plan {
						rb.DiscountAmount = totalAmount * rb.DiscountPercent / 100

					} else {
						rb.DiscountAmount = 0
					}

					RBs = append(RBs, rb)

				}
			}
		}
	}

	return RBs, nil
}

func GetRB14ThType(request model.RBRequest, contracts []model.Contract) ([]model.RbDTO, error) {
	//externalCodes := GetExternalCode(request.BIN)
	//contractsCode := JoinContractCode(externalCodes)

	var RBs []model.RbDTO
	for _, contract := range contracts {
		req := model.ReqBrand{
			ClientCode:     request.ClientCode,
			DateStart:      request.PeriodFrom,
			DateEnd:        request.PeriodTo,
			TypeValue:      "",
			TypeParameters: nil,
			SchemeType:     contract.View,
			//Contracts:      contractsCode, // необходимо получить коды контрактов
		}
		//purchase, _ := GetPurchase(req)
		sales, _ := GetSales(req)
		totalAmount := CountSales(sales)
		innerSalesTotal := totalAmount

		if contract.View == "PF" {
			innerSalesTotal = TotalRegionsSales(sales, contract.Regions)
		}
		for _, discount := range contract.Discounts {
			if discount.Code == RB14Code && discount.IsSelected == true {
				for _, product := range discount.Products {
					//total := GetTotalSalesForSku(sales, product.Sku)
					rb := model.RbDTO{
						ID:                   contract.ID,
						ContractNumber:       contract.ContractParameters.ContractNumber,
						StartDate:            contract.ContractParameters.StartDate,
						EndDate:              contract.ContractParameters.EndDate,
						ProductCode:          product.Sku,
						DiscountPercent:      float64(product.DiscountPercent),
						LeasePlan:            product.Plan,
						TotalWithoutDiscount: innerSalesTotal,
						DiscountType:         RB14Name,
					}
					if innerSalesTotal >= product.Plan {
						rb.DiscountAmount = innerSalesTotal * rb.DiscountPercent / 100
					} else {
						rb.DiscountAmount = 0
					}

					RBs = append(RBs, rb)

				}
			}
		}
	}

	return RBs, nil
}

func GetRB4thType(request model.RBRequest, contracts []model.Contract) (rbDTO []model.RbDTO, err error) {
	fmt.Println("ЗАПРОС", request)

	for _, contract := range contracts {

		for _, discount := range contract.Discounts {
			if discount.Code == RB4Code && discount.IsSelected == true {
				req := model.ReqBrand{
					ClientCode:  request.ClientCode,
					Beneficiary: request.ContractorName,
					DateStart:   request.PeriodFrom,
					DateEnd:     request.PeriodTo,
					SchemeType:  contract.View,
					//Contracts:   contractsCode,
				}
				fmt.Println("регион", contract.View)

				//sales, err := GetSales(req)
				purchase, _ := GetPurchase(req)
				//totalAmountPurchase := CountPurchaseByCode(purchase)
				amount := CountPurchase(purchase)

				//totalAmountPurchase := GetTotalAmountPurchase(purchase)

				//log.Printf("[CHECK PRES SAlES: %+v\n", purchase)
				//fmt.Println("Условия прошли")
				var discountAmount float64
				//if repository.DoubtedDiscountExecutionCheck(request, contract.ContractParameters.ContractNumber, discount.Code) {
				//for _, amount := range totalAmountPurchase {
				if amount > 0 {
					discountAmount = amount * discount.DiscountPercent / 100
					rbDTO = append(rbDTO, model.RbDTO{
						ContractNumber:       contract.ContractParameters.ContractNumber,
						StartDate:            contract.ContractParameters.StartDate,
						EndDate:              contract.ContractParameters.EndDate,
						DiscountPercent:      discount.DiscountPercent,
						DiscountAmount:       discountAmount,
						TotalWithoutDiscount: amount,
						DiscountType:         RB4Name,
					})

				} else {
					rbDTO = append(rbDTO, model.RbDTO{
						ContractNumber:       contract.ContractParameters.ContractNumber,
						StartDate:            contract.ContractParameters.StartDate,
						EndDate:              contract.ContractParameters.EndDate,
						DiscountPercent:      discount.DiscountPercent,
						DiscountAmount:       0,
						TotalWithoutDiscount: amount,
						DiscountType:         RB4Name,
					})
				}
				log.Printf("[CHECK PRES DISCOUNT AMOUNT]: %v\n", discountAmount)

				log.Printf("[CHECK PRES DISCOUNT PERCENT]: %f\n", discount.DiscountPercent)
				//log.Printf("[CHECK PRES TOTAL AMOUNT]: %f\n", totalAmountPurchase)

				//log.Println("[CHECK PRES TRUE/FALSE]: ", repository.DoubtedDiscountExecutionCheck(request, contract.ContractParameters.ContractNumber, discount.Code))

				log.Printf("CHECK PRES DISCOUNT rbDTO %+v\n", rbDTO)
			}
		}
	}
	//	}
	return rbDTO, nil

}

func GetRB5thType(request model.RBRequest, contracts []model.Contract) (rbDTO []model.RbDTO, err error) {
	for _, contract := range contracts {
		for _, discount := range contract.Discounts {
			if discount.Code == RB5Code && discount.IsSelected == true {
				rbDTO, err = RB5thTypeDetails(request, contract, discount)
				if err != nil {
					return nil, err
				}
			}
		}
	}

	return rbDTO, nil
}

func RB5thTypeDetails(request model.RBRequest, contract model.Contract, discount model.Discount) (rbDTO []model.RbDTO, err error) {
	log.Printf("\n[DISCOUNT_DETAILS] %+v\n", discount)
	for _, discountBrand := range discount.DiscountBrands {
		fmt.Println("data brands", discountBrand)

		fmt.Println("data brands", discountBrand)

		reqBrand := model.ReqBrand{
			ClientCode:     request.ClientCode,
			DateStart:      discountBrand.PeriodFrom,
			DateEnd:        discountBrand.PeriodTo,
			TypeValue:      "",
			TypeParameters: nil,
			SchemeType:     contract.View,
		}
		periodFrom, _ := ConvertStringTime(discountBrand.PeriodFrom)
		periodTo, _ := ConvertStringTime(discountBrand.PeriodTo)
		reqperiodFrom, _ := ConvertStringTime(request.PeriodFrom)
		reqperiodTo, _ := ConvertStringTime(request.PeriodTo)
		//GetPurchase(reqBrand)
		purchase, _ := GetPurchaseBrandOnly(reqBrand)
		totalAmount := GetPurchaseTotalAmount(purchase)
		fmt.Println("SUMMA PURCHASE", purchase)
		fmt.Println("OTOTAL", totalAmount)
		if reqperiodFrom.Before(periodFrom) || reqperiodFrom.Equal(periodFrom) && reqperiodTo.After(periodTo) || reqperiodTo.Equal(periodTo) {
			purchase, _ := GetPurchaseBrandOnly(reqBrand)
			totalAmount := GetPurchaseTotalAmount(purchase)
			fmt.Println("SUMMA PURCHASE", purchase)
			fmt.Println("OTOTAL", totalAmount)

			for _, brand := range discountBrand.Brands {

				//totalAmount := GetTotalPurchasesForBrands(sales, brand.BrandName)
				var discountAmount float64
				if totalAmount >= brand.PurchaseAmount {
					discountAmount = totalAmount * brand.DiscountPercent / 100
				}

				rbDTO = append(rbDTO, model.RbDTO{

					ContractNumber:       contract.ContractParameters.ContractNumber,
					StartDate:            discountBrand.PeriodFrom,
					EndDate:              discountBrand.PeriodTo,
					BrandName:            brand.BrandName,
					ProductCode:          brand.BrandCode,
					DiscountPercent:      brand.DiscountPercent,
					DiscountAmount:       discountAmount,
					TotalWithoutDiscount: totalAmount,
					DiscountType:         RB5Name,
				})
			}
		}
	}

	return rbDTO, nil
}

func GetRB18thType(request model.RBRequest, contracts []model.Contract) (rbDTO []model.RbDTO, err error) {
	discount, err := FillDiscount(request, contracts, RB18Code, RB18Name)
	if err != nil {
		return nil, err
	}

	return discount, nil

}

func GetRB6thType(request model.RBRequest, contracts []model.Contract) (rbDTO []model.RbDTO, err error) {
	discount, err := FillDiscount(request, contracts, RB6Code, RB6Name)
	if err != nil {
		return nil, err
	}

	return discount, nil

}

func FillDiscount(request model.RBRequest, contracts []model.Contract, RbCode string, RBName string) (rbDTO []model.RbDTO, err error) {
	fmt.Println("ЗАПРОС", request)

	for _, contract := range contracts {

		for _, discount := range contract.Discounts {
			if discount.Code == RbCode && discount.IsSelected == true {
				req := model.ReqBrand{
					ClientCode:  request.ClientCode,
					Beneficiary: request.ContractorName,
					DateStart:   request.PeriodFrom,
					DateEnd:     request.PeriodTo,
					SchemeType:  contract.View,
					//Contracts:   contractsCode,
				}
				fmt.Println("регион", contract.View)

				purchase, _ := GetPurchase(req)

				amount := CountPurchase(purchase)

				//totalAmountPurchase := GetTotalAmountPurchase(purchase)

				//log.Printf("[CHECK PRES SAlES: %+v\n", purchase)
				//fmt.Println("Условия прошли")

				//if repository.DoubtedDiscountExecutionCheck(request, contract.ContractParameters.ContractNumber, discount.Code) {
				//for _, amount := range totalAmountPurchase {
				if amount > 0 {
					discountAmount := amount * discount.DiscountPercent / 100
					rbDTO = append(rbDTO, model.RbDTO{
						ContractNumber:       contract.ContractParameters.ContractNumber,
						StartDate:            request.PeriodFrom,
						EndDate:              request.PeriodTo,
						DiscountPercent:      discount.DiscountPercent,
						DiscountAmount:       discountAmount,
						TotalWithoutDiscount: amount,
						DiscountType:         RBName,
					})

				} else {
					rbDTO = append(rbDTO, model.RbDTO{
						ContractNumber:       contract.ContractParameters.ContractNumber,
						StartDate:            contract.ContractParameters.StartDate,
						EndDate:              contract.ContractParameters.EndDate,
						DiscountPercent:      discount.DiscountPercent,
						DiscountAmount:       0,
						TotalWithoutDiscount: amount,
						DiscountType:         RBName,
					})
				}
				//log.Printf("[CHECK PRES DISCOUNT AMOUNT]: %v\n", discountAmount)

			}
			log.Printf("[CHECK PRES DISCOUNT PERCENT]: %f\n", discount.DiscountPercent)
			//log.Printf("[CHECK PRES TOTAL AMOUNT]: %f\n", totalAmountPurchase)

			//log.Println("[CHECK PRES TRUE/FALSE]: ", repository.DoubtedDiscountExecutionCheck(request, contract.ContractParameters.ContractNumber, discount.Code))

			log.Printf("CHECK PRES DISCOUNT rbDTO %+v\n", rbDTO)
		}
	}
	//}
	//	}
	return rbDTO, nil

}

//func GetRB6thType(rb model.RBRequest, contracts []model.Contract) (rbDTO []model.RbDTO, err error) {
//	req := model.ReqBrand{
//		ClientCode: rb.ClientCode,
//		DateStart:  rb.PeriodFrom,
//		DateEnd:    rb.PeriodTo,
//	}
//	sales, err := GetSales(req)
//	mapBrands := CountSalesByBrand(sales)
//
//	for _, contract := range contracts {
//		innerMapBrands := mapBrands
//		if contract.View == "PF" {
//			regionSales := GetRegionsSales(sales, contract.Regions)
//			innerMapBrands = CountSalesByBrand(regionSales)
//		}
//
//		for _, discount := range contract.Discounts {
//			if discount.Code == RB6Code && discount.IsSelected == true {
//
//				for _, discountBrand := range discount.DiscountBrands {
//					//   01.01.2022                01.01.2022       31.03.2022                 <= 31.03.2022
//					if discountBrand.PeriodFrom >= rb.PeriodFrom && discountBrand.PeriodTo <= rb.PeriodTo {
//						for _, dataBrand := range discountBrand.Brands {
//							for brand, total := range innerMapBrands {
//								if brand == dataBrand.BrandName {
//									var discountAmount float64
//									if total >= dataBrand.PurchaseAmount {
//										discountAmount = total * dataBrand.DiscountPercent / 100
//									}
//									rbDTO = append(rbDTO, model.RbDTO{
//										ContractNumber:       contract.ContractParameters.ContractNumber,
//										StartDate:            discountBrand.PeriodFrom,
//										EndDate:              discountBrand.PeriodTo,
//										BrandName:            dataBrand.BrandName,
//										ProductCode:          dataBrand.BrandCode,
//										DiscountPercent:      dataBrand.DiscountPercent,
//										TotalWithoutDiscount: total,
//										DiscountAmount:       discountAmount,
//										DiscountType:         RB6Name,
//									})
//								}
//
//							}
//
//						}
//					}
//				}
//			}
//		}
//	}
//
//	return rbDTO, nil
//}

func GetRB7thType(req model.RBRequest, contracts []model.Contract) (rbDTO []model.RbDTO, err error) {

	for _, contract := range contracts {
		//req := model.ReqBrand{
		//	ClientCode: rb.ClientCode,
		//	DateStart:  rb.PeriodFrom,
		//	DateEnd:    rb.PeriodTo,
		//	SchemeType: contract.View,
		//}

		for _, discount := range contract.Discounts {
			if discount.Code == RB7Code && discount.IsSelected == true {
				for _, discountBrand := range discount.DiscountBrands {
					reqBrand := model.ReqBrand{
						ClientCode:     req.ClientCode,
						DateStart:      discountBrand.PeriodFrom,
						DateEnd:        discountBrand.PeriodTo,
						TypeValue:      "",
						TypeParameters: nil,
						SchemeType:     contract.View,
						//Contracts:      contractsCode, // необходимо получить коды контрактов
					}
					sales, _ := GetSales(reqBrand)
					mapBrands := CountSalesByBrand(sales)
					innerMapBrands := mapBrands
					if contract.View == "PF" {
						regionSales := GetRegionsSales(sales, contract.Regions)
						innerMapBrands = CountSalesByBrand(regionSales)
					}

					periodFrom, _ := ConvertStringTime(discountBrand.PeriodFrom)
					periodTo, _ := ConvertStringTime(discountBrand.PeriodTo)
					reqperiodFrom, _ := ConvertStringTime(req.PeriodFrom)
					reqperiodTo, _ := ConvertStringTime(req.PeriodTo)
					if reqperiodFrom.Before(periodFrom) || reqperiodFrom.Equal(periodFrom) && reqperiodTo.After(periodTo) || reqperiodTo.Equal(periodTo) {
						for _, dataBrand := range discountBrand.Brands {
							for brand, total := range innerMapBrands {
								if brand == dataBrand.BrandName {
									var discountAmount float64
									if total >= dataBrand.PurchaseAmount {
										discountAmount = total * dataBrand.DiscountPercent / 100
									}
									rbDTO = append(rbDTO, model.RbDTO{
										ContractNumber:       contract.ContractParameters.ContractNumber,
										StartDate:            discountBrand.PeriodFrom,
										EndDate:              discountBrand.PeriodTo,
										BrandName:            dataBrand.BrandName,
										ProductCode:          dataBrand.BrandCode,
										DiscountPercent:      dataBrand.DiscountPercent,
										TotalWithoutDiscount: total,
										DiscountAmount:       discountAmount,
										DiscountType:         RB7Name,
									})
								}

							}

						}
					}
				}
			}
		}
	}

	return rbDTO, nil
}

func GetRB8thType(request model.RBRequest, contracts []model.Contract) ([]model.RbDTO, error) {
	var RBs []model.RbDTO
	fmt.Println("*********************************************")
	for _, contract := range contracts {
		reqBrand := model.ReqBrand{
			ClientCode:     request.ClientCode,
			DateStart:      request.PeriodFrom,
			DateEnd:        request.PeriodTo,
			TypeValue:      "",
			TypeParameters: nil,
			//Contracts:      contractsCode, // необходимо получить коды контрактов
			SchemeType: contract.View,
		}

		for _, discount := range contract.Discounts {
			if discount.Code == "DISCOUNT_FOR_LEASE_GENERAL" && discount.IsSelected == true {
				purchase, _ := GetPurchase(reqBrand)
				totalAmount := GetPurchaseTotalAmount(purchase)
				total := totalAmount * discount.DiscountPercent / 100
				sprintf := fmt.Sprintf("%8.2f", total)
				fmt.Println(sprintf, "SS")

				amount, _ := strconv.ParseFloat(sprintf, 64)
				rb := model.RbDTO{
					ID:                   contract.ID,
					ContractNumber:       contract.ContractParameters.ContractNumber,
					StartDate:            request.PeriodFrom,
					EndDate:              request.PeriodTo,
					DiscountPercent:      discount.DiscountPercent,
					DiscountAmount:       amount,
					TotalWithoutDiscount: totalAmount,
					DiscountType:         RB8Name,
				}

				RBs = append(RBs, rb)
			}
		}
	}
	fmt.Println("*********************************************")
	return RBs, nil
}

func GetRB9thType(request model.RBRequest, contracts []model.Contract) ([]model.RbDTO, error) {
	//req := model.ReqBrand{
	//	ClientCode:      request.BIN,
	//	Beneficiary:    request.ContractorName,
	//	DateStart:      request.PeriodFrom,
	//	DateEnd:        request.PeriodTo,
	//	Type:           "sales",
	//	TypeValue:      "",
	//	TypeParameters: nil,
	//}
	//
	//sales, err := GetBrandSales(req)
	//if err != nil {
	//	return nil, err
	//}

	//present := model.ReqBrand{
	//	ClientCode:      request.BIN,
	//	Beneficiary:    "",
	//	DateStart:      request.PeriodFrom,
	//	DateEnd:        request.PeriodTo,
	//	Type:           "",
	//	TypeValue:      "",
	//	TypeParameters: nil,
	//	Contracts:      nil,
	//}
	//
	//sales, err := GetSales1C(present, "sales_brand_only")
	////sales, err := GetSales(req)
	//if err != nil {
	//	return nil, err
	//}
	//
	//totalAmount := GetTotalAmount(sales)

	//externalCodes := GetExternalCode(request.BIN)
	//contractsCode := JoinContractCode(externalCodes)

	var RBs []model.RbDTO
	fmt.Println("*********************************************")
	for _, contract := range contracts {
		reqBrand := model.ReqBrand{
			ClientCode:     request.ClientCode,
			DateStart:      request.PeriodFrom,
			DateEnd:        request.PeriodTo,
			TypeValue:      "",
			TypeParameters: nil,
			SchemeType:     contract.View,
			//Contracts:      contractsCode, // необходимо получить коды контрактов
		}
		purchase, _ := GetPurchase(reqBrand)
		totalAmount := GetPurchaseTotalAmount(purchase)

		for _, discount := range contract.Discounts {
			if discount.Code == RB9Code && discount.IsSelected == true {
				rb := model.RbDTO{
					ID:                   contract.ID,
					ContractNumber:       contract.ContractParameters.ContractNumber,
					StartDate:            contract.ContractParameters.StartDate,
					EndDate:              contract.ContractParameters.EndDate,
					DiscountPercent:      discount.DiscountPercent,
					DiscountAmount:       totalAmount * discount.DiscountPercent / 100,
					TotalWithoutDiscount: totalAmount,
					DiscountType:         RB9Name,
				}

				RBs = append(RBs, rb)
			}
		}
	}
	fmt.Println("*********************************************")
	return RBs, nil
}

func GetRb10thType(request model.RBRequest, contracts []model.Contract) (rbDTO []model.RbDTO, err error) {

	for _, contract := range contracts {
		req := model.ReqBrand{
			ClientCode:  request.ClientCode,
			Beneficiary: request.ContractorName,
			DateStart:   request.PeriodFrom,
			DateEnd:     request.PeriodTo,
			Type:        "sales",
			SchemeType:  contract.View,
		}

		sales, err := GetSales(req)
		if err != nil {
			return nil, err
		}

		totalAmount := GetTotalAmount(sales)

		innerSalesTotal := totalAmount

		if contract.View == "PF" {
			innerSalesTotal = TotalRegionsSales(sales, contract.Regions)
		}
		for _, discount := range contract.Discounts {
			if discount.Code == RB10Code && discount.IsSelected == true {
				var discountAmount float64
				if innerSalesTotal == 0 {
					discountAmount = innerSalesTotal * discount.DiscountPercent / 100
				}
				rbDTO = append(rbDTO, model.RbDTO{
					ContractNumber:       contract.ContractParameters.ContractNumber,
					StartDate:            request.PeriodTo,
					EndDate:              request.PeriodFrom,
					DiscountPercent:      discount.DiscountPercent,
					DiscountAmount:       discountAmount,
					TotalWithoutDiscount: innerSalesTotal,
					DiscountType:         RB10Name,
				})

			}
		}
	}
	return rbDTO, nil
}

func GetTotalAmount(sales model.Sales) float64 {
	var amount float64
	for _, s := range sales.SalesArr {
		amount += s.Total
	}

	return amount
}

func GetRB11thType(req model.RBRequest, contracts []model.Contract) ([]model.RbDTO, error) {
	var rbDTOsl []model.RbDTO
	for _, contract := range contracts {
		fmt.Println("contract MESSAGE", contract.Discounts)
		for _, discount := range contract.Discounts {
			if discount.Code == RB11Code {
				for _, period := range discount.Periods {
					reqBrand := model.ReqBrand{
						ClientCode:     req.ClientCode,
						DateStart:      period.PeriodFrom,
						DateEnd:        period.PeriodTo,
						TypeValue:      "",
						TypeParameters: nil,
						SchemeType:     contract.View,
					}
					periodFrom, _ := ConvertStringTime(period.PeriodFrom)
					periodTo, _ := ConvertStringTime(period.PeriodTo)
					reqperiodFrom, _ := ConvertStringTime(req.PeriodFrom)
					reqperiodTo, _ := ConvertStringTime(req.PeriodTo)
					if reqperiodFrom.Before(periodFrom) || reqperiodFrom.Equal(periodFrom) && reqperiodTo.After(periodTo) || reqperiodTo.Equal(periodTo) {
						purchase, _ := GetPurchase(reqBrand)
						//totalPurchaseCode := CountPurchaseByCode(purchase)
						totalPurchaseAmount := GetTotalAmountPurchase(purchase)
						discountAmount := totalPurchaseAmount * period.DiscountPercent / 100

						rbDTOsl = append(rbDTOsl, model.RbDTO{
							ID:                   contract.ID,
							ContractNumber:       contract.ContractParameters.ContractNumber,
							StartDate:            period.PeriodFrom,
							EndDate:              period.PeriodTo,
							TypePeriod:           period.Name,
							DiscountPercent:      period.DiscountPercent,
							DiscountAmount:       discountAmount,
							TotalWithoutDiscount: totalPurchaseAmount,
							LeasePlan:            period.PurchaseAmount,
							DiscountType:         RB11Name,
						})
					}

				}

			}

		}
	}

	return rbDTOsl, nil
}
func GetTotalAmountPurchase(purchase model.Purchase) float64 {
	var amount float64
	for _, s := range purchase.PurchaseArr {
		amount += s.Total
	}

	return amount
}

func GetRB12thType(req model.RBRequest, contracts []model.Contract) ([]model.RbDTO, error) {
	var rbDTOsl []model.RbDTO
	for _, contract := range contracts {

		fmt.Println(contract.ContractParameters.ContractNumber, "номер договора")
		fmt.Println(contract.Discounts, "скидки")
		for _, discount := range contract.Discounts {
			fmt.Println(contract.ContractParameters.ContractNumber, "номер договора")
			if discount.Code == RB12Code && discount.IsSelected == true { // здесь сравниваешь тип скидки и берешь тот тип который тебе нужен
				for _, period := range discount.Periods {
					reqBrand := model.ReqBrand{
						ClientCode:     req.ClientCode,
						DateStart:      period.PeriodFrom,
						DateEnd:        period.PeriodTo,
						TypeValue:      "",
						TypeParameters: nil,
						SchemeType:     contract.View,
					}
					purchase, err := GetPurchase(reqBrand)
					if err != nil {
						return nil, err
					}
					amount := CountPurchase(purchase)

					periodFrom, _ := ConvertStringTime(period.PeriodFrom)
					periodTo, _ := ConvertStringTime(period.PeriodTo)
					reqperiodFrom, _ := ConvertStringTime(req.PeriodFrom)
					reqperiodTo, _ := ConvertStringTime(req.PeriodTo)
					if reqperiodFrom.Before(periodFrom) || reqperiodFrom.Equal(periodFrom) && reqperiodTo.After(periodTo) || reqperiodTo.Equal(periodTo) {
						if amount >= period.PurchaseAmount {
							total := amount * period.DiscountPercent / 100
							RbDTO := model.RbDTO{
								ContractNumber:       contract.ContractParameters.ContractNumber,
								StartDate:            period.PeriodFrom,
								EndDate:              period.PeriodTo,
								TypePeriod:           period.Name,
								DiscountPercent:      period.DiscountPercent,
								DiscountAmount:       total,
								TotalWithoutDiscount: amount,
								LeasePlan:            period.PurchaseAmount,
								DiscountType:         RB12Name,
							}
							rbDTOsl = append(rbDTOsl, RbDTO)

						} else {
							rbDTOsl, _ = GetNil12Rb(rbDTOsl, contract, period, RB12Name)
						}

					}

				}

			}

		}

	}

	return rbDTOsl, nil
}
func ConvertTime(date string) (string, error) {
	timeSplit := strings.Split(date, ".")
	if len(timeSplit) != 3 {
		return "", errors.New("len of time must be 3")
	}
	fmt.Println(timeSplit)
	convertYear, err := strconv.Atoi(timeSplit[2])
	if err != nil {
		log.Println(err)
		return "", err
	}
	convertYear -= 1
	updateTime := fmt.Sprintf("%s.%s.%d", timeSplit[0], timeSplit[1], convertYear)
	//fmt.Println(sprintf)

	return updateTime, nil
}

func GetRB13thType(rb model.RBRequest, contracts []model.Contract) ([]model.RbDTO, error) {
	var rbDTOsl []model.RbDTO

	// чтобы преобразоват дату в ввиде День.Месяц.Год
	layoutISO := "02.1.2006"

	// parsing string to Time
	reqPeriodFrom, _ := time.Parse(layoutISO, rb.PeriodFrom)
	reqPeriodTo, _ := time.Parse(layoutISO, rb.PeriodTo)
	fmt.Println(reqPeriodFrom)
	fmt.Println(reqPeriodTo)

	for _, contract := range contracts {
		fmt.Println("contract MESSAGE", contract.Discounts)

		// от сюда берем скидки и периоды
		for _, discount := range contract.Discounts {
			// после всех проверок логика начнется
			if discount.Code == RB13Code && discount.IsSelected == true {

				for _, period := range discount.Periods {
					periodFrom, _ := time.Parse(layoutISO, period.PeriodFrom)
					periodTo, _ := time.Parse(layoutISO, period.PeriodTo)
					fmt.Println(periodFrom)
					fmt.Println(periodTo)

					//if periodFrom.After(reqPeriodFrom) || periodTo.Before(reqPeriodTo) {
					pastTimeFrom, err := ConvertTime(period.PeriodFrom)
					if err != nil {
						return nil, err
					}
					pastTimeTo, err := ConvertTime(period.PeriodTo)
					if err != nil {
						return nil, err
					}

					// это чтобы брали на 1 год меньше
					pastPeriod := model.ReqBrand{
						ClientCode:     rb.ClientCode,
						DateStart:      pastTimeFrom,
						DateEnd:        pastTimeTo,
						Type:           "",
						TypeValue:      "",
						TypeParameters: nil,
						Contracts:      nil,
						SchemeType:     contract.View,
					}

					// Это необходимо, чтобы получить продажи за тек период
					present := model.ReqBrand{
						ClientCode:     rb.ClientCode,
						Beneficiary:    "",
						DateStart:      rb.PeriodFrom,
						DateEnd:        rb.PeriodTo,
						Type:           "",
						TypeValue:      "",
						TypeParameters: nil,
						Contracts:      nil,
						SchemeType:     contract.View,
					}
					// берем продажи за тек год и за 1 год меньше
					presentPeriod, err := GetSales1C(present, "sales_brand_only")
					fmt.Println("PRESENT", presentPeriod)
					if err != nil {
						return nil, err
					}

					oldPeriod, err := GetSales1C(pastPeriod, "sales_brand_only")

					if presentPeriod.SalesArr == nil || oldPeriod.SalesArr == nil {
						rbDTOsl, _ = CheckPeriodNullGrowth(contract, period, RB13Name)
						//rbDTO := model.RbDTO{
						//	ContractNumber:       contract.ContractParameters.ContractNumber,
						//	StartDate:            period.PeriodFrom,
						//	EndDate:              period.PeriodTo,
						//	TypePeriod:           "",
						//	BrandName:            "",
						//	ProductCode:          "",
						//	DiscountPercent:      period.DiscountPercent,
						//	DiscountAmount:       0,
						//	TotalWithoutDiscount: 0,
						//	DiscountType:         RB13Name,
						//}
						//rbDTOsl = append(rbDTOsl, rbDTO)
						return rbDTOsl, nil

					}

					fmt.Println("PAST ==================", oldPeriod)
					if err != nil {
						return nil, err
					}
					var preCount float64
					var pastCount float64
					// считаем за тек период
					for _, present := range presentPeriod.SalesArr {
						preCount += present.Total
					}

					// считаем за прошлый год
					for _, past := range oldPeriod.SalesArr {
						pastCount += past.Total

					}

					innerPreCount := preCount
					innerPastCount := pastCount
					fmt.Println("Сумма за настоящее", preCount)
					fmt.Println("Сумма за прошлый год", pastCount)

					if contract.View == "PF" {
						innerPreCount = TotalRegionsSales(presentPeriod, contract.Regions)
						innerPastCount = TotalRegionsSales(oldPeriod, contract.Regions)
					}

					// находим прирост в процентах
					//growthPercent := (pastCount * 100 / preCount) - 100

					// находим разницу за нынешний год
					diff := innerPreCount - innerPastCount
					growthPercent := (100 * diff) / innerPastCount
					fmt.Println("growthPercent", growthPercent)
					// проверяем разницу с тек по прошлогодний год, если процент прироста выше, логика выполнится

					fmt.Println("growth_percent", period.GrowthPercent)
					fmt.Println("discount percent", period.DiscountPercent)
					fmt.Println()
					if growthPercent > period.GrowthPercent {
						discountAmount := innerPreCount * period.DiscountPercent / 100

						fmt.Println("discountAmount", discountAmount)

						rbDTO := model.RbDTO{
							ContractNumber:       contract.ContractParameters.ContractNumber,
							StartDate:            period.PeriodFrom,
							EndDate:              period.PeriodTo,
							TypePeriod:           "",
							BrandName:            "",
							ProductCode:          "",
							DiscountPercent:      period.DiscountPercent,
							DiscountAmount:       discountAmount,
							TotalWithoutDiscount: innerPreCount,
							DiscountType:         RB13Name,
						}
						rbDTOsl = append(rbDTOsl, rbDTO)

					} else {
						rbDTOsl, _ = CheckPeriodNullGrowth(contract, period, RB13Name)

						//rbDTO := model.RbDTO{
						//	ContractNumber:       contract.ContractParameters.ContractNumber,
						//	StartDate:            period.PeriodFrom,
						//	EndDate:              period.PeriodTo,
						//	TypePeriod:           "",
						//	BrandName:            "",
						//	ProductCode:          "",
						//	DiscountPercent:      period.DiscountPercent,
						//	DiscountAmount:       0,
						//	TotalWithoutDiscount: preCount,
						//	DiscountType:         RB13Name,
						//}
						//rbDTOsl = append(rbDTOsl, rbDTO)
					}

				}

			}

		}
	}
	//}
	return rbDTOsl, nil
}

func CheckPeriodNullGrowth(contract model.Contract, period model.DiscountPeriod, discountType string) ([]model.RbDTO, error) {
	var rbDTOsl []model.RbDTO

	rbDTO := model.RbDTO{
		ContractNumber:       contract.ContractParameters.ContractNumber,
		StartDate:            period.PeriodFrom,
		EndDate:              period.PeriodTo,
		TypePeriod:           period.Name,
		DiscountPercent:      period.DiscountPercent,
		DiscountAmount:       0,
		TotalWithoutDiscount: 0,
		LeasePlan:            period.PurchaseAmount,
		DiscountType:         discountType,
	}
	rbDTOsl = append(rbDTOsl, rbDTO)
	return rbDTOsl, nil

}

func GetNil12Rb(rbDTOsl []model.RbDTO, contract model.Contract, period model.DiscountPeriod, discountType string) ([]model.RbDTO, error) {
	//var

	rbDTO := model.RbDTO{
		ContractNumber:       contract.ContractParameters.ContractNumber,
		StartDate:            period.PeriodFrom,
		EndDate:              period.PeriodTo,
		TypePeriod:           period.Name,
		DiscountPercent:      period.DiscountPercent,
		DiscountAmount:       0,
		TotalWithoutDiscount: 0,
		LeasePlan:            period.PurchaseAmount,
		DiscountType:         discountType,
	}
	rbDTOsl = append(rbDTOsl, rbDTO)
	return rbDTOsl, nil

}

// CountPurchaseByCode считываем итог по каждому контракт коду
func CountPurchaseByCode(purchase model.Purchase) map[string]float64 {
	totallyCode := map[string]float64{}
	for _, value := range purchase.PurchaseArr {
		if _, ok := totallyCode[value.ContractCode]; !ok {
			totallyCode[value.ContractCode] += value.Total
			//do something here
		}

	}
	return totallyCode

}

func CountSales(sales model.Sales) float64 {
	var amount float64
	for _, value := range sales.SalesArr {
		amount += value.Total
	}
	return amount
}

func CountSalesNew(sales model.SalesNew) float64 {
	var amount float64
	for _, value := range sales.SalesArrNEW {
		amount += value.Total
	}
	return amount
}

func CountPurchase(purchase model.Purchase) float64 {
	var amount float64
	for _, value := range purchase.PurchaseArr {
		amount += value.Total
	}
	return amount
}

// CountSalesByBrand считываем итог по каждому контракт коду
func CountSalesByBrand(sales model.Sales) map[string]float64 {
	totallyCode := map[string]float64{}
	for _, value := range sales.SalesArr {
		//if _, ok := totallyCode[value.ContractCode]; !ok {
		totallyCode[value.BrandName] += value.Total
		//do something here
		//}

	}
	return totallyCode
}

func TotalRegionsSales(sales model.Sales, regions []model.Regions) (totalAmount float64) {
	for _, region := range regions {
		for _, sale := range sales.SalesArr {
			if region.RegionCode == sale.RegionCode {
				totalAmount += float64(sale.Total)
			}
		}
	}

	return totalAmount
}

//// JoinContractCode собираем все контракт коды в слайс
//func JoinContractCode(externalCodes []model.ContractCode) []string {
//	var contractsCode []string
//	for _, value := range externalCodes {
//		fmt.Println("value.ExtContractCode===========================================================================", value.ExtContractCode)
//
//		if value.ExtContractCode == "" {
//			continue
//		}
//		contractsCode = append(contractsCode, value.ExtContractCode)
//	}
//	return contractsCode
//}
