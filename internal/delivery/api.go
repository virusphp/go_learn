package delivery

import (
	"go_learn/config"
	"go_learn/internal/controller"
	"go_learn/internal/repository"
	"go_learn/internal/service"
	"go_learn/middleware"
	"os"

	"github.com/gin-gonic/gin"
	// "github.com/golang-jwt/jwt"

	"gorm.io/gorm"
)

var server = config.Server{}

var (
	db *gorm.DB = server.SetupDatabaseConnection()

	// jwtService service.JWTService = service.NewJWTService()

)

func CORSMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {
		c.Writer.Header().Set("Access-Control-Allow-Origin", "*")
		c.Writer.Header().Set("Access-Control-Max-Age", "86400")
		c.Writer.Header().Set("Access-Control-Allow-Methods", "POST, GET, OPTIONS, PUT, DELETE, UPDATE")
		c.Writer.Header().Set("Access-Control-Allow-Headers", "Access-Control-Allow-Origin, Origin, Content-Type, Content-Length, Accept-Encoding, X-CSRF-Token, Authorization, _megono")
		c.Writer.Header().Set("Access-Control-Expose-Headers", "Content-Length")
		c.Writer.Header().Set("Access-Control-Allow-Credentials", "true")

		if c.Request.Method == "OPTIONS" {
			c.AbortWithStatus(200)
		} else {
			c.Next()
		}
	}
}

// subjectFromJWT parses a JWT and extract subject from sub claim.
// func subjectFromJWT(c *gin.Context) string {
// 	authHeader := c.Request.Header.Get("Authorization")
// 	aToken, err := jwtService.ValidateToken(authHeader)
// 	claims := aToken.Claims.(jwt.MapClaims)
// 	subject := fmt.Sprintf("%v", claims["email"])
// 	if err != nil {
// 		return ""
// 	}
// 	return subject
// }

func InitializeRoutes() {

	defer config.CloseDatabaseConnection(db)

	r := gin.Default()
	r.MaxMultipartMemory = 8 << 20
	r.Use(CORSMiddleware())

	// antrol.WsBpjsAntrolRoutes(r, db, jwtService)
	port := os.Getenv("APP_PORT")
	var (
		jwtService        service.JWTService           = service.NewJWTService()
		repo              repository.ContactRepository = repository.NewContactRepository(db)
		repoUser          repository.UserRepository    = repository.NewUserRepository(db)
		controllerContact controller.ContactController = controller.NewContactController(repo)
		controllerUser    controller.UserController    = controller.NewUserController(repoUser)
		controllerAuth    controller.AuthController    = controller.NewAuthController(repoUser, jwtService)
	)
	// r.GET("/api/contact/all", controller.All)

	contact := r.Group("/api/contact", middleware.AuthorizeJWT(jwtService))
	{
		contact.POST("/all", controllerContact.All)
		contact.POST("/insert", controllerContact.Insert)
		contact.PUT("/update/:id", controllerContact.Update)
		contact.DELETE("/delete/:id", controllerContact.Delete)
		contact.GET("/find/:id", controllerContact.FindByID)
	}

	r.POST("/api/user/insert", controllerUser.Insert)
	r.POST("/api/auth/login", controllerAuth.Login)
	r.Run(":" + port)

}
