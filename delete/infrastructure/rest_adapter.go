package infrastructure

import (
	"net/http"

	"github.com/gin-gonic/gin"

	"learajic/code-camp/1/config"
	"learajic/code-camp/1/service"
)

type Controller struct {
	router *gin.Engine
	service service.Todo
}

func NewRestController(service service.Todo) Controller {
	c := Controller{
		router: gin.Default(),
		service: service,
	}

	c.setupRoutes()
	return c
}

func (c *Controller) setupRoutes() {
	c.router.Use(gin.Logger())
	
	c.router.GET("/", c.health)

	//TODO implement functions for each route
	v1 := c.router.Group("/v1")
	{
		v1.POST("/task", c.createNewTask)
	}

}

func(c *Controller)createNewTask(ctx *gin.Context){
	ctx.JSON(http.StatusOK, gin.H{
		"message":"todo-cc is running...",
	})
}

func (c *Controller) health(ctx *gin.Context){
	ctx.JSON(http.StatusOK, gin.H{
		"message": "todo-cc is running...",
	})
}

func (c *Controller) Run() {
	err := c.router.Run(config.SERVER_PORT)
	if err != nil {
		panic(err)
	}
}
