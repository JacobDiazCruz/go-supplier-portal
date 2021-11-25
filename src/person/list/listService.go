package person

import (
	"context"
	"log"

	database "gitlab.com/JacobDCruz/supplier-portal/src/config"
	entity "gitlab.com/JacobDCruz/supplier-portal/src/person/entity"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
)

var personCollection *mongo.Collection = database.OpenCollection(database.Client, "newsletter")

// var validate = validator.New()

func GetAllUsers() []entity.Employee {
	cursor, err := personCollection.Find(context.TODO(), bson.M{})
	if err != nil {
		log.Fatal(err)
	}
	employees := []entity.Employee{}
	if err = cursor.All(context.TODO(), &employees); err != nil {
		log.Fatal(err)
	}
	return employees
}
