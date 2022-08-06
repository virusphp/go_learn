package controller

import (
	"fmt"
	"go_learn/internal/dto"
	"go_learn/internal/repository"
	"go_learn/pkg/responsebuilder"
	"net/http"
	"strconv"

	"go_learn/internal/model"

	"github.com/mashingan/smapping"

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
	var listcontactDTO dto.ListContactDTO
	errDTO := context.ShouldBind(&listcontactDTO)
	if errDTO != nil {
		res := responsebuilder.BuildErrorResponse("ERROR!", errDTO.Error(), nil)
		context.JSON(http.StatusInternalServerError, res)
		return
	}
	result, count_, err := c.ContactRepo.FindContactByName(listcontactDTO.Search, listcontactDTO.Limit, listcontactDTO.Page, listcontactDTO.Order)
	if err != nil {
		res := responsebuilder.BuildErrorResponse("Something Error", "Fatal Error", responsebuilder.EmptyObj{})
		context.JSON(http.StatusBadRequest, res)
		return
	}

	if *count_ == 0 {
		res := responsebuilder.BuildErrorResponse("Data not found", "No data with given id", responsebuilder.EmptyObj{})
		context.JSON(http.StatusNotFound, res)
		return
	}
	res := responsebuilder.BuildResponse_table(true, "OK", *count_, result)
	context.JSON(http.StatusOK, res)
}

// Delete implements ContactController
func (c *contactController) Delete(context *gin.Context) {
	id := context.Param("id")

	convertedId, err := strconv.Atoi(id)
	if err != nil {
		fmt.Println("error kampret := ", err)
		context.JSON(500, err)
		return
	}
	input := model.Contact{}

	input.ID = uint(convertedId)

	error := c.ContactRepo.DeleteContact(input)
	if error != nil {
		context.JSON(500, error)
		return
	}

	context.String(200, "SUKSES HAPUS!")
}

// FindByID implements ContactController
func (c *contactController) FindByID(context *gin.Context) {
	id := context.Param("id")
	//convert id from string to int
	convertedId, err := strconv.Atoi(id)

	if err != nil {
		context.JSON(404, err)
		return
	}

	proses, error := c.ContactRepo.FindContactByID(uint(convertedId))

	if error != nil {
		context.JSON(500, error)
		return
	}

	context.JSON(200, proses)

}

// Insert implements ContactController
func (c *contactController) Insert(context *gin.Context) {
	var dtoContact dto.ContactDTO
	var model model.Contact

	err := context.ShouldBind(&dtoContact)
	if err != nil {
		res := responsebuilder.BuildErrorResponse("ERROR!", err.Error(), nil)
		context.JSON(500, res)
		return
	}

	err = smapping.FillStruct(&model, smapping.MapFields(&dtoContact))
	if err != nil {
		fmt.Println("Failed map : ", err)
		return
	}

	proses, error := c.ContactRepo.InsertContact(model)
	if error != nil {
		res := responsebuilder.BuildErrorResponse("ERROR!", err.Error(), nil)
		context.JSON(500, res)
		return
	}

	res := responsebuilder.BuildResponse(true, "OK", proses)
	context.JSON(200, res)
}

// Update implements ContactController
func (c *contactController) Update(context *gin.Context) {
	id := context.Param("id")

	convertedId, err := strconv.Atoi(id)
	if err != nil {
		fmt.Println("error kampret := ", err)
		context.JSON(500, err)
		return
	}
	var dtoContact dto.ContactDTO

	err = context.ShouldBind(&dtoContact)
	if err != nil {
		fmt.Println("error kamvrett := ", err)
		context.JSON(500, err.Error())
		return
	}

	model := model.Contact{}
	err = smapping.FillStruct(&model, smapping.MapFields(&dtoContact))
	if err != nil {
		fmt.Println("Failed map : ", err)
		context.JSON(500, err)
		return
	}

	model.ID = uint(convertedId)
	proses, error := c.ContactRepo.UpdateContact(model)
	if error != nil {
		context.JSON(500, proses)
		return
	}

	context.JSON(200, proses)
}

func NewContactController(ContactRepo repository.ContactRepository) ContactController {
	return &contactController{
		ContactRepo: ContactRepo,
	}
}
