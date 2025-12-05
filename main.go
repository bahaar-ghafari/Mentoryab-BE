package main

import "github.com/gin-gonic/gin" // it imports the Gin framework

func main() {
	r := gin.Default() // it creates a new Gin instance(router or server) with the default middleware

	r.POST("/signup", func(c *gin.Context) {
		// get the data from the request
		var input struct {
			Name  string `json:"name"`
			Email string `json:"email"`
			Stack string `json:"stack"`
		}

		if err := c.ShouldBindJSON(&input); err != nil {
			c.JSON(400, gin.H{"error": "Invalid input"}) // if the input is invalid, return a 400 status code and an error message
			return
		}
		// you can use this command to test the API in your terminal after running the server:
		// curl -X POST http://localhost:8080/signup -H "Content-Type: application/json" -d '{"name":"Ali", "email":"ali@example.com", "stack":"Golang"}'

		// we can save the data to the database with GORM (for now we assume that we will add this part later)

		// return a success message
		c.JSON(200, gin.H{"message": "Mentor registered successfully"})
	})
	//
	r.Run(":8080") // it starts the server on port 8080 (default port is 8080)

}
