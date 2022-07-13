package controller

import (
	"contractor_panel/application/cerrors"
	"contractor_panel/application/middleware"
	"contractor_panel/application/respond"
	"contractor_panel/application/service"
	"contractor_panel/domain/model"
	"encoding/json"
	"fmt"
	"github.com/gorilla/mux"
	"net/http"
)

type RBReportController struct {
	s service.ReportTemplateService
}

func NewRBReportController(s service.ReportTemplateService) *RBReportController {
	return &RBReportController{s: s}
}

//func NewContractTemplateController(s service.ContractTemplateService) *ContractTemplateController {
//	return &ContractTemplateController{s}
//}

func (c *RBReportController) HandleRoutes(r *mux.Router) {

	r.HandleFunc("/rb", c.GetAllRBByContractorBIN).Methods(http.MethodOptions, http.MethodPost)
	r.HandleFunc("/rb/excel", c.FormExcelForRB).Methods(http.MethodOptions, http.MethodPost)
	//r.HandleFunc("/contractTemplates/{id}/file", c.DownloadFile).Methods(http.MethodOptions, http.MethodGet)
	//r.HandleFunc("/contracts", c.GetAllContracts).Methods(http.MethodOptions, http.MethodGet)

}

func (c *RBReportController) GetAllRBByContractorBIN(w http.ResponseWriter, r *http.Request) {
	tokenData, err := middleware.ExtractTokenData(r)
	if err != nil {
		//TODO handle error
	}
	rbRequest := model.RBRequest{}
	defer r.Body.Close()
	err = json.NewDecoder(r.Body).Decode(&rbRequest)
	if err != nil {
		respond.WithError(w, r, cerrors.ErrCouldNotDecodeBody(err))
		return
	}

	//TODO:
	RbDTOs, err := c.s.GetAllRBByContractorBIN(r.Context(), tokenData.UserId, rbRequest)
	if err != nil {
		// handle error
		fmt.Println("Ошибка при РБ", err)
	}
	for i := range RbDTOs {
		RbDTOs[i].Status = "Завершено"
	}

	respond.With(w, r, RbDTOs)

}

func (c *RBReportController) FormExcelForRB(w http.ResponseWriter, r *http.Request) {
	tokenData, err := middleware.ExtractTokenData(r)
	if err != nil {
		//TODO handle error
	}
	rbRequest := model.RBRequest{}
	defer r.Body.Close()
	err = json.NewDecoder(r.Body).Decode(&rbRequest)
	if err != nil {
		respond.WithError(w, r, cerrors.ErrCouldNotDecodeBody(err))
		return
	}
	w.Header().Set("Content-Type", "application/vnd.openxmlformats-officedocument.spreadsheetml.sheet")
	c.s.FormExcelForRBReport(r.Context(), tokenData.UserId, rbRequest)
	//w.WriteHeader
}
