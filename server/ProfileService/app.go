package main

import (
	"fmt"
	"log"
	"net"
	"os"
	"os/signal"

	"google.golang.org/grpc/reflection"

	"github.com/jmoiron/sqlx"
	_ "github.com/lib/pq"
	"google.golang.org/grpc"

	"github.com/getsentry/sentry-go"

	profiles_pb "github.com/stanlee321/demo-grpc-go-server-client/proto"
)

// App is the struct with app configuration values
type App struct {
	DB *sqlx.DB
	// Router *mux.Router

}

// Initialize create the DB connection and prepare all the routes
// func (a *App) Initialize(cache Cache, db *sqlx.DB) {
func (a *App) Initialize(db *sqlx.DB) {
	a.DB = db
	fmt.Println("Initialize Server...")
}

func (a *App) runGRPCServer(portAddr string) {

	sentry.CaptureMessage("[ User Service ]  it WORKS!!!")

	lis, err := net.Listen("tcp", portAddr)

	if err != nil {
		log.Fatalf("failed to listen: %v", err)
	}
	s := grpc.NewServer()
	// Register user service server to profiles_pb
	profiles_pb.RegisterProfileServiceServer(s, &userDataHandler{app: a})

	// Set reflection
	reflection.Register(s)

	if err := s.Serve(lis); err != nil {
		log.Fatalf("failed to serve: %v", err)
	}

	go func() {
		fmt.Println("Starting Server...")
		if err := s.Serve(lis); err != nil {
			log.Fatalf("failed to serve: %v", err)
		}
	}()

	// Wait for Control C to exit
	ch := make(chan os.Signal, 1)
	signal.Notify(ch, os.Interrupt)

	// Block until a signal is received
	<-ch

	fmt.Println("Closing POSGRES Connection")
	// client.Disconnect(context.TODO())
	if err := a.DB.Close(); err != nil {
		log.Fatalf("Error on disconnection with MongoDB : %v", err)
	}

	// Second step : closing the listener
	fmt.Println("Closing the listener")
	if err := lis.Close(); err != nil {
		log.Fatalf("Error on closing the listener : %v", err)
	}
	// Finally, we stop the server
	fmt.Println("Stopping the server")
	s.Stop()
	fmt.Println("End of Program")
}

type userDataHandler struct {
	app *App
}
