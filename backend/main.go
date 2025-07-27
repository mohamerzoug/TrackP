package main

import (
	"database/sql"
	"log"
	"net/http"
	"os"
	"strconv"
	"time"

	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
	_ "github.com/lib/pq"
)

// Models
type Project struct {
	ID          int    `json:"id" db:"id"`
	Title       string `json:"title" db:"title"`
	Description string `json:"description" db:"description"`
	CreatedAt   string `json:"created_at" db:"created_at"`
}

type Task struct {
	ID          int    `json:"id" db:"id"`
	Title       string `json:"title" db:"title"`
	Description string `json:"description" db:"description"`
	Status      string `json:"status" db:"status"`
	DueDate     string `json:"due_date" db:"due_date"`
	ProjectID   int    `json:"project_id" db:"project_id"`
	CreatedAt   string `json:"created_at" db:"created_at"`
}

var db *sql.DB

func main() {
	// Initialize database connection
	initDB()
	defer db.Close()

	// Create database tables
	createTables()

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

	log.Println("Server starting on :8080")
	r.Run(":8080")
}

func initDB() {
	// For development, use default PostgreSQL connection
	// In production, you would read these from environment variables
	dbHost := getEnv("DB_HOST", "localhost")
	dbPort := getEnv("DB_PORT", "5432")
	dbUser := getEnv("DB_USER", "postgres")
	dbPassword := getEnv("DB_PASSWORD", "040601") // Change this!
	dbName := getEnv("DB_NAME", "trackp")

	connStr := "host=" + dbHost + " port=" + dbPort + " user=" + dbUser + " password=" + dbPassword + " dbname=" + dbName + " sslmode=disable"

	var err error
	db, err = sql.Open("postgres", connStr)
	if err != nil {
		log.Fatal("Failed to connect to database:", err)
	}

	if err = db.Ping(); err != nil {
		log.Fatal("Failed to ping database:", err)
	}

	log.Println("Connected to PostgreSQL database")
}

func getEnv(key, defaultValue string) string {
	if value := os.Getenv(key); value != "" {
		return value
	}
	return defaultValue
}

func createTables() {
	projectsTable := `
	CREATE TABLE IF NOT EXISTS projects (
		id SERIAL PRIMARY KEY,
		title VARCHAR(255) NOT NULL,
		description TEXT,
		created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP
	)`

	tasksTable := `
	CREATE TABLE IF NOT EXISTS tasks (
		id SERIAL PRIMARY KEY,
		title VARCHAR(255) NOT NULL,
		description TEXT,
		status VARCHAR(50) DEFAULT 'To Do',
		due_date DATE,
		project_id INTEGER REFERENCES projects(id) ON DELETE CASCADE,
		created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP
	)`

	if _, err := db.Exec(projectsTable); err != nil {
		log.Fatal("Failed to create projects table:", err)
	}

	if _, err := db.Exec(tasksTable); err != nil {
		log.Fatal("Failed to create tasks table:", err)
	}

	log.Println("Database tables created successfully")
}

// Project handlers
func getProjects(c *gin.Context) {
	rows, err := db.Query("SELECT id, title, description, created_at FROM projects ORDER BY created_at DESC")
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	defer rows.Close()

	var projects []Project
	for rows.Next() {
		var project Project
		var createdAt time.Time
		err := rows.Scan(&project.ID, &project.Title, &project.Description, &createdAt)
		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
			return
		}
		project.CreatedAt = createdAt.Format("2006-01-02 15:04:05")
		projects = append(projects, project)
	}

	c.JSON(http.StatusOK, projects)
}

func createProject(c *gin.Context) {
	var project Project
	if err := c.ShouldBindJSON(&project); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	var id int
	err := db.QueryRow("INSERT INTO projects (title, description) VALUES ($1, $2) RETURNING id",
		project.Title, project.Description).Scan(&id)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	project.ID = id
	c.JSON(http.StatusCreated, project)
}

func getProject(c *gin.Context) {
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid project ID"})
		return
	}

	var project Project
	var createdAt time.Time
	err = db.QueryRow("SELECT id, title, description, created_at FROM projects WHERE id = $1", id).
		Scan(&project.ID, &project.Title, &project.Description, &createdAt)
	if err != nil {
		if err == sql.ErrNoRows {
			c.JSON(http.StatusNotFound, gin.H{"error": "Project not found"})
			return
		}
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	project.CreatedAt = createdAt.Format("2006-01-02 15:04:05")
	c.JSON(http.StatusOK, project)
}

func updateProject(c *gin.Context) {
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid project ID"})
		return
	}

	var project Project
	if err := c.ShouldBindJSON(&project); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	_, err = db.Exec("UPDATE projects SET title = $1, description = $2 WHERE id = $3",
		project.Title, project.Description, id)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	project.ID = id
	c.JSON(http.StatusOK, project)
}

func deleteProject(c *gin.Context) {
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid project ID"})
		return
	}

	_, err = db.Exec("DELETE FROM projects WHERE id = $1", id)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "Project deleted successfully"})
}

// Task handlers
func getProjectTasks(c *gin.Context) {
	projectID, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid project ID"})
		return
	}

	rows, err := db.Query("SELECT id, title, description, status, due_date, project_id, created_at FROM tasks WHERE project_id = $1 ORDER BY created_at DESC", projectID)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	defer rows.Close()

	var tasks []Task
	for rows.Next() {
		var task Task
		var createdAt time.Time
		var dueDate sql.NullTime
		err := rows.Scan(&task.ID, &task.Title, &task.Description, &task.Status, &dueDate, &task.ProjectID, &createdAt)
		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
			return
		}
		task.CreatedAt = createdAt.Format("2006-01-02 15:04:05")
		if dueDate.Valid {
			task.DueDate = dueDate.Time.Format("2006-01-02")
		}
		tasks = append(tasks, task)
	}

	c.JSON(http.StatusOK, tasks)
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

	task.ProjectID = projectID
	if task.Status == "" {
		task.Status = "To Do"
	}

	var id int
	var dueDate interface{}
	if task.DueDate != "" {
		dueDate = task.DueDate
	}

	err = db.QueryRow("INSERT INTO tasks (title, description, status, due_date, project_id) VALUES ($1, $2, $3, $4, $5) RETURNING id",
		task.Title, task.Description, task.Status, dueDate, task.ProjectID).Scan(&id)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	task.ID = id
	c.JSON(http.StatusCreated, task)
}

func updateTask(c *gin.Context) {
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid task ID"})
		return
	}

	var task Task
	if err := c.ShouldBindJSON(&task); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	var dueDate interface{}
	if task.DueDate != "" {
		dueDate = task.DueDate
	}

	_, err = db.Exec("UPDATE tasks SET title = $1, description = $2, status = $3, due_date = $4 WHERE id = $5",
		task.Title, task.Description, task.Status, dueDate, id)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	task.ID = id
	c.JSON(http.StatusOK, task)
}

func deleteTask(c *gin.Context) {
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid task ID"})
		return
	}

	_, err = db.Exec("DELETE FROM tasks WHERE id = $1", id)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "Task deleted successfully"})
}
