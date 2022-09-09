package main

import (
	"errors"
	"github.com/gin-gonic/gin"
	"net/http"
)

// tipe data json yang diperlukan
type todo struct {
	ID        string `json:"id"`
	Item      string `json:"item"`
	Completed bool   `json:"completed"`
}

// json atau data yang ingin di ubah / tinggal rubah ke database
var todos = []todo{
	{ID: "1", Item: "Clean Room", Completed: false},
	{ID: "2", Item: "Read Book", Completed: false},
	{ID: "3", Item: "Record Video", Completed: false},
}

// getTodos to handle Request
func getTodos(context *gin.Context) {
	context.IndentedJSON(http.StatusOK, todos)

}

func addTodo(context *gin.Context) {
	var newTodo todo

	if err := context.BindJSON(&newTodo); err != nil {
		return
	}

	todos = append(todos, newTodo)
	context.IndentedJSON(http.StatusCreated, newTodo)
}

func getTodo(context *gin.Context) {
	id := context.Param("id")
	todo, err := getTodoByid(id)

	if err != nil {
		context.IndentedJSON(http.StatusNotFound, gin.H{"massage": "TODO not Found"})
		return

	}

	context.IndentedJSON(http.StatusOK, todo)
}

func toggleTodoStatus(context *gin.Context) {
	id := context.Param("id")
	todo, err := getTodoByid(id)

	if err != nil {
		context.IndentedJSON(http.StatusNotFound, gin.H{"massage": "TODO not Found"})
		return

	}

	todo.Completed = !todo.Completed
	context.IndentedJSON(http.StatusOK, todo)
}

func getTodoByid(id string) (*todo, error) {
	for i, t := range todos {
		if t.ID == id {
			return &todos[i], nil
		}
	}
	return nil, errors.New("todo not found")
}

func main() {
	router := gin.Default()
	router.GET("/todos", getTodos)
	router.GET("/todos/:id", getTodo)
	router.PATCH("/todos/:id", toggleTodoStatus)
	router.POST("/todos", addTodo)
	router.Run("0.0.0.0:9099")
}
