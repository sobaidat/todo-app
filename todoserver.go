package main

import (
	"os"
	"strings"

	"github.com/labstack/echo"
	"github.com/labstack/echo/middleware"
	"github.com/sobaidat/todoapp/handlers"
	"github.com/sobaidat/todoapp/utils"
)

func getServerPort() string {
	port := "8000"
	if len(os.Args) > 1 {
		port = strings.TrimSpace(os.Args[1])
	}
	return port
}

func main() {
	db := utils.InitDB("storage.db")
	// create db main table if not already created
	utils.Migrate(db)

	e := echo.New()

	//e.File("/", "public/index.html")
	e.GET("/tasks", handlers.GetTasks(db))
	e.PUT("/tasks", handlers.PutTask(db))
	e.DELETE("/tasks/:id", handlers.DeleteTask(db))

	e.Use(middleware.CORSWithConfig(middleware.CORSConfig{
		AllowCredentials: true,
		AllowMethods:     []string{echo.GET, echo.PUT, echo.DELETE},
		AllowOrigins:     []string{"*"},
	}))

	port := getServerPort()
	e.Logger.Fatal(e.Start(":" + port))
}
