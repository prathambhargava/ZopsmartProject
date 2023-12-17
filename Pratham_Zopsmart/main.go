package main

import (
  "database/sql"
  "fmt"
  "net/http"

  "gofr.dev/pkg/gofr"
  "gofr.dev/pkg/log"

  "Pratham_Zopsmart/handler"
  "Pratham_Zopsmart/repository"
)

func main() {
  // Configure database connection (adapt this based on your implementation)
  db, err := database.ConfigureDB()
  if err != nil {
    log.Fatal(err)
  }

  // Create Gofr app instance
  app := gofr.New()

  // Store database connection in Gofr app context
  app.Set("db", db)

  // Create repository with dependency injection
  repo := repository.NewMySQLRepository(app)

  // Create book handler with repository dependency
  handler := handler.NewBookHandler(repo)

  // Register handler routes
  app.GET("/books", handler.GetBooks)
  app.POST("/books", handler.CreateBook)
  app.GET("/books/:id", handler.GetBook)
  app.PUT("/books/:id", handler.UpdateBook)
  app.DELETE("/books/:id", handler.DeleteBook)

  // Start server on port 8080
  fmt.Println("Server started on port 8080")
  http.ListenAndServe(":8080", app)
}
