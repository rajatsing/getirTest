package main

import (
	"context"
	"getir/getir"
	"getir/inmemory"
	"log"
	"net/http"
	"os"
	"os/signal"
	"syscall"
	"time"

	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

const (
	// DBURL is the url of the database
	DBURL = "mongodb+srv://challengeUser:WUMglwNBaydH8Yvu@challenge-xzwqd.mongodb.net/getircase-study?retryWrites=true"
)

type MongoConnection struct {
	Client *mongo.Client
}

/*
	InitDB function will be responsible for setting up the database connection
*/

func InitDB() (*mongo.Client, context.Context, error) {
	var err error
	ctx, cancel := context.WithTimeout(context.Background(), 20*time.Second)
	defer cancel()

	client, err := mongo.NewClient(options.Client().ApplyURI(DBURL))
	if err != nil {
		return nil, nil, err
	}
	err = client.Connect(ctx)
	if err != nil {
		return nil, nil, err
	}

	log.Println("Connected to MongoDB!")
	return client, ctx, nil
}

/*
InitLocalDB function will be responsible for setting up the in memory connection
*/

func InitLocalDB() map[string]string {
	return make(map[string]string)
}

/*
	main function will be responsible for setting up the database connection
	and the http server
*/

func main() {
	// Set up MongoDB connection
	client, ctx, err := InitDB()

	// Set up in memory connection
	localMemory := InitLocalDB()

	connection := getir.MongoConnection{
		Client: client,
	}

	localMem := inmemory.LocalMemory{
		Data: localMemory,
	}

	if err != nil {
		log.Fatal(err)
	}
	defer func() {
		if err = client.Disconnect(ctx); err != nil {
			log.Fatal(err)
		}
	}()

	// Set up HTTP server
	srv := &http.Server{Addr: ":8080"}
	http.HandleFunc("/in-memory", localMem.GetInMemoryHandler)
	http.HandleFunc("/getir", connection.GetirHandler)

	go func() {
		log.Println("Server is listening on port 8080")
		if err := srv.ListenAndServe(); err != nil {
			log.Fatal(err)
		}
	}()

	// Wait for SIGINT or SIGTERM signal
	sigchan := make(chan os.Signal, 1)
	signal.Notify(sigchan, os.Interrupt, syscall.SIGTERM)
	<-sigchan

	// Shut down server
	log.Println("Shutting down server...")
	if err := srv.Shutdown(context.Background()); err != nil {
		log.Fatal(err)
	}
}
