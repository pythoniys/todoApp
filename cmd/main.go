package main

import (
	todo_app "github.com/pythoniys/todoApp"
	"github.com/pythoniys/todoApp/pkg/handler"
	"log"
)

func main() {
	handlers := new(handler.Handler)

	server := new(todo_app.Server)
	if err := server.Run(":8080", handlers.InitRoutes()); err != nil {
		log.Fatalf("Failed to start server: %v", err.Error())
	}
}
