package main

import (
	"fmt"
	"log"
	"package-config-yaml/database"
)

func main() {
	log.Println("Starting application")
	fmt.Println(database.QueryURL())
}
