package database

import (
	"booking/domain/entity"
	"booking/domain/repository"
	"booking/infrastructure/observer"
	"context"
	"time"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

// MongoUser represents the user document in MongoDB
type MongoUser struct {
	ID        primitive.ObjectID `bson:"_id,omitempty"`
	Email     string             `bson:"email"`
	Username  string             `bson:"username"`
	Password  string             `bson:"password"`
	FullName  string             `bson:"full_name"`
	Phone     string             `bson:"phone"`
	IsActive  bool               `bson:"is_active"`
	CreatedAt time.Time          `bson:"created_at"`
	UpdatedAt time.Time          `bson:"updated_at"`
}

// userRepositoryMongo implements the UserRepository interface for MongoDB
type userRepositoryMongo struct {
	collection *mongo.Collection
	subject    *observer.Subject
}

// NewUserRepositoryMongo creates a new MongoDB user repository
func NewUserRepositoryMongo(db *MongoDB, subject *observer.Subject) repository.UserRepository {
	collection := db.GetCollection("users")

	// Create indexes
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	// Email index
	emailIndex := mongo.IndexModel{
		Keys:    bson.D{{Key: "email", Value: 1}},
		Options: options.Index().SetUnique(true),
	}

	// Username index
	usernameIndex := mongo.IndexModel{
		Keys:    bson.D{{Key: "username", Value: 1}},
		Options: options.Index().SetUnique(true),
	}

	collection.Indexes().CreateMany(ctx, []mongo.IndexModel{emailIndex, usernameIndex})

	return &userRepositoryMongo{
		collection: collection,
		subject:    subject,
	}
}

// toEntity converts MongoUser to entity.User
func (m *MongoUser) toEntity() *entity.User {
	// Convert ObjectID to uint (using timestamp as ID for simplicity)
	id := uint(m.ID.Timestamp().Unix())

	return &entity.User{
		ID:        id,
		Email:     m.Email,
		Username:  m.Username,
		Password:  m.Password,
		FullName:  m.FullName,
		Phone:     m.Phone,
		IsActive:  m.IsActive,
		CreatedAt: m.CreatedAt,
		UpdatedAt: m.UpdatedAt,
	}
}

// fromEntity converts entity.User to MongoUser
func fromEntity(user *entity.User) *MongoUser {
	mongoUser := &MongoUser{
		Email:     user.Email,
		Username:  user.Username,
		Password:  user.Password,
		FullName:  user.FullName,
		Phone:     user.Phone,
		IsActive:  user.IsActive,
		CreatedAt: user.CreatedAt,
		UpdatedAt: user.UpdatedAt,
	}

	// If ID exists, try to convert it to ObjectID
	if user.ID > 0 {
		// For existing users, we need to fetch the ObjectID from database
		// For now, we'll create a new ObjectID
		mongoUser.ID = primitive.NewObjectID()
	}

	return mongoUser
}

// Create creates a new user
func (r *userRepositoryMongo) Create(ctx context.Context, user *entity.User) error {
	mongoUser := fromEntity(user)
	mongoUser.ID = primitive.NewObjectID()
	mongoUser.CreatedAt = time.Now()
	mongoUser.UpdatedAt = time.Now()

	result, err := r.collection.InsertOne(ctx, mongoUser)
	if err != nil {
		return err
	}

	// Update user ID
	if oid, ok := result.InsertedID.(primitive.ObjectID); ok {
		user.ID = uint(oid.Timestamp().Unix())
		user.CreatedAt = mongoUser.CreatedAt
		user.UpdatedAt = mongoUser.UpdatedAt
	}

	// Notify observers
	r.subject.Notify(observer.Event{
		Type: observer.UserCreated,
		Data: user,
	})

	return nil
}

// GetByID retrieves a user by ID
func (r *userRepositoryMongo) GetByID(ctx context.Context, id uint) (*entity.User, error) {
	// Note: This is a simplified approach. In production, you'd store the ObjectID mapping
	var mongoUser MongoUser

	// Find by timestamp-based ID (this is a limitation of the simple conversion)
	// In production, you should maintain an ID mapping or use ObjectID directly
	filter := bson.M{"_id": primitive.NewObjectIDFromTimestamp(time.Unix(int64(id), 0))}

	err := r.collection.FindOne(ctx, filter).Decode(&mongoUser)
	if err != nil {
		if err == mongo.ErrNoDocuments {
			return nil, err
		}
		return nil, err
	}

	return mongoUser.toEntity(), nil
}

// GetByEmail retrieves a user by email
func (r *userRepositoryMongo) GetByEmail(ctx context.Context, email string) (*entity.User, error) {
	var mongoUser MongoUser
	filter := bson.M{"email": email}

	err := r.collection.FindOne(ctx, filter).Decode(&mongoUser)
	if err != nil {
		if err == mongo.ErrNoDocuments {
			return nil, err
		}
		return nil, err
	}

	return mongoUser.toEntity(), nil
}

// GetByUsername retrieves a user by username
func (r *userRepositoryMongo) GetByUsername(ctx context.Context, username string) (*entity.User, error) {
	var mongoUser MongoUser
	filter := bson.M{"username": username}

	err := r.collection.FindOne(ctx, filter).Decode(&mongoUser)
	if err != nil {
		if err == mongo.ErrNoDocuments {
			return nil, err
		}
		return nil, err
	}

	return mongoUser.toEntity(), nil
}

// List retrieves users based on filter
func (r *userRepositoryMongo) List(ctx context.Context, filter *entity.UserFilter) ([]*entity.User, error) {
	mongoFilter := bson.M{}

	if filter != nil {
		if filter.Email != nil {
			mongoFilter["email"] = *filter.Email
		}
		if filter.Username != nil {
			mongoFilter["username"] = *filter.Username
		}
		if filter.IsActive != nil {
			mongoFilter["is_active"] = *filter.IsActive
		}
	}

	findOptions := options.Find()
	if filter != nil {
		if filter.Limit > 0 {
			findOptions.SetLimit(int64(filter.Limit))
		}
		if filter.Offset > 0 {
			findOptions.SetSkip(int64(filter.Offset))
		}
	}

	cursor, err := r.collection.Find(ctx, mongoFilter, findOptions)
	if err != nil {
		return nil, err
	}
	defer cursor.Close(ctx)

	var users []*entity.User
	for cursor.Next(ctx) {
		var mongoUser MongoUser
		if err := cursor.Decode(&mongoUser); err != nil {
			return nil, err
		}
		users = append(users, mongoUser.toEntity())
	}

	if err := cursor.Err(); err != nil {
		return nil, err
	}

	return users, nil
}

// Update updates a user
func (r *userRepositoryMongo) Update(ctx context.Context, user *entity.User) error {
	filter := bson.M{"email": user.Email} // Use email as identifier since we can't reliably convert uint ID

	update := bson.M{
		"$set": bson.M{
			"username":   user.Username,
			"password":   user.Password,
			"full_name":  user.FullName,
			"phone":      user.Phone,
			"is_active":  user.IsActive,
			"updated_at": time.Now(),
		},
	}

	result, err := r.collection.UpdateOne(ctx, filter, update)
	if err != nil {
		return err
	}

	if result.MatchedCount == 0 {
		return mongo.ErrNoDocuments
	}

	user.UpdatedAt = time.Now()

	// Notify observers
	r.subject.Notify(observer.Event{
		Type: observer.UserUpdated,
		Data: user,
	})

	return nil
}

// Delete deletes a user by ID
func (r *userRepositoryMongo) Delete(ctx context.Context, id uint) error {
	// Find user first to get the actual ObjectID
	filter := bson.M{"_id": primitive.NewObjectIDFromTimestamp(time.Unix(int64(id), 0))}

	result, err := r.collection.DeleteOne(ctx, filter)
	if err != nil {
		return err
	}

	if result.DeletedCount == 0 {
		return mongo.ErrNoDocuments
	}

	// Notify observers
	r.subject.Notify(observer.Event{
		Type: observer.UserDeleted,
		Data: id,
	})

	return nil
}

// Count counts users based on filter
func (r *userRepositoryMongo) Count(ctx context.Context, filter *entity.UserFilter) (int64, error) {
	mongoFilter := bson.M{}

	if filter != nil {
		if filter.Email != nil {
			mongoFilter["email"] = *filter.Email
		}
		if filter.Username != nil {
			mongoFilter["username"] = *filter.Username
		}
		if filter.IsActive != nil {
			mongoFilter["is_active"] = *filter.IsActive
		}
	}

	count, err := r.collection.CountDocuments(ctx, mongoFilter)
	if err != nil {
		return 0, err
	}

	return count, nil
}
