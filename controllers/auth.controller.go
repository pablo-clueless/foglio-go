package controllers

import (
	"context"
	"foglio/common"
	"foglio/models"
	"foglio/responses"
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"gopkg.in/asaskevich/govalidator.v9"
)

var auth *mongo.Collection = common.GetCollection(common.DB, "users")

func Login() gin.HandlerFunc {
	return func(c *gin.Context) {
		ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
		var credentials *models.Auth
		var user *models.User
		defer cancel()

		if err := c.BindJSON(&credentials); err != nil {
			c.JSON(http.StatusUnprocessableEntity, responses.GenericResponse{
				Status:  http.StatusUnprocessableEntity,
				Message: "Invalid JSON body!",
				Data:    map[string]interface{}{"data": err.Error()},
			})
		}

		if !govalidator.IsEmail(credentials.Email) {
			c.JSON(http.StatusBadRequest, responses.GenericResponse{
				Status:  http.StatusBadRequest,
				Message: "Invalid email!",
				Data:    map[string]interface{}{},
			})
		}

		err := auth.FindOne(ctx, bson.M{"email": credentials.Email}).Decode(&user)
		if err != nil {
			c.JSON(http.StatusNotFound, responses.GenericResponse{
				Status:  http.StatusNotFound,
				Message: "User not found!",
				Data:    map[string]interface{}{},
			})
		}

		id := user.ID.Hex()
		token, jwtErr := common.JwtGenerateToken(id)
		if jwtErr != nil {
			c.JSON(http.StatusNotFound, responses.GenericResponse{
				Status:  http.StatusNotFound,
				Message: "User not found!",
				Data:    map[string]interface{}{},
			})
		}

		c.JSON(http.StatusOK, responses.GenericResponse{
			Status:  http.StatusOK,
			Message: "a magic link has been sent to your email!",
			Data:    map[string]interface{}{"data": token},
		})
	}
}

func Verify() gin.HandlerFunc {
	return func(c *gin.Context) {
		ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
		token := c.Query("token")
		var user *models.User
		defer cancel()

		claims, err := common.JwtVerifyToken(token)
		if err != nil {
			c.JSON(http.StatusUnauthorized, responses.GenericResponse{
				Status:  http.StatusUnauthorized,
				Message: "Invalid token!",
				Data:    map[string]interface{}{"data": err.Error()},
			})
		}

		id := claims.Id
		err = auth.FindOne(ctx, bson.M{"id": id}).Decode(&user)
		if err != nil {
			c.JSON(http.StatusNotFound, responses.GenericResponse{
				Status:  http.StatusNotFound,
				Message: "User not found!",
				Data:    map[string]interface{}{},
			})
		}

		c.JSON(http.StatusOK, responses.GenericResponse{
			Status:  http.StatusOK,
			Message: "a magic link has been sent to your email!",
			Data:    map[string]interface{}{"data": user},
		})
	}
}
