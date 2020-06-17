package main

import (
	"github.com/ryokuchama/gin_restapi/database"
	"github.com/ryokuchama/gin_restapi/server"
)

// エントリポイント
func main() {
	database.Init()
	server.Init()
}
