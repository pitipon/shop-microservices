package utils

import (
	"go.mongodb.org/mongo-driver/v2/bson"
)

func ConvertToObjectId(id string) bson.ObjectID {
	oid, err := bson.ObjectIDFromHex(id)
	if err != nil {
		return bson.ObjectID{}
	}
	return oid
}
