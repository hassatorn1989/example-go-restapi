package controllers

import (
	"inv-app/auth"
	"inv-app/configs"
	"inv-app/models"
	"inv-app/utils"

	"github.com/gin-gonic/gin"
)

func AuthController() *_AuthController {
	return &_AuthController{}
}

type _AuthController struct {
}

func (ctrl *_AuthController) Login(c *gin.Context) {

	var collection models.User
	if err := c.ShouldBind(&collection); err != nil {
		c.JSON(400, gin.H{"error": err.Error()})
		return
	}

	if err := configs.DB.Where("username = ?", c.PostForm("username")).First(&collection).Error; err != nil {
		c.JSON(400, gin.H{"error": err.Error()})
		return
	}

	if !utils.CheckPasswordHash(c.PostForm("password"), collection.Password) {
		c.JSON(400, gin.H{"error": "invalid password"})
		return
	}

	token, err := auth.GenerateJWT(collection.ID.String(), false)
	if err != nil {
		c.JSON(500, gin.H{"error": err.Error()})
		return
	}

	c.JSON(200, gin.H{
		"user_id": collection.ID,
		"token":   token,
	})

}

func (ctrl *_AuthController) Register(c *gin.Context) {
	// create user
	hashPassword, _ := utils.HashPassword(c.PostForm("password"))
	collection := models.User{
		Name:     c.PostForm("name"),
		Username: c.PostForm("username"),
		Password: hashPassword,
	}

	if err := configs.DB.Create(&collection).Error; err != nil {
		c.JSON(400, gin.H{"error": err.Error()})
		return
	}

	c.JSON(200, gin.H{
		"message": "success",
		"data":    collection,
	})
}

func (ctrl *_AuthController) Logout(c *gin.Context) {
	var collection models.User

	c.JSON(200, gin.H{
		"message": "success",
		"data":    collection,
	})
}
