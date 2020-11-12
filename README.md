# GO-REST

A small little webserver written in Go to learn about Go.

## Run

```go
go run main.go
```

## Available Routes

- [GET /](http://localhost:8081)
- [GET /recipes](http://localhost:8081/recipes)
- [GET /recipe/{slug}](http://localhost:8081/recipes/ham)
- POST /recipe
  - sample request body
  ```json
  {
    "slug": "popcorn",
    "title": "popcorn",
    "ingredients": "corn kernels",
    "directions": "cook over open flame until cooked"
  }
  ```
