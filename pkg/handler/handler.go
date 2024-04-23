package handler

import (
	"github.com/gin-gonic/gin"
	"github.com/pythoniys/todoApp/pkg/service"
)

type Handler struct {
	services *service.Service
}

func NewHandler(services *service.Service) *Handler {
	return &Handler{}
}

func (h *Handler) InitRoutes() *gin.Engine {
	router := gin.New()

	auth := router.Group("/auth")
	{
		auth.POST("/sign_in", h.SignIn)
		auth.POST("/sign_up", h.SignUp)
	}
	api := router.Group("/api")
	{
		todo := api.Group("/todo")
		{
			todo.GET("/", h.GetAllLists)
			todo.POST("/", h.CreateList)
			todo.GET("/:id", h.GetListById)
			todo.PUT("/:id", h.ChangeListById)
			todo.DELETE("/:id", h.DeleteList)

			items := todo.Group("/items")
			{
				items.GET("/", h.GetAllItems)
				items.POST("/", h.CreateItem)
				items.GET("/:item_id", h.GetItemsById)
				items.PUT("/:item_id", h.ChangeItemById)
				items.DELETE("/:item_id", h.DeleteItem)
			}
		}
	}
	return router
}
