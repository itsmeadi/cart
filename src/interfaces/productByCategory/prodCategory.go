package productByCategory


import (
"context"
"encoding/json"
"errors"
"fmt"
	"github.com/itsme/nuc/src/entities/models"
	"io/ioutil"
	"log"
	"net/http"
"time"
)

type ProductByCategory struct {
	Url    string
	Client *http.Client
}

var ERRINVALIDRESPONSE = errors.New("Invalid Response")


func InitService(url string, timeout int) ProductByCategory {
	return ProductByCategory{
		Url: url,
		Client: &http.Client{
			Timeout: time.Duration(timeout) * time.Millisecond,
		},
	}
}

func (prod *ProductByCategory) GetProductListByCategoryId(ctx context.Context, id int64) (models.ProductListResponse, error) {

	var pRes models.ProductListResponse

	url := fmt.Sprintf("%+v/%+v", prod.Url, id)
	req, err := http.NewRequest("GET", url, nil)
	if err != nil {
		return pRes, err
	}
	res, err := prod.Client.Do(req)
	if err != nil {
		return pRes, err
	}
	if res == nil || res.StatusCode != http.StatusOK {
		return pRes, ERRINVALIDRESPONSE
	}
	body, err := ioutil.ReadAll(res.Body)
	if body == nil || res.StatusCode != http.StatusOK {
		return pRes, ERRINVALIDRESPONSE
	}
	err = json.Unmarshal(body, &pRes)
	if err != nil {
		log.Println(string(body))
		return pRes, ERRINVALIDRESPONSE
	}
	return pRes, nil
}


func (prod *ProductByCategory) GetProductArrByCategoryId(ctx context.Context, id int64) ([]models.Product, error) {

	var products []models.Product
	var pRes1 models.ProductListResponse1


	url := fmt.Sprintf("%+v/%+v", prod.Url, id)
	req, err := http.NewRequest("GET", url, nil)
	if err != nil {
		return products, err
	}
	res, err := prod.Client.Do(req)
	if err != nil {
		return products, err
	}
	if res == nil || res.StatusCode != http.StatusOK {
		return products, ERRINVALIDRESPONSE
	}
	body, err := ioutil.ReadAll(res.Body)
	if body == nil || res.StatusCode != http.StatusOK {
		return products, ERRINVALIDRESPONSE
	}
	err = json.Unmarshal(body, &pRes1)
	if err != nil {
		log.Println(string(body))
		return products, ERRINVALIDRESPONSE
	}

	collectionInterface:=pRes1.Data.Page.Layouts[1].Value.Collection

	js,err:=json.Marshal(collectionInterface)

	var collection models.Collection

	err=json.Unmarshal(js, &collection)

	products=collection.Products
	//a:=collection["products"].

	return products, err
}

