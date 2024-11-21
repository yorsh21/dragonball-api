package service

import (
	"dragonball-api/pkg/domain"
	"dragonball-api/pkg/repository"
	"fmt"
)

type CharacterService struct {
	characterRepository repository.CharacterRepository
}

func NewCharacterService(characterRepository repository.CharacterRepository) CharacterService {
	return CharacterService{
		characterRepository: characterRepository,
	}
}

func (r CharacterService) FindOrCreate(name string) (*domain.Character, error) {
	char, err := r.characterRepository.GetFromDBByName(name)
	if err != nil {
		return nil, fmt.Errorf("error getting character from DB: %s", err.Error())
	}

	if char != nil && char.ID != 0 {
		return char, nil
	}

	char, err = r.characterRepository.GetFromAPI(name)
	if err != nil {
		return nil, fmt.Errorf("error getting character from API: %s", err.Error())
	}

	if err = r.characterRepository.CreateToDB(*char); err != nil {
		return nil, fmt.Errorf("error create character in DB: %s", err.Error())
	}

	return char, nil
}
