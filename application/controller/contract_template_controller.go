package controller

import (
	"contractor_panel/application/dto"
	"contractor_panel/application/middleware"
	"contractor_panel/application/respond"
	"contractor_panel/application/service"
	"fmt"
	"github.com/gorilla/mux"
	log "github.com/sirupsen/logrus"
	"net/http"
	"strconv"
)

type ContractTemplateController struct {
	s service.ContractTemplateService
}

func NewContractTemplateController(s service.ContractTemplateService) *ContractTemplateController {
	return &ContractTemplateController{s}
}

func (c *ContractTemplateController) HandleRoutes(r *mux.Router) {
	r.HandleFunc("/contractTemplates", c.GetAllContractTemplates).Methods(http.MethodOptions, http.MethodGet)
	r.HandleFunc("/contractTemplates/{id}/file", c.DownloadFile).Methods(http.MethodOptions, http.MethodGet)
	r.HandleFunc("/contracts", c.GetAllContracts).Methods(http.MethodOptions, http.MethodGet)
	r.HandleFunc("/contracts/{id}", c.GetContractByID).Methods(http.MethodOptions, http.MethodGet)

}
func (c *ContractTemplateController) GetContractByID(w http.ResponseWriter, r *http.Request) {
	id := mux.Vars(r)["id"]

	fmt.Println("ID", id)
	contractId, err := strconv.Atoi(id)
	if err != nil {
		respond.WithError(w, r, err)

	}

	contract, err := c.s.GetContractDetails(r.Context(), contractId)
	if err != nil {

		log.Println("[controller.GetContractDetails]|[service.GetContractDetails]| error is: ", err.Error())
		//c.JSON(http.StatusInternalServerError, gin.H{"reason": err.Error()})
		respond.WithError(w, r, err)
		return
	}

	respond.With(w, r, contract)

}

func (c *ContractTemplateController) GetAllContracts(w http.ResponseWriter, r *http.Request) {
	status := mux.Vars(r)["status"]
	//token := r.Header.Get("Authorization")
	data, _ := middleware.ExtractTokenData(r)
	fmt.Println("данные от токена", data)

	contracts, err := c.s.GetAllContracts(r.Context(), status, data.UserId)

	if err != nil {
		respond.WithError(w, r, err)
	}

	respond.With(w, r, contracts)
}

func (c *ContractTemplateController) GetAllContractTemplates(w http.ResponseWriter, r *http.Request) {
	err := r.ParseForm()
	if err != nil {
		respond.WithError(w, r, err)
		return
	}

	searchParameters, err := dto.ParseContractTemplateSearchParameters(r.Form)
	if err != nil {
		respond.WithError(w, r, err)
		return
	}

	res, total, err := c.s.FindContractTemplates(r.Context(), *searchParameters)
	if err != nil {
		respond.WithError(w, r, err)
		return
	}

	respond.WithPagination(w, r, dto.ConvertContractTemplateDtos(res), total)
}

func (c *ContractTemplateController) DownloadFile(w http.ResponseWriter, r *http.Request) {

}
