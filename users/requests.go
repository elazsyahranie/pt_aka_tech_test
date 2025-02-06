package users

import (
	"github.com/google/uuid"
)

type PostBody struct {
	ID          uuid.UUID 	
	Name		string 		`json:"name"`
	Password    string      `json:"password" binding:"required"`
	Email       string 		`json:"email" binding:"required"`
}

type DataForToken struct {
	ID 			uuid.UUID
	Email 		string 
}

type LogInRequest struct {
	Username 	string 		`json:"username"`
	Email 		string 		`json:"email"`
	Password    string      `json:"password" binding:"required"`
}