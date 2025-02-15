package database

import (
	"fmt"
	"intikom-test-go/model"
	"intikom-test-go/utils"
	"math/rand"

	"github.com/go-faker/faker/v4"
	"gorm.io/gorm"
)

func DatabaseSeeder(db *gorm.DB, count int) {

	for i := 0; i < count; i++ {
		firstName := faker.FirstName()
		lastName := faker.LastName()
		fullName := firstName + " " + lastName

		user := model.User{
			Name:     fullName,
			Email:    faker.Email(),
			Password: utils.GeneratePassword("password"),
		}

		err := db.Save(&user).Error
		if err != nil {
			fmt.Printf("Error when create user: %s\n", user.Name)
			return
		}

		for t := 0; t < 3; t++ {
			task := model.Task{
				Title:       faker.Sentence(),
				Description: faker.Paragraph(),
				Status:      randomTaskStatus(),
				UserID:      user.ID,
			}

			err = db.Save(&task).Error
			if err != nil {
				fmt.Printf("Error when create task: %s\n", task.Title)
			}
		}
	}
}

func randomTaskStatus() model.TaskStatus {
	status := []model.TaskStatus{model.TaskStatusPending, model.TaskStatusDone}
	return status[rand.Intn(len(status))]
}
