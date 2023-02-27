package api

import (
	"context"
	"encoding/json"
	"errors"
	"net/http"
	"strconv"
)

func DecodeUserRequest(c context.Context, r *http.Request) (interface{}, error) {
	if r.URL.Query().Get("uid") != "" {
		uid, _ := strconv.Atoi(r.URL.Query().Get("uid"))
		return UserRequest{
			UserID: uid,
		}, nil
	}
	return nil, errors.New("参数错误")
}

func EncodeUserResponse(c context.Context, w http.ResponseWriter, response interface{}) error {
	return json.NewEncoder(w).Encode(response)
}
