package controllers

import (
	"context"
    "errors"
	"log"
	"math/rand"
	"strconv"
	"time"
	"net/http"
	"github.com/labstack/echo/v4"
	"github.com/javiburn/AsteroidsDB/configs"
	"github.com/javiburn/AsteroidsDB/models"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
)

func UpdateOneAsteroid(c echo.Context, url string) (models.ExampleResponse, error) {
    var updateData bson.M
	if err := c.Bind(&updateData); err != nil {
		return  models.ExampleResponse{}, echo.NewHTTPError(http.StatusBadRequest, "Invalid update data")
	}

    
    db := configs.Client.Database("asteroids")
    collection := db.Collection("asteroids")
    filter := bson.D{{"id", url}}
    update := bson.D{{"$set", updateData}}
	result, _ := collection.UpdateOne(context.TODO(), filter, update)
    // Check if an error occurred
    if result.MatchedCount == 0 {
        return models.ExampleResponse{}, echo.NewHTTPError(http.StatusNotFound, "Asteroid not found") // Return the error as is
    }
    var asteroid models.ExampleResponse
    ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
    defer cancel() 
    collection.FindOne(ctx, filter).Decode(&asteroid)
    // Return nil and the document if no error occurred and document was updated
    return asteroid, nil
}

func UpdateAsteroid(c echo.Context) error{
    // Access the requested URL
    requestedURL := c.Request().URL.String()
    asteroids, err := UpdateOneAsteroid(c, requestedURL[1:])
    if err != nil {
        return c.JSON(http.StatusNotFound, map[string]string{"error": "not found"})
    }
    return c.JSON(http.StatusOK, asteroids)
}

func DeleteOneAsteroid(url string) error {
    db := configs.Client.Database("asteroids")
    collection := db.Collection("asteroids")
    filter := bson.D{{"id", url}}
    // Deletes the first document that matches the filter
    result, err := collection.DeleteOne(context.TODO(), filter)
    // Check if an error occurred
    if err != nil {
        return err // Return the error as is
    }
    // Check if the document was not found
    if result.DeletedCount == 0 {
        // If no document was deleted, return a 404 Not Found error
        return echo.NewHTTPError(http.StatusNotFound, "Asteroid not found")
    }
    // Return nil if no error occurred and document was deleted
    return nil
}

func DeleteAsteroid(c echo.Context) error{
    // Access the requested URL
    requestedURL := c.Request().URL.String()
    err := DeleteOneAsteroid(requestedURL[1:])
    if err != nil {
        return c.JSON(http.StatusNotFound, map[string]string{"error": "not found"})
    }
    return c.JSON(http.StatusOK, map[string]string{"result": "deleted"})
}

func  GetOneAsteroid(url string) (models.ExampleResponse, error){
    var asteroid models.ExampleResponse
    
    ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
    defer cancel() 
    db := configs.Client.Database("asteroids")
    collection := db.Collection("asteroids")
    filter := bson.M{"id": url}
    err := collection.FindOne(ctx, filter).Decode(&asteroid)
    if err != nil {
		if err == mongo.ErrNoDocuments {
			// If no document is found, return a 404 Not Found error
			return models.ExampleResponse{}, echo.NewHTTPError(http.StatusNotFound, "Asteroid not found")
		}
		return models.ExampleResponse{}, err
	}

	// Return the asteroid and no error
    return asteroid, nil
}

func GetAsteroidByID(c echo.Context) error{
    // Access the requested URL
    requestedURL := c.Request().URL.String()
    asteroids, err := GetOneAsteroid(requestedURL[1:])
    if err != nil {
        return c.JSON(http.StatusNotFound, map[string]string{"error": "not found"})
    }
    return c.JSON(http.StatusOK, asteroids)
}

func GetAllAsteroids() ([]models.ExampleResponse, error) {
    var asteroids []models.ExampleResponse

    ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
    defer cancel()

    db := configs.Client.Database("asteroids")
    collection := db.Collection("asteroids")

    if collection == nil {
        return nil, errors.New("failed to get collection: collection is nil")
    }

    cursor, err := collection.Find(ctx, bson.M{})
    if err != nil {
        return nil, err
    }
    defer cursor.Close(ctx)

    for cursor.Next(ctx) {
        var asteroid models.ExampleResponse
        if err := cursor.Decode(&asteroid); err != nil {
            return nil, err
        }
        asteroids = append(asteroids, asteroid)
    }

    if err := cursor.Err(); err != nil {
        return nil, err
    }

    return asteroids, nil
}

func GetAsteroids(c echo.Context) error {
    asteroids, err := GetAllAsteroids()
    if err != nil {
        return c.JSON(http.StatusInternalServerError, map[string]string{"error": err.Error()})
    }
    return c.JSON(http.StatusOK, asteroids)
}

func PostAsteroid(c echo.Context) error {
    // Bind the input data to ExampleRequest
    exampleRequest := new(models.ExampleRequest)
    if err := c.Bind(&exampleRequest); err != nil {
        return err
    }
    if len(exampleRequest.Name) <= 0 || len(exampleRequest.Discovery_date) <= 0 || exampleRequest.Diameter <= 0{
        return c.JSON(http.StatusUnprocessableEntity, map[string]string{"error": "invalid input"})
    }
    date, err := time.Parse("02-01-2006", string(exampleRequest.Discovery_date))
    if err != nil || !time.Now().After(date) {
        return c.JSON(http.StatusUnprocessableEntity, map[string]string{"error": "invalid input"})
    }
    ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
    defer cancel()
    db := configs.Client.Database("asteroids")
    collection := db.Collection("asteroids")
    id := strconv.Itoa(rand.Intn(99999))
    _, err = collection.InsertOne(ctx, bson.D{
        {Key: "id", Value: id},
        {Key: "name", Value: exampleRequest.Name},
        {Key: "diameter", Value: exampleRequest.Diameter},
        {Key: "discovery_date", Value: exampleRequest.Discovery_date},
        {Key: "observations", Value: exampleRequest.Observations},
        {Key: "distances", Value: exampleRequest.Distances},
    })
    if err != nil {
        log.Printf("Error inserting document: %v", err)
        return c.JSON(http.StatusInternalServerError, map[string]string{"error": err.Error()})
    }
    // GET the asteroid just created to return it
    var asteroid models.ExampleResponse
    filter := bson.M{"id": id}
    collection.FindOne(ctx, filter).Decode(&asteroid)
    return c.JSON(http.StatusCreated, asteroid)
}