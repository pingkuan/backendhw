package handler

import (
	"context"
	"log"
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/pingkuan/backendhw/server/db"
	"github.com/pingkuan/backendhw/server/model"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
)

func GetMessages(c *gin.Context) {
	id := c.Param("userid")
	result := model.User{}

	var userColl *mongo.Collection = db.GetCollection(db.DB, "users")
	ctx, cancel := context.WithTimeout(context.Background(), 2*time.Second)
	defer cancel()

	err := userColl.FindOne(ctx, bson.M{"userID": id}).Decode(&result)
	if err != nil {
		log.Println(err)
		return
	}

	c.JSON(http.StatusOK, gin.H{"messages": result.Messages})
}
