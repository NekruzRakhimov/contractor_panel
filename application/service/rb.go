package service

import (
	"context"
	"contractor_panel/domain/model"
	"contractor_panel/domain/repository"
	"contractor_panel/utils"
	"fmt"
	"github.com/xuri/excelize/v2"
	"log"
)

type ReportTemplateService interface {
	GetAllRBByContractorBIN(ctx context.Context, userID int64, request model.RBRequest) (rbDTOs []model.RbDTO, err error)
	// FormExcelForRBReport GetAllContractDetailByBIN(ctx context.Context, clientCode, PeriodFrom, PeriodTo string) (contracts []model.ContractWithJsonB, err error) {
	FormExcelForRBReport(ctx context.Context, userID int64, request model.RBRequest) error
}

type reportTemplateService struct {
	r repository.ReportTemplateRepository
}

func NewReportTemplateService(r repository.ReportTemplateRepository) *reportTemplateService {
	return &reportTemplateService{r: r}
}

func (s *reportTemplateService) GetAllRBByContractorBIN(ctx context.Context, userID int64, request model.RBRequest) (rbDTOs []model.RbDTO, err error) {

	contractsWithJson, err := s.r.GetAllContractDetailByBIN(ctx, userID, request)

	contracts, err := BulkConvertContractFromJsonB(contractsWithJson)
	for i, contract := range contracts {
		if contract.AdditionalAgreementNumber != 0 {
			var contractType string
			//ДС №1 к Договору маркетинговых услуг №1111 ИП  “Adal Trade“
			//marketing_services
			//supply
			switch contract.Type {
			case "marketing_services":
				contractType = "маркетинговых услуг"
			case "supply":
				contractType = "поставок"
			}

			contracts[i].ContractParameters.ContractNumber = fmt.Sprintf("ДС №%d к Договору %s №%s %s",
				contract.AdditionalAgreementNumber, contractType,
				contract.ContractParameters.ContractNumber,
				contract.Requisites.Beneficiary)
		}
	}
	// #1
	RB1stType, err := GetRB1stType(request, contracts)
	if err != nil {
		return
	}
	rbDTOs = append(rbDTOs, RB1stType...)

	// #2
	rb2ndType := GetRB2ndType(request, contracts)
	if err != nil {
		return
	}
	rbDTOs = append(rbDTOs, rb2ndType...)

	// #3
	rb3rdType, err := GetRB3rdType(request, contracts)
	if err != nil {
		return
	}
	rbDTOs = append(rbDTOs, rb3rdType...)

	// #4
	rbFourthType, err := GetRB4thType(request, contracts)
	if err != nil {
		return
	}
	rbDTOs = append(rbDTOs, rbFourthType...)

	// #5
	rb5thType, err := GetRB5thType(request, contracts)
	if err != nil {
		return
	}
	rbDTOs = append(rbDTOs, rb5thType...)

	// #6
	rb6thType, err := GetRB6thType(request, contracts)
	if err != nil {
		return
	}
	rbDTOs = append(rbDTOs, rb6thType...)

	// #7
	rb7thType, err := GetRB7thType(request, contracts)
	if err != nil {
		return
	}
	rbDTOs = append(rbDTOs, rb7thType...)

	// #8
	rb8thType, err := GetRB8thType(request, contracts)
	if err != nil {
		return
	}
	rbDTOs = append(rbDTOs, rb8thType...)

	// #9
	rb9thType, err := GetRB9thType(request, contracts)
	if err != nil {
		return
	}
	rbDTOs = append(rbDTOs, rb9thType...)

	// #10
	rbTenthType, err := GetRb10thType(request, contracts)
	if err != nil {
		return
	}
	rbDTOs = append(rbDTOs, rbTenthType...)

	// #11
	rb11thType, err := GetRB11thType(request, contracts)
	if err != nil {
		return
	}
	rbDTOs = append(rbDTOs, rb11thType...)

	// #12
	rb12thType, err := GetRB12thType(request, contracts)
	if err != nil {
		return
	}
	rbDTOs = append(rbDTOs, rb12thType...)

	// #13
	rb13thType, err := GetRB13thType(request, contracts)
	if err != nil {
		return
	}
	rbDTOs = append(rbDTOs, rb13thType...)

	// #14
	rb14ThType, err := GetRB14ThType(request, contracts)
	if err != nil {
		return
	}
	rbDTOs = append(rbDTOs, rb14ThType...)

	// #15
	rb15ThType, err := GetRB15ThType(request, contracts)
	if err != nil {
		return
	}
	rbDTOs = append(rbDTOs, rb15ThType...)

	// #16
	rb16ThType, err := GetRB16ThType(request, contracts)
	if err != nil {
		return
	}
	rbDTOs = append(rbDTOs, rb16ThType...)

	// #17
	rb17ThType, err := GetRB17ThType(request, contracts)
	if err != nil {
		return
	}
	rbDTOs = append(rbDTOs, rb17ThType...)

	// #18
	rb18ThType, err := GetRB18thType(request, contracts)
	if err != nil {
		return
	}
	rbDTOs = append(rbDTOs, rb18ThType...)

	return

}

func BulkConvertContractFromJsonB(contractsWithJson []model.ContractWithJsonB) (contracts []model.Contract, err error) {
	for i := range contractsWithJson {
		contract, err := ConvertContractFromJsonB(contractsWithJson[i])
		if err != nil {
			log.Println("Error: service.BulkConvertContractFromJsonB. Error is: ", err.Error())
			continue
		}
		contracts = append(contracts, contract)
	}

	return
}

