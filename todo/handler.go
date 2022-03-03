package todo

import (
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)

type Handler interface {
	GetAll(c *gin.Context)
	GetById(c *gin.Context)
	AddTodo(c *gin.Context)
	UpdateTodo(c *gin.Context)
	ToggleIsCompleteTodo(c *gin.Context)
	DeleteTodo(c *gin.Context)
}

type handler struct {
	service Service
}

func CreateNewHandler(service Service) Handler {
	return &handler{service}
}

func (h *handler) GetAll(c *gin.Context) {
	todos, err := h.service.GetAll()
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"message": err.Error()})
		return
	}

	responses := convertTodosToResponses(todos)
	c.JSON(http.StatusOK, gin.H{"data": responses})
}

func (h *handler) GetById(c *gin.Context) {
	id, _ := strconv.Atoi(c.Param("id"))
	todo, err := h.service.GetById(id)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"message": err.Error()})
		return
	}

	response := convertTodoToResponse(todo)
	c.JSON(http.StatusOK, gin.H{"data": response})
}

func (h *handler) AddTodo(c *gin.Context) {
	var request Request
	err := c.ShouldBindJSON(&request)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"message": err.Error()})
		return
	}

	todo, err := h.service.AddTodo(request)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"message": err.Error()})
		return
	}

	response := convertTodoToResponse(todo)
	c.JSON(http.StatusOK, gin.H{"data": response})
}

func (h *handler) UpdateTodo(c *gin.Context) {
	id, _ := strconv.Atoi(c.Param("id"))

	var request Request
	err := c.ShouldBindJSON(&request)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"message": err.Error()})
		return
	}

	todo, err := h.service.UpdateTodo(id, request)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"message": err.Error()})
		return
	}

	response := convertTodoToResponse(todo)
	c.JSON(http.StatusOK, gin.H{"data": response})
}

func (h *handler) ToggleIsCompleteTodo(c *gin.Context) {
	id, _ := strconv.Atoi(c.Param("id"))

	todo, err := h.service.ToggleIsCompleteTodo(id)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"message": err.Error()})
		return
	}

	response := convertTodoToResponse(todo)
	c.JSON(http.StatusOK, gin.H{"data": response})
}

func (h *handler) DeleteTodo(c *gin.Context) {
	id, _ := strconv.Atoi(c.Param("id"))
	todo, err := h.service.DeleteTodo(id)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"message": err.Error()})
		return
	}

	response := convertTodoToResponse(todo)
	c.JSON(http.StatusOK, gin.H{"data": response})
}

func convertTodoToResponse(todo Todo) Response {
	return Response{todo.ID, todo.Title, todo.IsCompleted}
}

func convertTodosToResponses(todos []Todo) []Response {
	var responses []Response
	for _, todo := range todos {
		responses = append(responses, convertTodoToResponse(todo))
	}

	return responses
}
