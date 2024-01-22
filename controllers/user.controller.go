package controllers

import (
	"context"
	"foglio/common"
	"foglio/models"
	"foglio/responses"
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/go-playground/validator/v10"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
)

var userCollection *mongo.Collection = common.GetCollection(common.DB, "users")
var validate = validator.New()

func GetAllUsers() gin.HandlerFunc {
	return func(c *gin.Context) {
		ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
		var users []models.User
		defer cancel()

		result, err := userCollection.Find(ctx, bson.M{})
		if err != nil {
			c.JSON(http.StatusInternalServerError, responses.GenericResponse{
				Status:  http.StatusInternalServerError,
				Message: "Internal server error!",
				Data:    map[string]interface{}{"data": err.Error()},
			})
		}

		defer result.Close(ctx)
		for result.Next(ctx) {
			var user models.User
			if err = result.Decode(&user); err != nil {
				c.JSON(http.StatusInternalServerError, responses.GenericResponse{
					Status:  http.StatusInternalServerError,
					Message: "Internal server error!",
					Data:    map[string]interface{}{"data": err.Error()},
				})
			}

			users = append(users, user)
		}

		c.JSON(http.StatusOK, responses.GenericResponse{
			Status:  http.StatusOK,
			Message: "All users retrieved!",
			Data:    map[string]interface{}{"data": users},
		})
	}
}

func GetUser() gin.HandlerFunc {
	return func(c *gin.Context) {
		ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
		id := c.Param("id")
		var user *models.User
		defer cancel()

		objectId, _ := primitive.ObjectIDFromHex(id)
		err := userCollection.FindOne(ctx, bson.M{"id": objectId}).Decode(&user)
		if err != nil {
			c.JSON(http.StatusInternalServerError, responses.GenericResponse{
				Status:  http.StatusInternalServerError,
				Message: "Internal server error!",
				Data:    map[string]interface{}{"data": err.Error()},
			})
		}

		c.JSON(http.StatusOK, responses.GenericResponse{
			Status:  http.StatusOK,
			Message: "Invalid JSON body!",
			Data:    map[string]interface{}{"data": user},
		})
	}
}

func UpdateUser() gin.HandlerFunc {
	return func(c *gin.Context) {
		ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
		id := c.Param("id")
		var user *models.User
		defer cancel()

		objectId, _ := primitive.ObjectIDFromHex(id)

		if err := c.BindJSON(&user); err != nil {
			c.JSON(http.StatusUnprocessableEntity, responses.GenericResponse{
				Status:  http.StatusUnprocessableEntity,
				Message: "Invalid JSON body!",
				Data:    map[string]interface{}{"data": err.Error()},
			})
		}

		if validatorErr := validate.Struct(&user); validatorErr != nil {
			c.JSON(http.StatusBadRequest, responses.GenericResponse{
				Status:  http.StatusBadRequest,
				Message: "Invalid JSON body!",
				Data:    map[string]interface{}{"data": validatorErr.Error()},
			})
		}

		update := bson.M{"$set": bson.M{
			"name":           user.Name,
			"role":           user.Role,
			"location":       user.Location,
			"about":          user.About,
			"website":        user.Website,
			"image":          user.Image,
			"contact":        user.Contact,
			"projects":       user.Projects,
			"side_projects":  user.SideProjects,
			"education":      user.Education,
			"experience":     user.Experience,
			"awards":         user.Awards,
			"volunteering":   user.Volunteering,
			"speakings":      user.Speaking,
			"exhibitions":    user.Exhibitions,
			"certifications": user.Certifications,
			"features":       user.Features,
		}}
		result, err := userCollection.UpdateOne(ctx, bson.M{"id": objectId}, bson.M{"$set": update})
		if err != nil {
			c.JSON(http.StatusInternalServerError, responses.GenericResponse{
				Status:  http.StatusInternalServerError,
				Message: "Internal server error!",
				Data:    map[string]interface{}{"data": err.Error()},
			})
		}

		var updatedUser models.User
		if result.MatchedCount == 1 {
			err := userCollection.FindOne(ctx, bson.M{"id": objectId}).Decode(&updatedUser)
			if err != nil {
				c.JSON(http.StatusInternalServerError, responses.GenericResponse{
					Status:  http.StatusInternalServerError,
					Message: "Internal server error!",
					Data:    map[string]interface{}{"data": err.Error()},
				})
			}
		}

		c.JSON(http.StatusOK, responses.GenericResponse{
			Status:  http.StatusOK,
			Message: "User updated!",
			Data:    map[string]interface{}{"data": user},
		})
	}
}

func DeleteUser() gin.HandlerFunc {
	return func(c *gin.Context) {
		ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
		id := c.Param("id")
		defer cancel()

		objectId, _ := primitive.ObjectIDFromHex(id)
		result, err := userCollection.DeleteOne(ctx, bson.M{"id": objectId})
		if err != nil {
			c.JSON(http.StatusInternalServerError, responses.GenericResponse{
				Status:  http.StatusInternalServerError,
				Message: "Internal server error!",
				Data:    map[string]interface{}{"data": err.Error()},
			})
		}

		if result.DeletedCount < 0 {
			c.JSON(http.StatusNotFound, responses.GenericResponse{
				Status:  http.StatusNotFound,
				Message: "User not found!",
				Data:    map[string]interface{}{},
			})
		}

		c.JSON(http.StatusOK, responses.GenericResponse{
			Status:  http.StatusOK,
			Message: "User deleted!",
			Data:    map[string]interface{}{},
		})
	}
}
