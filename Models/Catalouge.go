package model

import (
	"context"
	"log"
	"time"

	mongoinit "ProductService/Mongo"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo/options"
)

type Catalouge struct {
	Id            string    `json:"id" bson:"id"`
	Title         string    `json:"title" bson:"title"`
	Description   string    `json:"description" bson:"description"`
	Category      string    `json:"product_category" bson:"product_category"`
	Brand         string    `json:"brand" bson:"brand"`
	Availablility int       `json:"availability" bson:"availability"`
	Price         int64     `json:"price" bson:"price"`
	SalePrice     int64     `json:"sale_price" bson:"sale_price"`
	CreatedAt     time.Time `json:"created_at" bson:"created_at"`
	UpdatedAt     time.Time `json:"updated_at" bson:"updated_at"`
	IsActive      int       `json:"isActive" bson:"isActive"`
}

func (a *Catalouge) CollectionName() string {
	return "Catalouge"
}

func GetCatalouge(c *Catalouge, id, title string, availability int) error {
	if err := mongoinit.MongoDB.Collection(c.CollectionName()).FindOne(context.TODO(), bson.M{
		"id":           id,
		"title":        title,
		"availability": availability,
	}).Decode(&c); err != nil {
		return err
	}
	return nil
}

func InsertCatalouge(c *Catalouge) error {
	if _, err := mongoinit.MongoDB.Collection(c.CollectionName()).InsertOne(context.TODO(), c); err != nil {
		return err
	}
	return nil
}

func GetAllCatalouge() ([]Catalouge, error) {
	var c Catalouge
	cur, err := mongoinit.MongoDB.Collection(c.CollectionName()).Find(context.TODO(), bson.M{})
	if err != nil {
		log.Fatal(err)
	}
	var results []Catalouge
	for cur.Next(context.TODO()) {
		var elem Catalouge
		err := cur.Decode(&elem)
		if err != nil {
			log.Fatal(err)
		}

		results = append(results, elem)

	}
	return results, nil
}

func GetAllCatalougeByCategory(category string) ([]Catalouge, error) {
	var c Catalouge
	opts := options.Find()
	cur, err := mongoinit.MongoDB.Collection(c.CollectionName()).Find(context.TODO(), bson.M{"product_category": category}, opts)
	if err != nil {
		log.Fatal(err)
	}
	var results []Catalouge
	for cur.Next(context.TODO()) {
		var elem Catalouge
		err := cur.Decode(&elem)
		if err != nil {
			log.Fatal(err)
		}

		results = append(results, elem)

	}
	return results, nil
}

func UpdateCatalouge(id string, NoOfBuyers int) error {
	var c Catalouge
	filter := bson.D{{Key: "id", Value: id}}
	update := bson.D{{Key: "$set", Value: bson.D{{Key: "updated_at", Value: time.Now()}}}}
	_, err := mongoinit.MongoDB.Collection(c.CollectionName()).UpdateOne(context.TODO(), filter, update)
	if err != nil {
		log.Fatal(err)
	}
	return nil
}
