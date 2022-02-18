package utils

import (
	"fmt"
	"gunmurat7/todo-app-server/models"
	"strconv"

	"github.com/gofiber/fiber/v2"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo/options"
)

var Opt = options.Find()
var SFilter primitive.M

func PaginationWithFiber(c *fiber.Ctx, FindOptions *options.FindOptions) (int64, int64) {
	if c.Query("page") != "" && c.Query("limit") != "" {
		page, _ := strconv.ParseInt(c.Query("page"), 10, 32)
		limit, _ := strconv.ParseInt(c.Query("limit"), 10, 32)
		if page == 1 {
			FindOptions.SetSkip(0)
			FindOptions.SetLimit(limit)
			return page, limit
		}

		FindOptions.SetSkip((page - 1) * limit)
		FindOptions.SetLimit(limit)
		return page, limit
	}
	FindOptions.SetSkip(0)
	FindOptions.SetLimit(0)
	return 1, 0
}

func SearchFilter(key string) {
	if key != "" {
		filter := bson.M{
			"$or": GenerateSearchBsons(key, models.GetStringFields(models.Todo{})),
		}
		SFilter = filter
	} else {
		fmt.Println("key is empty")
		SFilter = bson.M{}
	}

}

func GenerateSearchBsons(key string, fields []string) []primitive.M {
	var bsons []bson.M
	for _, field := range fields {
		bsons = append(bsons, bson.M{field: bson.M{"$regex": primitive.Regex{Pattern: key, Options: "i"}}})
	}
	return bsons
}
