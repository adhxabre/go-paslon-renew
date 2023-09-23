package databases

import (
	"backend-micro-feature/models"
	"fmt"
)

func RunMigration() {
	err := DB.AutoMigrate(
		&models.Paslon{},
		&models.Party{},
		&models.Vote{},
	)

	if err != nil {
		panic(err)
	}

	fmt.Println("Migration success cuy")
}
