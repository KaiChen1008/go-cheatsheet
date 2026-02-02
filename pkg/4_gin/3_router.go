package ginutil

import "github.com/gin-gonic/gin"

func NewRouter(h Handler) *gin.Engine {
	r := gin.Default()

	r.GET("/health", func(c *gin.Context) {
		c.JSON(200, gin.H{"status": "ok"})
	})

	r.POST("create", h.Create)

	// grouping
	v1 := r.Group("/v1")
	{
		v1.GET("hi", h.HiV1)
		//..
	}
	return r
}
