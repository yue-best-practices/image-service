package main

import (
	"github.com/urfave/negroni"
	"github.com/gernest/alien"
	"net/http"
	"fmt"
)

func StartServer(){
	api:=negroni.New()
	router:=alien.New()
	router.Post("/upload",upload)
	router.Get("/:imgEncode",image)

	api.UseHandler(router)
	http.ListenAndServe(fmt.Sprintf(":%d",APP_PORT),api)
}