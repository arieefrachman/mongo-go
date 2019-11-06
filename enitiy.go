package entity

import (
	"go.mongodb.org/mongo-driver/bson/primitive"
)

// ID type 
type ID = primitive.ObjectID


// StringToID convert string to ID
func StringToID(s string) ID {
	result, err := primitive.ObjectIDFromHex(s)

	if err != nil {
		return primitive.NilObjectID
	}

	return result
}