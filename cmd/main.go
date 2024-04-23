package main

import (
	todo_app "github.com/pythoniys/todoApp"
	"github.com/pythoniys/todoApp/pkg/handler"
	"github.com/pythoniys/todoApp/pkg/repository"
	"github.com/pythoniys/todoApp/pkg/service"
	"log"
)

func main() {
	repo := repository.NewRepo()
	services := service.NewService(repo)
	handlers := handler.NewHandler(services)

	server := new(todo_app.Server)
	if err := server.Run("8000", handlers.InitRoutes()); err != nil {
		log.Fatalf("Failed to start server: %v", err.Error())
	}
}
