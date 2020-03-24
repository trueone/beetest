package api

import (
	"net"

	"github.com/labstack/echo"
	"github.com/labstack/gommon/log"
	"google.golang.org/grpc"

	"github.com/trueone/beetest/internal/app/secret"
)

func JSONStart() {
	e := echo.New()

	secretRegistry := secret.New()
	secretRegistry.RegisterJSONHandlers(e)

	e.Logger.Fatal(e.Start(":8080"))
}

func GRPCStart() {
	lis, err := net.Listen("tcp", ":8081")
	if err != nil {
		log.Fatal(err)
	}

	server := grpc.NewServer()

	secretRegistry := secret.New()
	secretRegistry.RegisterGRPCHandlers(server)

	log.Fatal(server.Serve(lis))
}
