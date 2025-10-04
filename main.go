package main

/*In Go, you can’t just import with "Todo/data" unless your module name is really Todo.

Imports must start with your module name (defined in go.mod), not just a folder name.*/
import (
	"Todo/data"
	"strconv"

	"github.com/gin-gonic/gin"
)

type todo = data.Todo

var todos = data.Todos

func gettodos(c *gin.Context) {
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
	r.GET("/gettodos", gettodos)
	r.POST("/addTodo", addTodo)
	r.PATCH("/over/:id", markComplete)

	r.Run("localhost:8080")
}
