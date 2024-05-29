package handlers

import (
	"context"
	"net/http"

	"github.com/dgrijalva/jwt-go"
	"github.com/gin-gonic/gin"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"golang.org/x/crypto/bcrypt"

	"task-manager/models"
	"task-manager/utils"
)

var userCollection *mongo.Collection

func init() {
    clientOptions := options.Client().ApplyURI("mongodb://localhost:27017")
    client, err := mongo.Connect(context.Background(), clientOptions)
    if err != nil {
        panic(err)
    }

    userCollection = client.Database("task_manager").Collection("users")
}

func Register(c *gin.Context) {
    var user models.User
    if err := c.ShouldBindJSON(&user); err != nil {
        utils.ErrorResponse(c.Writer, http.StatusBadRequest, "Invalid request payload")
        return
    }

    hash, err := bcrypt.GenerateFromPassword([]byte(user.Password), bcrypt.DefaultCost)
    if err != nil {
        utils.ErrorResponse(c.Writer, http.StatusInternalServerError, "Failed to generate password hash")
        return
    }

    user.Password = string(hash)

    result, err := userCollection.InsertOne(context.Background(), user)
    if err != nil {
        utils.ErrorResponse(c.Writer, http.StatusInternalServerError, "Failed to insert user")
        return
    }

    c.JSON(http.StatusCreated, result.InsertedID)
}

func Login(c *gin.Context) {
    var user models.User
    if err := c.ShouldBindJSON(&user); err != nil {
        utils.ErrorResponse(c.Writer, http.StatusBadRequest, "Invalid request payload")
        return
    }

    var result models.User
    err := userCollection.FindOne(context.Background(), bson.M{"username": user.Username}).Decode(&result)
    if err != nil {
        utils.ErrorResponse(c.Writer, http.StatusUnauthorized, "Invalid credentials")
        return
    }

    err = bcrypt.CompareHashAndPassword([]byte(result.Password), []byte(user.Password))
    if err != nil {
        utils.ErrorResponse(c.Writer, http.StatusUnauthorized, "Invalid credentials")
        return
    }

    // Generate JWT token
    token := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
        "user_id": result.ID.Hex(),
    })

    tokenString, err := token.SignedString([]byte("secret"))
    if err != nil {
        utils.ErrorResponse(c.Writer, http.StatusInternalServerError, "Failed to generate token")
        return
    }

    c.JSON(http.StatusOK, gin.H{"token": tokenString})
}
