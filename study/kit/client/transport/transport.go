package transport

import (
	"context"
	"encoding/json"
	"golang/study/kit/server/endpoint"
	"net/http"
	"net/url"
	"strconv"
)

func EncodeRequest(c context.Context, r *http.Request, rs interface{}) error {
	var res = rs.(endpoint.UserRequest)
	r.URL.Path += "/user"
	data := url.Values{}
	data.Set("id", strconv.Itoa(res.Id))
	r.URL.RawQuery = data.Encode()
	return nil
}

func DecodeResponse(c context.Context, r *http.Response) (response interface{}, err error) {
	res := endpoint.UserResponse{}
	err = json.NewDecoder(r.Body).Decode(&res)
	return res, err
}
