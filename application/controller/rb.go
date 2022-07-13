package controller

import (
	"contractor_panel/application/service"
	"github.com/gorilla/mux"
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

	//r.HandleFunc("/rb", c.).Methods(http.MethodOptions, http.MethodPost)
	//r.HandleFunc("/contractTemplates/{id}/file", c.DownloadFile).Methods(http.MethodOptions, http.MethodGet)
	//r.HandleFunc("/contracts", c.GetAllContracts).Methods(http.MethodOptions, http.MethodGet)

}

//func (c *RBReportController) GetAllRBByContractorBIN((w http.ResponseWriter, r *http.Request)  {
//	//c.s.GetAllRBByContractorBIN()
//
//}
