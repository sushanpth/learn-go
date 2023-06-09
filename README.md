# GO Language practice projects 
My go learning practice projects

# Folder Structure
```
Learn Go/
├── di-fx-1/
├── di-fx-modular/
│   ├── bundlefx/   
│   ├── handler/  
│   ├── health/
│   ├── main.go
│   ├── go.mod   
│   └── go.sum
├── go-jwt/
│   ├── controllers/
│   ├── initializers/
│   ├── middleware/
│   ├── models/
│   ├── main.go
│   ├── go.mod
│   └── go.sum
├── go-tour
├── simple-http-server/
├── task-crud/
│   ├── controllers/
│   ├── models/
│   ├── routes/
│   ├── main.go
│   ├── env-sample
│   ├── index.html
│   ├── go.mod    
│   └── go.sum
├── task-firebase/
│   ├── entity/
│   ├── initializers/
│   ├── repository/
│   ├── routes/
│   ├── main.go
│   ├── go.sum
│   └── go.mod
├── task-firebase-clean/
│   ├── controller/
│   ├── entity/
│   ├── initializers/
│   ├── service/
│   ├── repository/
│   ├── router/
│   ├── main.go
│   ├── go.mod
│   └── go.sum
├── try-gin/
├── go.work
├── go.work.sum
└── README.md
```

# Packages Used
Here are list of libraries being used in this project
- [Gin Framework](https://gin-gonic.com/)
- [Testify](https://github.com/stretchr/testify) for testing
- [GORM](https://gorm.io/) The fantastic ORM library for Golang
- [JWT](https://github.com/golang-jwt/jwt)
- [Uber FX](https://uber-go.github.io/fx/) Dependency injection system for Go.
- [dotenv library](https://github.com/joho/godotenv) to load environment variables from .env files
- [Firestore](https://cloud.google.com/go/firestore)

### Refrence Links
- Medium Article [JWT authentication In Golang with gin](https://articles.wesionary.team/jwt-authentication-in-golang-with-gin-63dbc0816d55)
- Medium Article [CRUD API using GO, GIN, MySql](https://articles.wesionary.team/crud-api-using-go-d55b0ace211e)
- Youtube Video [JWT Authentication in Go (Gin/Gorm)](https://www.youtube.com/watch?v=ma7rUS_vW9M)
- Medium Article [Dependency Injection with Go-Fx](https://articles.wesionary.team/dependency-injection-with-go-fx-b698a6585cf0)

- Youtube Video [Golang / Go Crash Course 02 | Connecting our REST API with Firebase Firestore Database](https://www.youtube.com/watch?v=RHa4D6aNVpg&list=PL3eAkoh7fypqUQUQPn-bXtfiYT_ZSVKmB&index=3)

- Youtube Video [Golang / Go Crash Course 03 | Implementing Clean Architecture principles in our REST API](https://www.youtube.com/watch?v=Yg_ae0UvCv4&list=PL3eAkoh7fypqUQUQPn-bXtfiYT_ZSVKmB&index=3)

- [Golang / Go Crash Course 05 | Building an API Mashup using Goroutines and Channels](https://www.youtube.com/watch?v=dihX12GkBnc&list=PL3eAkoh7fypqUQUQPn-bXtfiYT_ZSVKmB&index=8&ab_channel=PragmaticReviews)