# Go-GitHub 

Go-GitHub is a scallable API which provides endpoints to retrieve the latest updated 100 repositories.
As a user you can filter this result by `language` and/or `licences`
It also responds some statistics about those repositories

You can see a diagram of usecase under `/diagrams/usecase.jpeg`

## Architecture

This API will be separate in two different pieces

You can see a diagram of the architecture under `/diagrams/architecture.jpeg`


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


## Execution

To start the API, docker is required.
If it's not installed yet please follow this link : [docker](https://docs.docker.com/get-docker/)

When installed, execute these command in your terminal :

```shell
docker-compose up --build
```

The API will be accessible at this url : [http://go-github.localhost](http://go-github.localhost)

To help you, you can visit the dashboard of Traefik [here](http://localhost:8080)

Please consider the Traefik dashboard only for development purpose and not for production.