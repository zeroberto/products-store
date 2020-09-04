package handler

import (
	"encoding/json"
	"errors"
	"fmt"
	"log"
	"net/http"
	"strconv"

	"github.com/zeroberto/products-store/products-list-endpoint/api"
	"github.com/zeroberto/products-store/products-list-endpoint/chrono"
	"github.com/zeroberto/products-store/products-list-endpoint/container"
	"github.com/zeroberto/products-store/products-list-endpoint/container/factory/apifactory"
	"github.com/zeroberto/products-store/products-list-endpoint/container/factory/chronofactory"
)

const (
	// ProductPath represents the path for products
	ProductPath = "/product/"
)

// RestHTTPHandler is responsible for providing routines with HTTP1.1 methods
type RestHTTPHandler struct {
	Container container.Container
}

// ProductRootHandle is responsible for handling http requests to the product root path
func (handler *RestHTTPHandler) ProductRootHandle(w http.ResponseWriter, r *http.Request) {
	ts := chronofactory.MakeTimeStamp(handler.Container)

	af := &apifactory.APIFactory{}
	api, _ := af.MakeProductAPI(handler.Container)

	switch r.Method {
	case http.MethodGet:
		handler.productGetRequest(api, ts, w, r)
	default:
		reportErrorMethodNotAllowed(ts, w, r)
	}
}

func (handler *RestHTTPHandler) productGetRequest(api api.ProductAPI, ts chrono.TimeStamp, w http.ResponseWriter, r *http.Request) {
	var userID int64

	if idParam := r.Header.Get("X-USER-ID"); len(idParam) > 0 {
		convertedID, err := strconv.ParseInt(idParam, 10, 64)
		if err != nil {
			reportErrorBadRequest("ID parameter is invalid", ts, w, r)
			return
		}
		userID = convertedID
	}
	response := api.Get(userID)
	respond(response, ts, w, r)
}

func respond(response api.Response, ts chrono.TimeStamp, w http.ResponseWriter, r *http.Request) {
	if response.Path != 0 {
		w.Header().Set("Location", fmt.Sprintf("%s%s%d", r.Host, r.URL.Path, 1))
	}
	if response.Body != nil {
		respondWithJSONBody(response.Body, response.Code, ts, w, r)
	} else {
		w.WriteHeader(response.Code)
	}
}

func respondWithJSONBody(body interface{}, code int, ts chrono.TimeStamp, w http.ResponseWriter, r *http.Request) {
	content, err := json.Marshal(body)
	if err != nil {
		reportError(http.StatusInternalServerError, "respondWithJSONBody", err, ts, w, r)
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(code)
	w.Write(content)
}

func reportErrorMethodNotAllowed(ts chrono.TimeStamp, w http.ResponseWriter, r *http.Request) {
	reportError(http.StatusMethodNotAllowed, "ExampleRootHandle", errors.New("Method not allowed"), ts, w, r)
}

func reportErrorBadRequest(message string, ts chrono.TimeStamp, w http.ResponseWriter, r *http.Request) {
	reportError(http.StatusBadRequest, "ExampleRootHandle", errors.New(message), ts, w, r)
}

func reportError(
	status int, method string, err error, ts chrono.TimeStamp,
	w http.ResponseWriter, r *http.Request,
) {
	log.Printf("context=httphandler, m=%s, err=%v", method, err)
	respond(createResponse(status, err.Error(), ts), ts, w, r)
}

func createResponse(code int, message string, ts chrono.TimeStamp) api.Response {
	return api.Response{
		Code: code,
		Body: api.ResponseBody{
			Code:    code,
			Message: message,
			Time:    ts.GetCurrentTime(),
		},
	}
}
