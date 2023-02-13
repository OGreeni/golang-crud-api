package controllers

import (
	"context"
	"fmt"
	"net/http"
	"time"

	"example.com/ogreeni/configs"
	"example.com/ogreeni/models"
	"example.com/ogreeni/responses"
	"github.com/gin-gonic/gin"
	"github.com/go-playground/validator/v10"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
)

var userCollection *mongo.Collection = configs.GetCollection(configs.DB, "users")
var validate = validator.New()

func CreateUser() gin.HandlerFunc {
	return func(c *gin.Context) {
		ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
		var user models.User
		defer cancel()

		// validate request body
		if err := c.BindJSON(&user); err != nil {
			c.JSON(http.StatusBadRequest, responses.UserResponse{
				Status:  http.StatusBadRequest,
				Message: "invalid request body",
				Data:    nil,
			})
			return
		}

		// use the validator library to validate required fields
		if err := validate.Struct(user); err != nil {
			c.JSON(http.StatusBadRequest, responses.UserResponse{
				Status:  http.StatusBadRequest,
				Message: "invalid request body",
				Data:    nil,
			})
			return
		}

		newUser := models.User{
			ID:        user.ID,
			FirstName: user.FirstName,
			LastName:  user.LastName,
			Gender:     user.Gender,
			BirthDate: user.BirthDate,
		}

		// insert user into database
		result, err := userCollection.InsertOne(ctx, newUser)
		if err != nil {
			c.JSON(http.StatusInternalServerError, responses.UserResponse{
				Status:  http.StatusInternalServerError,
				Message: "error creating user",
				Data:    nil,
			})
			fmt.Println(err)
			return
		}

		c.JSON(http.StatusCreated, responses.UserResponse{
			Status:  http.StatusCreated,
			Message: "user created successfully",
			Data:    gin.H{"data": result},
		})

	}
}

func GetUsers() gin.HandlerFunc {
	return func(c *gin.Context) {
		ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
		var users []models.User
		defer cancel()


		filter := bson.M{}
		if (c.Query("gender") == "male" || c.Query("gender") == "female") {
			filter = bson.M{"gender": c.Query("gender")}
		}

		results, err := userCollection.Find(ctx, filter)
		if err != nil {
			c.JSON(http.StatusInternalServerError, responses.UserResponse{
				Status:  http.StatusInternalServerError,
				Message: "error fetching users",
				Data:    nil,
			})
			return
		}

		defer results.Close(ctx)
		for results.Next(ctx) {
			var singleUser models.User
			if err = results.Decode(&singleUser); err != nil {
				c.JSON(http.StatusInternalServerError, responses.UserResponse{
					Status:  http.StatusInternalServerError,
					Message: "error fetching users",
					Data:    nil,
				})
				return
			}
			users = append(users, singleUser)
		}

		c.JSON(http.StatusOK, responses.UserResponse{
			Status:  http.StatusOK,
			Message: "users fetched successfully",
			Data:    gin.H{"data": users},
		})
	}
}
