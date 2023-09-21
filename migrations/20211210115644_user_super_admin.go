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
	var coll = "users"
	var up migrate.MigrationFunc = func(db *mongo.Database) error {
		cnt, _ := db.Collection(coll).CountDocuments(ctx, nil)
		if cnt == 0 {
			data := &domain.Users{
				ID:           uuid.Generate(),
				FirstName:    "Super Admin",
				LastName:     "iDev",
				Email:        "admin@idev.com",
				Username:     "admin@idev.com",
				Password:     encrypt.HashPassword("adminidev1234"),
				Role:         "1400000000000000001",
				ImageProfile: "https://www.pngmart.com/files/21/Admin-Profile-PNG-Clipart.png",
				Status:       "active",
				CreatedAt:    carbon.Now().Unix(),
				UpdatedAt:    carbon.Now().Unix(),
			}
			_, err := db.Collection(coll).InsertOne(ctx, data)
			if err != nil {
				return err
			}
			log.Info(fmt.Sprintf("- Migrate user \"%s\" success", data.FirstName))
		}
		return nil
	}
	var down migrate.MigrationFunc = func(db *mongo.Database) error {
		filter := bson.M{"username": "admin@idev.com"}
		_, err := db.Collection(coll).DeleteOne(ctx, filter)
		if err != nil {
			return err
		}
		log.Info(fmt.Sprintf("- Rollback user \"%s\" success", "admin@idev.com"))
		return nil
	}

	if err := migrate.Register(up, down); err != nil {
		log.Fatal(err)
	}
}
