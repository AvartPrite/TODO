package main

import (
	"fmt"
	"net/http"
	"strconv"

	"github.com/labstack/echo/v4"
)

type Todo struct {
	ID    int    `json:"id"`
	Title string `json:"title"`
}

var todos []Todo

func main() {
	e := echo.New()
	fmt.Println("Zdarova")
	e.GET("/todos", getTodos)
	e.POST("/todos/add", addTodos)
	e.DELETE("/todos/delete/:id", deleteTodos)

	e.Start(":8080")
}

func getTodos(c echo.Context) error {
	return c.JSON(http.StatusOK, todos)
}

func addTodos(c echo.Context) error {
	var newTodo Todo
	if err := c.Bind(&newTodo); err != nil {
		return c.JSON(http.StatusBadRequest, map[string]string{"error": err.Error()})
	}
	newTodo.ID = len(todos) + 1
	todos = append(todos, newTodo)
	return c.JSON(http.StatusCreated, newTodo)
}

func deleteTodos(c echo.Context) error {
	idParam := c.Param("id")
	id, err := strconv.Atoi(idParam)
	if err != nil {
		return c.JSON(http.StatusBadRequest, map[string]string{"error": err.Error()})
	}
	for i, todo := range todos {
		if todo.ID == id {
			todos = append(todos[:i], todos[i+1:]...)
			return c.NoContent(http.StatusOK)
		}
	}
	return c.NoContent(http.StatusNotFound)
}
