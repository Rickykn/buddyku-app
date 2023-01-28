package handlers

import (
	"net/http"

	"github.com/Rickykn/buddyku-app.git/dtos"
	"github.com/gin-gonic/gin"
)

func (h *Handler) RegisterAdmin(c *gin.Context) {
	var registerAdminInput *dtos.AdminRegisterDTO

	err := c.ShouldBindJSON(&registerAdminInput)

	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"message": "Bad Request",
		})
		return

	}

	admin, response, _ := h.adminService.RegisterAdmin(registerAdminInput)

	if response.Error {
		c.JSON(response.Code, gin.H{
			"message":     response.Message,
			"status code": response.Code,
			"data":        response.Data,
		})
	} else {
		newAdmin := &dtos.AdminRegisterResponse{
			Name: admin.Name,
			Role: admin.Role,
		}

		c.JSON(response.Code, gin.H{
			"message":     response.Message,
			"status code": response.Code,
			"data":        newAdmin,
		})
	}

}
