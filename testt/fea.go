package main

import (
	"fmt"

	"go.uber.org/zap"
)

func main() {
	fmt.Println("fea")
	logger, _ := zap.NewProduction()
	logger.Warn("eat")

}
