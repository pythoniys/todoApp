package main

import (
	todo_app "github.com/pythoniys/todoApp"
	"log"
)

func main() {
	server := new(todo_app.Server)
	if err := server.Run(":8080"); err != nil {
		log.Fatalf("Failed to start server: %v", err)
	}

}
