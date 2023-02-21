package getir

import (
	"context"
	"encoding/json"
	"getir/models"
	writeresponse "getir/writeResponse"
	"log"
	"net/http"
	"time"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
)

const (
	// DBName is the name of the database
	DBName = "getircase-study"
	// CollectionName is the name of the collection
	CollectionName = "records"
)

type MongoConnection struct {
	Client *mongo.Client
}

// GetirHandler is the handler function for the getir endpoint
func (c *MongoConnection) GetirHandler(w http.ResponseWriter, r *http.Request) {
	if r.Method != "POST" {
		http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
		return
	}

	var params models.RequestParams

	err := json.NewDecoder(r.Body).Decode(&params)
	if err != nil {
		writeresponse.WriteResponse(
			w,
			models.Response{
				Code:    1,
				Message: err.Error(),
				Records: nil,
			},
		)
		return
	}

	collection := c.Client.Database(DBName).Collection(CollectionName)
	filter := &bson.M{
		"createdAt": bson.M{
			"$gte": params.StartDate, "$lte": params.EndDate,
		},
		"totalCount": bson.M{
			"$gte": params.MinCount, "$lte": params.MaxCount,
		},
	}

	var records []models.ResponseItem
	ctx, cancel := context.WithTimeout(context.Background(), 1000*time.Second)
	defer cancel()

	cursor, err := collection.Find(ctx, filter)

	defer func() {
		cursor.Close(ctx)
		cancel()
	}()

	if err != nil {
		writeresponse.WriteResponse(
			w,
			models.Response{
				Code:    1,
				Message: err.Error(),
				Records: nil,
			},
		)
		log.Fatal(err)
	}

	for cursor.Next(ctx) {
		var record models.ResponseItem
		err := cursor.Decode(&record)
		if err != nil {
			response := models.Response{
				Code:    1,
				Message: err.Error(),
				Records: nil,
			}
			writeresponse.WriteResponse(w, response)
			return
		}
		records = append(records, record)
	}
	response := models.Response{
		Code:    0,
		Message: "Success",
		Records: records,
	}
	writeresponse.WriteResponse(w, response)
}
