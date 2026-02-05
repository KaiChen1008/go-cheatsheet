package ginutil

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

type Handler interface {
	Create(c *gin.Context)
	HiV1(c *gin.Context)
}

type handler struct{}

// for get
func (h *handler) Create(c *gin.Context) {

	// e.g. items/:id
	param := c.Param("id")
	if param == "" {
		c.JSON(400, gin.H{"msg": "param is required"})
	}

	// e.g. url?name=""
	q := c.Query("name")
	if q == "" {
		c.JSON(400, gin.H{"msg": "query is required"})
	}
	// ...
}

type hiReq struct {
	ID   int    `json:"id"`
	Name string `json:"name"`
}

// for post, put...
func (h *handler) HiV1(c *gin.Context) {
	form := c.PostForm("abc")
	if form == "" {
		c.JSON(400, gin.H{"msg": "form is required"})
		return
	}

	req := hiReq{}
	// BindJson: will return 400 directly
	// ShouldBindJson: you have to return a response by yourself
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(400, gin.H{"msg": err.Error()})
		return
	}
	//....
}

// patch example
type User struct {
	Name  string `json:"name"`
	Email string `json:"email"`
}

func PATCH() {
	r := gin.Default()

	// handle patch
	// r.PATCH("/user/:id", func(c *gin.Context) {
	r.PATCH("/user/:id/email", func(c *gin.Context) {
		id := c.Param("id")
		var user User

		// bind json
		if err := c.ShouldBindJSON(&user); err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			return
		}

		// partial updat (e.g. only update email)
		c.JSON(http.StatusOK, gin.H{
			"id":      id,
			"message": "User updated (PATCH)",
			"updated": user,
		})
	})

	r.Run(":8080")
}
