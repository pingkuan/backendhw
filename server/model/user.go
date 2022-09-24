package model

import (
	"github.com/line/line-bot-sdk-go/v7/linebot"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

type User struct {
	UID      primitive.ObjectID `bson:"_id,omitempty"`
	UserID   string             `bson:"userID,omitempty"`
	Messages []Message          `bson:"messages,omitempty"`
}

type Message struct {
	ID     string           `bson:"id,omitempty"`
	Text   string           `bson:"text,omitempty"`
	Emojis []*linebot.Emoji `bson:"emojis,omitempty"`
}
