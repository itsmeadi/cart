package product

import (
	"context"
	"encoding/json"
	"errors"
	"fmt"
	"github.com/itsmeadi/cart/src/entities/models"
	"github.com/itsmeadi/cart/src/interfaces/db/cache"
	"io/ioutil"
	"net/http"
	"time"
)

type Product struct {
	Url    string
	Client *http.Client
	Cache  *cache.ProductCache
}

var ERRINVALIDRESPONSE = errors.New("invalid response")

func InitService(url string, timeout int, cache *cache.ProductCache) Product {
	return Product{
		Url: url,
		Client: &http.Client{
			Timeout: time.Duration(timeout) * time.Millisecond,
		},
		Cache: cache,
	}
}

func (prod *Product) GetProductDetailById(ctx context.Context, id int64) (models.Product, error) {

	var pRes models.ProductDetailResponse
	var product models.Product

	product = prod.Cache.GetCache(ctx, id)
	if product.ID != 0 {
		return product, nil
	}

	url := fmt.Sprintf("%+v/%+v", prod.Url, id)
	req, err := http.NewRequest("GET", url, nil)
	if err != nil {
		return product, err
	}
	res, err := prod.Client.Do(req)
	if err != nil {
		return product, err
	}
	if res == nil || res.StatusCode != http.StatusOK {
		return product, ERRINVALIDRESPONSE
	}
	body, err := ioutil.ReadAll(res.Body)
	if body == nil || res.StatusCode != http.StatusOK {
		return product, ERRINVALIDRESPONSE
	}
	err = json.Unmarshal(body, &pRes)
	if err != nil {
		return product, ERRINVALIDRESPONSE
	}

	product = pRes.Data.Page.Entity.Product

	go prod.Cache.SaveCache(ctx, product)
	//pRes.Data.Page.Entity.Product.Price=3.4
	return product, nil
}
