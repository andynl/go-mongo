package repository

import (
	"time"

	"github.com/andynl/go-mongo/config"
	"github.com/andynl/go-mongo/src/modules/profile/model"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
)

//profileRepositoryMongo
type profileRepositoryMongo struct {
	db         *mongo.Database
	collection string
}

//NewProfileRepositoryMongo
func NewProfileRepositoryMongo(db *mongo.Database, collection string) *profileRepositoryMongo {
	return &profileRepositoryMongo{
		db:         db,
		collection: collection,
	}
}

//Save
func (r *profileRepositoryMongo) Save(profile *model.Profile) error {
	ctx, cancel := config.NewMongoContext()
	defer cancel()

	_, err := r.db.Collection(r.collection).InsertOne(ctx, profile)
	return err
}

//Update
func (r *profileRepositoryMongo) Update(id string, profile *model.Profile) error {
	ctx, cancel := config.NewMongoContext()
	defer cancel()

	profile.UpdatedAt = time.Now()
	_, err := r.db.Collection(r.collection).UpdateOne(ctx, bson.M{"id": id}, profile)
	return err
}

//Delete
func (r *profileRepositoryMongo) Delete(id string) error {
	ctx, cancel := config.NewMongoContext()
	defer cancel()

	_, err := r.db.Collection(r.collection).DeleteOne(ctx, bson.M{"id": id})
	return err
}

//FindByID
func (r *profileRepositoryMongo) FindByID(id string) (*model.Profile, error) {
	ctx, cancel := config.NewMongoContext()
	defer cancel()

	var profile model.Profile

	err := r.db.Collection(r.collection).FindOne(ctx, bson.M{"id": id}).Decode(&profile)

	if err != nil {
		return nil, err
	}

	return &profile, nil
}

//FindAll
func (r *profileRepositoryMongo) FindAll() (model.Profiles, error) {
	ctx, cancel := config.NewMongoContext()
	defer cancel()

	var profiles model.Profiles

	_, err := r.db.Collection(r.collection).Find(ctx, bson.D{})

	if err != nil {
		return nil, err
	}

	return profiles, nil
}
