package main

import (
	"strconv"

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

/*
If we dont put binding and wrong json or no json is provided still it will add with default values.
*/
type todo struct {
	Id        int    `json:"id" binding:"required"`
	Task      string `json:"task" binding:"required"`
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

func addTodo(c *gin.Context) {
	var newTodo todo

	if Err := c.ShouldBindJSON(&newTodo); Err != nil {
		c.JSON(400, gin.H{"message": Err.Error()})
		return
	}

	todos = append(todos, newTodo)

	c.JSON(201, todos)

}

func markComplete(c *gin.Context) {
	newId, _ := strconv.Atoi(c.Param("id"))
	for i := 0; i < len(todos); i++ {
		if todos[i].Id == newId {
			todos[i].Completed = true
			c.JSON(202, todos)
			return
		}
	}
	c.JSON(400, gin.H{"message": "No task found"})
}

func main() {
	r := gin.Default()
	r.GET("/getTodos", getTodos)
	r.POST("/addTodo", addTodo)
	r.PATCH("/over/:id", markComplete)

	r.Run("localhost:8080")
}
