package controller

import (
	"log"
	"net/http"
	"strings"

	"github.com/access-module/api/db"
	"github.com/access-module/api/model"
	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
	"golang.org/x/crypto/bcrypt"
)

// https://github.com/go-playground/validator
type UserRequest struct {
	Username string `json:"username" binding:"required,email"`
	Password string `json:"password" binding:"required,alphanum"`
}

func CreateUser(ctx *gin.Context) {

	var req UserRequest
	err := ctx.ShouldBindJSON(&req)
	if err != nil {
		log.Println(err)
		ctx.JSON(http.StatusBadRequest, gin.H{
			"response": "bad request",
		})
		return
	}

	username := strings.ToLower(req.Username)
	db, err := db.Connect()

	if err != nil {
		log.Println(err)
		ctx.JSON(http.StatusBadRequest, gin.H{
			"response": "bad request",
		})
		return
	}

	var user model.User
	queryResults := db.Where("username = ?", username).Limit(1).Find(&user)
	if queryResults.RowsAffected > 0 {
		log.Println("Username already exists")
		ctx.JSON(http.StatusBadRequest, gin.H{"response": "username already exists"})
		return
	}

	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(req.Password), bcrypt.DefaultCost) // 10
	if err != nil {
		log.Println(err)
		ctx.JSON(http.StatusBadRequest, gin.H{
			"response": "bad request",
		})
		return
	}

	user = model.User{
		Id:             uuid.NewString(),
		Username:       username,
		HashedPassword: string(hashedPassword),
	}

	result := db.Create(&user)
	ctx.JSON(http.StatusOK, result)

}
func ListUser(ctx *gin.Context) {
	ctx.JSON(http.StatusOK, "you are so smart")
}
