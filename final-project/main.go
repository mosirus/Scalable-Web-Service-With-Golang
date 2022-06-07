package main

import (
	"final-project/databases"
	"final-project/routers"
)

func main() {
	databases.StartDB()
	r := routers.StartApp()

	r.Run(":4444")
}
