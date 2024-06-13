package main

import (
	"flag"
	"fmt"
	"github.com/joho/godotenv"
	"log"
	"technoscape-hackathon-7.0-backend/cmd"
	"technoscape-hackathon-7.0-backend/helper"
)

const defaultPort = "8080"

var (
	runMigration bool
	runSeeding   bool
)

func init() {
	flag.BoolVar(&runMigration, "migrate", false, "Run database migrations")
	flag.BoolVar(&runSeeding, "seed", false, "Run database seeding")
}

func main() {
	flag.Parse()

	if runMigration || runSeeding {
		if runMigration {
			cmd.Migrate()
		}
		if runSeeding {
			cmd.Seed()
		}
		return
	}

	if err := godotenv.Load(); err != nil {
		log.Fatalf("Error loading .env file: %s", err.Error())
	}

	port := helper.Getenv("API_PORT", defaultPort)

	fmt.Println(port)
}
