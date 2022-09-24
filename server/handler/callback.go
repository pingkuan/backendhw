package handler

import (
	"context"
	"log"
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/line/line-bot-sdk-go/v7/linebot"
	"github.com/pingkuan/backendhw/server/config"
	"github.com/pingkuan/backendhw/server/db"
	"github.com/pingkuan/backendhw/server/model"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
)

func Callback(c *gin.Context) {
	var bot = config.Bot
	events, err := bot.ParseRequest(c.Request)

	if err != nil {
		if err == linebot.ErrInvalidSignature {
			c.JSON(http.StatusBadRequest, gin.H{
				"message": "signature validation error",
			})
			return
		}
		c.JSON(http.StatusInternalServerError, gin.H{
			"message": "events parsing error",
		})
		return
	}

	for _, event := range events {
		if event.Type == linebot.EventTypeMessage {
			switch message := event.Message.(type) {
			case *linebot.TextMessage:
				_, err := bot.GetMessageQuota().Do()
				if err != nil {
					log.Println("Quota err:", err)
					return
				}

				var user = model.User{}
				//收到之訊息
				var newMessage = model.Message{
					ID:     message.ID,
					Text:   message.Text,
					Emojis: message.Emojis,
				}

				var userColl *mongo.Collection = db.GetCollection(db.DB, "users")
				ctx, cancel := context.WithTimeout(context.Background(), 2*time.Second)
				defer cancel()
				//訊息傳送者的id
				userid := event.Source.UserID
				//查詢使用者是否已存在於資料庫
				err = userColl.FindOne(ctx, bson.M{"userID": userid}).Decode(&user)
				if err != nil {
					//若不存在，則創建一新使用者，並回傳收到第一則訊息
					if err == mongo.ErrNoDocuments {
						var sliceMessage []model.Message
						sliceMessage = append(sliceMessage, newMessage)

						user.UID = primitive.NewObjectID()
						user.UserID = userid
						user.Messages = sliceMessage

						_, err := userColl.InsertOne(ctx, user)
						if err != nil {
							log.Println(err)
						}

						if _, err = bot.ReplyMessage(event.ReplyToken, linebot.NewTextMessage("收到第一則訊息，已輸入資料庫")).Do(); err != nil {
							log.Print(err)
						}
						return
					}
					log.Println(err)
					return
				}
				//若使用者已存在，則新增訊息於資料庫
				updated, err := userColl.UpdateOne(ctx, bson.M{"userID": userid}, bson.M{"$push": bson.M{"messages": newMessage}})
				if err != nil {
					log.Println(err)
				}
				log.Println(updated)
				if _, err = bot.ReplyMessage(event.ReplyToken, linebot.NewTextMessage("收到訊息，已輸入資料庫")).Do(); err != nil {
					log.Print(err)
				}
				return
			}
		}

	}
}
