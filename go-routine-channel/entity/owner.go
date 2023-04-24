package entity

/*
"User": {
		"id": 1,
		"first_name": "Alyosha",
		"last_name": "Caldero",
		"email": "acaldero0@behance.net",
		"gender": "Male",
		"birthdate": "29/12/1997",
		"company_name": "King and Sons",
		"department": "Sales",
		"job_title": "Senior Editor",
		"address": [
			....
			]
			.....
*/

type Owner struct {
	OwnerData `json:"User"`
}

type OwnerData struct {
	ID        int    `json:"id"`
	FirstName string `json:"first_name"`
	LastName  string `json:"last_name"`
	Email     string `json:"email"`
}
