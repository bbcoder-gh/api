package api

import (
	"net/http"

	"github.com/bbcoder-gh/api/config"
	"github.com/bbcoder-gh/api/router"
)

func GetServer() (srv *http.Server) {
	router := router.Register()

	//Configuring the web server
	srv = &http.Server{
		Addr:    ":" + config.Configuration.ServerPort,
		Handler: router,
	}

	return
}
