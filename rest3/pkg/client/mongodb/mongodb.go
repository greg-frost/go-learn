package mongodb

import (
	"context"

	"go.mongodb.org/mongo-driver/mongo"
)

func NewClient(ctx context.Context, host, port, username, password, database, authDB string) (*mongo.Database, error) {
	// var mongoDBURL string
	// if username == "" && password == "" {
	// 	mongoDBURL = "mongodb://%s:%s"
	// } else {
	// 	mongoDBURL = "mongodb://%s:%s@%s:%s"
	// }

	return nil, nil
}
