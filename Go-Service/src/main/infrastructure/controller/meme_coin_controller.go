package controller

import (
	"Go-Service/src/main/application/DTO"
	"Go-Service/src/main/application/usecase"
	"github.com/gin-gonic/gin"
	"net/http"
)

type MemeCoinController struct {
	usecase *usecase.MemeCoinUsecase
}

func NewMemeCoinController(usecase *usecase.MemeCoinUsecase) *MemeCoinController {
	return &MemeCoinController{usecase: usecase}
}

func (c *MemeCoinController) Create(ctx *gin.Context) {
	var req DTO.CreateMemeCoinRequestDTO
	if err := ctx.ShouldBindJSON(&req); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	err := c.usecase.Create(ctx, &req)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	ctx.JSON(http.StatusCreated, gin.H{"message": "Meme coin created successfully"})
}
