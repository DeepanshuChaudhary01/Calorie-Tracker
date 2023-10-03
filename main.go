package main

import (
	"os"

	// "go-react-calorie-tracker/routes"
	"go-react-calorie-tracker/server/routes"

	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
)

func main() {
	port := os.Getenv("PORT")
	//using the value of environment variable name 'PORT' on which our application listen
	if port == "" { //fall-back value if port value is empty
		port = "8000"
	}

	router := gin.New()
	//This line creates a new instance of the Gin router. Gin is a popular web framework for Go that simplifies the process of building web applications and APIs. The gin.New() function initializes a new Gin router for your application to handle incoming HTTP requests and responses.
	router.Use(gin.Logger())
	//This line adds a middleware to the Gin router. Middleware functions in Gin are used to perform operations on HTTP requests and responses before they reach the main route handlers. In this case, gin.Logger() is a built-in Gin middleware that logs information about incoming requests, such as the HTTP method, route, and response status code. This is useful for debugging and monitoring purposes.
	router.Use(cors.Default())
	//This line adds another middleware to the Gin router. The cors.Default() middleware is provided by the "github.com/gin-contrib/cors" package, and it handles Cross-Origin Resource Sharing (CORS) configuration. CORS is a security feature implemented by web browsers that controls access to resources on a web page served from a different domain.

	router.POST("/entry/create", routes.AddEntry)
	//These handler functions are in routes folder
	router.GET("/entries", routes.GetEntries)
	router.GET("/entry/:id/", routes.GetEntryById)
	router.GET("/ingredient/:ingredient", routes.GetEntriesByIngredient)

	router.PUT("/entry/update/:id", routes.UpdateEntry)
	router.PUT("/ingredient/update/:id", routes.UpdateIngredient)
	router.DELETE("/entry/delete/:id", routes.DeleteEntry)
	router.Run(":" + port)
}
