package main

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"net/http"
)

func main() {

	fmt.Println("Welcome to the Ping-Pong API")

	//Initializing the gin package
	router := gin.Default()

	//Endpoint we are going to use today
	router.GET("/ping", handlePing)

	//Server runs on everywhere IP and on 8080 port
	router.Run("0.0.0.0:5000")
}

//Function to handle the /ping endpoint
func handlePing(g *gin.Context) {
	g.IndentedJSON(http.StatusOK, gin.H{"message": "Pong!"})
}