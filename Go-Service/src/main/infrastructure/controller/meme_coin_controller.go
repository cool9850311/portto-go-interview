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

	id, err := c.usecase.Create(ctx, &req)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	ctx.JSON(http.StatusCreated, gin.H{"id": id})
}

func (c *MemeCoinController) GetByID(ctx *gin.Context) {
	var req DTO.GetMemeCoinRequestDTO
	if err := ctx.ShouldBindJSON(&req); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	response, err := c.usecase.GetByID(ctx, &req)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	ctx.JSON(http.StatusOK, response)
}

func (c *MemeCoinController) Update(ctx *gin.Context) {
	var req DTO.UpdateMemeCoinRequestDTO
	if err := ctx.ShouldBindJSON(&req); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	err := c.usecase.Update(ctx, &req)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	ctx.JSON(http.StatusOK, gin.H{"message": "Meme coin updated successfully"})
}

func (c *MemeCoinController) Delete(ctx *gin.Context) {
	var req DTO.DeleteMemeCoinRequestDTO
	if err := ctx.ShouldBindJSON(&req); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	err := c.usecase.Delete(ctx, &req)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	ctx.JSON(http.StatusOK, gin.H{"message": "Meme coin deleted successfully"})
}

func (c *MemeCoinController) Poke(ctx *gin.Context) {
	var req DTO.PokeMemeCoinRequestDTO
	if err := ctx.ShouldBindJSON(&req); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	err := c.usecase.Poke(ctx, &req)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	ctx.JSON(http.StatusOK, gin.H{"message": "Meme coin poked successfully"})
}
