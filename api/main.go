package main

import (
	"github.com/ryokuchama/gin_restapi/api/database"
	"github.com/ryokuchama/gin_restapi/api/server"
)

func main() {
	database.Init()
	server.Init()
}
