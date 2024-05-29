package handlers

import (
	"context"
	"net/http"

	"github.com/gin-gonic/gin"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"

	"task-manager/config"
	"task-manager/models"
	"task-manager/utils"
)

var taskCollection *mongo.Collection

func init() {
	clientOptions := options.Client().ApplyURI("mongodb://localhost:27017")
	client, err := mongo.Connect(context.Background(), clientOptions)
	if err != nil {
		panic(err)
	}

	taskCollection = client.Database(config.DBName).Collection(config.Collection)
}

func CreateTask(c *gin.Context) {
	var task models.Task
	if err := c.ShouldBindJSON(&task); err != nil {
		utils.ErrorResponse(c.Writer, http.StatusBadRequest, "Invalid request payload")
		return
	}

	
	// Convert string user ID to primitive.ObjectID
	userID, err := primitive.ObjectIDFromHex(c.GetString("user_id"))
	if err != nil {
		utils.ErrorResponse(c.Writer, http.StatusInternalServerError, "Failed to convert user ID")
		return
	}

	task.UserID = userID
	task.ID = primitive.NewObjectID()
	task.Status = "Pending"

	_, err = taskCollection.InsertOne(context.Background(), task)
	if err != nil {
		utils.ErrorResponse(c.Writer, http.StatusInternalServerError, "Failed to create task")
		return
	}

	c.JSON(http.StatusCreated, task)

	// task.UserID = c.GetString("user_id")
	// task.ID = primitive.NewObjectID()
	// task.Status = "Pending"

	// _, err := taskCollection.InsertOne(context.Background(), task)
	// if err != nil {
	// 	utils.ErrorResponse(c.Writer, http.StatusInternalServerError, "Failed to create task")
	// 	return
	// }

	// c.JSON(http.StatusCreated, task)
}

func GetTask(c *gin.Context) {
	taskID := c.Param("id")

	var task models.Task
	err := taskCollection.FindOne(context.Background(), bson.M{"_id": taskID, "user_id": c.GetString("user_id")}).Decode(&task)
	if err != nil {
		utils.ErrorResponse(c.Writer, http.StatusNotFound, "Task not found")
		return
	}

	c.JSON(http.StatusOK, task)
}

func UpdateTask(c *gin.Context) {
	taskID := c.Param("id")

	var task models.Task
	if err := c.ShouldBindJSON(&task); err != nil {
		utils.ErrorResponse(c.Writer, http.StatusBadRequest, "Invalid request payload")
		return
	}

	update := bson.M{"$set": task}
	_, err := taskCollection.UpdateOne(context.Background(), bson.M{"_id": taskID, "user_id": c.GetString("user_id")}, update)
	if err != nil {
		utils.ErrorResponse(c.Writer, http.StatusInternalServerError, "Failed to update task")
		return
	}

	c.Status(http.StatusNoContent)
}

func DeleteTask(c *gin.Context) {
	taskID := c.Param("id")

	_, err := taskCollection.DeleteOne(context.Background(), bson.M{"_id": taskID, "user_id": c.GetString("user_id")})
	if err != nil {
		utils.ErrorResponse(c.Writer, http.StatusInternalServerError, "Failed to delete task")
		return
	}

	c.Status(http.StatusNoContent)
}

func ListTasks(c *gin.Context) {
	options := options.Find()
	options.SetLimit(10)

	cursor, err := taskCollection.Find(context.Background(), bson.M{"user_id": c.GetString("user_id")}, options)
	if err != nil {
		utils.ErrorResponse(c.Writer, http.StatusInternalServerError, "Failed to fetch tasks")
		return
	}
	defer cursor.Close(context.Background())

	var tasks []models.Task
	if err := cursor.All(context.Background(), &tasks); err != nil {
		utils.ErrorResponse(c.Writer, http.StatusInternalServerError, "Failed to fetch tasks")
		return
	}

	c.JSON(http.StatusOK, tasks)
}
