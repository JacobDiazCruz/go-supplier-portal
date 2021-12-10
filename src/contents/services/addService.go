package contents

import (
	"context"
	"fmt"

	entity "gitlab.com/JacobDCruz/supplier-portal/src/contents/entity"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

func AddService(content entity.Content) string {
	fmt.Println(content)
	fmt.Println("add service here id ^")
	// query
	result, err := contentCollection.InsertOne(context.TODO(), content)
	if err != nil {
		panic(err)
	}

	oid := result.InsertedID.(primitive.ObjectID)
	return oid.Hex()
}
