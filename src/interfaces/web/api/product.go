package api

import (
	"errors"
	"github.com/gorilla/mux"
	"github.com/itsmeadi/cart/src/entities/models"
	"github.com/itsmeadi/cart/src/templatego"
	"log"
	"net/http"
	"strconv"
)

func (api *API) ProductDetail(w http.ResponseWriter, r *http.Request) {

	ctx := r.Context()

	w.Header().Set("Content-Type", "text/html")

	vars := mux.Vars(r)

	idStr := vars["id"]

	id, err := strconv.ParseInt(idStr, 10, 64)
	if err != nil || id == 0 {
		w.WriteHeader(http.StatusUnauthorized)
		errors.New("invalid category id")
	}

	products, err := api.Interactor.Product.GetProductDetailById(ctx, id)

	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	qtemplate := struct {
		Product models.Product
	}{
		Product: products,
	}
	if err := templatego.TemplateMap["product_detail"].Execute(w, qtemplate); err != nil {
		log.Printf("[ERROR] [Question] Render page error: %s\n", err)

	}

}

func (api *API) GetProductList(w http.ResponseWriter, r *http.Request) (interface{}, error) {
	ctx := r.Context()

	categoryIdStr := r.FormValue("category_id")
	categoryId, err := strconv.ParseInt(categoryIdStr, 10, 64)
	if err != nil || categoryId == 0 {
		return nil, errors.New("invalid category id")
	}

	list, err := api.Interactor.ProductByCategory.GetProductArrByCategoryId(ctx, categoryId)
	return list, err
}

func (api *API) GetProductDetail(w http.ResponseWriter, r *http.Request) (interface{}, error) {
	ctx := r.Context()

	idStr := r.FormValue("id")
	id, err := strconv.ParseInt(idStr, 10, 64)
	if err != nil || id == 0 {
		return nil, errors.New("invalid product id")
	}

	list, err := api.Interactor.Product.GetProductDetailById(ctx, id)
	return list, err
}
