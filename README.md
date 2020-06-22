# go-service
Outline of a simple HTTP service in Go

I'll be using the [gin](https://github.com/gin-gonic/gin) framework for the bones of this REST API. The purpose is to create
a starting point for any HTTP service such that the dockerization, configuration and general service layout is predefined.
This will hopefully make it easier to get started or at least be a good example of what a service using the gin framework
might look like.

## Getting Started
Not much for now, but might become more involved as we go along.
* Install dependencies: `go get -u ./...`
* Run the service: `docker-compose up --build`

## Configuration
This service is setup to use `<env_name>.env` files to load its configuration (from both `env_file` in docker 
compose & directly). These files can be passed into docker, depending on which environment you are deploying to.
 
The environment variable `SERVICES_PROFILE` should be set to `docker` (e.g. `SERVICES_PROFILE=docker`) if the service is
running within a container. The reason for this is because this tells the service to grab the environment variables as 
its configuration. This would allow someone to change the config by injecting new values for these variables, which is sometimes
desired over a redeploy with the changed `.env` file.

In a local instance we could set `SERVICES_PROFILE=local`  and it will read the `local.env` file for the configuration. 

_Note: this configuration is very opinionated and can be easily changed._ 
