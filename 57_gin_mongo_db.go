package main

import (
	"fmt"
	"log"
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

type Task struct {
	ID        string    `json:"id" bson:"_id"`
	Title     string    `json:"title" bson:"title"`
	Completed bool      `json:"completed" bson:"completed"`
	CreatedAt time.Time `json:"created_at" bson:"created_at"`
}

var tasksCollection *mongo.Collection

func main() {
	clientOptions := options.Client().ApplyURI("mongodb://localhost:27017")
	client, err := mongo.Connect(context.TODO(), clientOptions)
	if err != nil {
		log.Fatal(err)
	}

	err = client.Ping(context.TODO(), nil)
	if err != nil {
		log.Fatal(err)
	}

	tasksCollection = client.Database("myapp").Collection("tasks")

	r := gin.Default()

	r.GET("/tasks", getTasks)
	r.POST("/tasks", createTask)

	r.Run(":8080")
}

func getTasks(c *gin.Context) {
	var tasks []Task
	cursor, err := tasksCollection.Find(context.TODO(), bson.M{})
	if err != nil {
		c.JSON(http.StatusInternalServerError, err)
		return
	}
	defer cursor.Close(context.TODO())

	for cursor.Next(context.TODO()) {
		var task Task
		if err := cursor.Decode(&task); err != nil {
			c.JSON(http.StatusInternalServerError, err)
			return
		}
		tasks = append(tasks, task)
	}

	c.JSON(http.StatusOK, tasks)
}

func createTask(c *gin.Context) {
	var task Task
	if err := c.ShouldBindJSON(&task); err != nil {
		c.JSON(http.StatusBadRequest, err)
		return
	}

	task.ID = bson.NewObjectId().Hex()
	task.Completed = false
	task.CreatedAt = time.Now()

	_, err := tasksCollection.InsertOne(context.TODO(), task)
	if err != nil {
		c.JSON(http.StatusInternalServerError, err)
		return
	}

	c.JSON(http.StatusCreated, task)
}
