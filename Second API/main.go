package main

import (
	"errors"
	"net/http"

	"github.com/gin-gonic/gin"
)

type todo struct {
	ID        string `json:"id"`
	Work      string `jspn:"title"`
	Completed bool   `json:"done"`
}

var todos = []todo{
	{"1", "Build API", false},
	{"2", "Build API", false},
	{"3", "Build API", false},
	{"4", "Build API", false},
	{"5", "Build API", true},
}

// RUNNING THE SERVER
func main() {
	router := gin.Default()
	router.GET("/", getTodos)
	router.GET("/:id", getTodo)
	router.POST("/addTodo", addTodo)
	router.Run("localhost:4000")

}

// GET LIST OF TODOS
func getTodos(context *gin.Context) {
	context.IndentedJSON(http.StatusOK, todos)
}

// GET ONLY ONE TODO BY ITS ID
func getTodo(context *gin.Context) {
	id := context.Param("id")
	todo, err := getTodoById(id)
	if err != nil {
		context.IndentedJSON(http.StatusNotFound, gin.H{"message": "TODO NOT FOUND"})
		return
	}
	context.IndentedJSON(http.StatusOK, todo)
}
func getTodoById(Id string) (*todo, error) {
	for i, t := range todos {
		if t.ID == Id {
			return &todos[i], nil
		}
	}
	return nil, errors.New("TODO NOT FOUND")
}

// ADD TODO IN THE LIST (POST)
func addTodo(context *gin.Context) {
	var newTodo todo

	if err := context.BindJSON(&newTodo); err != nil {
		return
	}

	todos = append(todos, newTodo)
	context.IndentedJSON(http.StatusCreated, newTodo)
}

//
