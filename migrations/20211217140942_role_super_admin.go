package migrations

import (
	"fmt"

	"idev-cms-service/domain"

	"github.com/uniplaces/carbon"
	migrate "github.com/xakep666/mongo-migrate"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
)

func init() {
	var coll = "roles"
	data := &domain.Roles{
		ID:          "1400000000000000001",
		Name:        "Super Administrator",
		Description: "For Supper Admin Only",
		Permissions: []domain.MenuPermissions{
			{
				KeySlug:    "ga",
				Permission: 2,
			},
			{
				KeySlug:    "monitoring",
				Permission: 2,
			},
			{
				KeySlug:    "slides",
				Permission: 15,
			},
			{
				KeySlug:    "content",
				Permission: 15,
			},
			{
				KeySlug:    "category",
				Permission: 15,
			},
			{
				KeySlug:    "users",
				Permission: 15,
			},
			{
				KeySlug:    "roles-and-permissions",
				Permission: 15,
			},
		},
		Status:    "active",
		CreatedAt: carbon.Now().Unix(),
		UpdatedAt: carbon.Now().Unix(),
	}

	var up migrate.MigrationFunc = func(db *mongo.Database) error {
		filter := bson.M{"_id": data.ID}
		cnt, _ := db.Collection(coll).CountDocuments(ctx, filter)
		if cnt == 0 {
			_, err := db.Collection(coll).InsertOne(ctx, data)
			if err != nil {
				return err
			}
			log.Info(fmt.Sprintf("- Migrate role \"%s\" success", data.Name))
		}
		return nil
	}
	var down migrate.MigrationFunc = func(db *mongo.Database) error {
		filter := bson.M{"_id": "1400000000000000001"}
		_, err := db.Collection(coll).DeleteOne(ctx, filter)
		if err != nil {
			return err
		}
		log.Info(fmt.Sprintf("- Rollback role \"%s\" success", data.Name))
		return nil
	}

	if err := migrate.Register(up, down); err != nil {
		log.Fatal(err)
	}
}
