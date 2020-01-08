package api

import (
	"errors"
	"github.com/itsmeadi/cart/src/entities/constants"
	"github.com/itsmeadi/cart/src/entities/models"
	"github.com/itsmeadi/cart/src/templatego"
	"log"
	"net/http"
	"strconv"
)

func (api *API) AddToCart(w http.ResponseWriter, r *http.Request) (interface{}, error) {
	ctx := r.Context()

	userId := GetUserId(ctx)
	if userId == 0 {
		w.WriteHeader(http.StatusUnauthorized)
		return nil, ErrUnAuthorized
	}
	productIdStr := r.FormValue("product_id")
	productId, err := strconv.ParseInt(productIdStr, 10, 64)
	if err != nil || productId == 0 {
		return nil, errors.New("invalid product id")
	}

	qtyStr := r.FormValue("quantity")
	qty, err := strconv.ParseInt(qtyStr, 10, 64)
	if err != nil || qty == 0 {
		qty = 1
	}
	//productId = 4		//TODO

	var errMsg string
	cartUpdateStr := r.FormValue("cart_update")
	cartUpdate, err := strconv.ParseInt(cartUpdateStr, 10, 64)
	if cartUpdate != 0 {
		err = api.Interactor.Cart.UpdateCart(ctx, userId, productId, qty)
	} else {
		err = api.Interactor.Cart.AddToCart(ctx, userId, productId, qty)
	}
	if err != nil && err == constants.ERRPRODUCTUNAVAILABLE {
		errMsg = err.Error()
	} else if err != nil {
		log.Println("[API][AddToCart][AddToCart]", err)
	}

	http.Redirect(w, r, "/cart?err_msg="+errMsg, http.StatusSeeOther)
	return nil, err
}

func (api *API) RemoveFromCart(w http.ResponseWriter, r *http.Request) (interface{}, error) {
	ctx := r.Context()

	userId := GetUserId(ctx)
	if userId == 0 {
		w.WriteHeader(http.StatusUnauthorized)
		return nil, ErrUnAuthorized
	}
	productIdStr := r.FormValue("product_id")
	productId, err := strconv.ParseInt(productIdStr, 10, 64)
	if err != nil || productId == 0 {
		return nil, errors.New("invalid product id")
	}

	err = api.Interactor.Cart.RemoveProductFromCart(ctx, productId, userId)

	if err != nil {
		log.Println("[API][RemoveFromCart][RemoveProductFromCart]", err)
	}
	http.Redirect(w, r, "/cart", http.StatusSeeOther)
	return nil, err
}

func (api *API) GetCart(w http.ResponseWriter, r *http.Request) (interface{}, error) {
	ctx := r.Context()
	userId := GetUserId(ctx)
	if userId == 0 {
		w.WriteHeader(http.StatusUnauthorized)
		return nil, ErrUnAuthorized
	}

	cart, err := api.Interactor.Cart.GetCart(ctx, userId)
	return models.CartResponse{Cart: cart}, err
}

func (api *API) ShowCart(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "text/html")

	ctx := r.Context()
	userId := GetUserId(ctx)
	email := GetUserEmail(ctx)
	if userId == 0 {
		w.WriteHeader(http.StatusUnauthorized)
		return
	}

	cart, err := api.Interactor.Cart.GetCart(ctx, userId)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	qtemplate := struct {
		Cart      models.CartDetail
		UserEmail string
		ErrorMsg  string
	}{
		Cart:      cart,
		UserEmail: email,
		ErrorMsg:  r.FormValue("err_msg"),
	}
	if err := templatego.TemplateMap["cart"].Execute(w, qtemplate); err != nil {
		log.Printf("[ERROR] [Question] Render page error: %s\n", err)

	}
}
