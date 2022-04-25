package main

import (
	"net/http"

	h "tdd.com/v1/handlers"

	"tdd.com/v1/middlewares"
	"tdd.com/v1/utils"

	"github.com/julienschmidt/httprouter"
)

func main() {
	connector := utils.InitDBConnector()
	db := connector.DB

	router := httprouter.New()

	router.GlobalOPTIONS = http.HandlerFunc(middlewares.OptionsRequestHandler)
	router.GET("/users", middlewares.MiddlewareChain(h.ListUsers(db), middlewares.RefreshRDSConnectToken(connector)))
	router.POST("/users", middlewares.MiddlewareChain(h.InsertUser(db), middlewares.RefreshRDSConnectToken(connector)))
	router.GET("/health", h.GetServiceHealth)

	http.ListenAndServe(":9000", router)
}
