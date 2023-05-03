package controllers

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/programmer-for-good/flashcardApi/Auth"
	"github.com/programmer-for-good/flashcardApi/models"
	"gorm.io/gorm"
	// "../models/models"
)

type CardsController struct {
	DB *gorm.DB
}

func (c *CardsController) Index(ctx *gin.Context) {
	var cards []models.Card
	result := c.DB.Find(&cards)
	if result.Error != nil {
		ctx.AbortWithStatus(http.StatusInternalServerError)
		return
	}
	ctx.JSON(http.StatusOK, cards)
}

func (c *CardsController) Show(ctx *gin.Context) {
	id := ctx.Param("id")
	var card models.Card
	result := c.DB.First(&card, id)
	if result.Error != nil {
		ctx.AbortWithStatus(http.StatusNotFound)
		return
	}
	ctx.JSON(http.StatusOK, card)
}

func (c *CardsController) CreateCard(ctx *gin.Context) {
	// id := ctx.Param("id")

	var flashcard models.Card
	if err := ctx.ShouldBindJSON(&flashcard); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	if err := c.DB.Create(&flashcard).Error; err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	// c.JSON(http.StatusOK, gin.H{"data": flashcard})
	// var card models.Card
	// result := c.DB.First(&card, id)
	// if result.Error != nil {
	// 	ctx.AbortWithStatus(http.StatusNotFound)
	// 	return
	// }
	ctx.JSON(http.StatusOK, flashcard)
}

func (c *CardsController) RegisterUser(ctx *gin.Context) {
	Auth.Register(ctx, c.DB)
}

func (c *CardsController) LoginUser(ctx *gin.Context) {
	Auth.LoginHandler(ctx, c.DB)
}
