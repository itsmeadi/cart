package models

type Product struct { //ProductDetail
	ID     int64    `json:"id"`
	Images []string `json:"images"`
	Name   string   `json:"name"`
	Slug   string   `json:"slug"`
	Status string   `json:"status"`
	Price  float64  `json:"price"`
	Offers []struct {
		Price float64 `json:"price"`
	} `json:"offers"`
}

type ProductDetailResponse struct {
	Code int `json:"code"`
	Data struct {
		Page struct {
			Entity struct {
				Product
			} `json:"entity"`
		} `json:"page"`
	} `json:"data"`
	Status string `json:"status"`
}

type ProductListResponse struct {
	Code int `json:"code"`
	Data struct {
		Page struct {
			Layouts []struct {
				Value struct {
					Collection []struct {
						Products []Product   `json:"product"`
						Category interface{} `json:"category"`
					} `json:"collection"`
				} `json:"value"`
			} `json:"layouts"`
		} `json:"page"`
	} `json:"data"`
	Status string `json:"status"`
}

type ProductListResponse1 struct {
	Code int `json:"code"`
	Data struct {
		Page struct {
			Layouts []struct {
				Value struct {
					Collection interface{} `json:"collection"`
				} `json:"value"`
			} `json:"layouts"`
		} `json:"page"`
	} `json:"data"`
	Status string `json:"status"`
}
type ProductListResponse2 struct {
	Code int `json:"code"`
	Data struct {
		Page struct {
			Layouts []struct {
				Value struct {
					Collection []struct {
						Products []Product   `json:"product"`
						Category interface{} `json:"category"`
					} `json:"collection"`
				} `json:"value"`
			} `json:"layouts"`
		} `json:"page"`
	} `json:"data"`
	Status string `json:"status"`
}

type Collection struct {
	Count interface {
	} `json:"count"`
	Filters interface {
	} `json:"filters"`
	Limit interface {
	} `json:"limit"`
	Offset interface {
	} `json:"offset"`
	Products []Product `json:"product"`
}

type UserGoogle struct {
	Sub           string `json:"sub"`
	Picture       string `json:"picture"`
	Email         string `json:"email"`
	EmailVerified bool   `json:"email_verified"`
}
