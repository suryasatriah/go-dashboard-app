package controller

import (
	"log"

	"github.com/surysatriah/go-dashboard-app/internal/database"
	"github.com/surysatriah/go-dashboard-app/internal/model"
)

func InsertPayload(payload model.Payload) {

	db := database.GetDatabase()

	err := db.Debug().Create(&payload).Error
	if err != nil {
		log.Printf("Error inserting payload: %v", err)
	}

}
