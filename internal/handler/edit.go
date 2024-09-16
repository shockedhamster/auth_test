package handler

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func (h *Handler) hello(c *gin.Context) {

	c.JSON(http.StatusOK, gin.H{
		"hello": "hello",
	})
}

type deleteUser struct {
	Username string `json:"username"`
}

func (h *Handler) deleteUser(c *gin.Context) {
	var input deleteUser

	if err := c.BindJSON(&input); err != nil {
		newErrorResponce(c, http.StatusBadRequest, err.Error())
		return
	}

	err := h.services.Edit.DeleteUser(input.Username)
	if err != nil {
		newErrorResponce(c, http.StatusInternalServerError, err.Error())
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"status": "user is deleted",
	})

}
