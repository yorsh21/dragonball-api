package repository_test

import (
	"database/sql"
	"dragonball-api/pkg/domain"
	"dragonball-api/pkg/repository"
	"dragonball-api/pkg/utils"
	"encoding/json"
	"fmt"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/DATA-DOG/go-sqlmock"
	"github.com/stretchr/testify/assert"
)

func TestGetFromAPI(t *testing.T) {
	t.Run("successfully fetches character from API", func(t *testing.T) {
		character := domain.Character{
			ID:   1,
			Name: "Goku",
		}
		mockResponse, _ := json.Marshal([]domain.Character{character})

		server := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			assert.Equal(t, "/?name=Goku", r.URL.String())
			w.WriteHeader(http.StatusOK)
			w.Write(mockResponse)
		}))
		defer server.Close()

		r := repository.NewCharacterRepository(nil, server.URL)

		result, err := r.GetFromAPI("Goku")
		assert.NoError(t, err)
		assert.NotNil(t, result)
		assert.Equal(t, "Goku", result.Name)
	})

	t.Run("returns error when API responds with non 200 status", func(t *testing.T) {
		server := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			w.WriteHeader(http.StatusInternalServerError)
		}))
		defer server.Close()

		r := repository.NewCharacterRepository(nil, server.URL)

		result, err := r.GetFromAPI("Goku")
		assert.Error(t, err)
		assert.Nil(t, result)
	})

	t.Run("returns error when API returns no characters", func(t *testing.T) {
		server := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			w.WriteHeader(http.StatusOK)
			w.Write([]byte("[]"))
		}))
		defer server.Close()

		repo := repository.NewCharacterRepository(nil, server.URL)

		result, err := repo.GetFromAPI("NonExistent")
		assert.Error(t, err)
		assert.Nil(t, result)
		assert.Equal(t, "character not found", err.Error())
	})

	t.Run("returns error when API response is invalid JSON", func(t *testing.T) {
		server := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			w.WriteHeader(http.StatusOK)
			w.Write([]byte("{invalid json}"))
		}))
		defer server.Close()

		r := repository.NewCharacterRepository(nil, server.URL)

		result, err := r.GetFromAPI("Goku")
		assert.Error(t, err)
		assert.Nil(t, result)
	})
}

func TestGetFromDBByName(t *testing.T) {
	t.Run("successfully fetches character from DB by name", func(t *testing.T) {
		db, mock, err := sqlmock.New()
		assert.NoError(t, err)
		defer db.Close()

		name := "Goku"
		query := "SELECT \\* FROM characters WHERE name LIKE \\?"
		searchName := "%" + utils.CapitalizeWords(name) + "%"
		rows := sqlmock.NewRows([]string{"id", "name", "ki"}).
			AddRow(1, "Goku", 9001)

		mock.ExpectQuery(query).WithArgs(searchName).WillReturnRows(rows)

		repo := repository.NewCharacterRepository(db, "")

		result, err := repo.GetFromDBByName(name)
		assert.NoError(t, err)
		assert.NotNil(t, result)
		assert.Equal(t, "Goku", result.Name)
		assert.Equal(t, "9001", result.Ki)

		assert.NoError(t, mock.ExpectationsWereMet())
	})

	t.Run("returns nil when no characters found in DB", func(t *testing.T) {
		db, mock, err := sqlmock.New()
		assert.NoError(t, err)
		defer db.Close()

		name := "NonExistent"
		query := "SELECT \\* FROM characters WHERE name LIKE \\?"
		searchName := "%" + utils.CapitalizeWords(name) + "%"

		mock.ExpectQuery(query).WithArgs(searchName).WillReturnRows(sqlmock.NewRows([]string{"id", "name", "ki"}))

		repo := repository.NewCharacterRepository(db, "")

		result, err := repo.GetFromDBByName(name)
		assert.NoError(t, err)
		assert.Nil(t, result)

		assert.NoError(t, mock.ExpectationsWereMet())
	})

	t.Run("returns error when DB query fails", func(t *testing.T) {
		db, mock, err := sqlmock.New()
		assert.NoError(t, err)
		defer db.Close()

		name := "Goku"
		query := "SELECT \\* FROM characters WHERE name LIKE \\?"
		searchName := "%" + utils.CapitalizeWords(name) + "%"

		mock.ExpectQuery(query).WithArgs(searchName).WillReturnError(sql.ErrConnDone)

		repo := repository.NewCharacterRepository(db, "")

		result, err := repo.GetFromDBByName(name)
		fmt.Println(result, err)
		assert.Error(t, err)
		assert.Nil(t, result)

		assert.NoError(t, mock.ExpectationsWereMet())
	})
}

func TestCreateToDB(t *testing.T) {
	t.Run("successfully inserts character into DB", func(t *testing.T) {
		db, mock, err := sqlmock.New()
		assert.NoError(t, err)
		defer db.Close()

		character := domain.Character{
			ID:   1,
			Name: "Goku",
			Ki:   "9001",
		}
		query := "INSERT INTO characters \\(id, name, ki\\) VALUES\\(\\?, \\?, \\?\\)"

		mock.ExpectPrepare(query)
		mock.ExpectExec(query).
			WithArgs(character.ID, character.Name, character.Ki).
			WillReturnResult(sqlmock.NewResult(1, 1))

		repo := repository.NewCharacterRepository(db, "")

		err = repo.CreateToDB(character)
		assert.NoError(t, err)

		assert.NoError(t, mock.ExpectationsWereMet())
	})

	t.Run("returns error when prepare statement fails", func(t *testing.T) {
		db, mock, err := sqlmock.New()
		assert.NoError(t, err)
		defer db.Close()

		query := "INSERT INTO characters \\(id, name, ki\\) VALUES\\(\\?, \\?, \\?\\)"
		mock.ExpectPrepare(query).WillReturnError(sql.ErrConnDone)

		repo := repository.NewCharacterRepository(db, "")

		character := domain.Character{
			ID:   1,
			Name: "Goku",
			Ki:   "9001",
		}

		err = repo.CreateToDB(character)
		assert.Error(t, err)

		assert.NoError(t, mock.ExpectationsWereMet())
	})

	t.Run("returns error when exec fails", func(t *testing.T) {
		db, mock, err := sqlmock.New()
		assert.NoError(t, err)
		defer db.Close()

		// Configurar mock para que falle la ejecución
		query := "INSERT INTO characters \\(id, name, ki\\) VALUES\\(\\?, \\?, \\?\\)"
		mock.ExpectPrepare(query)
		mock.ExpectExec(query).
			WithArgs(1, "Goku", "9001").
			WillReturnError(sql.ErrNoRows)

		repo := repository.NewCharacterRepository(db, "")

		character := domain.Character{
			ID:   1,
			Name: "Goku",
			Ki:   "9001",
		}

		err = repo.CreateToDB(character)
		assert.Error(t, err)

		// Verificar que todas las expectativas del mock se cumplieron
		assert.NoError(t, mock.ExpectationsWereMet())
	})
}
