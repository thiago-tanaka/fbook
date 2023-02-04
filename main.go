package main

import (
	"fmt"
	"log"
	"net/http"
	"webapp/src/router"
	"webapp/src/utils"
)

func main() {
	utils.LoadTemplates()
	r := router.Generate()

	fmt.Println("Hello, World!")
	log.Fatal(http.ListenAndServe(":3000", r))
}
