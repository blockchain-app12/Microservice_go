package service

import(
  "Golang-Microservice/domain/items"
  "Golang-Microservice/utils"
  "go.mongodb.org/mongo-driver/bson/primitive"
  "go.mongodb.org/mongo-driver/mongo"
  "time"
)

var (
	ItemsService ItemsServiceInterface = &itemsService{}
)

type ItemsServiceInterface interface {
	Create(items.Item) (*items.Item, *utils.RestErr)
	Get() ([]primitive.M, *utils.RestErr)
	GetByID(string) (*items.Item, *utils.RestErr)
	Update(bool,items.Item) (*mongo.UpdateResult, *utils.RestErr)
	Delete(string) (*mongo.DeleteResult, *utils.RestErr)
}

type itemsService struct{}

func (s *itemsService)Create(item items.Item) (*items.Item, *utils.RestErr) {
	if err := item.Validate(); err != nil {
		return nil, err
	}
	item.ID = primitive.NewObjectID()
	item.Created_at = time.Now()
	item.Updated_at = time.Now()
	item.Item_ID = item.ID.Hex()

	if err := item.Save(); err != nil {
		return nil, err
	}
	return &item, nil

}

func (s *itemsService) Get() ([]primitive.M, *utils.RestErr) {
	result := &items.Item{}
	result1, err := result.GetAll();
	if err != nil {
		return nil, err
	}
	return result1, nil
  }

 func (s *itemsService) GetByID(itemID string) (*items.Item, *utils.RestErr) {
	result := &items.Item{Item_ID: itemID}
	if err := result.GetById(); err != nil {
		return nil, err
	}
	return result, nil
  }

  func (s *itemsService) Update(isPartial bool,item items.Item) (*mongo.UpdateResult, *utils.RestErr) {
	result := &items.Item{Item_ID: item.Item_ID}

	if isPartial {
		if item.Title != "" {
			result.Title = item.Title
		}
		if item.Description != "" {
			result.Description = item.Description
		}

	} else {
		result.Title = item.Title
		result.Description = item.Description
	}

	updateID, updateErr := result.Update(); 
	  if updateErr != nil {
		return nil, updateErr
	}
	return updateID, nil
  }

  func (s *itemsService) Delete(ItemID string) (*mongo.DeleteResult, *utils.RestErr) {
	result := &items.Item{Item_ID: ItemID}
	deleteID, deleteErr := result.Delete()
	if deleteErr != nil {
		return nil, deleteErr
	}
	return deleteID, nil
  }