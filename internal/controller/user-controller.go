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

type UserController interface {
	All(context *gin.Context)
	FindByID(context *gin.Context)
	Insert(context *gin.Context)
	Update(context *gin.Context)
	Delete(context *gin.Context)
}

type userController struct {
	UserRepo repository.UserRepository
}

// All implements UserController
func (c *userController) All(context *gin.Context) {
	var listuserDTO dto.ListUserDTO
	errDTO := context.ShouldBind(&listuserDTO)
	if errDTO != nil {
		res := responsebuilder.BuildErrorResponse("ERROR!", errDTO.Error(), nil)
		context.JSON(http.StatusInternalServerError, res)
		return
	}
	result, count_, err := c.UserRepo.FindUserByName(listuserDTO.Search, listuserDTO.Limit, listuserDTO.Page, listuserDTO.Order)
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

// Delete implements UserController
func (c *userController) Delete(context *gin.Context) {
	id := context.Param("id")

	convertedId, err := strconv.Atoi(id)
	if err != nil {
		fmt.Println("error kampret := ", err)
		context.JSON(500, err)
		return
	}
	input := model.User{}

	input.ID = uint(convertedId)

	error := c.UserRepo.DeleteUser(input)
	if error != nil {
		context.JSON(500, error)
		return
	}

	context.String(200, "SUKSES HAPUS!")
}

// FindByID implements UserController
func (c *userController) FindByID(context *gin.Context) {
	id := context.Param("id")
	//convert id from string to int
	convertedId, err := strconv.Atoi(id)

	if err != nil {
		context.JSON(404, err)
		return
	}

	proses, error := c.UserRepo.FindUserByID(uint(convertedId))

	if error != nil {
		context.JSON(500, error)
		return
	}

	context.JSON(200, proses)

}

// Insert implements UserController
func (c *userController) Insert(context *gin.Context) {
	var dtoUser dto.UserDTO
	var model model.User

	err := context.ShouldBind(&dtoUser)
	if err != nil {
		res := responsebuilder.BuildErrorResponse("ERROR!", err.Error(), nil)
		context.JSON(500, res)
		return
	}

	err = smapping.FillStruct(&model, smapping.MapFields(&dtoUser))
	if err != nil {
		fmt.Println("Failed map : ", err)
		return
	}

	proses, error := c.UserRepo.InsertUser(model)
	if error != nil {
		res := responsebuilder.BuildErrorResponse("ERROR!", err.Error(), nil)
		context.JSON(500, res)
		return
	}

	res := responsebuilder.BuildResponse(true, "OK", proses)
	context.JSON(200, res)
}

// Update implements UserController
func (c *userController) Update(context *gin.Context) {
	id := context.Param("id")

	convertedId, err := strconv.Atoi(id)
	if err != nil {
		fmt.Println("error kampret := ", err)
		context.JSON(500, err)
		return
	}
	var dtoUser dto.UserDTO

	err = context.ShouldBind(&dtoUser)
	if err != nil {
		fmt.Println("error kamvrett := ", err)
		context.JSON(500, err.Error())
		return
	}

	model := model.User{}
	err = smapping.FillStruct(&model, smapping.MapFields(&dtoUser))
	if err != nil {
		fmt.Println("Failed map : ", err)
		context.JSON(500, err)
		return
	}

	model.ID = uint(convertedId)
	proses, error := c.UserRepo.UpdateUser(model)
	if error != nil {
		context.JSON(500, proses)
		return
	}

	context.JSON(200, proses)
}

func NewUserController(UserRepo repository.UserRepository) UserController {
	return &userController{
		UserRepo: UserRepo,
	}
}
