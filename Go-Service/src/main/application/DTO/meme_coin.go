package DTO

import "time"

type CreateMemeCoinRequestDTO struct {
	Name        string `json:"name" validate:"required,unique"`
	Description string `json:"description"`
}

type GetMemeCoinDTO struct {
	ID string `json:"id" validate:"required"`
}

type GetMemeCoinResponseDTO struct {
	Name            string    `json:"name"`
	Description     string    `json:"description"`
	CreatedAt       time.Time `json:"created_at"`
	PopularityScore int       `json:"popularity_score"`
}

type UpdateMemeCoinRequestDTO struct {
	ID          string `json:"id" validate:"required"`
	Description string `json:"description"`
}

type DeleteMemeCoinRequestDTO struct {
	ID string `json:"id" validate:"required"`
}

// poke meme coin
type PokeMemeCoinRequestDTO struct {
	ID string `json:"id" validate:"required"`
}
