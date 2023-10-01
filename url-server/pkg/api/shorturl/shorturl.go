package shorturl

import (
	"database/sql"

	v1 "url-server/pkg/api/shorturl/v1"
	"url-server/pkg/http_multiplexer"
)

func AllRoutes(
	multiplexer http_multiplexer.IMultiplexer,
	dbConnection *sql.DB,
) {
	urlService := v1.NewService(dbConnection)
	v1.Routes(multiplexer, urlService)
}
