package main

import (
	"github.com/google/wire"
	"github.com/unyooon/prompt-sharing/controller"
	"github.com/unyooon/prompt-sharing/domain"
	"github.com/unyooon/prompt-sharing/infrastructure/db"
	"github.com/unyooon/prompt-sharing/infrastructure/repository"
	"github.com/unyooon/prompt-sharing/usecase"
)

var SuperSet = wire.NewSet(
	// controller
	controller.NewPromptController,
	controller.NewUserController,

	// usecase
	domain.NewCreatePromptInteractor,
	domain.NewReadPromptsInteractor,
	domain.NewCreateUserInteractor,

	// bind
	wire.Bind(new(usecase.PromptCreater), new(*domain.CreatePromptInteractor)),
	wire.Bind(new(usecase.PromptsReader), new(*domain.ReadPromptsInteractor)),
	wire.Bind(new(usecase.UserCreater), new(*domain.CreateUserInteractor)),

	// repository
	repository.NewLlmMasterRepository,
	repository.NewParameterMasterRepository,
	repository.NewPromptRepository,
	repository.NewUserRepository,
	wire.Bind(new(repository.LlmMasterRepositoryInterface), new(*repository.LlmMasterRepository)),
	wire.Bind(new(repository.ParameterMasterRepositoryInterface), new(*repository.ParameterMasterRepository)),
	wire.Bind(new(repository.PromptRepositoryInterface), new(*repository.PromptRepository)),
	wire.Bind(new(repository.UserRepositoryInterface), new(*repository.UserRepository)),

	//client

	// db
	db.NewDb,
	wire.Bind(new(db.DbInterface), new(*db.Db)),
)
