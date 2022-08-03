package controller

import (
	"go_learn/internal/repository"

	"go_learn/internal/model"

	"github.com/gin-gonic/gin"
)

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

// All implements ContactController
func (c *contactController) All(context *gin.Context) {
	proses, error := c.ContactRepo.GetContacts()
	if error != nil {
		context.JSON(500, proses)
	}
	context.JSON(200, proses)
}

// Delete implements ContactController
func (*contactController) Delete(context *gin.Context) {
	panic("unimplemented")
}

// FindByID implements ContactController
func (*contactController) FindByID(context *gin.Context) {
	panic("unimplemented")
}

// Insert implements ContactController
func (c *contactController) Insert(context *gin.Context) {
	var input model.Contact
	// nama := context.Param("nama")
	// alamat := context.Param("alamat")
	// no_telp := context.Param("no_telp")

	proses, error := c.ContactRepo.InsertContact(input)
	if error != nil {
		context.JSON(500, proses)
	}
	// fmt.Println(context.Params)
	context.JSON(200, proses)
}

// Update implements ContactController
func (*contactController) Update(context *gin.Context) {
	panic("unimplemented")
}

func NewContactController(ContactRepo repository.ContactRepository) ContactController {
	return &contactController{
		ContactRepo: ContactRepo,
	}
}
