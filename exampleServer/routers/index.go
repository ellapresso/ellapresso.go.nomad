package v1

import (
	"github.com/gin-gonic/gin"
)

func InitRoutes(g *gin.RouterGroup) {
	// g.Use(appid.AppIDMiddleWare())
	SetExampleRoutes(g)

	// SetHelloRoutes(g)
	// SetAuthRoutes(g) // SetAuthRoutes invoked
	// g.Use(token.TokenAuthMiddleWare())  //secure the API From this line to bottom with JSON Auth
	// g.Use(appid.ValidateAppIDMiddleWare())
	// SetTaskRoutes(g)
	// SetUserRoutes(g)
}
