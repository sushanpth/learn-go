package service

import (
	"fmt"
	"net/http"
	// "github.com/sushanpth/learn-go/go-routine-channel/repository"
)

type OwnerService interface {
	FetchData()
}

const (
	ownerServiceURL = "https://myfakeapi.com/api/users/1"
)

type fetchOwnerDataService struct{}

// NewCarService returns a new CarService
func NewOwnerService() CarService {
	return &fetchOwnerDataService{}
}

func (*fetchOwnerDataService) FetchData() {
	client := http.Client{}
	fmt.Println("Fetching owner URL: ", ownerServiceURL)
	// call the api

	resp, _ := client.Get(ownerServiceURL)

	// Write response to channel
	ownerDataChannel <- resp

}
