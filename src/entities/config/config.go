package config

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

func InitConfig() Conf {
	return Conf{
		Product: Product{
			Url:     `https://raw.githubusercontent.com/ntuc-social-enterprises/api-mock/master/challenge-1/product`,
			Timeout: 50000,
		},
		ProductList: ProductList{
			Url:     `https://raw.githubusercontent.com/ntuc-social-enterprises/api-mock/master/challenge-1/category`,
			Timeout: 50000,
		},
		DB: DBConfig{
			DBStr: "root:@/NTUC?parseTime=true&loc=Local",
		},
		PrdCache: ProductCache{
			Timeout:           10,
			CacheResetTimeOut: 100,
		},
	}
}
