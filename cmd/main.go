package main

import (
	todo_app "github.com/pythoniys/todoApp"
	"github.com/pythoniys/todoApp/pkg/handler"
	"github.com/pythoniys/todoApp/pkg/repository"
	"github.com/pythoniys/todoApp/pkg/service"
	"github.com/spf13/viper"
	"log"
)

func main() {
	if err := viper.ReadInConfig(); err != nil {
		log.Fatalf("Error reading config file, %s", err)
	}
	repo := repository.NewRepo()
	services := service.NewService(repo)
	handlers := handler.NewHandler(services)

	server := new(todo_app.Server)
	if err := server.Run(viper.GetString("8000"), handlers.InitRoutes()); err != nil {
		log.Fatalf("Failed to start server: %v", err.Error())
	}
}

func initConfig() error {
	viper.AddConfigPath("configs")
	viper.SetConfigName("config")
	return viper.ReadInConfig()
}
