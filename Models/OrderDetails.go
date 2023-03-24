package model

import (
	"context"
	"log"
	"time"

	mongoinit "ProductService/Mongo"

	"go.mongodb.org/mongo-driver/bson"
)

type OrderDetails struct {
	OrderNo      string    `json:"order_no" bson:"order_no"`
	OrderStatus  string    `json:"order_status" bson:"order_status"`
	DispatchDate time.Time `json:"dispatch_date" bson:"dispatch_date"`
	OrderPrice   int64     `json:"order_price" bson:"order_price"`
}

func (a *OrderDetails) CollectionName() string {
	return "OrderDetails"
}

func InsertOrderDetails(c *OrderDetails) error {
	if _, err := mongoinit.MongoDB.Collection(c.CollectionName()).InsertOne(context.TODO(), c); err != nil {
		return err
	}
	return nil
}

func UpdateOrderState(orderNo, orderStatus string, DispatchDate time.Time) error {
	var c OrderDetails
	filter := bson.D{{Key: "order_no", Value: orderNo}}
	update := bson.D{{Key: "$set", Value: bson.D{{Key: "order_status", Value: orderStatus}}}, {Key: "$set", Value: bson.D{{Key: "dispatch_date", Value: DispatchDate}}}}
	_, err := mongoinit.MongoDB.Collection(c.CollectionName()).UpdateOne(context.TODO(), filter, update)
	if err != nil {
		log.Fatal(err)
	}
	return nil
}
