# Go-GitHub 

Go-GitHub is a scallable API which provides endpoints to retrieve the latest updated 100 repositories.
As a user you can filter this result by `language` and/or `licences`
It also responds some statistics about those repositories

You can see a diagram of usecase under `/docs/diagrams/usecase.jpeg`

-----------------

## Architecture

This API will be separate in two different pieces described bellow.
You can see a diagram of the architecture under `/docs/diagrams/architecture.jpeg`

-----------------

### Project Parts

#### Go-GitHub-Fetcher

This part will request the GitHub API every 6 minutes to retrieve the data about repositories.
To not exceed the rate limit of the GitHub API which is 10 requests per hour.

It can be increased to 30 by passing a personal token in the requests
This will not be covered in this version but can be a great upgrade in the future.

All the data retrieved will be saved in a Master MongoDB will replicate into multiple Slave MongoDB.


#### Go-Github-Api

This part will provide an API to the client.
The API while retrieve data from the Slave MongoDB and sent it back to the client.
It can be scalled horizontally coupled with new Slave MongoDB to increase performance / time request. 

You can find a full documentation about all the endpoint under `/docs/api/Go-Github.postman.json`
To see it and test the endpoint, you can use [postman](https://www.postman.com/) and import the json file.

### Why this architecture

I choose to separate this projects for more scalability and availability.
Because even if the GitHub API is down or has a change, clients can continue to request our API and get data.
If the GitHub API has a change, we just have to update the Go-Github-Fetcher.

On the other part, the Go-Github-Api is easily horizontaly scalable to cope with the increased workload.
Same can be applied to the MongoDB, with replicates.

By the way, Go-Github-Api and Go-Github-Fetcher are loosely couple because they are just coupled with the MongoDB.

-----------------

## Execution

To start the API, docker is required.
If it's not installed yet please follow this link : [docker](https://docs.docker.com/get-docker/)

When installed, execute these command in your terminal :

```shell
docker-compose up --build
```

The API will be accessible at this url : [http://go-github.localhost](http://go-github.localhost)

To watch all the services and your status, you can visit the dashboard of Traefik [here](http://localhost:8080)
Please consider the Traefik dashboard only for development purpose and not for production.

-----------------

## What's next

#### Go-Github-Api
- [] Set Repository design pattern
- [] Add unit tests
- [] Add e2e tests

#### Go-Github-Fetcher
- [] Set Repository design pattern
- [] Add environment variable "GITHUB_API_TOKEN" increase request to GitHub API per hour
- [] Add unit tests
- [] Add e2e tests
- [] Refactorize helpers package

#### Whole project
- [] Improve horizontal scalability by setting more instance of MongoDB as replicates into docker-compose.yaml
- [] Set CI/CD for testing project before push it or merge it