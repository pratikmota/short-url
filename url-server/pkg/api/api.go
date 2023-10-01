package api

import (
	"database/sql"
	"fmt"

	"url-server/config"
	"url-server/pkg/api/shorturl"
	"url-server/pkg/db/postgres"

	"url-server/pkg/http_multiplexer"
	gorilla_multiplexer "url-server/pkg/http_multiplexer/gorilla"
	"url-server/pkg/server"
)

type registerRoutesFn func(http_multiplexer.IMultiplexer, *sql.DB)

// Start is used to initialize server
func Start(cfg *config.Config) {
	mux := gorilla_multiplexer.New(cfg.Server.Host, cfg.Server.Port)

	database, err := postgres.NewDatabase(cfg.DB)
	if err != nil {
		panic(fmt.Errorf("problem occured while creating database. %s", err.Error()))
	}

	dbConnection, err := database.Connect()
	if err != nil {
		panic(fmt.Errorf("problem occured while connecting to database. %s", err.Error()))
	}

	registerPublicRoutes(mux, dbConnection)
	srv := server.New(mux)
	srv.Serve()

}

func registerPublicRoutes(
	mux http_multiplexer.IMultiplexer,
	dbConnection *sql.DB,
) {
	shorturl.AllRoutes(mux, dbConnection)
}
