package controller

import (
	"go_learn/internal/repository"

	"github.com/gin-gonic/gin"
)

//ContactController is a ...
type ContactController interface {
	All(context *gin.Context)
	FindByID(context *gin.Context)
	Insert(context *gin.Context)
	Update(context *gin.Context)
	Delete(context *gin.Context)
}

type contactController struct {
	ContactRepo repository.ContactRepository
}

//NewContactController create a new instances of BoookController
func NewContactController(ContactRepo repository.ContactRepository) ContactController {
	return &contactController{
		ContactRepo: ContactRepo,
	}
}

func (c *contactController) All(context *gin.Context) {
	proses, err := c.ContactRepo.AllContact()
	if err != nil {
		context.JSON(500, proses)
		return
	}
	context.JSON(200, proses)
}
func (c *contactController) FindByID(context *gin.Context) {

}
func (c *contactController) Insert(context *gin.Context) {

	// context.JSON(200,)
}
func (c *contactController) Update(context *gin.Context) {

}
func (c *contactController) Delete(context *gin.Context) {

}
