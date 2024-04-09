package main

import (
	"github.com/gorilla/mux"
	"github.com/itsmeadi/cart/src/entities/config"
	"github.com/itsmeadi/cart/src/interfaces/db/cache"
	"github.com/itsmeadi/cart/src/interfaces/db/mysql"
	"github.com/itsmeadi/cart/src/interfaces/product"
	"github.com/itsmeadi/cart/src/interfaces/productByCategory"
	"github.com/itsmeadi/cart/src/interfaces/web/api"
	"github.com/itsmeadi/cart/src/usecase/cart"
	"github.com/itsmeadi/cart/src/usecase/product_usecase"
	"github.com/itsmeadi/cart/src/usecase/user_usecase"
	"log"
	"net/http"
)

func main() {

	conf := config.CF
	mysql.InitDb(conf.DB)

	productCache := cache.InitProductCache(conf.PrdCache.Timeout, conf.PrdCache.CacheResetTimeOut)
	productSer := product.InitService(conf.Product.Url, conf.Product.Timeout, productCache)
	productList := productByCategory.InitService(conf.ProductList.Url, conf.ProductList.Timeout)

	db := mysql.GetDb()
	cartUC := cart.InitCartUseCase(cart.CartUseCase{
		CartItemsRepo:     db,
		CartXCartItemRepo: db,
		CartRepo:          db,
		Products:          &productSer,
	})
	userUC := user_usecase.InitUsecase(db)

	productUC := product_usecase.InitProductUseCase(&productSer)
	productListUC := product_usecase.InitProductListUseCase(&productList)
	apiStr := api.API{
		Interactor: &api.Interactor{
			Cart:              &cartUC,
			Product:           &productUC,
			ProductByCategory: &productListUC,
			User:              userUC,
		},
	}
	api := api.New(&apiStr)

	api.InitRoutes(mux.NewRouter())
	log.Fatal(http.ListenAndServe(":9090", nil))
}
