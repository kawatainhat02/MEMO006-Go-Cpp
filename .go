package main

import "fmt"

import "rsc.io/quote"

func main() {
    fmt.Println(quote.Go())
}

$ go mod tidy
go: finding module for package rsc.io/quote
go: found rsc.io/quote in rsc.io/quote v1.5.2


$ go run .
Don't communicate by sharing memory, share memory by communicating.

type atomicInt struct {
    value int
    lock chan int = make(chan int, 1)
}

func (n *atomicInt) Write(value int) {
    n.lock <- 1
    defer func() { <- n.lock } ()
    n.value = value
}

package route

import (
	"net/http"

	"github.com/rs/zerolog"

	"github.com/me/myproject/api/response"
	"github.com/me/myproject/logging"
)

func GetAppVersion(appVersion string, logger zerolog.Logger) http.HandlerFunc {
	return func(responseWriter http.ResponseWriter, request *http.Request) {
		response.SetHeaderContentTypeToJson(responseWriter)

		if appVersion == "" {
                        // construct a failure struct with some details
			apiResponse := response.NewFailureApiResponse("APP_VERSION_UNAVAILABLE", "The app version is unavailable.")
			err := response.WriteJsonResponse(responseWriter, apiResponse)

			if err != nil {
				logging.LogError(logger, err)

                                // construct an error struct with some details
				apiResponse = response.NewErrorApiResponse()
				err = response.WriteJsonResponse(responseWriter, apiResponse)

				if err != nil {
					logging.LogError(logger, err)
				}
			}

			return
		}
 
                // construct a success struct
		apiResponse := response.NewSuccessApiResponse(appVersion)
		err := response.WriteJsonResponse(responseWriter, apiResponse)

		if err != nil {
			logging.LogError(logger, err)

                        // construct an error struct with some details
			apiResponse := response.NewErrorApiResponse()
			err = response.WriteJsonResponse(responseWriter, apiResponse)

			if err != nil {
				logging.LogError(logger, err)
			}
		}
	}
}
