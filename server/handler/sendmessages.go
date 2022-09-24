package handler

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/line/line-bot-sdk-go/v7/linebot"
	"github.com/pingkuan/backendhw/server/config"
)

type SendbackRequestBody struct {
	Message string `json:"message"`
}

func Sendmessages(c *gin.Context) {
	id := c.Param("userid")
	body := SendbackRequestBody{}
	var bot = config.Bot

	if err := c.BindJSON(&body); err != nil {
		c.AbortWithError(http.StatusBadRequest, err)
		return
	}

	message := linebot.NewTextMessage(body.Message)

	_, err := bot.PushMessage(id, message).Do()
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err})
	}

	c.JSON(http.StatusOK, gin.H{"message": "傳送成功"})
}