func (s *reportTemplateService) FormExcelForRBReport(ctx context.Context, userID int64, request model.RBRequest) error {
	contractsWithJson, err := s.r.GetAllContractDetailByBIN(ctx, userID, request)
	if err != nil {
		fmt.Println(">> 1")
		return err
	}

	contracts, err := BulkConvertContractFromJsonB(contractsWithJson)
	if err != nil {
		fmt.Println(">> 2")
		return err
	}

	for i, contract := range contracts {
		if contract.AdditionalAgreementNumber != 0 {
			var contractType string
			//ДС №1 к Договору маркетинговых услуг №1111 ИП  “Adal Trade“
			//marketing_services
			//supply
			switch contract.Type {
			case "marketing_services":
				contractType = "маркетинговых услуг"
			case "supply":
				contractType = "поставок"
			}

			contracts[i].ContractParameters.ContractNumber = fmt.Sprintf("ДС №%d к Договору %s №%s %s",
				contract.AdditionalAgreementNumber, contractType,
				contract.ContractParameters.ContractNumber,
				contract.Requisites.Beneficiary)
		}
	}

	var (
		isRB1  bool
		isRB2  bool
		isRB3  bool
		isRB4  bool
		isRB5  bool
		isRB6  bool
		isRB7  bool
		isRB8  bool
		isRB9  bool
		isRB10 bool
		isRB11 bool
		isRB12 bool
		isRB13 bool
		isRB14 bool
		isRB15 bool
		isRB16 bool
		isRB17 bool
		isRB18 bool
	)

	for _, contract := range contracts {
		for _, discount := range contract.Discounts {
			if discount.Code == RB1Code && discount.IsSelected {
				isRB1 = true
			}
			if discount.Code == RB2Code && discount.IsSelected {
				isRB2 = true
			}
			if discount.Code == RB3Code && discount.IsSelected {
				isRB3 = true
			}
			if discount.Code == RB4Code && discount.IsSelected {
				isRB4 = true
			}
			if discount.Code == RB5Code && discount.IsSelected {
				isRB5 = true
			}
			if discount.Code == RB6Code && discount.IsSelected {
				isRB6 = true
			}
			if discount.Code == RB7Code && discount.IsSelected {
				isRB7 = true
			}
			if discount.Code == RB8Code && discount.IsSelected {
				isRB8 = true
			}
			if discount.Code == RB9Code && discount.IsSelected {
				isRB9 = true
			}
			if discount.Code == RB10Code && discount.IsSelected {
				isRB10 = true
			}
			if discount.Code == RB11Code && discount.IsSelected {
				isRB11 = true
			}
			if discount.Code == RB12Code && discount.IsSelected {
				isRB12 = true
			}
			if discount.Code == RB13Code && discount.IsSelected {
				isRB13 = true
			}
			if discount.Code == RB14Code && discount.IsSelected {
				isRB14 = true
			}
			if discount.Code == RB15Code && discount.IsSelected {
				isRB15 = true
			}
			if discount.Code == RB16Code && discount.IsSelected {
				isRB16 = true
			}
			if discount.Code == RB17Code && discount.IsSelected {
				isRB17 = true
			}
			if discount.Code == RB18Code && discount.IsSelected {
				isRB18 = true
			}
		}
	}

	//totalAmount := GetPurchaseTotalAmount(sales)

	fmt.Println(contracts)
	//fmt.Println(totalAmount)
	//var conTotalAmount float32
	//var rewardAmount int

	f := excelize.NewFile()

	//var discount int
	//if conTotalAmount <= totalAmount {
	//	discount = rewardAmount
	//}
	moneyStyle, _ := f.NewStyle(`{"number_format": 4}`)

	coloredMoneyStyle, _ := f.NewStyle(&excelize.Style{
		Fill: excelize.Fill{Type: "pattern", Color: []string{"#F5DEB3"}, Pattern: 1},
		Border: []excelize.Border{
			{Type: "left", Color: "000000", Style: 1},
			{Type: "right", Color: "000000", Style: 1},
			{Type: "top", Color: "000000", Style: 1},
			{Type: "bottom", Color: "000000", Style: 1},
		},
		NumFmt: 4,
	})

	style, err := f.NewStyle(&excelize.Style{
		Fill: excelize.Fill{Type: "pattern", Color: []string{"#F5DEB3"}, Pattern: 1},
		Border: []excelize.Border{
			{Type: "left", Color: "000000", Style: 1},
			{Type: "right", Color: "000000", Style: 1},
			{Type: "top", Color: "000000", Style: 1},
			{Type: "bottom", Color: "000000", Style: 1},
		},
	})
	if err != nil {
		fmt.Println(err)
	}

	f.NewSheet(sheet)
	ineration := 1
	var TotalAmount float64
	f.SetCellValue(sheet, "A1", "Вид скидки:")
	f.SetColWidth(sheet, "A", "A", 42)
	f.SetCellValue(sheet, "B1", "Сумма скидки:")
	f.SetColWidth(sheet, "B", "B", 17)

	//f.SetCellValue(sheet, "C1", "Итог:")

	//f.SetCellValue(sheet, "A1", "Бренд")
	//f.SetCellValue(sheet, "B1", "Номер бренда")
	//f.SetCellValue(sheet, "C1", "Период")
	//f.SetCellValue(sheet, "D1", "Стоимость")
	//f.SetCellValue(sheet, "E1", "Количество")
	//f.SetCellValue(sheet, "F1", "Итог:")
	//f.SetCellValue(sheet, "G1", "Сумма скидки:")
	//f.SetCellValue(sheet, "H1", "Вид скидки:")

	//fmt.Printf(">>arr>>%+v", sales.SalesArr)

	var lastRow int

	period := fmt.Sprintf("%s-%s", request.PeriodFrom, request.PeriodTo)
	fmt.Println("<request>: ", period)

	//for i, s := range purchases.PurchaseArr {
	//	f.SetCellValue(sheet, fmt.Sprintf("%s%d", "A", i+2), s.BrandName)
	//	f.SetCellValue(sheet, fmt.Sprintf("%s%d", "B", i+2), s.BrandCode)
	//	f.SetCellValue(sheet, fmt.Sprintf("%s%d", "С", i+2), period)
	//	f.SetCellValue(sheet, fmt.Sprintf("%s%d", "D", i+2), utils.FloatToMoneyFormat(s.Total/s.QntTotal))
	//	f.SetCellValue(sheet, fmt.Sprintf("%s%d", "E", i+2), utils.FloatToMoneyFormat(s.QntTotal))
	//	f.SetCellValue(sheet, fmt.Sprintf("%s%d", "F", i+2), utils.FloatToMoneyFormat(s.Total))
	//	lastRow = i
	//}

	lastRow += 3

	//f.SetCellValue(sheet, fmt.Sprintf("%s%d", "E", lastRow), "Итог:")
	//f.SetCellValue(sheet, fmt.Sprintf("%s%d", "F", lastRow), discount)
	//f.SetCellValue(sheet, fmt.Sprintf("%s%d", "F", lastRow), utils.FloatToMoneyFormat(totalAmount))
	//_ = f.MergeCell(sheet, fmt.Sprintf("%s%d", "A", lastRow), fmt.Sprintf("%s%d", "B", lastRow))
	//err = f.SetCellStyle(sheet, fmt.Sprintf("%s%d", "A", lastRow), fmt.Sprintf("%s%d", "F", lastRow), style)
	//err = f.SetCellStyle(sheet, fmt.Sprintf("%s%d", "A", 1), fmt.Sprintf("%s%d", "F", 1), style)
	err = f.SetCellStyle(sheet, "A1", "B1", style)
	//f.SetColWidth(sheet, "A", "B", 20)
	//f.SetCellValue("Sheet1", "D102", discount)
	//RB1

	if isRB1 {
		contract1stType, err := GetRB1stType(request, contracts)
		if err != nil {
			return err
		}

		f.NewSheet(RB1Name)
		f.SetCellValue(RB1Name, "A1", "Период")
		f.SetCellValue(RB1Name, "B1", "Номер договора/ДС")
		f.SetCellValue(RB1Name, "C1", "Тип скидки")
		f.SetCellValue(RB1Name, "D1", "Сумма вознаграждения")
		f.SetCellValue(RB1Name, "E1", "Сумма скидки")
		//f.SetCellValue(RB1Name, "F1", "Общая сумма")
		err = f.SetCellStyle(RB1Name, "A1", "E1", style)

		var totalDiscountsSum float64
		fmt.Printf("CHECK \n%+v\n CHECK", contracts)
		var i int // вроде бы занова создается
		for _, contract := range contract1stType {
			f.SetCellValue(RB1Name, fmt.Sprintf("%s%d", "A", i+2), fmt.Sprintf("%s-%s", contract.StartDate, contract.EndDate))
			f.SetCellValue(RB1Name, fmt.Sprintf("%s%d", "B", i+2), contract.ContractNumber)
			f.SetCellValue(RB1Name, fmt.Sprintf("%s%d", "C", i+2), "Скидка за объем закупа")
			f.SetCellValue(RB1Name, fmt.Sprintf("%s%d", "D", i+2), utils.FloatToMoneyFormat(float64(contract.RewardAmount)))
			f.SetCellValue(RB1Name, fmt.Sprintf("%s%d", "E", i+2), contract.DiscountAmount)
			//f.SetCellValue(RB1Name, fmt.Sprintf("%s%d", "F", i+2), totalDiscountsSum)
			totalDiscountsSum += contract.DiscountAmount
			lastRow = i + 2
			i++
		}
		// вот что странно...
		lastRow += 1
		f.SetCellValue(RB1Name, fmt.Sprintf("%s%d", "D", lastRow), "Итог:")
		f.SetCellValue(RB1Name, fmt.Sprintf("%s%d", "E", lastRow), totalDiscountsSum)
		err = f.SetCellStyle(RB1Name, fmt.Sprintf("%s%d", "D", lastRow), fmt.Sprintf("%s%d", "D", lastRow), style)
		err = f.SetCellStyle(RB1Name, fmt.Sprintf("%s%d", "E", lastRow), fmt.Sprintf("%s%d", "D", lastRow), style)
	}

	if isRB2 {
		rbSecondType := GetRB2ndType(request, contracts)
		if err != nil {
			return err
		}

		f.NewSheet(RB2Name)
		f.SetCellValue(RB2Name, "A1", "Период")
		f.SetCellValue(RB2Name, "B1", "Номер договора/ДС")
		f.SetCellValue(RB2Name, "C1", "Тип скидки")
		f.SetCellValue(RB2Name, "D1", "Бренд")
		f.SetCellValue(RB2Name, "E1", "Скидка %")
		f.SetCellValue(RB2Name, "F1", "Сумма скидки")
		f.SetCellValue(RB2Name, "G1", "Общая сумма")
		f.SetCellValue(RB2Name, "H1", "Тип")
		f.SetColWidth(RB2Name, "A", "H", 22)
		f.SetColWidth(RB2Name, "C", "C", 40)
		err = f.SetCellStyle(RB2Name, "A1", "H1", style)

		var totalDiscountsSum float64
		var i int
		for _, contract := range rbSecondType {
			f.SetCellValue(RB2Name, fmt.Sprintf("%s%d", "A", i+2), fmt.Sprintf("%s-%s", contract.StartDate, contract.EndDate))
			f.SetCellValue(RB2Name, fmt.Sprintf("%s%d", "B", i+2), contract.ContractNumber)
			f.SetCellValue(RB2Name, fmt.Sprintf("%s%d", "C", i+2), RB2Name)
			f.SetCellValue(RB2Name, fmt.Sprintf("%s%d", "D", i+2), contract.BrandName)
			f.SetCellValue(RB2Name, fmt.Sprintf("%s%d", "E", i+2), contract.DiscountPercent)
			f.SetCellValue(RB2Name, fmt.Sprintf("%s%d", "F", i+2), contract.DiscountAmount)
			f.SetCellValue(RB2Name, fmt.Sprintf("%s%d", "G", i+2), contract.TotalWithoutDiscount)
			f.SetCellValue(RB2Name, fmt.Sprintf("%s%d", "H", i+2), "по продажам")
			totalDiscountsSum += contract.DiscountAmount

			lastRow = i + 2
			f.SetCellStyle(RB2Name, fmt.Sprintf("%s%d", "F", i+2), fmt.Sprintf("%s%d", "G", i+2), moneyStyle)
			i++
		}
		lastRow += 1

		// для вкладки итогов
		TotalAmount += totalDiscountsSum

		f.SetCellValue(RB2Name, fmt.Sprintf("%s%d", "E", lastRow), "Итог:")
		f.SetCellValue(RB2Name, fmt.Sprintf("%s%d", "F", lastRow), totalDiscountsSum)
		err = f.SetCellStyle(RB2Name, fmt.Sprintf("%s%d", "E", lastRow), fmt.Sprintf("%s%d", "F", lastRow), coloredMoneyStyle)
		//err = f.SetCellStyle(RB2Name, fmt.Sprintf("%s%d", "D", lastRow), fmt.Sprintf("%s%d", "D", lastRow), style)
		//err = f.SetCellStyle(RB2Name, fmt.Sprintf("%s%d", "F", lastRow), fmt.Sprintf("%s%d", "G", lastRow), moneyStyle)
		ineration++

		f.SetCellValue(sheet, fmt.Sprintf("%s%d", "A", ineration), RB2Name)
		f.SetCellValue(sheet, fmt.Sprintf("%s%d", "B", ineration), totalDiscountsSum)
		f.SetCellStyle(sheet, fmt.Sprintf("%s%d", "B", ineration), fmt.Sprintf("%s%d", "B", ineration), moneyStyle)
	}

	if isRB3 {
		rbThirdType, err := GetRB3rdType(request, contracts)
		if err != nil {
			return err
		}

		f.NewSheet(RB3Name)
		f.SetCellValue(RB3Name, "A1", "Период")
		f.SetCellValue(RB3Name, "B1", "Номер договора/ДС")
		f.SetCellValue(RB3Name, "C1", "Тип скидки")
		f.SetCellValue(RB3Name, "D1", "Код товара")
		f.SetCellValue(RB3Name, "E1", "План закупа")
		f.SetCellValue(RB3Name, "F1", "Скидка %")
		f.SetCellValue(RB3Name, "G1", "Сумма скидки")
		f.SetCellValue(RB3Name, "H1", "Общая сумма")
		f.SetCellValue(RB3Name, "I1", "Тип")
		f.SetColWidth(RB3Name, "A", "I", 22)
		f.SetColWidth(RB3Name, "C", "C", 40)
		err = f.SetCellStyle(RB3Name, "A1", "I1", style)

		var totalDiscountsSum float64
		fmt.Printf("CHECK \n%+v\n CHECK", contracts)
		var i int
		for _, contract := range rbThirdType {
			f.SetCellValue(RB3Name, fmt.Sprintf("%s%d", "A", i+2), fmt.Sprintf("%s-%s", contract.StartDate, contract.EndDate))
			f.SetCellValue(RB3Name, fmt.Sprintf("%s%d", "B", i+2), contract.ContractNumber)
			f.SetCellValue(RB3Name, fmt.Sprintf("%s%d", "C", i+2), contract.DiscountType)
			f.SetCellValue(RB3Name, fmt.Sprintf("%s%d", "D", i+2), contract.ProductCode)
			f.SetCellValue(RB3Name, fmt.Sprintf("%s%d", "E", i+2), contract.LeasePlan)
			f.SetCellValue(RB3Name, fmt.Sprintf("%s%d", "F", i+2), contract.DiscountPercent)
			f.SetCellValue(RB3Name, fmt.Sprintf("%s%d", "G", i+2), contract.DiscountAmount)
			f.SetCellValue(RB3Name, fmt.Sprintf("%s%d", "H", i+2), contract.TotalWithoutDiscount)
			f.SetCellValue(RB3Name, fmt.Sprintf("%s%d", "I", i+2), "по закупам")
			totalDiscountsSum += contract.DiscountAmount

			lastRow = i + 2
			f.SetCellStyle(RB3Name, fmt.Sprintf("%s%d", "G", i+2), fmt.Sprintf("%s%d", "H", i+2), moneyStyle)
			i++
		}
		lastRow += 1

		//TODO: - вкладка итогов
		TotalAmount += totalDiscountsSum
		f.SetCellValue(RB3Name, fmt.Sprintf("%s%d", "F", lastRow), "Итог:")
		f.SetCellValue(RB3Name, fmt.Sprintf("%s%d", "G", lastRow), totalDiscountsSum)
		err = f.SetCellStyle(RB3Name, fmt.Sprintf("%s%d", "F", lastRow), fmt.Sprintf("%s%d", "G", lastRow), coloredMoneyStyle)
		//err = f.SetCellStyle(RB3Name, fmt.Sprintf("%s%d", "F", lastRow), fmt.Sprintf("%s%d", "G", lastRow), style)
		//err = f.SetCellStyle(RB3Name, fmt.Sprintf("%s%d", "F", lastRow), fmt.Sprintf("%s%d", "H", lastRow), moneyStyle)
		//err = f.SetCellStyle(RB2Name, fmt.Sprintf("%s%d", "G", lastRow), fmt.Sprintf("%s%d", "G", lastRow), style)

		ineration++
		//err = f.SetCellStyle(sheet, fmt.Sprintf("%s%d", "A", ineration), fmt.Sprintf("%s%d", "D", lastRow), style)
		f.SetCellValue(sheet, fmt.Sprintf("%s%d", "A", ineration), RB3Name)
		f.SetCellValue(sheet, fmt.Sprintf("%s%d", "B", ineration), totalDiscountsSum)
		f.SetCellStyle(sheet, fmt.Sprintf("%s%d", "B", ineration), fmt.Sprintf("%s%d", "B", ineration), moneyStyle)
	}

	if isRB4 {
		rbFourthType, err := GetRB4thType(request, contracts)
		if err != nil {
			return err
		}

		f.NewSheet(RB4Name)
		f.SetCellValue(RB4Name, "A1", "Период")
		f.SetCellValue(RB4Name, "B1", "Номер договора/ДС")
		f.SetCellValue(RB4Name, "C1", "Тип скидки")
		f.SetCellValue(RB4Name, "D1", "Скидка %")
		f.SetCellValue(RB4Name, "E1", "Сумма скидки")
		f.SetCellValue(RB4Name, "F1", "Общая сумма")
		f.SetCellValue(RB4Name, "G1", "Тип")
		f.SetColWidth(RB4Name, "A", "G", 22)
		f.SetColWidth(RB4Name, "C", "C", 40)
		f.SetCellStyle(RB4Name, "A1", "G1", style)

		var totalDiscountsSum float64
		fmt.Printf("CHECK \n%+v\n CHECK", contracts)
		var i int
		for _, contract := range rbFourthType {
			f.SetCellValue(RB4Name, fmt.Sprintf("%s%d", "A", i+2), fmt.Sprintf("%s-%s", contract.StartDate, contract.EndDate))
			f.SetCellValue(RB4Name, fmt.Sprintf("%s%d", "B", i+2), contract.ContractNumber)
			f.SetCellValue(RB4Name, fmt.Sprintf("%s%d", "C", i+2), RB4Name)
			f.SetCellValue(RB4Name, fmt.Sprintf("%s%d", "D", i+2), contract.DiscountPercent)
			f.SetCellValue(RB4Name, fmt.Sprintf("%s%d", "E", i+2), contract.DiscountAmount)
			f.SetCellValue(RB4Name, fmt.Sprintf("%s%d", "F", i+2), contract.TotalWithoutDiscount)
			f.SetCellValue(RB4Name, fmt.Sprintf("%s%d", "G", i+2), "по закупам")
			totalDiscountsSum += contract.DiscountAmount

			lastRow = i + 2
			f.SetCellStyle(RB4Name, fmt.Sprintf("%s%d", "E", i+2), fmt.Sprintf("%s%d", "F", i+2), moneyStyle)
			i++
		}
		lastRow += 1

		//TODO: - вкладка итогов
		TotalAmount += totalDiscountsSum
		f.SetCellValue(RB4Name, fmt.Sprintf("%s%d", "D", lastRow), "Итог:")
		f.SetCellValue(RB4Name, fmt.Sprintf("%s%d", "E", lastRow), totalDiscountsSum)
		err = f.SetCellStyle(RB4Name, fmt.Sprintf("%s%d", "D", lastRow), fmt.Sprintf("%s%d", "E", lastRow), coloredMoneyStyle)
		//err = f.SetCellStyle(RB4Name, fmt.Sprintf("%s%d", "D", lastRow), fmt.Sprintf("%s%d", "E", lastRow), style)
		//err = f.SetCellStyle(RB4Name, fmt.Sprintf("%s%d", "D", lastRow), fmt.Sprintf("%s%d", "F", lastRow), moneyStyle)

		ineration++
		//err = f.SetCellStyle(sheet, fmt.Sprintf("%s%d", "A", ineration), fmt.Sprintf("%s%d", "D", lastRow), style)
		f.SetCellValue(sheet, fmt.Sprintf("%s%d", "A", ineration), RB4Name)
		f.SetCellValue(sheet, fmt.Sprintf("%s%d", "B", ineration), totalDiscountsSum)
		f.SetCellStyle(sheet, fmt.Sprintf("%s%d", "B", ineration), fmt.Sprintf("%s%d", "B", ineration), moneyStyle)
	}

	if isRB5 {
		rb5thType, err := GetRB5thType(request, contracts)
		if err != nil {
			return err
		}

		if err != nil {
			return err
		}

		f.NewSheet(RB5Name)
		f.SetCellValue(RB5Name, "A1", "Период")
		f.SetCellValue(RB5Name, "B1", "Номер договора/ДС")
		f.SetCellValue(RB5Name, "C1", "Тип скидки")
		f.SetCellValue(RB5Name, "D1", "Бренд")
		f.SetCellValue(RB5Name, "E1", "Скидка %")
		f.SetCellValue(RB5Name, "F1", "Сумма скидки")
		f.SetCellValue(RB5Name, "G1", "Общая сумма")
		f.SetCellValue(RB5Name, "H1", "Тип")
		f.SetColWidth(RB5Name, "A", "H", 22)
		f.SetColWidth(RB5Name, "C", "C", 40)

		err = f.SetCellStyle(RB5Name, "A1", "H1", style)

		var totalDiscountsSum float64
		for i, contract := range rb5thType {
			f.SetCellValue(RB5Name, fmt.Sprintf("%s%d", "A", i+2), fmt.Sprintf("%s-%s", contract.StartDate, contract.EndDate))
			f.SetCellValue(RB5Name, fmt.Sprintf("%s%d", "B", i+2), contract.ContractNumber)
			f.SetCellValue(RB5Name, fmt.Sprintf("%s%d", "C", i+2), contract.DiscountType)
			f.SetCellValue(RB5Name, fmt.Sprintf("%s%d", "D", i+2), contract.BrandName)
			f.SetCellValue(RB5Name, fmt.Sprintf("%s%d", "E", i+2), contract.DiscountPercent)
			f.SetCellValue(RB5Name, fmt.Sprintf("%s%d", "F", i+2), contract.DiscountAmount)
			f.SetCellValue(RB5Name, fmt.Sprintf("%s%d", "G", i+2), contract.TotalWithoutDiscount)
			f.SetCellValue(RB5Name, fmt.Sprintf("%s%d", "H", i+2), "по закупам")
			totalDiscountsSum += contract.DiscountAmount

			lastRow = i + 2
			f.SetCellStyle(RB5Name, fmt.Sprintf("%s%d", "F", i+2), fmt.Sprintf("%s%d", "G", i+2), moneyStyle)
			i++
		}
		lastRow += 1

		//TODO: - вкладка итогов
		TotalAmount += totalDiscountsSum

		f.SetCellValue(RB5Name, fmt.Sprintf("%s%d", "E", lastRow), "Итог:")
		f.SetCellValue(RB5Name, fmt.Sprintf("%s%d", "F", lastRow), totalDiscountsSum)
		err = f.SetCellStyle(RB5Name, fmt.Sprintf("%s%d", "D", lastRow), fmt.Sprintf("%s%d", "E", lastRow), coloredMoneyStyle)
		//err = f.SetCellStyle(RB5Name, fmt.Sprintf("%s%d", "E", lastRow), fmt.Sprintf("%s%d", "F", lastRow), style)
		//err = f.SetCellStyle(RB5Name, fmt.Sprintf("%s%d", "E", lastRow), fmt.Sprintf("%s%d", "D", lastRow), style)
		//err = f.SetCellStyle(RB5Name, fmt.Sprintf("%s%d", "E", lastRow), fmt.Sprintf("%s%d", "G", lastRow), moneyStyle)

		ineration++
		//err = f.SetCellStyle(sheet, fmt.Sprintf("%s%d", "A", ineration), fmt.Sprintf("%s%d", "D", lastRow), style)
		f.SetCellValue(sheet, fmt.Sprintf("%s%d", "A", ineration), RB5Name)
		f.SetCellValue(sheet, fmt.Sprintf("%s%d", "B", ineration), totalDiscountsSum)
		f.SetCellStyle(sheet, fmt.Sprintf("%s%d", "B", ineration), fmt.Sprintf("%s%d", "B", ineration), moneyStyle)

	}

	if isRB6 {
		rb6thType, err := GetRB6thType(request, contracts)
		if err != nil {
			return err
		}

		if err != nil {
			return err
		}
		f.NewSheet(RB6Name)
		f.SetCellValue(RB6Name, "A1", "Период")
		f.SetCellValue(RB6Name, "B1", "Номер договора/ДС")

		f.SetCellValue(RB6Name, "C1", "Тип скидки")
		f.SetCellValue(RB6Name, "D1", "Бренд")
		f.SetCellValue(RB6Name, "E1", "Скидка %")
		f.SetCellValue(RB6Name, "F1", "Сумма скидки")
		f.SetCellValue(RB6Name, "G1", "Общая сумма")
		f.SetCellValue(RB6Name, "H1", "Тип")
		f.SetColWidth(RB6Name, "A", "H", 22)
		f.SetColWidth(RB6Name, "C", "C", 40)

		err = f.SetCellStyle(RB6Name, "A1", "H1", style)

		var totalDiscountsSum float64
		var i int
		for _, contract := range rb6thType {
			f.SetCellValue(RB6Name, fmt.Sprintf("%s%d", "A", i+2), fmt.Sprintf("%s-%s", contract.StartDate, contract.EndDate))
			f.SetCellValue(RB6Name, fmt.Sprintf("%s%d", "B", i+2), contract.ContractNumber)
			f.SetCellValue(RB6Name, fmt.Sprintf("%s%d", "C", i+2), RB6Name)
			f.SetCellValue(RB6Name, fmt.Sprintf("%s%d", "D", i+2), contract.BrandName)
			f.SetCellValue(RB6Name, fmt.Sprintf("%s%d", "E", i+2), contract.DiscountPercent)
			f.SetCellValue(RB6Name, fmt.Sprintf("%s%d", "F", i+2), contract.DiscountAmount)
			f.SetCellValue(RB6Name, fmt.Sprintf("%s%d", "G", i+2), contract.TotalWithoutDiscount)
			f.SetCellValue(RB6Name, fmt.Sprintf("%s%d", "H", i+2), "по закупам")
			totalDiscountsSum += contract.DiscountAmount

			lastRow = i + 2
			f.SetCellStyle(RB6Name, fmt.Sprintf("%s%d", "F", i+2), fmt.Sprintf("%s%d", "G", i+2), moneyStyle)

			i++
		}
		lastRow += 1
		//f.SetCellValue(sheet, fmt.Sprintf("%s%d", "H1", 2), "МТЗ")

		//TODO: - вкладка итогов
		TotalAmount += totalDiscountsSum
		f.SetCellValue(RB6Name, fmt.Sprintf("%s%d", "E", lastRow), "Итог:")
		f.SetCellValue(RB6Name, fmt.Sprintf("%s%d", "F", lastRow), totalDiscountsSum)

		//f.SetCellStyle(RB6Name, fmt.Sprintf("%s%d", "D", lastRow), fmt.Sprintf("%s%d", "D", lastRow), style)
		//f.SetCellStyle(RB6Name, fmt.Sprintf("%s%d", "E", lastRow), fmt.Sprintf("%s%d", "D", lastRow), style)

		err = f.SetCellStyle(RB6Name, fmt.Sprintf("%s%d", "E", lastRow), fmt.Sprintf("%s%d", "F", lastRow), coloredMoneyStyle)

		//f.SetCellValue(sheet, "F1", "Вид скидки:")
		//f.SetCellValue(sheet, "B1", "Сумма скидки:")
		//f.SetCellValue(sheet, "C1", "Итог:")
		ineration++
		//err = f.SetCellStyle(sheet, fmt.Sprintf("%s%d", "A", ineration), fmt.Sprintf("%s%d", "D", lastRow), style)
		f.SetCellValue(sheet, fmt.Sprintf("%s%d", "A", ineration), RB6Name)
		f.SetCellValue(sheet, fmt.Sprintf("%s%d", "B", ineration), totalDiscountsSum)
		f.SetCellStyle(sheet, fmt.Sprintf("%s%d", "B", ineration), fmt.Sprintf("%s%d", "B", ineration), moneyStyle)
		//f.SetCellValue(sheet, fmt.Sprintf("%s%d", "C", ineration), RB6Name)

	}

	if isRB7 {
		rb7thType, err := GetRB7thType(request, contracts)
		if err != nil {
			return err
		}

		if err != nil {
			return err
		}
		f.NewSheet(RB7Name)

		f.SetCellValue(RB7Name, "A1", "Период")
		f.SetCellValue(RB7Name, "B1", "Номер договора/ДС")
		f.SetCellValue(RB7Name, "C1", "Тип скидки")
		f.SetCellValue(RB7Name, "D1", "Бренд")
		f.SetCellValue(RB7Name, "E1", "Скидка %")
		f.SetCellValue(RB7Name, "F1", "Сумма скидки")
		f.SetCellValue(RB7Name, "G1", "Общая сумма")
		f.SetCellValue(RB7Name, "H1", "Тип")

		f.SetColWidth(RB7Name, "A", "H", 22)
		f.SetColWidth(RB7Name, "С", "С", 40)

		err = f.SetCellStyle(RB7Name, "A1", "H1", style)
		var i int
		var totalDiscountsSum float64
		for _, contract := range rb7thType {
			f.SetCellValue(RB7Name, fmt.Sprintf("%s%d", "A", i+2), fmt.Sprintf("%s-%s", contract.StartDate, contract.EndDate))
			f.SetCellValue(RB7Name, fmt.Sprintf("%s%d", "B", i+2), contract.ContractNumber)
			f.SetCellValue(RB7Name, fmt.Sprintf("%s%d", "C", i+2), RB7Name)
			f.SetCellValue(RB7Name, fmt.Sprintf("%s%d", "D", i+2), contract.BrandName)
			f.SetCellValue(RB7Name, fmt.Sprintf("%s%d", "E", i+2), contract.DiscountPercent)
			f.SetCellValue(RB7Name, fmt.Sprintf("%s%d", "F", i+2), contract.DiscountAmount)

			f.SetCellValue(RB7Name, fmt.Sprintf("%s%d", "G", i+2), contract.TotalWithoutDiscount)
			f.SetCellValue(RB7Name, fmt.Sprintf("%s%d", "H", i+2), "по продажам")
			totalDiscountsSum += contract.DiscountAmount
			lastRow = i + 2
			f.SetCellStyle(RB7Name, fmt.Sprintf("%s%d", "F", i+2), fmt.Sprintf("%s%d", "H", i+2), moneyStyle)
			i++

		}
		lastRow += 1
		TotalAmount += totalDiscountsSum
		f.SetCellValue(RB7Name, fmt.Sprintf("%s%d", "E", lastRow), "Итог:")
		f.SetCellValue(RB7Name, fmt.Sprintf("%s%d", "F", lastRow), totalDiscountsSum)
		err = f.SetCellStyle(RB7Name, fmt.Sprintf("%s%d", "E", lastRow), fmt.Sprintf("%s%d", "F", lastRow), coloredMoneyStyle)
		//err = f.SetCellStyle(RB7Name, fmt.Sprintf("%s%d", "D", lastRow), fmt.Sprintf("%s%d", "D", lastRow), style)
		//err = f.SetCellStyle(RB7Name, fmt.Sprintf("%s%d", "E", lastRow), fmt.Sprintf("%s%d", "D", lastRow), style)

		err = f.SetCellStyle(RB7Name, fmt.Sprintf("%s%d", "E", lastRow), fmt.Sprintf("%s%d", "G", lastRow), moneyStyle)

		ineration++
		//err = f.SetCellStyle(sheet, fmt.Sprintf("%s%d", "A", ineration), fmt.Sprintf("%s%d", "D", lastRow), style)
		f.SetCellValue(sheet, fmt.Sprintf("%s%d", "A", ineration), RB7Name)
		f.SetCellValue(sheet, fmt.Sprintf("%s%d", "B", ineration), totalDiscountsSum)
		f.SetCellStyle(sheet, fmt.Sprintf("%s%d", "B", ineration), fmt.Sprintf("%s%d", "B", ineration), moneyStyle)

	}

	if isRB8 {
		rbEighthType, err := GetRB8thType(request, contracts)
		if err != nil {
			return err
		}

		f.NewSheet(RB8Name)
		f.SetCellValue(RB8Name, "A1", "Период")
		f.SetCellValue(RB8Name, "B1", "Номер договора/ДС")
		f.SetCellValue(RB8Name, "C1", "Тип скидки")
		f.SetCellValue(RB8Name, "D1", "Скидка %")
		f.SetCellValue(RB8Name, "E1", "Сумма скидки")
		f.SetCellValue(RB8Name, "F1", "Общая сумма")
		f.SetCellValue(RB8Name, "G1", "Тип")
		f.SetColWidth(RB8Name, "A", "G", 22)
		f.SetColWidth(RB8Name, "C", "C", 40)

		err = f.SetCellStyle(RB8Name, "A1", "G1", style)

		var totalDiscountsSum float64
		fmt.Printf("CHECK \n%+v\n CHECK", contracts)
		var i int
		for _, contract := range rbEighthType {
			f.SetCellValue(RB8Name, fmt.Sprintf("%s%d", "A", i+2), fmt.Sprintf("%s-%s", contract.StartDate, contract.EndDate))
			f.SetCellValue(RB8Name, fmt.Sprintf("%s%d", "B", i+2), contract.ContractNumber)
			f.SetCellValue(RB8Name, fmt.Sprintf("%s%d", "C", i+2), RB8Name)
			f.SetCellValue(RB8Name, fmt.Sprintf("%s%d", "D", i+2), contract.DiscountPercent)
			f.SetCellValue(RB8Name, fmt.Sprintf("%s%d", "E", i+2), contract.DiscountAmount)
			f.SetCellValue(RB8Name, fmt.Sprintf("%s%d", "F", i+2), contract.TotalWithoutDiscount)
			f.SetCellValue(RB8Name, fmt.Sprintf("%s%d", "G", i+2), "по закупам")
			totalDiscountsSum += contract.DiscountAmount

			lastRow = i + 2
			f.SetCellStyle(RB8Name, fmt.Sprintf("%s%d", "E", i+2), fmt.Sprintf("%s%d", "F", i+2), moneyStyle)
			i++
		}
		lastRow += 1
		TotalAmount = totalDiscountsSum
		f.SetCellValue(RB8Name, fmt.Sprintf("%s%d", "D", lastRow), "Итог:")
		f.SetCellValue(RB8Name, fmt.Sprintf("%s%d", "E", lastRow), totalDiscountsSum)
		err = f.SetCellStyle(RB8Name, fmt.Sprintf("%s%d", "D", lastRow), fmt.Sprintf("%s%d", "E", lastRow), coloredMoneyStyle)
		//err = f.SetCellStyle(RB8Name, fmt.Sprintf("%s%d", "D", lastRow), fmt.Sprintf("%s%d", "E", lastRow), style)

		//err = f.SetCellStyle(RB8Name, fmt.Sprintf("%s%d", "D", lastRow), fmt.Sprintf("%s%d", "F", lastRow), moneyStyle)

		ineration++
		//err = f.SetCellStyle(sheet, fmt.Sprintf("%s%d", "A", ineration), fmt.Sprintf("%s%d", "D", lastRow), style)
		f.SetCellValue(sheet, fmt.Sprintf("%s%d", "A", ineration), RB8Name)
		f.SetCellValue(sheet, fmt.Sprintf("%s%d", "B", ineration), totalDiscountsSum)
		f.SetCellStyle(sheet, fmt.Sprintf("%s%d", "B", ineration), fmt.Sprintf("%s%d", "B", ineration), moneyStyle)
	}

	if isRB9 {
		rb9thType, err := GetRB9thType(request, contracts)
		if err != nil {
			return err
		}

		f.NewSheet(RB9Name)
		f.SetCellValue(RB9Name, "A1", "Период")
		f.SetCellValue(RB9Name, "B1", "Номер договора/ДС")
		f.SetCellValue(RB9Name, "C1", "Тип скидки")
		f.SetCellValue(RB9Name, "D1", "Скидка %")
		f.SetCellValue(RB9Name, "E1", "Сумма скидки")

		f.SetCellValue(RB9Name, "F1", "Общая сумма")
		f.SetCellValue(RB9Name, "G1", "Тип")

		f.SetColWidth(RB9Name, "A", "G", 22)
		f.SetColWidth(RB9Name, "С", "С", 40)
		err = f.SetCellStyle(RB9Name, "A1", "G1", style)

		var totalDiscountsSum float64
		fmt.Printf("CHECK \n%+v\n CHECK", contracts)
		var i int
		for _, contract := range rb9thType {
			f.SetCellValue(RB9Name, fmt.Sprintf("%s%d", "A", i+2), fmt.Sprintf("%s-%s", contract.StartDate, contract.EndDate))
			f.SetCellValue(RB9Name, fmt.Sprintf("%s%d", "B", i+2), contract.ContractNumber)
			f.SetCellValue(RB9Name, fmt.Sprintf("%s%d", "C", i+2), RB9Name)
			f.SetCellValue(RB9Name, fmt.Sprintf("%s%d", "D", i+2), contract.DiscountPercent)
			f.SetCellValue(RB9Name, fmt.Sprintf("%s%d", "E", i+2), contract.DiscountAmount)

			f.SetCellValue(RB9Name, fmt.Sprintf("%s%d", "F", i+2), contract.TotalWithoutDiscount)
			f.SetCellValue(RB9Name, fmt.Sprintf("%s%d", "G", i+2), "по закупам")
			totalDiscountsSum += contract.DiscountAmount

			lastRow = i + 2
			f.SetCellStyle(RB9Name, fmt.Sprintf("%s%d", "E", i+2), fmt.Sprintf("%s%d", "F", i+2), moneyStyle)
			i++
		}
		lastRow += 1
		TotalAmount += totalDiscountsSum
		f.SetCellValue(RB9Name, fmt.Sprintf("%s%d", "D", lastRow), "Итог:")
		f.SetCellValue(RB9Name, fmt.Sprintf("%s%d", "E", lastRow), totalDiscountsSum)
		err = f.SetCellStyle(RB9Name, fmt.Sprintf("%s%d", "D", lastRow), fmt.Sprintf("%s%d", "E", lastRow), coloredMoneyStyle)
		//err = f.SetCellStyle(RB9Name, fmt.Sprintf("%s%d", "D", lastRow), fmt.Sprintf("%s%d", "E", lastRow), style)
		//f.SetCellStyle(RB9Name, fmt.Sprintf("%s%d", "D", lastRow), fmt.Sprintf("%s%d", "F", lastRow), moneyStyle)

		ineration++
		//err = f.SetCellStyle(sheet, fmt.Sprintf("%s%d", "A", ineration), fmt.Sprintf("%s%d", "D", lastRow), style)
		f.SetCellValue(sheet, fmt.Sprintf("%s%d", "A", ineration), RB9Name)
		f.SetCellValue(sheet, fmt.Sprintf("%s%d", "B", ineration), totalDiscountsSum)
		f.SetCellStyle(sheet, fmt.Sprintf("%s%d", "B", ineration), fmt.Sprintf("%s%d", "B", ineration), moneyStyle)
	}

	if isRB10 {
		rbFourthType, err := GetRb10thType(request, contracts)
		if err != nil {
			return err
		}

		f.NewSheet(RB10Name)
		f.SetCellValue(RB10Name, "A1", "Период")
		f.SetCellValue(RB10Name, "B1", "Номер договора/ДС")
		f.SetCellValue(RB10Name, "C1", "Тип скидки")
		f.SetCellValue(RB10Name, "D1", "Скидка %")
		f.SetCellValue(RB10Name, "E1", "Сумма скидки")
		f.SetCellValue(RB10Name, "F1", "Общая сумма")
		f.SetCellValue(RB10Name, "G1", "Тип")

		f.SetColWidth(RB10Name, "A", "G", 22)
		f.SetColWidth(RB10Name, "С", "С", 40)
		err = f.SetCellStyle(RB10Name, "A1", "G1", style)

		var totalDiscountsSum float64
		fmt.Printf("CHECK \n%+v\n CHECK", contracts)
		var i int
		for _, contract := range rbFourthType {
			f.SetCellValue(RB10Name, fmt.Sprintf("%s%d", "A", i+2), fmt.Sprintf("%s-%s", contract.StartDate, contract.EndDate))
			f.SetCellValue(RB10Name, fmt.Sprintf("%s%d", "B", i+2), contract.ContractNumber)
			f.SetCellValue(RB10Name, fmt.Sprintf("%s%d", "C", i+2), RB10Name)
			f.SetCellValue(RB10Name, fmt.Sprintf("%s%d", "D", i+2), contract.DiscountPercent)
			f.SetCellValue(RB10Name, fmt.Sprintf("%s%d", "E", i+2), contract.DiscountAmount)
			f.SetCellValue(RB10Name, fmt.Sprintf("%s%d", "F", i+2), contract.TotalWithoutDiscount)
			f.SetCellValue(RB10Name, fmt.Sprintf("%s%d", "G", i+2), "по продажам")
			totalDiscountsSum += contract.DiscountAmount

			lastRow = i + 2
			f.SetCellStyle(RB10Name, fmt.Sprintf("%s%d", "E", i+2), fmt.Sprintf("%s%d", "F", i+2), moneyStyle)
			i++
		}
		lastRow += 1
		TotalAmount += totalDiscountsSum
		f.SetCellValue(RB10Name, fmt.Sprintf("%s%d", "D", lastRow), "Итог:")
		f.SetCellValue(RB10Name, fmt.Sprintf("%s%d", "E", lastRow), totalDiscountsSum)
		err = f.SetCellStyle(RB10Name, fmt.Sprintf("%s%d", "D", lastRow), fmt.Sprintf("%s%d", "E", lastRow), coloredMoneyStyle)
		//err = f.SetCellStyle(RB10Name, fmt.Sprintf("%s%d", "D", lastRow), fmt.Sprintf("%s%d", "E", lastRow), style)
		//f.SetCellStyle(RB10Name, fmt.Sprintf("%s%d", "D", lastRow), fmt.Sprintf("%s%d", "F", lastRow), moneyStyle)

		ineration++
		//err = f.SetCellStyle(sheet, fmt.Sprintf("%s%d", "A", ineration), fmt.Sprintf("%s%d", "D", lastRow), style)
		f.SetCellValue(sheet, fmt.Sprintf("%s%d", "A", ineration), RB10Name)
		f.SetCellValue(sheet, fmt.Sprintf("%s%d", "B", ineration), totalDiscountsSum)
		f.SetCellStyle(sheet, fmt.Sprintf("%s%d", "B", ineration), fmt.Sprintf("%s%d", "B", ineration), moneyStyle)
	}

	if isRB11 {

		rb11thType, err := GetRB11thType(request, contracts)
		if err != nil {
			return err
		}

		f.NewSheet(RB11Name)
		f.SetCellValue(RB11Name, "A1", "Период")
		f.SetCellValue(RB11Name, "B1", "Номер договора/ДС")
		f.SetCellValue(RB11Name, "C1", "Тип скидки")
		f.SetCellValue(RB11Name, "D1", "Скидка %")
		f.SetCellValue(RB11Name, "E1", "Сумма скидки")
		f.SetCellValue(RB11Name, "F1", "Общая сумма")
		f.SetCellValue(RB11Name, "G1", "Тип")

		//f.SetCellValue(RB4Name, "D1", "Код товара")
		//f.SetCellValue(RB4Name, "E1", "План закупа")
		f.SetColWidth(RB11Name, "A", "G", 22)
		f.SetColWidth(RB11Name, "С", "С", 40)
		err = f.SetCellStyle(RB11Name, "A1", "G1", style)

		var totalDiscountsSum float64
		fmt.Printf("CHECK \n%+v\n CHECK", contracts)
		var i int
		for _, contract := range rb11thType {
			f.SetCellValue(RB11Name, fmt.Sprintf("%s%d", "A", i+2), fmt.Sprintf("%s-%s", contract.StartDate, contract.EndDate))
			f.SetCellValue(RB11Name, fmt.Sprintf("%s%d", "B", i+2), contract.ContractNumber)
			f.SetCellValue(RB11Name, fmt.Sprintf("%s%d", "C", i+2), contract.DiscountType)
			f.SetCellValue(RB11Name, fmt.Sprintf("%s%d", "D", i+2), contract.DiscountPercent)
			f.SetCellValue(RB11Name, fmt.Sprintf("%s%d", "E", i+2), contract.DiscountAmount)
			f.SetCellValue(RB11Name, fmt.Sprintf("%s%d", "F", i+2), contract.TotalWithoutDiscount)
			f.SetCellValue(RB11Name, fmt.Sprintf("%s%d", "G", i+2), "по закупам")
			//f.SetCellValue(RB4Name, fmt.Sprintf("%s%d", "D", i+2), contract.ProductCode)
			//f.SetCellValue(RB4Name, fmt.Sprintf("%s%d", "E", i+2), contract.LeasePlan)
			totalDiscountsSum += contract.DiscountAmount

			lastRow = i + 2
			f.SetCellStyle(RB11Name, fmt.Sprintf("%s%d", "E", i+2), fmt.Sprintf("%s%d", "F", i+2), moneyStyle)
			i++
		}
		lastRow += 1
		f.SetCellValue(RB11Name, fmt.Sprintf("%s%d", "D", lastRow), "Итог:")
		f.SetCellValue(RB11Name, fmt.Sprintf("%s%d", "E", lastRow), totalDiscountsSum)
		err = f.SetCellStyle(RB11Name, fmt.Sprintf("%s%d", "D", lastRow), fmt.Sprintf("%s%d", "E", lastRow), coloredMoneyStyle)
		//err = f.SetCellStyle(RB11Name, fmt.Sprintf("%s%d", "D", lastRow), fmt.Sprintf("%s%d", "E", lastRow), style)
		//f.SetCellStyle(RB11Name, fmt.Sprintf("%s%d", "D", lastRow), fmt.Sprintf("%s%d", "F", lastRow), moneyStyle)
		if err != nil {
			return err
		}

		ineration++
		TotalAmount += totalDiscountsSum
		//err = f.SetCellStyle(sheet, fmt.Sprintf("%s%d", "A", ineration), fmt.Sprintf("%s%d", "D", lastRow), style)
		f.SetCellValue(sheet, fmt.Sprintf("%s%d", "A", ineration), RB11Name)
		f.SetCellValue(sheet, fmt.Sprintf("%s%d", "B", ineration), totalDiscountsSum)
		f.SetCellStyle(sheet, fmt.Sprintf("%s%d", "B", ineration), fmt.Sprintf("%s%d", "B", ineration), moneyStyle)
	}

	if isRB12 {

		log.Println("RB12->>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>")
		rb12thType, err := GetRB12thType(request, contracts)
		if err != nil {
			return err
		}

		f.NewSheet(RB12Name)
		f.SetCellValue(RB12Name, "A1", "Период")
		f.SetCellValue(RB12Name, "B1", "Номер договора/ДС")
		f.SetCellValue(RB12Name, "C1", "Тип скидки")
		f.SetCellValue(RB12Name, "D1", "Скидка %")
		f.SetCellValue(RB12Name, "E1", "Сумма скидки")

		f.SetCellValue(RB12Name, "F1", "Общая сумма")

		f.SetCellValue(RB12Name, "G1", "Тип")
		//f.SetCellValue(RB4Name, "D1", "Код товара")
		//f.SetCellValue(RB4Name, "E1", "План закупа")
		f.SetColWidth(RB12Name, "A", "G", 22)
		f.SetColWidth(RB12Name, "С", "С", 40)
		err = f.SetCellStyle(RB12Name, "A1", "G1", style)

		var totalDiscountsSum float64
		fmt.Printf("CHECK \n%+v\n CHECK", contracts)
		var i int
		for _, contract := range rb12thType {
			f.SetCellValue(RB12Name, fmt.Sprintf("%s%d", "A", i+2), fmt.Sprintf("%s-%s", contract.StartDate, contract.EndDate))
			f.SetCellValue(RB12Name, fmt.Sprintf("%s%d", "B", i+2), contract.ContractNumber)
			f.SetCellValue(RB12Name, fmt.Sprintf("%s%d", "C", i+2), contract.DiscountType)
			f.SetCellValue(RB12Name, fmt.Sprintf("%s%d", "D", i+2), contract.DiscountPercent)
			f.SetCellValue(RB12Name, fmt.Sprintf("%s%d", "E", i+2), contract.DiscountAmount)
			f.SetCellValue(RB12Name, fmt.Sprintf("%s%d", "F", i+2), contract.TotalWithoutDiscount)
			f.SetCellValue(RB12Name, fmt.Sprintf("%s%d", "G", i+2), "по закупам")
			//f.SetCellValue(RB4Name, fmt.Sprintf("%s%d", "D", i+2), contract.ProductCode)
			//f.SetCellValue(RB4Name, fmt.Sprintf("%s%d", "E", i+2), contract.LeasePlan)
			totalDiscountsSum += contract.DiscountAmount

			lastRow = i + 2
			f.SetCellStyle(RB12Name, fmt.Sprintf("%s%d", "E", i+2), fmt.Sprintf("%s%d", "F", i+2), moneyStyle)
			i++

		}
		lastRow += 1
		TotalAmount += totalDiscountsSum
		f.SetCellValue(RB12Name, fmt.Sprintf("%s%d", "D", lastRow), "Итог:")
		f.SetCellValue(RB12Name, fmt.Sprintf("%s%d", "E", lastRow), totalDiscountsSum)
		err = f.SetCellStyle(RB12Name, fmt.Sprintf("%s%d", "D", lastRow), fmt.Sprintf("%s%d", "E", lastRow), coloredMoneyStyle)
		//err = f.SetCellStyle(RB12Name, fmt.Sprintf("%s%d", "D", lastRow), fmt.Sprintf("%s%d", "E", lastRow), style)
		//f.SetCellStyle(RB12Name, fmt.Sprintf("%s%d", "D", lastRow), fmt.Sprintf("%s%d", "F", lastRow), moneyStyle)
		//err = f.SetCellStyle(RB2Name, fmt.Sprintf("%s%d", "G", lastRow), fmt.Sprintf("%s%d", "G", lastRow), style)

		ineration++
		//err = f.SetCellStyle(sheet, fmt.Sprintf("%s%d", "A", ineration), fmt.Sprintf("%s%d", "D", lastRow), style)
		f.SetCellValue(sheet, fmt.Sprintf("%s%d", "A", ineration), RB12Name)
		f.SetCellValue(sheet, fmt.Sprintf("%s%d", "B", ineration), totalDiscountsSum)
		f.SetCellStyle(sheet, fmt.Sprintf("%s%d", "B", ineration), fmt.Sprintf("%s%d", "B", ineration), moneyStyle)

		//if isRB13 {
		//	rb12thType, err := GetRB12thType(request)
		//	if err != nil {
		//		return err
		//	}
		//
		//	f.NewSheet(RB13Name)
		//	f.SetCellValue(RB13Name, "A1", "Период")
		//	f.SetCellValue(RB13Name, "B1", "Номер договора/ДС")
		//	f.SetCellValue(RB13Name, "C1", "Тип скидки")
		//	f.SetCellValue(RB13Name, "D1", "Скидка %")
		//	f.SetCellValue(RB13Name, "E1", "Сумма скидки")
		//	//f.SetCellValue(RB4Name, "D1", "Код товара")
		//	//f.SetCellValue(RB4Name, "E1", "План закупа")
		//	err = f.SetCellStyle(RB13Name, "A1", "E1", style)
		//
		//	var totalDiscountsSum float32
		//	fmt.Printf("CHECK \n%+v\n CHECK", contracts)
		//	var i int
		//	for _, contract := range rb12thType {
		//		f.SetCellValue(RB13Name, fmt.Sprintf("%s%d", "A", i+2), fmt.Sprintf("%s-%s", contract.StartDate, contract.EndDate))
		//		f.SetCellValue(RB13Name, fmt.Sprintf("%s%d", "B", i+2), contract.ContractNumber)
		//		f.SetCellValue(RB13Name, fmt.Sprintf("%s%d", "C", i+2), RB13Name)
		//		f.SetCellValue(RB13Name, fmt.Sprintf("%s%d", "D", i+2), contract.DiscountPercent)
		//		f.SetCellValue(RB13Name, fmt.Sprintf("%s%d", "E", i+2), contract.DiscountAmount)
		//		//f.SetCellValue(RB4Name, fmt.Sprintf("%s%d", "D", i+2), contract.ProductCode)
		//		//f.SetCellValue(RB4Name, fmt.Sprintf("%s%d", "E", i+2), contract.LeasePlan)
		//		totalDiscountsSum += contract.DiscountAmount
		//		lastRow = i + 2
		//		i++
		//	}
		//	lastRow += 1
		//	f.SetCellValue(RB13Name, fmt.Sprintf("%s%d", "D", lastRow), "Итог:")
		//	f.SetCellValue(RB13Name, fmt.Sprintf("%s%d", "E", lastRow), totalDiscountsSum)
		//	err = f.SetCellStyle(RB13Name, fmt.Sprintf("%s%d", "D", lastRow), fmt.Sprintf("%s%d", "E", lastRow), style)
		//	//err = f.SetCellStyle(RB2Name, fmt.Sprintf("%s%d", "G", lastRow), fmt.Sprintf("%s%d", "G", lastRow), style)
	}

	if isRB13 {
		log.Println("13 отчет генерировался--------------------------------------------------------------------")

		rb13thType, err := GetRB13thType(request, contracts)
		if err != nil {
			return err
		}

		f.NewSheet(RB13Name)
		f.SetCellValue(RB13Name, "A1", "Период")
		f.SetCellValue(RB13Name, "B1", "Номер договора/ДС")
		f.SetCellValue(RB13Name, "C1", "Тип скидки")
		f.SetCellValue(RB13Name, "D1", "Скидка %")
		f.SetCellValue(RB13Name, "E1", "Сумма скидки")

		//f.SetCellValue(RB4Name, "D1", "Код товара")
		//f.SetCellValue(RB4Name, "E1", "План закупа")
		f.SetColWidth(RB13Name, "A", "E", 22)
		f.SetColWidth(RB13Name, "C", "C", 40)
		err = f.SetCellStyle(RB13Name, "A1", "E1", style)

		var totalDiscountsSum float64
		fmt.Printf("CHECK \n%+v\n CHECK", contracts)
		var i int
		for _, contract := range rb13thType {
			f.SetCellValue(RB13Name, fmt.Sprintf("%s%d", "A", i+2), fmt.Sprintf("%s-%s", contract.StartDate, contract.EndDate))
			f.SetCellValue(RB13Name, fmt.Sprintf("%s%d", "B", i+2), contract.ContractNumber)
			f.SetCellValue(RB13Name, fmt.Sprintf("%s%d", "C", i+2), RB13Name)
			f.SetCellValue(RB13Name, fmt.Sprintf("%s%d", "D", i+2), contract.DiscountPercent)
			f.SetCellValue(RB13Name, fmt.Sprintf("%s%d", "E", i+2), contract.DiscountAmount)
			//f.SetCellValue(RB4Name, fmt.Sprintf("%s%d", "D", i+2), contract.ProductCode)
			//f.SetCellValue(RB4Name, fmt.Sprintf("%s%d", "E", i+2), contract.LeasePlan)
			totalDiscountsSum += contract.DiscountAmount
			TotalAmount += totalDiscountsSum
			lastRow = i + 2

			i++
			//f.SetCellStyle(RB13Name, fmt.Sprintf("%s%d", "E", i+2), fmt.Sprintf("%s%d", "H", i+2), moneyStyle)
		}
		lastRow += 1
		f.SetCellValue(RB13Name, fmt.Sprintf("%s%d", "D", lastRow), "Итог:")
		f.SetCellValue(RB13Name, fmt.Sprintf("%s%d", "E", lastRow), totalDiscountsSum)
		err = f.SetCellStyle(RB13Name, fmt.Sprintf("%s%d", "D", lastRow), fmt.Sprintf("%s%d", "E", lastRow), coloredMoneyStyle)
		//err = f.SetCellStyle(RB13Name, fmt.Sprintf("%s%d", "D", lastRow), fmt.Sprintf("%s%d", "E", lastRow), style)
		//err = f.SetCellStyle(RB2Name, fmt.Sprintf("%s%d", "G", lastRow), fmt.Sprintf("%s%d", "G", lastRow), style)

		ineration++
		//err = f.SetCellStyle(sheet, fmt.Sprintf("%s%d", "A", ineration), fmt.Sprintf("%s%d", "D", lastRow), style)
		f.SetCellValue(sheet, fmt.Sprintf("%s%d", "A", ineration), RB13Name)
		f.SetCellValue(sheet, fmt.Sprintf("%s%d", "B", ineration), totalDiscountsSum)
		f.SetCellStyle(sheet, fmt.Sprintf("%s%d", "B", ineration), fmt.Sprintf("%s%d", "B", ineration), moneyStyle)
	}

	if isRB14 {
		rb14ThType, err := GetRB14ThType(request, contracts)
		if err != nil {
			return err
		}

		f.NewSheet(RB14Name)
		f.SetCellValue(RB14Name, "A1", "Период")
		f.SetCellValue(RB14Name, "B1", "Номер договора/ДС")
		f.SetCellValue(RB14Name, "C1", "Тип скидки")
		f.SetCellValue(RB14Name, "D1", "Код товара")
		f.SetCellValue(RB14Name, "E1", "План закупа")
		f.SetCellValue(RB14Name, "F1", "Скидка %")
		f.SetCellValue(RB14Name, "G1", "Сумма скидки")
		f.SetCellValue(RB14Name, "H1", "Общая сумма")
		f.SetCellValue(RB14Name, "I1", "Тип")
		f.SetColWidth(RB14Name, "A", "I", 22)
		f.SetColWidth(RB14Name, "С", "С", 40)
		err = f.SetCellStyle(RB14Name, "A1", "I1", style)

		var totalDiscountsSum float64
		fmt.Printf("CHECK \n%+v\n CHECK", contracts)
		var i int
		for _, contract := range rb14ThType {
			f.SetCellValue(RB14Name, fmt.Sprintf("%s%d", "A", i+2), fmt.Sprintf("%s-%s", contract.StartDate, contract.EndDate))
			f.SetCellValue(RB14Name, fmt.Sprintf("%s%d", "B", i+2), contract.ContractNumber)
			f.SetCellValue(RB14Name, fmt.Sprintf("%s%d", "C", i+2), RB14Name)
			f.SetCellValue(RB14Name, fmt.Sprintf("%s%d", "D", i+2), contract.ProductCode)
			f.SetCellValue(RB14Name, fmt.Sprintf("%s%d", "E", i+2), contract.LeasePlan)
			f.SetCellValue(RB14Name, fmt.Sprintf("%s%d", "F", i+2), contract.DiscountPercent)
			f.SetCellValue(RB14Name, fmt.Sprintf("%s%d", "G", i+2), contract.DiscountAmount)
			f.SetCellValue(RB14Name, fmt.Sprintf("%s%d", "H", i+2), contract.TotalWithoutDiscount)
			f.SetCellValue(RB14Name, fmt.Sprintf("%s%d", "I", i+2), "по продажам")
			totalDiscountsSum += contract.DiscountAmount

			lastRow = i + 2
			f.SetCellStyle(RB14Name, fmt.Sprintf("%s%d", "E", i+2), fmt.Sprintf("%s%d", "H", i+2), moneyStyle)
			i++
		}
		lastRow += 1
		TotalAmount += totalDiscountsSum
		f.SetCellValue(RB14Name, fmt.Sprintf("%s%d", "F", lastRow), "Итог:")
		f.SetCellValue(RB14Name, fmt.Sprintf("%s%d", "G", lastRow), totalDiscountsSum)
		err = f.SetCellStyle(RB14Name, fmt.Sprintf("%s%d", "F", lastRow), fmt.Sprintf("%s%d", "G", lastRow), coloredMoneyStyle)
		//err = f.SetCellStyle(RB14Name, fmt.Sprintf("%s%d", "F", lastRow), fmt.Sprintf("%s%d", "G", lastRow), style)
		//err = f.SetCellStyle(RB2Name, fmt.Sprintf("%s%d", "G", lastRow), fmt.Sprintf("%s%d", "G", lastRow), style)

		ineration++
		//err = f.SetCellStyle(sheet, fmt.Sprintf("%s%d", "A", ineration), fmt.Sprintf("%s%d", "D", lastRow), style)
		f.SetCellValue(sheet, fmt.Sprintf("%s%d", "A", ineration), RB14Name)
		f.SetCellValue(sheet, fmt.Sprintf("%s%d", "B", ineration), totalDiscountsSum)
		f.SetCellStyle(sheet, fmt.Sprintf("%s%d", "B", ineration), fmt.Sprintf("%s%d", "B", ineration), moneyStyle)
	}

	if isRB15 {
		contract1stType, err := GetRB15ThType(request, contracts)
		if err != nil {
			return err
		}

		f.NewSheet(RB15Name)
		f.SetCellValue(RB15Name, "A1", "Период")
		f.SetCellValue(RB15Name, "B1", "Номер договора/ДС")
		f.SetCellValue(RB15Name, "C1", "Тип скидки")
		f.SetCellValue(RB15Name, "D1", "План продаж")
		f.SetCellValue(RB15Name, "E1", "Сумма вознаграждения")
		f.SetCellValue(RB15Name, "F1", "Сумма скидки")
		f.SetCellValue(RB15Name, "G1", "Общая сумма")
		f.SetCellValue(RB15Name, "H1", "Тип")

		f.SetColWidth(RB15Name, "A", "G", 22)
		f.SetColWidth(RB15Name, "C", "H", 40)
		err = f.SetCellStyle(RB15Name, "A1", "H1", style)

		var totalDiscountsSum float64
		fmt.Printf("CHECK \n%+v\n CHECK", contracts)
		var i int
		for _, contract := range contract1stType {
			f.SetCellValue(RB15Name, fmt.Sprintf("%s%d", "A", i+2), fmt.Sprintf("%s-%s", contract.StartDate, contract.EndDate))
			f.SetCellValue(RB15Name, fmt.Sprintf("%s%d", "B", i+2), contract.ContractNumber)
			f.SetCellValue(RB15Name, fmt.Sprintf("%s%d", "C", i+2), RB15Name)
			f.SetCellValue(RB15Name, fmt.Sprintf("%s%d", "D", i+2), contract.LeasePlan)
			f.SetCellValue(RB15Name, fmt.Sprintf("%s%d", "E", i+2), contract.RewardAmount)
			f.SetCellValue(RB15Name, fmt.Sprintf("%s%d", "F", i+2), contract.DiscountAmount)
			f.SetCellValue(RB15Name, fmt.Sprintf("%s%d", "G", i+2), contract.TotalWithoutDiscount)
			f.SetCellValue(RB15Name, fmt.Sprintf("%s%d", "H", i+2), "по продажам")
			//TODO: добавить сумму без скидки
			totalDiscountsSum += contract.DiscountAmount

			lastRow = i + 2
			f.SetCellStyle(RB15Name, fmt.Sprintf("%s%d", "D", i+2), fmt.Sprintf("%s%d", "G", i+2), moneyStyle)
			i++

		}
		lastRow += 1
		TotalAmount += totalDiscountsSum
		f.SetCellValue(RB15Name, fmt.Sprintf("%s%d", "D", lastRow), "Итог:")
		f.SetCellValue(RB15Name, fmt.Sprintf("%s%d", "E", lastRow), totalDiscountsSum)
		err = f.SetCellStyle(RB15Name, fmt.Sprintf("%s%d", "D", lastRow), fmt.Sprintf("%s%d", "E", lastRow), coloredMoneyStyle)
		//err = f.SetCellStyle(RB15Name, fmt.Sprintf("%s%d", "D", lastRow), fmt.Sprintf("%s%d", "D", lastRow), style)
		//err = f.SetCellStyle(RB15Name, fmt.Sprintf("%s%d", "E", lastRow), fmt.Sprintf("%s%d", "D", lastRow), style)

		ineration++
		//err = f.SetCellStyle(sheet, fmt.Sprintf("%s%d", "A", ineration), fmt.Sprintf("%s%d", "D", lastRow), style)
		f.SetCellValue(sheet, fmt.Sprintf("%s%d", "A", ineration), RB15Name)
		f.SetCellValue(sheet, fmt.Sprintf("%s%d", "B", ineration), totalDiscountsSum)
		f.SetCellStyle(sheet, fmt.Sprintf("%s%d", "B", ineration), fmt.Sprintf("%s%d", "B", ineration), moneyStyle)
	}

	if isRB16 {
		rb16thType, err := GetRB16ThType(request, contracts)
		if err != nil {
			return err
		}

		f.NewSheet(RB16Name)
		f.SetCellValue(RB16Name, "A1", "Период")
		f.SetCellValue(RB16Name, "B1", "Номер договора/ДС")
		f.SetCellValue(RB16Name, "C1", "Тип скидки")
		f.SetCellValue(RB16Name, "D1", "Скидка %")
		f.SetCellValue(RB16Name, "E1", "План продаж")
		f.SetCellValue(RB16Name, "F1", "Сумма скидки")
		f.SetCellValue(RB16Name, "G1", "Общая сумма")
		f.SetCellValue(RB16Name, "H1", "Тип")
		f.SetColWidth(RB16Name, "A", "H", 22)
		f.SetColWidth(RB16Name, "С", "С", 40)

		err = f.SetCellStyle(RB16Name, "A1", "H1", style)

		var totalDiscountsSum float64
		fmt.Printf("CHECK \n%+v\n CHECK", contracts)
		var i int
		for _, contract := range rb16thType {
			f.SetCellValue(RB16Name, fmt.Sprintf("%s%d", "A", i+2), fmt.Sprintf("%s-%s", contract.StartDate, contract.EndDate))
			f.SetCellValue(RB16Name, fmt.Sprintf("%s%d", "B", i+2), contract.ContractNumber)
			f.SetCellValue(RB16Name, fmt.Sprintf("%s%d", "C", i+2), RB16Name)
			f.SetCellValue(RB16Name, fmt.Sprintf("%s%d", "D", i+2), contract.DiscountPercent)
			f.SetCellValue(RB16Name, fmt.Sprintf("%s%d", "E", i+2), contract.LeasePlan)
			f.SetCellValue(RB16Name, fmt.Sprintf("%s%d", "F", i+2), contract.DiscountAmount)
			f.SetCellValue(RB16Name, fmt.Sprintf("%s%d", "G", i+2), contract.TotalWithoutDiscount)
			f.SetCellValue(RB16Name, fmt.Sprintf("%s%d", "H", i+2), "по продажам")
			totalDiscountsSum += contract.DiscountAmount

			lastRow = i + 2
			i++
			f.SetCellStyle(RB16Name, fmt.Sprintf("%s%d", "E", i+2), fmt.Sprintf("%s%d", "G", i+2), moneyStyle)
		}
		lastRow += 1
		TotalAmount += totalDiscountsSum
		f.SetCellValue(RB16Name, fmt.Sprintf("%s%d", "D", lastRow), "Итог:")
		f.SetCellValue(RB16Name, fmt.Sprintf("%s%d", "E", lastRow), totalDiscountsSum)
		err = f.SetCellStyle(RB16Name, fmt.Sprintf("%s%d", "D", lastRow), fmt.Sprintf("%s%d", "E", lastRow), coloredMoneyStyle)
		//err = f.SetCellStyle(RB16Name, fmt.Sprintf("%s%d", "D", lastRow), fmt.Sprintf("%s%d", "E", lastRow), style)
		//f.SetCellStyle(RB16Name, fmt.Sprintf("%s%d", "E", lastRow), fmt.Sprintf("%s%d", "G", lastRow), moneyStyle)

		ineration++
		//err = f.SetCellStyle(sheet, fmt.Sprintf("%s%d", "A", ineration), fmt.Sprintf("%s%d", "D", lastRow), style)
		f.SetCellValue(sheet, fmt.Sprintf("%s%d", "A", ineration), RB16Name)
		f.SetCellValue(sheet, fmt.Sprintf("%s%d", "B", ineration), totalDiscountsSum)
		f.SetCellStyle(sheet, fmt.Sprintf("%s%d", "B", ineration), fmt.Sprintf("%s%d", "B", ineration), moneyStyle)

	}

	if isRB17 {
		contract17ThType, err := GetRB17ThType(request, contracts)
		if err != nil {
			return err
		}

		f.NewSheet(RB17Name)
		f.SetCellValue(RB17Name, "A1", "Период")
		f.SetCellValue(RB17Name, "B1", "Номер договора/ДС")
		f.SetCellValue(RB17Name, "C1", "Тип скидки")
		f.SetCellValue(RB17Name, "D1", "Скидка %")
		f.SetCellValue(RB17Name, "E1", "Сумма скидки")
		f.SetCellValue(RB17Name, "F1", "Общая сумма")
		f.SetCellValue(RB17Name, "G1", "Общая сумма")
		f.SetColWidth(RB17Name, "A", "G", 22)
		f.SetColWidth(RB17Name, "C", "C", 40)
		err = f.SetCellStyle(RB17Name, "A1", "G1", style)

		var totalDiscountsSum float64
		fmt.Printf("CHECK \n%+v\n CHECK", contracts)
		var i int
		for _, contract := range contract17ThType {
			f.SetCellValue(RB17Name, fmt.Sprintf("%s%d", "A", i+2), fmt.Sprintf("%s-%s", contract.StartDate, contract.EndDate))
			f.SetCellValue(RB17Name, fmt.Sprintf("%s%d", "B", i+2), contract.ContractNumber)
			f.SetCellValue(RB17Name, fmt.Sprintf("%s%d", "C", i+2), RB17Name)
			f.SetCellValue(RB17Name, fmt.Sprintf("%s%d", "D", i+2), contract.DiscountPercent)
			f.SetCellValue(RB17Name, fmt.Sprintf("%s%d", "E", i+2), contract.DiscountAmount)
			f.SetCellValue(RB17Name, fmt.Sprintf("%s%d", "F", i+2), contract.TotalWithoutDiscount)
			f.SetCellValue(RB17Name, fmt.Sprintf("%s%d", "G", i+2), "по продажам")
			totalDiscountsSum += contract.DiscountAmount

			lastRow = i + 2
			f.SetCellStyle(RB17Name, fmt.Sprintf("%s%d", "E", i+2), fmt.Sprintf("%s%d", "F", i+2), moneyStyle)
			i++

		}
		lastRow += 1

		TotalAmount += totalDiscountsSum
		f.SetCellValue(RB17Name, fmt.Sprintf("%s%d", "D", lastRow), "Итог:")
		f.SetCellValue(RB17Name, fmt.Sprintf("%s%d", "E", lastRow), totalDiscountsSum)
		//err = f.SetCellStyle(RB17Name, fmt.Sprintf("%s%d", "D", lastRow), fmt.Sprintf("%s%d", "E", lastRow), style)
		err = f.SetCellStyle(RB17Name, fmt.Sprintf("%s%d", "D", lastRow), fmt.Sprintf("%s%d", "E", lastRow), coloredMoneyStyle)
		//f.SetCellStyle(RB17Name, fmt.Sprintf("%s%d", "D", lastRow), fmt.Sprintf("%s%d", "F", lastRow), moneyStyle)

		ineration++
		//err = f.SetCellStyle(sheet, fmt.Sprintf("%s%d", "A", ineration), fmt.Sprintf("%s%d", "D", lastRow), style)
		f.SetCellValue(sheet, fmt.Sprintf("%s%d", "A", ineration), RB17Name)
		f.SetCellValue(sheet, fmt.Sprintf("%s%d", "B", ineration), totalDiscountsSum)
		f.SetCellStyle(sheet, fmt.Sprintf("%s%d", "B", ineration), fmt.Sprintf("%s%d", "B", ineration), moneyStyle)

	}

	if isRB18 {
		rb18thType, err := GetRB18thType(request, contracts)
		if err != nil {
			return err
		}

		f.NewSheet(RB18Name)
		f.SetCellValue(RB18Name, "A1", "Период")
		f.SetCellValue(RB18Name, "B1", "Номер договора/ДС")
		f.SetCellValue(RB18Name, "C1", "Тип скидки")
		f.SetCellValue(RB18Name, "D1", "Скидка %")
		f.SetCellValue(RB18Name, "E1", "Сумма скидки")
		f.SetCellValue(RB18Name, "F1", "Общая сумма")
		f.SetCellValue(RB18Name, "G1", "Тип")
		f.SetColWidth(RB18Name, "A", "G", 22)
		f.SetColWidth(RB18Name, "C", "C", 40)
		err = f.SetCellStyle(RB18Name, "A1", "G1", style)

		var totalDiscountsSum float64
		fmt.Printf("CHECK \n%+v\n CHECK", contracts)
		var i int
		for _, contract := range rb18thType {
			f.SetCellValue(RB18Name, fmt.Sprintf("%s%d", "A", i+2), fmt.Sprintf("%s-%s", contract.StartDate, contract.EndDate))
			f.SetCellValue(RB18Name, fmt.Sprintf("%s%d", "B", i+2), contract.ContractNumber)
			f.SetCellValue(RB18Name, fmt.Sprintf("%s%d", "C", i+2), RB18Name)
			f.SetCellValue(RB18Name, fmt.Sprintf("%s%d", "D", i+2), contract.DiscountPercent)
			f.SetCellValue(RB18Name, fmt.Sprintf("%s%d", "E", i+2), contract.DiscountAmount)
			f.SetCellValue(RB18Name, fmt.Sprintf("%s%d", "F", i+2), contract.TotalWithoutDiscount)
			f.SetCellValue(RB18Name, fmt.Sprintf("%s%d", "G", i+2), "по закупам")
			totalDiscountsSum += contract.DiscountAmount

			lastRow = i + 2
			f.SetCellStyle(RB18Name, fmt.Sprintf("%s%d", "E", i+2), fmt.Sprintf("%s%d", "F", i+2), moneyStyle)
			i++
		}
		lastRow += 1
		TotalAmount += totalDiscountsSum
		f.SetCellValue(RB18Name, fmt.Sprintf("%s%d", "D", lastRow), "Итог:")
		f.SetCellValue(RB18Name, fmt.Sprintf("%s%d", "E", lastRow), totalDiscountsSum)
		err = f.SetCellStyle(RB18Name, fmt.Sprintf("%s%d", "D", lastRow), fmt.Sprintf("%s%d", "E", lastRow), coloredMoneyStyle)
		//err = f.SetCellStyle(RB18Name, fmt.Sprintf("%s%d", "D", lastRow), fmt.Sprintf("%s%d", "E", lastRow), style)
		//f.SetCellStyle(RB18Name, fmt.Sprintf("%s%d", "D", lastRow), fmt.Sprintf("%s%d", "F", lastRow), moneyStyle)

		ineration++
		//err = f.SetCellStyle(sheet, fmt.Sprintf("%s%d", "A", ineration), fmt.Sprintf("%s%d", "D", lastRow), style)
		f.SetCellValue(sheet, fmt.Sprintf("%s%d", "A", ineration), RB18Name)
		f.SetCellValue(sheet, fmt.Sprintf("%s%d", "B", ineration), totalDiscountsSum)
		f.SetCellStyle(sheet, fmt.Sprintf("%s%d", "B", ineration), fmt.Sprintf("%s%d", "B", ineration), moneyStyle)
	}
	ineration++
	f.SetCellStyle(sheet, fmt.Sprintf("%s%d", "A", ineration), fmt.Sprintf("%s%d", "B", ineration), coloredMoneyStyle)
	f.SetCellValue(sheet, fmt.Sprintf("%s%d", "A", ineration), "Итог")
	f.SetCellValue(sheet, fmt.Sprintf("%s%d", "B", ineration), TotalAmount)

	f.DeleteSheet("Sheet1")
	f.SaveAs("files/reports/rb/rb_report.xlsx")
	return nil
}
