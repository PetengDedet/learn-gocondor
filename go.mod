module learn-gocondor

replace (
	learn-gocondor/config => ./config
	learn-gocondor/http => ./http
	learn-gocondor/http/handlers => ./http/handlers
	learn-gocondor/http/middlewares => ./http/middlewares
	learn-gocondor/models => ./models
)

go 1.16

require (
	github.com/gin-gonic/gin v1.7.1
	github.com/gocondor/core v1.4.4
	github.com/gosimple/slug v1.10.0 // indirect
	github.com/joho/godotenv v1.3.0
	golang.org/x/crypto v0.0.0-20210421170649-83a5a9bb288b
	gorm.io/gorm v1.21.6
)
