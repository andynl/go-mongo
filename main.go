package main

import (
	"fmt"
	"time"

	"github.com/andynl/go-mongo/config"
	"github.com/andynl/go-mongo/src/modules/profile/model"
	"github.com/andynl/go-mongo/src/modules/profile/repository"
)

func main() {
	fmt.Println("Go Mongo")

	db, err := config.GetMongoDB()

	if err != nil {
		fmt.Println(err)
	}

	profileRepository := repository.NewProfileRepositoryMongo(db, "profile")

	// saveProfile(profileRepository)
	updateProfile(profileRepository)
}

func saveProfile(profileRepository repository.ProfileRepository) {
	var p model.Profile
	p.ID = "1"
	p.FirstName = "Andy"
	p.LastName = "Natalino"
	p.Email = "andy.natalino@gmail.com"
	p.Password = "123456"
	p.CreatedAt = time.Now()
	p.UpdatedAt = time.Now()

	err := profileRepository.Save(&p)

	if err != nil {
		fmt.Println(err)
	} else {
		fmt.Println("Profile saved..")
	}
}

func updateProfile(profileRepository repository.ProfileRepository) {
	var p model.Profile
	p.ID = "1"
	p.FirstName = "Lihan"
	p.LastName = "Andy"
	p.Email = "andy@gmail.com"
	p.Password = "123123"
	p.CreatedAt = time.Now()
	p.UpdatedAt = time.Now()

	err := profileRepository.Update("1", &p)

	if err != nil {
		fmt.Println(err)
	} else {
		fmt.Println("Profile updated..")
	}

}
