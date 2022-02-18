package models

import (
	"reflect"
	"time"

	"go.mongodb.org/mongo-driver/bson/primitive"
)

type Todo struct {
	ID          primitive.ObjectID `json:"_id,omitempty" bson:"_id,omitempty"`
	Title       string             `json:"title" bson:"title"`
	Description string             `json:"description" bson:"description"`
	Completed   bool               `json:"completed" bson:"completed"`
	CreatedAt   time.Time          `json:"created_at" bson:"created_at"`
}

func GetStringFields(st interface{}) []string {
	var fields []string
	t := reflect.TypeOf(st)

	for i := 0; i < t.NumField(); i++ {
		if t.Field(i).Type.Kind() == reflect.String {
			fields = append(fields, t.Field(i).Tag.Get("bson"))
		}
	}

	return fields
}
