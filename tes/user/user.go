package user

import (
	"encoding/json"
	"net/http"

	"github.com/gin-gonic/gin"
)

func RegisterRoutes(r *gin.Engine) {
	r.GET("/users/:id", getUserProfile)
}

func getUserProfile(c *gin.Context) {
	userId := c.Param("id")

	// Mengambil data profil dari Profile Module
	resp, err := http.Get("http://localhost:8080/profiles/" + userId)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"message": "Error fetching profile"})
		return
	}
	defer resp.Body.Close()

	if resp.StatusCode == http.StatusOK {
		var profile map[string]interface{}
		if err := json.NewDecoder(resp.Body).Decode(&profile); err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"message": "Error decoding profile"})
			return
		}
		c.JSON(http.StatusOK, profile)
	} else {
		c.JSON(resp.StatusCode, gin.H{"message": "Profile not found"})
	}
}
