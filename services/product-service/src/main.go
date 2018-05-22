package main

import "github.com/gin-gonic/gin"

type product struct {
	ID          string
	Title       string
	Description string
	Price       float64
}

func main() {

	router := gin.Default()

	// Respond to GET requests
	router.GET("/products", func(c *gin.Context) {

		products := []product{
			product{
				ID:          "0000-0000-0001",
				Title:       "CICD",
				Description: "CD CDE DEMO ",
				Price:       100,
			},
			product{
				ID:          "0000-0000-0002",
				Title:       "CI CD",
				Description: "CI CD DEMO bundle",
				Price:       3.75,
			},
			product{
				ID:          "0000-0000-0003",
				Title:       "CICD",
				Description: "CI CD DEMO    s",
				Price:       9.99,
			},
		}

		c.IndentedJSON(200, products)

	})

	// Serve all of the things..
	router.Run(":8001")

}
