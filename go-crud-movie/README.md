# Movie API
This is a simple RESTful API built with Golang, utilizing the Gorilla Mux router. The API allows you to perform CRUD (Create, Read, Update, Delete) operations on a collection of movies.

## Features

- GET /movies: Retrieve a list of all movies
- GET /movies/{id}: Retrieve a single movie by its ID
- POST /movies: Create a new movie
- PUT /movies/{id}: Update an existing movie by its ID
- DELETE /movies/{id}: Delete a movie by its ID

## Prerequisites
- Golang installed
- Git installed

## Installation
Clone the repository:
```sh
git clone https://github.com/yourusername/movie-api.git
```
Navigate to the project directory:
``` sh
cd movie-api
```
Install the required packages:
``` sh 
go get -u github.com/gorilla/mux
```
Run the application:
``` sh
go build
go run main.go
```
The server will start on port 8000
``` sh
Starting the server at port 8000
```
# API Endpoints
### GET /movies
Retrieve a list of all movies.
Request:
``` sh
curl -X GET http://localhost:8000/movies
```
Response:
``` sh
[
    {
        "id": "1",
        "isbn": "45464",
        "title": "Movie One",
        "director": {
            "firstname": "John",
            "lastname": "Smith"
        }
    },
    {
        "id": "2",
        "isbn": "43364",
        "title": "Movie Two",
        "director": {
            "firstname": "Wood",
            "lastname": "Joe"
        }
    }
]
```

### GET /movies/{id}
Retrieve a single movie by its ID.
Request:
``` sh
curl -X GET http://localhost:8000/movies/1
```
Response:
``` sh
{
    "id": "1",
    "isbn": "45464",
    "title": "Movie One",
    "director": {
        "firstname": "John",
        "lastname": "Smith"
    }
```
### POST /movies
Create a new movie.
Request:
``` sh
curl -X POST -H "Content-Type: application/json" -d '{
    "isbn": "12345",
    "title": "New Movie",
    "director": {
        "firstname": "New",
        "lastname": "Director"
    }
}' http://localhost:8000/movies
```
Response:
``` sh
{
    "id": "3",
    "isbn": "12345",
    "title": "New Movie",
    "director": {
        "firstname": "New",
        "lastname": "Director"
    }

```

### PUT /movies/{id}
Update an existing movie by its ID.
Request:
``` sh
curl -X PUT -H "Content-Type: application/json" -d '{
    "isbn": "12345",
    "title": "Updated Movie",
    "director": {
        "firstname": "Updated",
        "lastname": "Director"
    }
}' http://localhost:8000/movies/1
```
Response:
``` sh
{
    "id": "1",
    "isbn": "12345",
    "title": "Updated Movie",
    "director": {
        "firstname": "Updated",
        "lastname": "Director"
    }

```
### DELETE /movies/{id}
Delete a movie by its ID.
Request:
``` sh
curl -X DELETE http://localhost:8000/movies/1
```
Response:
``` sh
[
    {
        "id": "2",
        "isbn": "43364",
        "title": "Movie Two",
        "director": {
            "firstname": "Wood",
            "lastname": "Joe"
        }
    }
]
```

# Code Explanation
Here's a brief overview of the main parts of the code:
### Struct Definitions
``` sh
type Movie struct {
    ID       string    `json:"id"`
    Isbn     string    `json:"isbn"`
    Title    string    `json:"title"`
    Director *Director `json:"director"`
}

type Director struct {
    Firstname string `json:"firstname"`
    Lastname  string `json:"lastname"`
}
```
# Main Function

##### The main function initializes the movies slice with some data, sets up the router, and starts the server.

## Handler Functions

- getMovies: Returns all movies in JSON format.
- getMovie: Returns a single movie by ID.
- createMovie: Adds a new movie to the list.
- updateMovie: Updates an existing movie by ID.
- deleteMovie: Deletes a movie by ID.

# Contact
You can reach me on LinkedIn: https://www.linkedin.com/in/suleman-masih/

