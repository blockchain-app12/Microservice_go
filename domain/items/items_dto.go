package items

import (
	"time"
	"strings"
	"Golang-Microservice/utils"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

type Item struct {
	ID           primitive.ObjectID    `json:"id"`
	Title        string                `json:"title"`
	Description  string                `json:"description"`
	Created_at   time.Time             `json:"created_at"`
	Updated_at   time.Time             `json:"updated_at"`
	Item_ID      string                `json:"item_id`
}

type Items []Item

func(item *Item) Validate() *utils.RestErr{
	item.Title = strings.TrimSpace(item.Title)
	item.Description = strings.TrimSpace(item.Description)
	if item.Title == "" {
		utils.NewBadRequestError("invalid title")
	}
	if item.Description == "" {
		utils.NewBadRequestError("invalid description")
	}
	return nil
}