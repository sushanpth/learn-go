package service

import (
	"fmt"
	"net/http"
	// "github.com/sushanpth/learn-go/go-routine-channel/repository"
)

type CarService interface {
	FetchData()
}

const (
	carServiceURL = "https://myfakeapi.com/api/cars/1"
)

type fetchCarDataService struct{}

// NewCarService returns a new CarService
func NewCarService() CarService {
	return &fetchCarDataService{}
}

func (*fetchCarDataService) FetchData() {
	client := http.Client{}
	fmt.Println("Fetching car URL: ", carServiceURL)
	// call the api

	resp, _ := client.Get(carServiceURL)

	// Write response to channel
	carDataChannel <- resp

}
