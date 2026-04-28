package main

import (
	"context"
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

type User struct {
	ID    primitive.ObjectID `json:"id,omitempty" bosn:"_id,omitempty"`
	Name  string             `json:"name"`
	Email string             `json:"email"`
}

var userCollection *mongo.Collection

func connectDB() {
	client, _ := mongo.NewClient(options.Client().ApplyURI("mongodb://localhost:27017"))

	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	client.Connect(ctx)
	userCollection = client.Database("user_db").Collection("users")
}

// Create a user
func createUser(c *gin.Context) {
	var user User
	if err := c.BindJSON(&user); err != nil {
		c.JSON(http.StatusBadRequest, err.Error())
		return
	}

	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	result, _ := userCollection.InsertOne(ctx, user)
	c.JSON(http.StatusOK, result)
}

// Get all users
func getUsers(c *gin.Context) {
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	cursor, _ := userCollection.Find(ctx, bson.M{})

	var users []User
	cursor.All(ctx, &users)

	c.JSON(http.StatusOK, users)

}

func getUser(c *gin.Context) {
	id := c.Param("id")
	objID, _ := primitive.ObjectIDFromHex(id)

	var user User
	userCollection.FindOne(context.Background(), bson.M{"_id": objID}).Decode(&user)

	c.JSON(http.StatusOK, user)
}

// Update user
func updateUser(c *gin.Context) {
	id := c.Param("id")
	objID, _ := primitive.ObjectIDFromHex(id)

	var user User
	c.BindJSON(&user)

	update := bson.M{"$set": user}
	userCollection.UpdateOne(context.Background(), bson.M{"_id": objID}, update)
	c.JSON(http.StatusOK, gin.H{"message": "updated"})
}

// Delete user
func deleteUser(c *gin.Context) {
	id := c.Param("id")
	objID, _ := primitive.ObjectIDFromHex(id)

	userCollection.DeleteOne(context.Background(), bson.M{"_id": objID})

	c.JSON(http.StatusOK, gin.H{"message": "deleted"})
}

func main() {
	connectDB()

	r := gin.Default()

	r.POST("/users", createUser)
	r.GET("/users", getUsers)
	r.GET("/users/:id", getUser)
	r.PUT("/users/:id", updateUser)
	r.DELETE("/users/:id", deleteUser)

	r.Run(":8080")
}
