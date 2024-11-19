package productshandler

import (
	"ecommers/helper"
	"ecommers/service"
	"ecommers/util"
	"net/http"
	"strconv"

	"go.uber.org/zap"
)

type ProductHandler struct {
	Service service.AllService
	Log     *zap.Logger
	Config  util.Configuration
}

func NewProductsHandler(service service.AllService, log *zap.Logger, config util.Configuration) ProductHandler {
	return ProductHandler{
		Service: service,
		Log:     log,
		Config:  config,
	}
}

func (ph *ProductHandler) GetAll(w http.ResponseWriter, r *http.Request) {

	pageStr := r.URL.Query().Get("page")
	category := r.URL.Query().Get("category")
	name := r.URL.Query().Get("name")

	page, _ := strconv.Atoi(pageStr)

	if page < 1 {
		page = 1
	}

	products, totalData, totalPage, err := ph.Service.ProductService.GetAll(page, category, name)
	if err != nil {
		ph.Log.Error("Error Handler Product: " + err.Error())
		helper.Responses(w, http.StatusNotFound, "Data tidak tersedia", nil)
		return
	}

	helper.SuccessWithPage(w, http.StatusOK, page, 10, totalPage, totalData, "Successfully", products)
}
