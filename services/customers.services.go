package services

import (
	"context"
	"token-eg/interfaces"
	"token-eg/models"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
)

type CustomerService struct {
	ctx             context.Context
	mongoCollection *mongo.Collection
	client          *mongo.Client
}

func InitializeCustomerService(ctx context.Context, collection *mongo.Collection, client *mongo.Client) interfaces.ICustomer {
	return &CustomerService{ctx, collection, client}
}

func (c *CustomerService) CreateCustomer(customer *models.UserSignIn) (*models.DBResponse, error) {
	res, err := c.mongoCollection.InsertOne(c.ctx, &customer)

	if err != nil {
		return nil, err
	}

	var newUser *models.DBResponse
	query := bson.M{"_id": res.InsertedID}

	err = c.mongoCollection.FindOne(c.ctx, query).Decode(&newUser)
	if err != nil {
		return nil, err
	}
	return newUser, nil
}
