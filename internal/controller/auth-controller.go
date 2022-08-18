package controller

import (
	"go_learn/internal/dto"
	"go_learn/internal/model"
	"go_learn/internal/repository"
	"go_learn/internal/service"
	"go_learn/pkg/responsebuilder"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
	// "strings"
)

//AuthController interface is a contract what this controller can do
type AuthController interface {
	Login(ctx *gin.Context)
	// Register(ctx *gin.Context)
	// LoginBpjs(ctx *gin.Context)
}

type authController struct {
	userRepo   repository.UserRepository
	jwtService service.JWTService
}

//NewAuthController creates a new instance of AuthController
func NewAuthController(userRepo repository.UserRepository, jwtService service.JWTService) AuthController {
	return &authController{
		userRepo:   userRepo,
		jwtService: jwtService,
	}
}

func (c *authController) Login(ctx *gin.Context) {
	var loginDTO dto.LoginDTO
	errDTO := ctx.ShouldBind(&loginDTO)
	if errDTO != nil {
		ctx.String(419, "Parameter tidak sesuai")
		return
	}
	authResult := c.userRepo.VerifyCredential(loginDTO.Email, loginDTO.Password)
	if v, ok := authResult.(model.User); ok {
		generatedToken := c.jwtService.GenerateToken(strconv.FormatUint(uint64(v.ID), 10), v.First_name+" "+v.Last_name, strconv.FormatUint(uint64(v.Otoritas), 10), v.Email)
		// fmt.Printf(generatedToken)
		Data := struct {
			ID         uint   `json:"id" form:"id"`
			Nickname   string ` json:"nickname" `
			First_name string ` json:"first_name" binding:"required"`
			Last_name  string ` json:"last_name" binding:"required"`
			Email      string ` json:"email" binding:"required"`
			Phone      string ` json:"phone" binding:"required"`
			Otoritas   uint32 ` json:"otoritas" binding:"required"`
			Status     string `json:"status" binding:"required"`
		}{
			ID:         v.ID,
			Nickname:   v.Nickname,
			First_name: v.First_name,
			Last_name:  v.Last_name,
			Email:      v.Email,
			Phone:      v.Phone,
			Otoritas:   v.Otoritas,
			Status:     v.Status,
		}
		// v.Token = generatedToken
		response := responsebuilder.BuildResponseLogin(true, "OK!", generatedToken, Data)
		ctx.JSON(http.StatusOK, response)
		return
	}
	response := responsebuilder.BuildErrorResponse("Please check again your credential", "Invalid Credential", nil)
	ctx.AbortWithStatusJSON(http.StatusUnauthorized, response)
}

// func (c *authController) Register(ctx *gin.Context) {
// 	var registerDTO dto.RegisterDTO
// 	errDTO := ctx.ShouldBind(&registerDTO)
// 	if errDTO != nil {
// 		response := helper.BuildErrorResponse("Failed to process request", errDTO.Error(), helper.EmptyObj{})
// 		ctx.AbortWithStatusJSON(http.StatusBadRequest, response)
// 		return
// 	}

// 	if !c.authService.IsDuplicateEmail(registerDTO.Email) {
// 		response := helper.BuildErrorResponse("Failed to process request", "Duplicate email", helper.EmptyObj{})
// 		ctx.JSON(http.StatusConflict, response)
// 	} else {
// 		createdUser := c.authService.CreateUser(registerDTO)
// 		// token := c.jwtService.GenerateToken(strconv.FormatUint(uint64(createdUser.ID), 10))
// 		// fmt.Printf(token)
// 		// createdUser.Token = token
// 		response := helper.BuildResponse(true, "OK!", createdUser)
// 		ctx.JSON(http.StatusCreated, response)
// 	}
// }
