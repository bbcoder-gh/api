package main

import (
	"errors"
	"log"
	"net/http"

	"github.com/bbcoder-gh/api"
	"github.com/bbcoder-gh/api/common/httputils"
	"github.com/bbcoder-gh/api/pgxhelper"
)

func main() {

	//Connecting to Database
	conn, err := pgxhelper.Connect()
	if conn == nil {
		log.Println("Could not connect with database")
		if err != nil {
			log.Println(err)
			return
		}
	}

	srv := api.GetServer(9090)
	// Run deferred graceful shutdown routine
	defer httputils.GracefulShutdown(srv)

	// Initializing the server in a goroutine so that
	// it won't block the graceful shutdown handling below
	go func() {
		if err := srv.ListenAndServe(); err != nil && errors.Is(err, http.ErrServerClosed) {
			log.Printf("listen: %s\n%s", srv.Addr, err)
		}
	}()

}
