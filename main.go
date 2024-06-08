package main

//	@title			Swagger Example API
//	@version		1.0
//	@description	This is a sample server celler server.
//	@termsOfService	http://swagger.io/terms/

//	@contact.name	API Support
//	@contact.url	http://www.swagger.io/support
//	@contact.email	support@swagger.io

//	@license.name	Apache 2.0
//	@license.url	http://www.apache.org/licenses/LICENSE-2.0.html

// @host		localhost:8080
// @BasePath	/api/v1

import (
	"github.com/gin-gonic/gin"
	swaggerFiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"
	"swago/controller"
	_ "swago/docs"
)

func main() {

	g := gin.Default()
c:= controller.NewController()
	v1 := g.Group("/api/v1")
	{
		ant := v1.Group("/ant")
		{
			ant.GET("/",c.GetCollonies)
		}


	}

	g.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))
	err := g.Run()

	if err != nil {
		panic(err)

	}
}
