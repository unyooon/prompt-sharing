package main

import (
	"github.com/gin-gonic/gin"
	swaggerFiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"
	"github.com/unyooon/prompt-sharing/controller"
	"github.com/unyooon/prompt-sharing/domain/setting"
	"github.com/unyooon/prompt-sharing/middleware"
)

type Routing struct {
	Gin                 *gin.Engine
	Setting             setting.Setting
	PromptController    *controller.PromptController
	UserController      *controller.UserController
	LlmMasterController *controller.LlmMasterController
}

func NewRouting(
	PromptController *controller.PromptController,
	UserController *controller.UserController,
	LlmMasterController *controller.LlmMasterController,
	Setting setting.Setting,
) *Routing {
	r := &Routing{
		Gin:                 gin.New(),
		Setting:             Setting,
		PromptController:    PromptController,
		UserController:      UserController,
		LlmMasterController: LlmMasterController,
	}

	r.Gin.Use(gin.Recovery())
	r.Gin.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))
	r.Gin.Use(middleware.ValidateToken(r.Setting))
	r.Gin.Use(middleware.HttpLogger)
	r.setRouting()

	return r
}

func (r *Routing) setRouting() {
	v1 := r.Gin.Group("")
	{
		v1.GET("/prompts", func(c *gin.Context) { r.PromptController.ReadPrompts(c) })
		v1.POST("/prompts", func(c *gin.Context) { r.PromptController.CreatePrompt(c) })
		v1.POST("/users", func(c *gin.Context) { r.UserController.CreateUser(c) })
		v1.GET("/llm-masters", func(c *gin.Context) { r.LlmMasterController.ReadLlmMasters(c) })
	}
}

func (r *Routing) Run() {
	r.Gin.Run(r.Setting.Port)
}
