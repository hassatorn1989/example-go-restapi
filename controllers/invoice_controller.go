package controllers

import "github.com/gin-gonic/gin"

func InvoiceController() *_InvoiceController {
	return &_InvoiceController{}
}

type _InvoiceController struct{}

func (ctrl *_InvoiceController) Index(c *gin.Context) {
	c.JSON(200, gin.H{"message": "users"})
}

func (ctrl *_InvoiceController) Create(c *gin.Context) {

}

func (ctrl *_InvoiceController) Store(c *gin.Context) {

}

func (ctrl *_InvoiceController) Show(c *gin.Context) {

}

func (ctrl *_InvoiceController) Update(c *gin.Context) {

}

func (ctrl *_InvoiceController) UpdateStatus(c *gin.Context) {

}
