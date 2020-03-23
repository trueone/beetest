package main

import (
	"fmt"

	"github.com/trueone/beetest/internal/app/api/json"
)

func main() {
	fmt.Println("Starting json api...")
	json.Start()
}
