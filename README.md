# Car Rental API in GO

Basic api for car rentals to study Golang

## Requirements

- The [Go](https://go.dev) programming language
- A [PostgreSQL](https://www.postgresql.org) database

## Instructions

1. Clone the project
1. Run the command `mv .env.example .env` to rename the `.env.example` file to `.env`
1. Fill in with the environment variables in the `.env` file
1. Run the server with the command `go run main.go`

### Environment Variables

- PORT: The port you want the server to run in. Example: `4000` (defaults to 8080)
- DSN: A string with the information to connect to the database. Example: `"user=gorm password=gorm dbname=gorm port=9920 sslmode=disable"`
- JWT_SECRET: A secret string to use to hash the jwt. Example: `uidgqw78dgqw78giodcqwuih`

## Routes

All application routes

### Brands

- `GET /brands` used to get the list of all brands
- `POST /brands` used to create a new brand
  ###### Data (json):
  - Name `[string, required]`
    <br />
- `GET /brands/:id` used to get a brand by id
- `PUT /brands/:id` used to update a brand by id
  ###### Data (json):
  - Name `[string, required]`
    <br />
- `DELETE /brands/:id` used to delete a brand by id

### Cars

- `GET /cars` used to get the list of all cars
- `POST /cars` used to create a new car
  ###### Data (json):
  - BrandId `[uuid, required]`
  - Model `[string, required]`
  - RentalPriceDailyInUsd `[float, required]`
  - HorsePower `[int, required]`
  - TorqueInLb `[float, required]`
  - TopSpeedInKm `[int, required]`
  - AccelerationSpeedInKm `[float, required]`
  - WeightInKg `[int, required]`
    <br />
- `GET /cars/:id` used to get a car by id
- `PUT /cars/:id` used to update a car by id
  ###### Data (json):
  - BrandId `[uuid, required]`
  - Model `[string, required]`
  - RentalPriceDailyInUsd `[float, required]`
  - HorsePower `[int, required]`
  - TorqueInLb `[float, required]`
  - TopSpeedInKm `[int, required]`
  - AccelerationSpeedInKm `[float, required]`
  - WeightInKg `[int, required]`
    <br />
- `DELETE /cars/:id` used to delete a car by id

### Auth

- `POST /auth/register` used to register new user
  ###### Data (json):
  - Name `[string, required]`
  - Email `[string, required]`
  - Password `[string, required]`
    <br />
- `POST /auth/login` used to login
  ###### Data (json):
  - Email `[string, required]`
  - Password `[string, required]`
    <br />
- `GET /auth/me` used to get logged in user data `[requires auth]`
- `GET /auth/me/rentals` used to get logged in user rentals `[requires auth]`
  ###### Filters (query params):
  - CarId - if specified it will return only the rentals for this CarId `[optional]`
  - StartsAt - if specified it will return only the rentals that have a StartsAt greater than or equal to the specified date `[optional]`
  - EndsAt - if specified it will return only the rentals that have an EndsAt lesser than or equal to the specified date `[optional]`

### Rentals

- `GET /rentals` used to get the list of all rentals
  ###### Filters (query params):
  - CarId - if specified it will return only the rentals for this CarId `[optional]`
  - StartsAt - if specified it will return only the rentals that have a StartsAt greater than or equal to the specified date `[optional]`
  - EndsAt - if specified it will return only the rentals that have an EndsAt lesser than or equal to the specified date `[optional]`
    <br />
- `POST /rentals` used to rent a car `[requires auth]`
  ###### Data (json)
  - CarId - The id of the car you want to rent `[required]`
  - StartsAt - The start date for which you want to rent the car `[required]`
  - EndsAt - The end date for which you want to rent the car `[required]`
    <br />
- `PATCH /rentals/:id/cancel` used to cancel a rental `[requires auth]`

[Insomnia import file](./insomnia.json)

## How to Authenticate

The `/auth/register` and the `/auth/login` endpoints return an `accessToken`, send it as a Bearer token in the request, as the Authorization header, for example: `Bearer accessToken_here`

## Tech Stack

- [Gin](https://gin-gonic.com)
- [Gorm](https://gorm.io)
- [GoDotEnv](https://github.com/joho/godotenv)
- [JWT-Go](https://github.com/golang-jwt/jwt)
