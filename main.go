package main

import (
	"github.com/gin-gonic/gin"
)

/*
	type todo struct {
		id        int
		task      string
		completed bool
	}

if use small letter then it will become private and c.json will not able to export it and api response will be empty.
*/
/*
Without json tag it will also work but mainly in DB camelcase or snakecase so that is why  we explicitly provide names for it.
*/
type todo struct {
	Id        int    `json:"id"`
	Task      string `json:"task"`
	Completed bool   `json:"completed"`
}

var todos = []todo{
	{Id: 1,
		Task:      "go",
		Completed: false},
}

func getTodos(c *gin.Context) {
	c.JSON(200, todos)
}

func main() {
	r := gin.Default()
	r.GET("/getTodos", getTodos)
	r.Run("localhost:8080")
}
