package api

import (
	"context"
	"encoding/json"
	"errors"
	"log"
	"net/http"
)

type Handler func(w http.ResponseWriter, r *http.Request)

type Response struct {
	Res  interface{} `json:"response"`
	Base Base        `json:"base"`
}

type Base struct {
	Error string
}

var ErrUnAuthorized = errors.New("unauthorized")

func (api *API) Wrapper(hand func(w http.ResponseWriter, r *http.Request) (interface{}, error)) Handler {

	return Handler(
		func(w http.ResponseWriter, r *http.Request) {
			w.Header().Set("Content-Type", "application/json")

			res, err := hand(w, r)

			resp := Response{}
			if err != nil {
				//resp.Base.Error = err.Error()
				log.Println("Error in api=", err)
				resp.Base.Error = "Something Went Wrong"
			}
			resp.Res = res
			j, err := json.Marshal(resp)
			if err != nil {
				log.Println("[Error while Marshall]Error in api=", err)
			}
			_, err = w.Write(j)

			if err != nil {
				log.Println("Error while Writing Response=", err)
			}
		})
}

func GetUserId(ctx context.Context) int64 {
	userId, _ := ctx.Value("user_id").(int64)

	return userId
}
func GetUserEmail(ctx context.Context) string {
	email, _ := ctx.Value("user_email").(string)

	return email
}
