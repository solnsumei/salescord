package main

import (
	"fmt"

	"github.com/solnsumei/api-starter-template/initializers"
)

func init() {
	initializers.LoadEnvVariables()
}

func main() {
	fmt.Println("Welcome to API Starter template")
}
