package handler

import (
	"dragonball-api/pkg/service"
	"net/http"

	"github.com/gin-gonic/gin"
)

type CharacterHandler struct {
	characterService service.CharacterService
}

type Data struct {
	Name string `json:"namr" binding:"required"`
	Ki   string `json:"ki" binding:"required"`
}

func NewCharacterHandler(characterService service.CharacterService) CharacterHandler {
	return CharacterHandler{
		characterService: characterService,
	}
}

func (r CharacterHandler) Create(c *gin.Context) {
	name := c.Query("name")
	if name == "" {
		c.JSON(http.StatusBadRequest, gin.H{"message": "name param is required"})
		return
	}
	char, err := r.characterService.FindOrCreate(name)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"message": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"data": char})
}
