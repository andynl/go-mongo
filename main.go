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
	// updateProfile(profileRepository)
	// deleteProfile(profileRepository)
	getProfile("2", profileRepository)
}

func saveProfile(profileRepository repository.ProfileRepository) {
	var p model.Profile
	p.ID = "2"
	p.FirstName = "Legion"
	p.LastName = "Lenovo"
	p.Email = "legion@gmail.com"
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

func deleteProfile(profileRepository repository.ProfileRepository) {
	err := profileRepository.Delete("1")

	if err != nil {
		fmt.Println(err)
	} else {
		fmt.Println("Profile deleted..")
	}
}

func getProfile(id string, profileRepository repository.ProfileRepository) {
	profile, err := profileRepository.FindByID(id)

	if err != nil {
		fmt.Println(err)
	}

	fmt.Println(profile.ID)
	fmt.Println(profile.FirstName)
	fmt.Println(profile.LastName)
	fmt.Println(profile.Email)
}
