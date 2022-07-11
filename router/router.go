package router

import "github.com/julienschmidt/httprouter"

var Router = *httprouter.New()

func Register() *httprouter.Router {
	Router.POST("/user/signup", nil)
	Router.POST("/client/signup", nil)
	Router.GET("/login", nil)
	Router.GET("/otp-verification", nil)

	return &Router
}
