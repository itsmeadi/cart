package api

import (
	"github.com/gorilla/mux"
	//"github.com/itsmeadi/cart/src/entities/models"
	"github.com/itsmeadi/cart/src/usecase/definitions"
	"net/http"
)

type Interactor struct {
	Cart              definitions.Cart
	CartItem          definitions.CartItems
	Product           definitions.Product
	ProductByCategory definitions.ProductByCategory
	User              definitions.User
}

type API struct {
	Interactor *Interactor
}

//New is the api initializer
func New(this *API) *API {
	return &API{Interactor: this.Interactor}
}

type HandlerFunc func(rw http.ResponseWriter, r *http.Request) (interface{}, error)

func (api *API) InitRoutes(r *mux.Router) {

	//Rest APIs
	r.HandleFunc("/product-detail/{id}", api.Wrapper(api.GetProductDetail))
	r.HandleFunc("/product-list", api.Wrapper(api.GetProductList))
	r.HandleFunc("/get-cart", api.Auth(api.Wrapper(api.GetCart)))

	r.HandleFunc("/add-to-cart", api.Auth(api.Wrapper(api.AddToCart)))
	r.HandleFunc("/remove-from-cart", api.Auth(api.Wrapper(api.RemoveFromCart)))

	r.HandleFunc("/login", api.googleLoginHandler)
	r.HandleFunc("/signout", api.logOut)
	r.HandleFunc("/auth", api.authHandler)

	//FrontEnd APIs
	r.HandleFunc("/detail", api.ProductDetail)
	r.HandleFunc("/cart", api.Auth(api.ShowCart))
	r.HandleFunc("/products", api.Auth(api.ProductList))

	r.PathPrefix("/").Handler(http.StripPrefix("/", http.FileServer(http.Dir("./frontend/vegefoods"))))

	http.Handle("/", r)

}
