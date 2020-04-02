// +build !wasm

package main

import (
	"net/http"

	"github.com/maxence-charriere/go-app/v6/pkg/app"
	"github.com/maxence-charriere/go-app/v6/pkg/log"
)

const port = "8080"

func main() {

	err := http.ListenAndServe(":"+port, &app.Handler{
		Title: "oschina",
	})

	if err != nil {
		log.Error("server crashed").T("error", err)
	}
}
