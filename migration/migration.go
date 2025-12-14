package migration

import (
	"fmt"

	"SentiService/db"
	"SentiService/internal/modules/address"
)

func Migration() {
	database := db.Database()

	err := database.AutoMigrate(
		&address.Address{},
	)

	if err != nil {
		panic(err)
	}

	fmt.Println("Migraciones de SentiService ejecutadas")
}
