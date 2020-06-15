package main

import (
	"github.com/ryokuchama/gin_restapi/database"
	"github.com/ryokuchama/gin_reatapi/server"
)

func main() {
	database.Init()
	server.Init()
}
