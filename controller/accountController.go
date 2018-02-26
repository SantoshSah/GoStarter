package controller

import (
	"GoStarter/app"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/sirupsen/logrus"
)

func (c *Controller) signUp(ctx *gin.Context) {
	var req map[string]struct {
		ID           int64  `json:"-"`
		Email        string `json:"email"`
		FirstName    string `json:"firstName"`
		LastName     string `json:"lastName"`
		MobileNumber string `json:"mobileNumber"`
	}

	if err := ctx.BindJSON(&req); err != nil {
		logrus.Error(err)
		ctx.JSON(http.StatusBadRequest, gin.H{
			"error": "invalid body",
		})
		return
	}
	if _, ok := req["data"]; !ok {
		ctx.JSON(http.StatusBadRequest, gin.H{
			"error": "not founded \"data\"",
		})
		return
	}
	logrus.Debug(req["data"].ID)
	logrus.Debug(req["data"].Email)

	if req["data"].Email == "" {
		ctx.JSON(http.StatusBadRequest, gin.H{
			"error": "email\" must be set",
		})
		return
	}
	u := &app.User{
		ID:           0,
		Email:        req["data"].Email,
		FirstName:    req["data"].FirstName,
		LastName:     req["data"].LastName,
		MobileNumber: req["data"].MobileNumber,
	}
	/*
		if err := c.App.DB.Create(u).Error; err != nil {
			ctx.JSON(http.StatusInternalServerError, gin.H{
				"error": err.Error(),
			})
			return
		}
	*/
	logrus.Debug(u)

	ctx.JSON(http.StatusOK, gin.H{
		"response": "ok",
	})
}
