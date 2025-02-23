package db

import (
	
	"log"
	"os"

	"github.com/joho/godotenv"
	supa "github.com/nedpals/supabase-go"
	"go.mongodb.org/mongo-driver/v2/mongo"
	"go.mongodb.org/mongo-driver/v2/mongo/options"
	//"go.mongodb.org/mongo-driver/v2/mongo/readpref"
)

var Supabase *supa.Client
var MongoClient   *mongo.Client


func CreateMongoClient() error{
	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error loading .env file")
		return err
		}
		uri := os.Getenv("MONGO_KEY")
		serverAPI := options.ServerAPI(options.ServerAPIVersion1)
	opts := options.Client().ApplyURI(uri).SetServerAPIOptions(serverAPI)
		MongoClient, err = mongo.Connect(opts)
		if err != nil {
			log.Fatal(err)
			return err
		}
		return nil
}
func CreateClient() error {
	// Load environment variables from .env file
	err := godotenv.Load(".env")
	if err != nil {
		log.Println("Warning: Could not load .env file. Using system environment variables.")
	}

	// Get Supabase URL and key from environment variables
	url := os.Getenv("SUPABASE_URL")
	key := os.Getenv("SUPABASE_KEY")

	// Check if URL and key are set
	if url == "" || key == "" {
		log.Fatal("Supabase URL and key must be set in environment variables")
	}

	// Create Supabase client
	Supabase = supa.CreateClient(url, key)
	
	log.Println("Supabase client created successfully")
	return nil
}
