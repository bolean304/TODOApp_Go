// main.go
package main

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

// Task struct represents a task in our task list
type Task struct {
	Description string `json:"description"`
}

var tasks []Task

func main() {
	// Initialize tasks slice
	tasks = make([]Task, 0)

	router := gin.Default()

	// Set up routes
	router.GET("/", showTasks)
	router.POST("/add", addTask)
	// Load HTML templates
	router.LoadHTMLGlob("templates/*")
	// Start the server
	router.Run(":8080")
}

func showTasks(c *gin.Context) {
	// Ensure tasks slice is not nil
	if tasks == nil {
		tasks = make([]Task, 0)
	}

	c.HTML(http.StatusOK, "index.html", gin.H{
		"Tasks": tasks,
	})
}

func addTask(c *gin.Context) {
	var task Task
	if err := c.ShouldBindJSON(&task); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	tasks = append(tasks, task)
	c.JSON(http.StatusOK, gin.H{"message": "Task added successfully"})
}
