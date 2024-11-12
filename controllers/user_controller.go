package controllers

import (
	"github.com/gin-gonic/gin"
)

type _UserController struct{}

func UserController() *_UserController {
	return &_UserController{}
}

func (ctrl *_UserController) Index(c *gin.Context) {
	c.JSON(200, gin.H{"message": "users"})
}

func (ctrl *_UserController) Create(c *gin.Context) {

}

func (ctrl *_UserController) Store(c *gin.Context) {

}

func (ctrl *_UserController) Show(c *gin.Context) {

}

func (ctrl *_UserController) Update(c *gin.Context) {

}

func (ctrl *_UserController) UpdateStatus(c *gin.Context) {

}
