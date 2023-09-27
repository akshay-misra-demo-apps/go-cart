package services

import (
	"context"
	json2 "encoding/json"
	"errors"
	"fmt"
	"log"
	"strings"
	"time"

	"github.com/akshay-misra-demo-apps/go-cart/interfaces"
	"github.com/akshay-misra-demo-apps/go-cart/models"
	"github.com/akshay-misra-demo-apps/go-cart/repositories"
	"github.com/akshay-misra-demo-apps/go-cart/utils"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"golang.org/x/crypto/bcrypt"
)

type UserService struct {
	repository repositories.IRepository
}

func GetUserService(repository repositories.IRepository) interfaces.IUser {
	return &UserService{
		repository: repository,
	}
}

func (u *UserService) Register(user *models.User) (*models.SignupResponse, error) {

	user.Id = primitive.NewObjectID()
	user.Email = strings.ToLower(user.Email)
	user.PasswordConfirm = ""
	user.Verified = false
	user.Role = "user"
	user.Status = "active"
	user.CreatedAt = time.Now()
	user.UpdatedAt = time.Now()

	hashed, err := utils.Encrypt(user.Password, bcrypt.DefaultCost)
	if err != nil {
		return nil, fmt.Errorf("error while inserting new user to database: %v", err)
	}
	user.Password = hashed

	_, err = u.repository.GetCollection().InsertOne(context.TODO(), user)
	if err != nil {
		if er, ok := err.(mongo.WriteException); ok && er.WriteErrors[0].Code == 11000 {
			return nil, errors.New("user with that email already exist")
		}
		return nil, fmt.Errorf("error while inserting new user to database: %v", err)
	}

	// Create a unique index for the email field
	opt := options.Index()
	opt.SetUnique(true)
	index := mongo.IndexModel{Keys: bson.M{"email": 1}, Options: opt}

	if _, err := u.repository.GetCollection().Indexes().CreateOne(context.TODO(), index); err != nil {
		return nil, errors.New("could not create index for email")
	}

	response := models.SignupResponse{
		Name:      user.Name,
		Email:     user.Email,
		CreatedAt: user.CreatedAt,
	}

	return &response, nil

}

func (u *UserService) Authenticate(user *models.Login) (auth *models.LoginResponse, err error) {
	out := &models.User{}

	result := u.repository.GetCollection().FindOne(context.TODO(),
		bson.M{
			"email": user.Email,
		})

	if result.Err() == mongo.ErrNoDocuments {
		return nil, fmt.Errorf("user with email %v not found", user.Email)
	}

	if result.Err() != nil {
		return nil, result.Err()
	}

	result.Decode(out)

	if !utils.IsValidPassword(user.Password, out.Password) {
		return nil, fmt.Errorf("wrong password for user with email: %v", user.Email)
	}

	token, refreshToken, err := utils.GenerateTokenWithRefresh(out.Email, out.Name, out.Role, out.Id.Hex())
	if err != nil {
		return nil, fmt.Errorf("error while generating JWT token for user with email: %v", user.Email)
	}
	auth = &models.LoginResponse{
		Token:        token,
		RefreshToken: refreshToken,
	}

	return
}

func (u *UserService) Logout(id string) error {
	return nil
}
func (u *UserService) Get(id string) (*models.User, error) {
	out := &models.User{}
	objectId, err := primitive.ObjectIDFromHex(id)
	if err != nil {
		return nil, err
	}
	result := u.repository.GetCollection().FindOne(context.TODO(), bson.M{"_id": objectId})

	if result.Err() == mongo.ErrNoDocuments {
		return nil, nil
	}

	if result.Err() != nil {
		return nil, result.Err()
	}

	result.Decode(out)

	return out, nil
}

func (u *UserService) Patch(user *models.User) (*models.User, error) {
	if json, err := json2.Marshal(user); err == nil {
		log.Printf("patched existing user: %v", string(json))
	}

	var updates map[string]interface{}
	json, err := json2.Marshal(user)
	if err != nil {
		return nil, err
	}

	err = json2.Unmarshal(json, &updates)
	if err != nil {
		return nil, err
	}

	patch := bson.M{"$set": updates}
	filter := bson.D{{"_id", user.Id}}

	result, err := u.repository.GetCollection().UpdateOne(context.TODO(), filter, patch)
	if err != nil {
		return nil, fmt.Errorf("error while patching existing product to database: %v", err)
	}

	if result.MatchedCount == 0 {
		return nil, fmt.Errorf("user with id %v not found", user.Id.String())
	}

	return user, nil
}

func (u *UserService) Delete(id string) error {
	objectId, err := primitive.ObjectIDFromHex(id)
	if err != nil {
		return err
	}
	result, err := u.repository.GetCollection().DeleteOne(context.TODO(), bson.M{"_id": objectId})
	if err == mongo.ErrNoDocuments || result.DeletedCount == 0 {
		return fmt.Errorf("user with id %v not found", id)
	}

	if err != nil {
		return err
	}

	return nil
}
