package configs

import (
	"context"
	"log"
	"os"
	"path/filepath"
	"runtime"
	"time"
	"crypto/tls"
	//"fmt"

	"github.com/joho/godotenv"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

var Client *mongo.Client

func LoadEnv() string {
	// Load the environment
	_, filename, _, ok := runtime.Caller(0)
	if !ok {
		log.Fatal("Error: Unable to get the directory of the current file.")
	}
	dir := filepath.Dir(filename)

	// Construct the absolute path to the .env file in the root directory
	envFilePath := filepath.Join(dir, "..", ".env")

	// Load .env file
	if err := godotenv.Load(envFilePath); err != nil {
		log.Fatal("Error loading .env file")
	}
	// Get the MongoDB URI from environment variables
	uri := os.Getenv("MONGODB_URI")
	if uri == "" {
		log.Fatal("MONGODB_URI environment variable not set")
	}
	return uri
}

func Init(){
	serverAPI := options.ServerAPI(options.ServerAPIVersion1)
	uri := LoadEnv()
	clientOptions := options.Client().ApplyURI(uri).
		SetServerAPIOptions(serverAPI).
		SetTLSConfig(&tls.Config{InsecureSkipVerify: true})  // Use InsecureSkipVerify for self-signed certificates only
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

    var err error
    Client, err = mongo.Connect(ctx, clientOptions)
    if err != nil {
        log.Fatalf("Failed to connect to MongoDB: %v", err)
    }

    err = Client.Ping(ctx, nil)
    if err != nil {
        log.Fatalf("Failed to ping MongoDB: %v", err)
    }

    log.Println("Connected to MongoDB!")
}

