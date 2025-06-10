package main

import (
	"database/sql"

	"github.com/gofiber/fiber/v2"
	_ "github.com/lib/pq"
)

func initDB() (*sql.DB, error) {
	dns := "user=postgres.lniqicukthxfhxcchmky password=Z1ZvqpWtO9L8c2BA host=aws-0-ap-southeast-1.pooler.supabase.com port=6543 dbname=postgres"
	db, err := sql.Open("postgres", dns)
	if err != nil {
		return nil, err
	}

	err = db.Ping()
	if err != nil {
		db.Close()
		return nil, err
	}
	return db, nil
}

func main()  {
	db, err := initDB()
	if err != nil {
		panic(err)
	}
	defer db.Close()

	app := fiber.New()

	app.Listen(":8081")
}