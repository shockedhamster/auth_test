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

type userToMod struct {
	Username    string `json:"username"`
	NewUsername string `json:"newusername"`
}

func (h *Handler) deleteUser(c *gin.Context) {
	var input userToMod

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

func (h *Handler) updateUsername(c *gin.Context) {
	var input userToMod

	if err := c.BindJSON(&input); err != nil {
		newErrorResponce(c, http.StatusBadRequest, err.Error())
		return
	}

	id, err := h.services.Edit.UpdateUsername(input.Username, input.NewUsername)
	if err != nil {
		newErrorResponce(c, http.StatusInternalServerError, err.Error())
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"status": "New username is set",
		"id":     id,
	})
}
