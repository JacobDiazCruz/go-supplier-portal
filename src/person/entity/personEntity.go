package person

import "go.mongodb.org/mongo-driver/bson/primitive"

type Employee struct {
	ID      primitive.ObjectID `json:"id" bson:"_id,omitempty"`
	Email   string             `json:"email"`
	Contact string             `json:"contact"`
	Company string             `json:"company"`
}
