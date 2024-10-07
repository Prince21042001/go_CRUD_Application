# Go Car CRUD API

This is a simple RESTful API written in Go that allows for basic CRUD (Create, Read, Update, Delete) operations on a collection of cars. The API uses an in-memory map to store cars and allows clients to interact with the data through HTTP requests.

## Features
- **Create a Car** (POST `/cars`): Add a new car to the system.
- **List all Cars** (GET `/cars`): Retrieve a list of all cars.
- **Get a Car by ID** (GET `/cars/{id}`): Retrieve details of a specific car by its ID.
- **Update a Car by ID** (PUT `/cars/{id}`): Modify details of a car.
- **Delete a Car by ID** (DELETE `/cars/{id}`): Remove a car from the system.

## Data Structure
Each car has the following fields:
- `ID` (int): Auto-incremented unique identifier for each car (assigned by the system).
- `Company` (string): The manufacturer of the car (e.g., Toyota, Ford).
- `Model` (string): The car's model (e.g., Corolla, Mustang).
- `Year` (int): The year the car was manufactured.

Example:
```json
{
    "id": 1,
    "company": "Toyota",
    "model": "Corolla",
    "year": 2020
}
