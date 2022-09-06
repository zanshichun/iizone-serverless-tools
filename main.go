package main

import (
	"iizone-serverless-tools/function"
	"net/http"

	"github.com/julienschmidt/httprouter"
)

func main() {

	router := httprouter.New()

	var iizoneTools *function.IIZoneTools

	// API地址
	router.GET("/ping", iizoneTools.PingHandler)
	router.POST("", iizoneTools.IIZoneToolsBase64Decode)
	router.POST("/api/iizone-tool/api/iizone-tools/base64-decodes/base64-encode", iizoneTools.IIZoneToolsBase64Encode)

	http.ListenAndServe(":9000", router)
}
