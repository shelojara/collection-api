package main

import (
	"log/slog"
	"net/http"

	"connectrpc.com/connect"
	"github.com/shelojara/collection-api/handler/rpc"
	"github.com/shelojara/collection-api/handler/rpc/interceptor"
	"github.com/shelojara/collection-api/proto/gen/v1/genv1connect"
	"golang.org/x/net/http2"
	"golang.org/x/net/http2/h2c"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

func main() {
	db, err := gorm.Open(
		postgres.Open("host=localhost user=postgres password=postgres dbname=postgres port=5432 sslmode=disable"),
		&gorm.Config{
			PrepareStmt: true,
		},
	)
	if err != nil {
		panic(err)
	}

	logger := slog.Default().With("service", "collection-api")

	mux := http.NewServeMux()
	mux.Handle(genv1connect.NewListsHandler(&rpc.Lists{DB: db},
		connect.WithInterceptors(interceptor.NewLoggerInterceptor(logger)),
	))

	http.ListenAndServe("localhost:8000", h2c.NewHandler(mux, &http2.Server{}))
}
