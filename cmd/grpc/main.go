package main

import (
	"fmt"

	"github.com/trueone/beetest/internal/app/api"
)

func main() {
	fmt.Println("Starting grpc...")
	api.GRPCStart()
}
