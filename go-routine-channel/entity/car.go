package entity

/*
"cars": [
	{
			"id": 1,
			"car": "Mitsubishi",
			"car_model": "Montero",
			"car_color": "Yellow",
			"car_model_year": 2002,
			"car_vin": "SAJWJ0FF3F8321657",
			"price": "$2814.46",
			"availability": false
	},
]
*/

type Car struct {
	CarData `json:"car"`
}

type CarData struct {
	ID    int    `json:"id"`
	Brand string `json:"car"`
	Model string `json:"car_model"`
	Year  int    `json:"car_model_year"`
	// Color     string `json:"car_color"`
	// CarVin       string `json:"car_vin"`
	// Price        string `json:"price"`
	// Availability bool   `json:"availability"`
}
