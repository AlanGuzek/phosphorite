package main

import (
	"github.com/go-pg/pg/v10"
	"log"
)

var Database *pg.DB

func main() {
	Database = DatabaseConnect()
	defer func() {
		if err := Database.Close(); err != nil {
			log.Fatalln("Error closing database: ", err)
		}
	}()
	r := CreateHTTPEngine()
	DatabaseInitSchema(Database)
	amqpConn := InitMessagingService()
	defer func() {
		if err := amqpConn.Close(); err != nil {
			log.Fatalln("Error closing RabbitMQ connection: ", err)
		}
	}()

	log.Fatalln(r.Run(":" + GetEnvVariable("PHO_PORT", "8321")))
}
