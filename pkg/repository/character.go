package repository

import (
	"database/sql"
	"dragonball-api/pkg/domain"
	"dragonball-api/pkg/utils"
	"encoding/json"
	"errors"
	"fmt"
	"io"
	"net/http"
	"net/url"
)

type CharacterRepository struct {
	DB        *sql.DB
	ApiUrl    string
	TableName string
}

func NewCharacterRepository(db *sql.DB, apiUrl string) CharacterRepository {
	return CharacterRepository{
		DB:        db,
		ApiUrl:    apiUrl,
		TableName: "characters",
	}
}

func (r CharacterRepository) GetFromAPI(name string) (*domain.Character, error) {
	u, err := url.Parse(r.ApiUrl)
	if err != nil {
		return nil, err
	}

	q := u.Query()
	q.Set("name", name)
	u.RawQuery = q.Encode()

	resp, err := http.Get(u.String())
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	body, err := io.ReadAll(resp.Body)
	if err != nil {
		return nil, err
	}

	var data []domain.Character
	if err := json.Unmarshal(body, &data); err != nil {
		return nil, err
	}

	if len(data) > 0 {
		return &data[0], nil
	}

	return nil, errors.New("character not found")
}

func (r CharacterRepository) GetFromDBByName(name string) (*domain.Character, error) {
	q := fmt.Sprintf("SELECT * FROM %s WHERE name LIKE ?", r.TableName)
	seachName := "%" + utils.CapitalizeWords(name) + "%"
	rows, err := r.DB.Query(q, seachName)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var characters []domain.Character
	for rows.Next() {
		var character domain.Character
		err = rows.Scan(&character.ID, &character.Name, &character.Ki)
		if err != nil {
			return nil, err
		}
		characters = append(characters, character)
	}

	if len(characters) > 0 {
		return &characters[0], nil
	}

	return nil, nil
}

func (r CharacterRepository) CreateToDB(character domain.Character) error {
	q := fmt.Sprintf("INSERT INTO %s (id, name, ki) VALUES(?, ?, ?)", r.TableName)
	stmt, err := r.DB.Prepare(q)
	if err != nil {
		return err
	}
	defer stmt.Close()

	_, err = stmt.Exec(character.ID, character.Name, character.Ki)
	if err != nil {
		return err
	}

	return nil
}
