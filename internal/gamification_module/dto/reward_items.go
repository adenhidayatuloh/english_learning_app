package dto

import "github.com/google/uuid"

type RewardItemsResponse struct {
	ID          uuid.UUID `json:"id"`
	Name        string    `json:"name"`
	Points      int       `json:"points"`
	Description string    `json:"description,omitempty"`
	Terms       string    `json:"terms,omitempty"`
}

type ReedemItemResponse struct {
	Name    string `json:"name"`
	Message string `json:"message"`
}
