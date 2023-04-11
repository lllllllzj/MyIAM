package apiserver

import "github.com/gin-gonic/gin"

func initRouter(g *gin.Engine) {
	installMiddleware(g)
	installAPI(g)
}

func installMiddleware(g *gin.Engine) {

}

func installAPI(g *gin.Engine) {
	v1 := g.Group("/v1")
	{
		userv1 := v1.Group("/users")
		{
			userv1.POST("")
		}
	}
}
