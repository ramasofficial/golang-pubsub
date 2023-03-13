package main

import (
	"context"
	"log"
	"os"

	"cloud.google.com/go/pubsub"
	"github.com/joho/godotenv"
)

var customLoggerInfo *log.Logger
var customLoggerError *log.Logger

func main() {
	err := godotenv.Load()

	// log.SetOutput(os.Stdout)
	customLoggerInfo = log.New(os.Stdout, getEnv("ENV")+".INFO: ", log.Ldate|log.Ltime|log.Lshortfile)
	customLoggerError = log.New(os.Stdout, getEnv("ENV")+".ERROR: ", log.Ldate|log.Ltime|log.Lshortfile)

	if err != nil {
		customLoggerError.Fatal("Error loading .env file")
	}

	ctx := context.Background()

	pubsubClient, err := pubsub.NewClient(ctx, getEnv("GCP_PROJECT_ID"))

	if err != nil {
		customLoggerError.Fatal(err)
	}

	defer pubsubClient.Close()

	customLoggerInfo.Println("Starting to read subscription...")

	subscription := pubsubClient.Subscription(getEnv("SUBSCRIPTION"))

	err = subscription.Receive(ctx, func(ctx context.Context, m *pubsub.Message) {
		customLoggerInfo.Printf("Received message: %s", m.PublishTime)
		m.Ack()
	})

	if err != nil {
		customLoggerError.Printf("Error: subscription.Receive: %v", err)
	}
}

func getEnv(k string) string {
	v := os.Getenv(k)

	if v == "" {
		customLoggerError.Fatalf("%s environment variable is not set.", k)
	}

	return v
}
