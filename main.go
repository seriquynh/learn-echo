package main

import (
	"context"
	"fmt"
	"github.com/google/uuid"
	"github.com/labstack/echo/v4"
	"go.mongodb.org/mongo-driver/v2/bson"
	"go.mongodb.org/mongo-driver/v2/mongo"
	"go.mongodb.org/mongo-driver/v2/mongo/options"
	"net/http"
)

func MongoConnect() (*mongo.Client, error) {
	MongoDsn := "mongodb://root:root@127.0.0.1:28001/learn-echo?authSource=admin"

	serverAPI := options.ServerAPI(options.ServerAPIVersion1)
	opts := options.Client().ApplyURI(MongoDsn).SetServerAPIOptions(serverAPI)

	client, err := mongo.Connect(opts)
	if err != nil {
		panic(err)
	}
	//defer func() {
	//      if err = client.Disconnect(context.TODO()); err != nil {
	//              panic(err)
	//      }
	//}()

	// Send a ping to confirm a successful connection
	var result bson.M
	if err := client.Database("admin").RunCommand(context.TODO(), bson.D{{"ping", 1}}).Decode(&result); err != nil {
		panic(err)
	}
	fmt.Println("Pinged your deployment. You successfully connected to MongoDB!")

	return client, err
}

type User struct {
	Uuid  string
	Name  string
	Email string
}

func main() {
	var mongoClient, _ = MongoConnect()

	e := echo.New()
	e.POST("/api/users", func(c echo.Context) error {
		name := c.FormValue("name")
		email := c.FormValue("email")

		coll := mongoClient.Database("learn-echo").Collection("users")
		user := User{Uuid: uuid.New().String(), Name: name, Email: email}

		result, err := coll.InsertOne(context.TODO(), user)

		fmt.Println(result, err)

		data := map[string]interface{}{
			"id":    user.Uuid,
			"name":  user.Name,
			"email": user.Email,
		}

		body := map[string]interface{}{
			"data": data,
		}

		err2 := c.JSON(http.StatusOK, body)
		if err2 != nil {
			return err2
		}

		return nil
	})

	e.GET("/api/users/:user", func(c echo.Context) error {
		coll := mongoClient.Database("learn-echo").Collection("users")

		userUuid := c.Param("user")

		filter := bson.D{{"uuid", userUuid}}
		var result User
		err := coll.FindOne(context.TODO(), filter).Decode(&result)

		if err != nil {
			return c.JSON(http.StatusNotFound, map[string]string{"message": "User not found"})
		}

		data := map[string]interface{}{
			"id":    result.Uuid,
			"name":  result.Name,
			"email": result.Email,
		}

		body := map[string]interface{}{
			"data": data,
		}

		return c.JSON(http.StatusOK, body)
	})

	e.Logger.Fatal(e.Start(":1323"))
}
