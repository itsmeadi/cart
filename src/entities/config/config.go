package config

import (
	"gopkg.in/gcfg.v1"
	"io/ioutil"
	"log"
	"os"
)

type Product struct {
	Url     string
	Timeout int //millisecond
}

type ProductList struct {
	Url     string
	Timeout int //millisecond
}

type ProductCache struct {
	CacheResetTimeOut int64 //second
	Timeout           int64 //second
}

type Conf struct {
	Product     Product
	PrdCache    ProductCache
	ProductList ProductList
	DB          DBConfig
}

type DBConfig struct {
	DBStr string
}

var CF *Conf

func init() {
	CF = &Conf{}
	GOPATH := os.Getenv("GOPATH")
	fname := GOPATH + "/src/github.com/itsmeadi/cart/files/config.ini"

	ok := ReadConfig(CF, fname)
	if !ok {
		log.Fatal("Failed to read config file")
	}
}

func ReadConfig(cfg *Conf, path string) bool {

	configString, err := ioutil.ReadFile(path)
	if err != nil {
		log.Println("config.go [ReadFile] function ReadConfig", err)
		return false
	}
	err = gcfg.ReadStringInto(cfg, string(configString))

	if err != nil {
		log.Println("config.go [ReadStringInto] function ReadConfig", err)
		return false
	}

	return true
}

func GetConfig() *Conf {
	return CF
}
