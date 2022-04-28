package items

import(
	"Golang-Microservice/utils"
	"Golang-Microservice/database"
	"Golang-Microservice/logger"
	"context"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo/options"
	"time"
	"log"
)

var itemsCollection *mongo.Collection = database.OpenCollection(database.Client, "items")

func(item *Item) Save() *utils.RestErr{
	var ctx, cancel = context.WithTimeout(context.Background(), 100*time.Second)
	_, insertErr := itemsCollection.InsertOne(ctx, item)
	if insertErr != nil {
		logger.Error("Inserted Error::::", insertErr)
	}
	defer cancel()
	return nil
}

func(item *Item) GetAll() ([]bson.M, *utils.RestErr){
	var ctx, cancel = context.WithTimeout(context.Background(), 100*time.Second)
	var allItems []bson.M
	result, getErr := itemsCollection.Find(context.TODO(), bson.M{})
	if getErr != nil {
		logger.Error("error finding", getErr)
	}
	
	if err := result.All(ctx, &allItems); err != nil {
		logger.Error("ERROR::::", err)
	}
	defer cancel()
	return allItems, nil
}

func (item *Item) GetById() *utils.RestErr{
	var ctx, cancel = context.WithTimeout(context.Background(), 100*time.Second)
	err := itemsCollection.FindOne(ctx, bson.M{"item_id": item.Item_ID}).Decode(&item)
	defer cancel()
	if err != nil {
		logger.Error("error finding", err)
	}
	return nil
}

func (item *Item) Update() (*mongo.UpdateResult, *utils.RestErr){
	var ctx, cancel = context.WithTimeout(context.Background(), 100*time.Second)

	item.Updated_at, _ = time.Parse(time.RFC3339, time.Now().Format(time.RFC3339))

	filter := bson.M{"item_id": item.Item_ID}
	opts := options.Update().SetUpsert(true)
	update := bson.M{"$set":bson.M{"title":item.Title,"description": item.Description,"updated_at": item.Updated_at}}

	result , err := itemsCollection.UpdateOne(
		ctx,
		filter,
		update,
		opts,
	)
	defer cancel()
	if err != nil {
		logger.Error("error finding", err)
	 }
	return result, nil
}

func (item *Item) Delete() (*mongo.DeleteResult, *utils.RestErr){
	var ctx, cancel = context.WithTimeout(context.Background(), 100*time.Second)
	filter := bson.M{"item_id": item.Item_ID}
	result, err := itemsCollection.DeleteOne(ctx, filter)
	defer cancel()
	if err != nil {
		log.Println("ERROR:::", err)
	}
	return result, nil
}
