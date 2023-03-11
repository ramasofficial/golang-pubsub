package main

import (
	"context"
	"fmt"
	"log"
	"os"

	"cloud.google.com/go/pubsub"
	"github.com/joho/godotenv"
)

func main() {
	err := godotenv.Load()

	if err != nil {
		log.Fatal("Error loading .env file")
	}

	// log.SetOutput(os.Stdout)
	customLogger := log.New(os.Stdout, getEnv("ENV")+".INFO: ", log.Ldate|log.Ltime|log.Lshortfile)

	ctx := context.Background()

	pubsubClient, err := pubsub.NewClient(ctx, getEnv("GCP_PROJECT_ID"))

	if err != nil {
		log.Fatal(err)
	}

	defer pubsubClient.Close()

	customLogger.Println("Starting to read subscription...")

	subscription := pubsubClient.Subscription(getEnv("SUBSCRIPTION"))

	subErr := subscription.Receive(context.Background(), func(ctx context.Context, m *pubsub.Message) {
		customLogger.Printf("Received message: %s", m.PublishTime)
		m.Ack()
	})

	if subErr != nil {
		fmt.Printf("Error: sub.Receive: %v", subErr)
	}
}

func getEnv(k string) string {
	v := os.Getenv(k)

	if v == "" {
		log.Fatalf("%s environment variable is not set.", k)
	}

	return v
}
