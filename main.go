package main

import (
	"database/sql"
	"time"

	"github.com/gofiber/fiber/v2"
	_ "github.com/lib/pq"
)

type Activity struct {
	ID int `json:"id"`
	Title string `json:"title"`
	Category string `json:"category"`
	Description string `json:"description"`
	ActivityDate time.Time `json:"activity_date"`
	Status string `json:"status"`
	CreatedAt time.Time `json:"created_at"`
}

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

	// GET //activities
	app.Get("/activities", func (c *fiber.Ctx) error {
		rows, err := db.Query("SELECT * FROM activities")
		if err != nil {
			return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"message":err.Error()})
		}
		defer rows.Close()

		var activities []Activity
		for rows.Next() {
			var activity Activity
			err = rows.Scan(&activity.ID, &activity.Title, &activity.Category, &activity.Description, &activity.ActivityDate, &activity.Status, &activity.CreatedAt)

			activities = append(activities, activity)
		}

		return c.JSON(activities)
	})

	app.Listen(":8081")
}