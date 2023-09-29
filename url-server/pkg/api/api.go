// package api

// import (
// 	"database/sql"
// 	"fmt"

// 	"github.com/motapratik/golang-rest-api-demo/config"
// 	"github.com/motapratik/golang-rest-api-demo/pkg/api/registration"
// 	"github.com/motapratik/golang-rest-api-demo/pkg/db/postgres"
// 	"github.com/motapratik/golang-rest-api-demo/pkg/middleware/jwt"

// 	"github.com/motapratik/golang-rest-api-demo/pkg/http_multiplexer"
// 	gorillamultiplexer "github.com/motapratik/golang-rest-api-demo/pkg/http_multiplexer/gorilla"
// 	"github.com/motapratik/golang-rest-api-demo/pkg/server"
// )

// type registerRoutesFn func(http_multiplexer.IMultiplexer, *sql.DB)

// // Start is used to initialize server
// func Start(cfg *config.Config) {
// 	mux := gorillamultiplexer.New(cfg.Server.Host, cfg.Server.Port)

// 	database, err := postgres.NewDatabase(cfg.DB)
// 	if err != nil {
// 		panic(fmt.Errorf("problem occured while creating database. %s", err.Error()))
// 	}

// 	dbConnection, err := database.Connect()
// 	if err != nil {
// 		panic(fmt.Errorf("problem occured while connecting to database. %s", err.Error()))
// 	}

// 	jwt, err := jwt.New(cfg.JWT.SigningKeyEnv, cfg.JWT.SigningAlgorithm, cfg.JWT.Duration)
// 	if err != nil {
// 		panic(fmt.Errorf("problem occured while creating JWT middleware. %s", err.Error()))
// 	}

// 	registerPublicRoutes(mux, dbConnection, *jwt)
// 	registerPrivateRoutes(mux, dbConnection, *jwt)

// 	mux.Use(jwt.MWFunc)

// 	srv := server.New(mux)
// 	srv.Serve()

// }

// func registerPrivateRoutes(
// 	mux http_multiplexer.IMultiplexer,
// 	dbConnection *sql.DB,
// 	jwt jwt.JWTService,
// ) {
// 	//user.AllRoutes(mux, dbConnection, jwt)
// }

// func registerPublicRoutes(
// 	mux http_multiplexer.IMultiplexer,
// 	dbConnection *sql.DB,
// 	jwt jwt.JWTService,
// ) {
// 	//login.AllRoutes(mux, dbConnection, jwt)
// 	registration.AllRoutes(mux, dbConnection)
// }