package main

import (
	"log"
	"net/http"
	"strconv"
	"sync"
	"time"

	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
)

// Models (same as main.go)
type Project struct {
	ID          int    `json:"id"`
	Title       string `json:"title"`
	Description string `json:"description"`
	CreatedAt   string `json:"created_at"`
}

type Task struct {
	ID          int    `json:"id"`
	Title       string `json:"title"`
	Description string `json:"description"`
	Status      string `json:"status"`
	DueDate     string `json:"due_date"`
	ProjectID   int    `json:"project_id"`
	CreatedAt   string `json:"created_at"`
}

// In-memory storage
var (
	projects     []Project
	tasks        []Task
	projectIDSeq = 1
	taskIDSeq    = 1
	mu           sync.RWMutex
)

func main() {
	// Initialize demo data
	initDemoData()

	// Initialize Gin router
	r := gin.Default()

	// Configure CORS
	config := cors.DefaultConfig()
	config.AllowOrigins = []string{"http://localhost:3000"}
	config.AllowMethods = []string{"GET", "POST", "PUT", "DELETE", "OPTIONS"}
	config.AllowHeaders = []string{"Origin", "Content-Type", "Accept", "Authorization"}
	r.Use(cors.New(config))

	// API routes
	api := r.Group("/api")
	{
		// Project routes
		api.GET("/projects", getProjects)
		api.POST("/projects", createProject)
		api.GET("/projects/:id", getProject)
		api.PUT("/projects/:id", updateProject)
		api.DELETE("/projects/:id", deleteProject)

		// Task routes
		api.GET("/projects/:id/tasks", getProjectTasks)
		api.POST("/projects/:id/tasks", createTask)
		api.PUT("/tasks/:id", updateTask)
		api.DELETE("/tasks/:id", deleteTask)
	}

	log.Println("TrackP Demo Server starting on :8080")
	log.Println("Note: Using in-memory storage - data will be lost on restart")
	r.Run(":8080")
}

func initDemoData() {
	now := time.Now().Format("2006-01-02 15:04:05")

	projects = []Project{
		{ID: 1, Title: "Sample Project", Description: "This is a sample project to get you started with TrackP", CreatedAt: now},
		{ID: 2, Title: "Website Redesign", Description: "Complete redesign of the company website with modern UI/UX", CreatedAt: now},
	}

	tasks = []Task{
		{ID: 1, Title: "Set up development environment", Description: "Install all necessary tools and dependencies", Status: "Done", ProjectID: 1, CreatedAt: now},
		{ID: 2, Title: "Create database schema", Description: "Design and implement the database structure", Status: "Done", ProjectID: 1, CreatedAt: now},
		{ID: 3, Title: "Implement user interface", Description: "Build the React frontend components", Status: "In Progress", ProjectID: 1, CreatedAt: now},
		{ID: 4, Title: "API development", Description: "Create RESTful API endpoints", Status: "To Do", ProjectID: 1, CreatedAt: now},
		{ID: 5, Title: "Research design trends", Description: "Look into current web design trends and best practices", Status: "To Do", ProjectID: 2, CreatedAt: now},
		{ID: 6, Title: "Create wireframes", Description: "Design the layout and structure of new pages", Status: "To Do", ProjectID: 2, CreatedAt: now},
	}

	projectIDSeq = 3
	taskIDSeq = 7
}

// Project handlers
func getProjects(c *gin.Context) {
	mu.RLock()
	defer mu.RUnlock()
	c.JSON(http.StatusOK, projects)
}

func createProject(c *gin.Context) {
	var project Project
	if err := c.ShouldBindJSON(&project); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	mu.Lock()
	project.ID = projectIDSeq
	project.CreatedAt = time.Now().Format("2006-01-02 15:04:05")
	projects = append([]Project{project}, projects...)
	projectIDSeq++
	mu.Unlock()

	c.JSON(http.StatusCreated, project)
}

func getProject(c *gin.Context) {
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid project ID"})
		return
	}

	mu.RLock()
	defer mu.RUnlock()

	for _, project := range projects {
		if project.ID == id {
			c.JSON(http.StatusOK, project)
			return
		}
	}

	c.JSON(http.StatusNotFound, gin.H{"error": "Project not found"})
}

func updateProject(c *gin.Context) {
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid project ID"})
		return
	}

	var updateData Project
	if err := c.ShouldBindJSON(&updateData); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	mu.Lock()
	defer mu.Unlock()

	for i, project := range projects {
		if project.ID == id {
			projects[i].Title = updateData.Title
			projects[i].Description = updateData.Description
			updateData.ID = id
			updateData.CreatedAt = project.CreatedAt
			c.JSON(http.StatusOK, updateData)
			return
		}
	}

	c.JSON(http.StatusNotFound, gin.H{"error": "Project not found"})
}

func deleteProject(c *gin.Context) {
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid project ID"})
		return
	}

	mu.Lock()
	defer mu.Unlock()

	// Remove project
	for i, project := range projects {
		if project.ID == id {
			projects = append(projects[:i], projects[i+1:]...)
			break
		}
	}

	// Remove associated tasks
	var filteredTasks []Task
	for _, task := range tasks {
		if task.ProjectID != id {
			filteredTasks = append(filteredTasks, task)
		}
	}
	tasks = filteredTasks

	c.JSON(http.StatusOK, gin.H{"message": "Project deleted successfully"})
}

// Task handlers
func getProjectTasks(c *gin.Context) {
	projectID, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid project ID"})
		return
	}

	mu.RLock()
	defer mu.RUnlock()

	var projectTasks []Task
	for _, task := range tasks {
		if task.ProjectID == projectID {
			projectTasks = append(projectTasks, task)
		}
	}

	c.JSON(http.StatusOK, projectTasks)
}

func createTask(c *gin.Context) {
	projectID, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid project ID"})
		return
	}

	var task Task
	if err := c.ShouldBindJSON(&task); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	mu.Lock()
	task.ID = taskIDSeq
	task.ProjectID = projectID
	task.CreatedAt = time.Now().Format("2006-01-02 15:04:05")
	if task.Status == "" {
		task.Status = "To Do"
	}
	tasks = append([]Task{task}, tasks...)
	taskIDSeq++
	mu.Unlock()

	c.JSON(http.StatusCreated, task)
}

func updateTask(c *gin.Context) {
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid task ID"})
		return
	}

	var updateData Task
	if err := c.ShouldBindJSON(&updateData); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	mu.Lock()
	defer mu.Unlock()

	for i, task := range tasks {
		if task.ID == id {
			if updateData.Title != "" {
				tasks[i].Title = updateData.Title
			}
			if updateData.Description != "" {
				tasks[i].Description = updateData.Description
			}
			if updateData.Status != "" {
				tasks[i].Status = updateData.Status
			}
			if updateData.DueDate != "" {
				tasks[i].DueDate = updateData.DueDate
			}
			updateData = tasks[i]
			c.JSON(http.StatusOK, updateData)
			return
		}
	}

	c.JSON(http.StatusNotFound, gin.H{"error": "Task not found"})
}

func deleteTask(c *gin.Context) {
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid task ID"})
		return
	}

	mu.Lock()
	defer mu.Unlock()

	for i, task := range tasks {
		if task.ID == id {
			tasks = append(tasks[:i], tasks[i+1:]...)
			c.JSON(http.StatusOK, gin.H{"message": "Task deleted successfully"})
			return
		}
	}

	c.JSON(http.StatusNotFound, gin.H{"error": "Task not found"})
}
