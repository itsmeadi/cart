package constants

import "errors"

var (
	StatusActive          = 1
	StatusInActive        = 0
	GetItemsLimit         = int64(100)
	ERRPRODUCTUNAVAILABLE = errors.New("sorry this product is unavailable")
)
