package service

import (
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/sushanpth/learn-go/go-routine-channel/entity"
	// "github.com/sushanpth/learn-go/go-routine-channel/repository"
)

type CarDetailService interface {
	GetDetails() entity.CarDetails
}

var (
	carService       CarService   = NewCarService()
	ownerService     OwnerService = NewOwnerService()
	carDataChannel                = make(chan *http.Response)
	ownerDataChannel              = make(chan *http.Response)
)

type service struct{}

// NewCarService returns a new CarService
func NewCarDetailService() CarDetailService {
	return &service{}
}

func (*service) GetDetails() entity.CarDetails {
	// go routine to get data from first API
	go carService.FetchData()

	// another go routine to get data from second API
	go ownerService.FetchData()

	// create car channel to get the data from first API
	// create owner channel to get the data from second API
	car, _ := getCarData()
	owner, _ := getOwnerData()

	return entity.CarDetails{
		ID:        car.ID,
		Brand:     car.Brand,
		Model:     car.Model,
		Year:      car.Year,
		FirstName: owner.FirstName,
		LastName:  owner.LastName,
		Email:     owner.Email,
	}
}

func getCarData() (entity.Car, error) {
	r1 := <-carDataChannel

	var car entity.Car
	err := json.NewDecoder(r1.Body).Decode(&car)
	if err != nil {
		fmt.Println(err.Error())
		return car, err
	}
	return car, nil
}

func getOwnerData() (entity.Owner, error) {
	r1 := <-ownerDataChannel

	var owner entity.Owner
	err := json.NewDecoder(r1.Body).Decode(&owner)
	if err != nil {
		fmt.Println(err.Error())
		return owner, err
	}
	return owner, nil
}
