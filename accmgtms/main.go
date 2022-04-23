package main

import (
	"fmt"

	"github.com/ashokelumalai1997/TravelPartner/clientcommon"
	"github.com/gin-gonic/gin"
)

func main() {
	fmt.Println("Hey, there!")
	port := clientcommon.CliPort(3000)
	router := gin.New()
	router.run(port)
}
