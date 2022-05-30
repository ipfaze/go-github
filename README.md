# Go-GitHub 

Go-GitHub is a scallable API which provides endpoints to retrieve the most recently updated 100 repositories.

As a user you can filter the results by `language` or `licences`.

It also provides some statistics about those repositories such as number of repositories by language or number of repositories by license.

You can see a diagram of use case under `/docs/diagrams/usecase.jpeg`

-----------------

## Architecture

This API will be separated in two different pieces described bellow.
You can see a diagram of the architecture under `/docs/diagrams/architecture.jpeg`

### Project Parts

#### Go-GitHub-Fetcher

This part will query the GitHub API every 6 minutes to fetch any data relating to the repositories.
The 6 minutes timeframe has be chosen so as not to exceed the maximum request limit of the GitHub API which is 10 requests per hour.

Although it can be increased to 30 by passing a personal token in the requests it will not be covered in this version but could be a beneficial upgrade in the future.

All the data retrieved will be saved in a Master MongoDB which will be replicated into multiple Slave MongoDB.

#### Go-Github-Api

This part will provide a client API.
The API is designed to retrieve data from the Slave MongoDB and send it back to the client.

It can be scalled horizontally coupled with new Slave MongoDB to increase performance / time request. 

You can find full documentation about all the endpoints under `/docs/api/Go-Github.postman.json`
To see and test the endpoint, you can use [postman](https://www.postman.com/) and import the json file.

### Why this architecture ?

I chose to divide this project in 2 distinct parts being Go-Github-Api to read the MongoDB data and Go-Github-Fetcher to query the GitHub Api and save the data into MongoDB.

This architecture design choice has been made in order to achieve maximum scalability and availability.

On the one hand, even if the GitHub API is down or is modified, clients can continue to query our API and get data.
If the GitHub API is modified, we just have to update the Go-Github-Fetcher project.

On the other hand, the Go-Github-Api is easily horizontaly scalable to cope with the increased workload.
Same can be applied to the MongoDB, with replicates.

By the way, Go-Github-Api and Go-Github-Fetcher are loosely coupled because they are just linked by the MongoDB instances.

-----------------

## Execution

To start the API, docker is required.
If it's not installed yet please follow this link : [docker](https://docs.docker.com/get-docker/)

When installed, execute these command in your terminal :

```shell
docker-compose up --build
```

The API will be accessible at this url : [http://go-github.localhost](http://go-github.localhost)

To watch all the services and their statuses, you can visit the dashboard of Traefik [here](http://localhost:8080)
Please consider the Traefik dashboard for development purposes only and disregard for production.

-----------------

## Tests

Unit tests are only available at this moment for the Go-Github-Fetcher.
To run it, please start a terminal in the `go-github-fetcher` and run :

### Without coverage report
```shell
go test ./...
```

### With coverage report
```shell
go test -cover ./... -v -coverpkg=./...
```

### With coverage report by function
```shell
go test -coverprofile=coverage.out ./... -v -coverpkg=./...
go tool cover -func=coverage.out
```

### With coverage report as html
```shell
go test -coverprofile=coverage.out ./... -v -coverpkg=./...
go tool cover -html=coverage.out
```

-----------------

## What's next

#### Go-Github-Api
- [] Add unit tests
- [] Add e2e tests

#### Go-Github-Fetcher
- [] Add environment variable "GITHUB_API_TOKEN" to allow for an increase in the number of requests allowance to GitHub API per hour
- [] Add more unit tests
- [] Add e2e tests
- [] Refactorize helpers package (maybe use generics)

#### Whole project
- [] Improve horizontal scalability by setting more instances of MongoDB as replicates into docker-compose.yaml
- [] Set CI/CD to test the project before merging it