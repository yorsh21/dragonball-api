package initialize

import (
	"database/sql"

	"dragonball-api/pkg/http/handler"
	"dragonball-api/pkg/repository"
	"dragonball-api/pkg/service"
)

const dragonballApiUrl = "https://dragonball-api.com/api/characters"

type Resources struct {
	CharacterHandler    handler.CharacterHandler
	CharacterService    service.CharacterService
	CharacterRepository repository.CharacterRepository
}

func InitResources(db *sql.DB) Resources {
	characterRepository := repository.NewCharacterRepository(db, dragonballApiUrl)
	characterService := service.NewCharacterService(characterRepository)
	characterHandler := handler.NewCharacterHandler(characterService)

	return Resources{
		CharacterHandler:    characterHandler,
		CharacterService:    characterService,
		CharacterRepository: characterRepository,
	}
}
