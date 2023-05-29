package controller

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/surysatriah/go-dashboard-app/internal/database"
	"github.com/surysatriah/go-dashboard-app/internal/helper"
	"github.com/surysatriah/go-dashboard-app/internal/model"
)

// @Summary Register a new user
// @Description Register a new user with the provided details
// @ID register-user
// @Accept json
// @Produce json
// @Param user body model.User true "User object containing registration details"
// @Success 201 {object} UserResponse
// @Failure 400 {object} ErrorResponse
// @Router /register [post]
func UserRegister(c *gin.Context) {

	db := database.GetDatabase()
	contentType := helper.GetContentType(c)

	User := model.User{}

	if contentType == "application/json" {
		c.ShouldBindJSON(&User)
	} else {
		c.ShouldBind(&User)
	}

	err := db.Debug().Create(&User).Error
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error":   "Bad Request",
			"message": err.Error(),
		})
		return
	}

	c.JSON(http.StatusCreated, gin.H{
		"Email":    User.Email,
		"Username": User.Username,
	})

}

func UserLogin(c *gin.Context) {

	db := database.GetDatabase()
	contentType := helper.GetContentType(c)

	User := model.User{}

	if contentType == "application/json" {
		c.ShouldBindJSON(&User)
	} else {
		c.ShouldBind(&User)
	}

	Password := User.Password

	err := db.Debug().Where("username = ?", User.Username).Take(&User).Error
	if err != nil {
		c.JSON(http.StatusUnauthorized, gin.H{
			"error":   "Unauthorized",
			"message": "invalid username/password",
		})
		return
	}

	comparePassword := helper.CompareHashedPassword([]byte(User.Password), []byte(Password))
	if !comparePassword {
		c.JSON(http.StatusUnauthorized, gin.H{
			"error":   "Unauthorized",
			"message": "invalid username/password",
		})
		return
	}

	token := helper.GenerateJWT(User.ID, User.Email)

	c.JSON(http.StatusOK, gin.H{
		"token": token,
	})

}

func Protected(c *gin.Context) {
	c.JSON(http.StatusOK, "This is a protected route")
}
