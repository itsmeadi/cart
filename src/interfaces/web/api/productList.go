package api

import (
	"github.com/itsmeadi/cart/src/entities/models"
	"github.com/itsmeadi/cart/src/templatego"
	"log"
	"net/http"
	"strconv"
)

func (api *API) ProductList(w http.ResponseWriter, r *http.Request) {

	ctx := r.Context()

	//r.se
	w.Header().Set("Content-Type", "text/html")

	categoryIdStr := r.FormValue("category_id")
	categoryId, err := strconv.ParseInt(categoryIdStr, 10, 64)
	if err != nil || categoryId == 0 {
		categoryId = 3
		//AbortWithError(http.StatusBadRequest, errors.New("invalid category id"), &w)
		//return
	}

	products, err := api.Interactor.ProductByCategory.GetProductArrByCategoryId(ctx, categoryId)

	if err != nil {
		log.Println("[API][ProductList][GetProductArrByCategoryId] Error=", err)
		w.WriteHeader(http.StatusInternalServerError)
		return
	}
	email := GetUserEmail(ctx)
	qtemplate := struct {
		Response  []models.Product
		UserEmail string
	}{
		Response:  products,
		UserEmail: email,
	}

	if err := templatego.TemplateMap["index"].Execute(w, qtemplate); err != nil {
		log.Printf("[ERROR] [Question] Render page error: %s\n", err)

	}

}
