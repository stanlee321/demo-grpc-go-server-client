package main

import (
	"fmt"
	"log"
	"os"
	"github.com/jmoiron/sqlx"
	_ "github.com/lib/pq"
	"github.com/getsentry/sentry-go"
)

const (
	portAddr = ":50051"
)

func init() {
	err := sentry.Init(sentry.ClientOptions{
		Dsn: "https://6b32e5da222a42f4b3a5a0795171e35c@o447510.ingest.sentry.io/5427463",
	})
	if err != nil {
		log.Fatalf("sentry.Init: %s", err)
	}
}

func main() {

	fmt.Println("INITING.... UserService gRPC server")

	connectionString := os.Getenv("DATABASE_DEV_URL")

	db, err := sqlx.Open("postgres", connectionString)
	if err != nil {
		// log.Fatal(err)
		fmt.Println(err)
		log.Fatalf("Failed to listen: %v", err)

	}

	a := App{}
	// a.Initialize(cache, db)
	a.Initialize(db)
	a.runGRPCServer(portAddr)
	// a.Run(":3000")
}
