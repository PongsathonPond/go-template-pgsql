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
	var coll = "menus"
	data := []*domain.Menu{
		{
			KeySlug: "dashboard",
			Name:    "Dashboard",
			Icon:    "dashboard",
			SubMenus: []*domain.SubMenu{
				{
					KeySlug:          "ga",
					Name:             "Google Analytics",
					Icon:             "ga",
					Path:             "/ga",
					Sort:             1,
					CanSetPermission: 2,
					Status:           "active",
				},
				{
					KeySlug:          "monitoring",
					Name:             "Monitoring",
					Icon:             "monitor",
					Path:             "/monitoring",
					Sort:             2,
					CanSetPermission: 2,
					Status:           "active",
				},
			},
			Sort:   1,
			Status: "active",
		},
		{
			KeySlug: "cms",
			Name:    "Contents Management",
			Icon:    "cms",
			SubMenus: []*domain.SubMenu{
				{
					KeySlug:          "slides",
					Name:             "Slides",
					Icon:             "slides",
					Path:             "/slides",
					Sort:             1,
					CanSetPermission: 15,
					Status:           "active",
				},
				{
					KeySlug:          "content",
					Name:             "Content",
					Icon:             "content",
					Path:             "/content",
					Sort:             2,
					CanSetPermission: 15,
					Status:           "active",
				},
				{
					KeySlug:          "category",
					Name:             "Category",
					Icon:             "category",
					Path:             "/category",
					Sort:             3,
					CanSetPermission: 15,
					Status:           "active",
				},
			},
			Sort:   2,
			Status: "active",
		},
		{
			KeySlug: "users-management",
			Name:    "Users Management",
			Icon:    "user-management",
			SubMenus: []*domain.SubMenu{
				{
					KeySlug:          "users",
					Name:             "Users",
					Icon:             "user",
					Path:             "/users",
					Sort:             1,
					CanSetPermission: 15,
					Status:           "active",
				},
				{
					KeySlug:          "roles-and-permissions",
					Name:             "Roles & Permissions",
					Icon:             "roles-permission-management",
					Path:             "/roles-and-permissions",
					Sort:             2,
					CanSetPermission: 15,
					Status:           "active",
				},
			},
			Sort:   3,
			Status: "active",
		},
	}

	var up migrate.MigrationFunc = func(db *mongo.Database) error {
		for _, v := range data {
			filter := bson.M{"key_slug": v.KeySlug}
			cnt, _ := db.Collection(coll).CountDocuments(ctx, filter)
			if cnt == 0 {
				v.ID = uuid.Generate()
				v.CreatedAt = carbon.Now().Unix()
				v.UpdatedAt = carbon.Now().Unix()
				_, err := db.Collection(coll).InsertOne(ctx, v)
				if err != nil {
					return err
				}
				log.Info(fmt.Sprintf("- Migrate menu \"%s\" success", v.Name))
			}
		}
		return nil
	}
	var down migrate.MigrationFunc = func(db *mongo.Database) error {
		for _, v := range data {
			filter := bson.M{"key_slug": v.KeySlug}
			_, err := db.Collection(coll).DeleteOne(ctx, filter)
			if err != nil {
				return err
			}
			log.Info(fmt.Sprintf("- Rollback menu \"%s\" success", v.Name))
		}
		return nil
	}

	if err := migrate.Register(up, down); err != nil {
		log.Fatal(err)
	}

}
