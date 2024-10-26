package profile

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

type Profile struct {
	ID   string `json:"id"`
	Name string `json:"name"`
}

var profiles = []Profile{
	{ID: "1", Name: "Alice"},
	{ID: "2", Name: "Bob"},
}

func RegisterRoutes(r *gin.Engine) {
	r.GET("/profiles/:id", getProfile)
}

func getProfile(c *gin.Context) {
	id := c.Param("id")
	for _, profile := range profiles {
		if profile.ID == id {
			c.JSON(http.StatusOK, profile)
			return
		}
	}
	c.JSON(http.StatusNotFound, gin.H{"message": "Profile not found"})
}
