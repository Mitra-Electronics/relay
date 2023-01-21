package main

import (
	"github.com/Mitra-Electronics/relay/account"
	"github.com/gin-gonic/gin"
)

func main() {
	m := gin.Default()
	account.Router(m)
	m.Run(":8080")
}
