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

	// usecase
	domain.NewCreatePromptInteractor,
	domain.NewReadPromptsInteractor,

	// bind
	wire.Bind(new(usecase.PromptCreater), new(*domain.CreatePromptInteractor)),
	wire.Bind(new(usecase.PromptsReader), new(*domain.ReadPromptsInteractor)),

	// repository
	repository.NewLlmMasterRepository,
	repository.NewParameterMasterRepository,
	repository.NewPromptRepository,
	wire.Bind(new(repository.LlmMasterRepositoryInterface), new(*repository.LlmMasterRepository)),
	wire.Bind(new(repository.ParameterMasterRepositoryInterface), new(*repository.ParameterMasterRepository)),
	wire.Bind(new(repository.PromptRepositoryInterface), new(*repository.PromptRepository)),

	//client

	// db
	db.NewDb,
	wire.Bind(new(db.DbInterface), new(*db.Db)),
)
