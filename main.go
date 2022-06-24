package main

import (
	"errors"
	"log"
	"net/http"

	"github.com/bbcoder-gh/api/common/httputils"
	"github.com/bbcoder-gh/api/config"
	"github.com/bbcoder-gh/api/pgxhelper"
	"github.com/bbcoder-gh/api/router"
)

func main() {

	pgxhelper.Connect()
	router := router.Register()

	srv := &http.Server{
		Addr:    ":" + config.Configuration.ServerPort,
		Handler: router,
	}

	defer httputils.GracefulShutdown(srv)

	// Initializing the server in a goroutine so that
	// it won't block the graceful shutdown handling below
	go func() {
		if err := srv.ListenAndServe(); err != nil && errors.Is(err, http.ErrServerClosed) {
			log.Printf("listen: %s\n%s", srv.Addr, err)
		}
	}()

}
