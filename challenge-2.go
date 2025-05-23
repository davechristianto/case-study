package main

import (
	"net/http"
	"strconv"
	"strings"

	"github.com/gin-gonic/gin"
)

type Movie struct {
	ID          int      `json:"id"`
	Title       string   `json:"title"`
	Description string   `json:"description"`
	Duration    int      `json:"duration"`
	Artists     []string `json:"artists"`
	Genres      []string `json:"genres"`
}

// CONTOH INPUT API

// CREATE
//curl -X POST http://localhost:8080/movies 
// \-H "Content-Type: application/json" 
// \-d '{"title":"Avengers Infinity Wars","description":"Soul Stone","duration":134,"artists":["Chris Hemsworth","RDJ"],"genres":["Drama","Action"]}'

// GET WITH PAGINATION LIMIT
//curl "http://localhost:8080/movies?page=1&limit=10"

// GET WITH SEARCHING BY TITLE/GENRES/ARTIST/DESCRIPTION
//curl "http://localhost:8080/movies/search?q=drama"

// UPDATE
//curl -X PUT http://localhost:8080/movies/1 \
//-H "Content-Type: application/json" \
//-d '{"title":"Avengers End Game","description":"Time Stone","duration":188,"artists":["Chris Evans"],"genres":["Action"]}'



var movies []Movie
var nextID = 1

func main() {
	r := gin.Default()

	// Create
	r.POST("/movies", func(c *gin.Context) {
		var newMovie Movie
		if err := c.ShouldBindJSON(&newMovie); err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			return
		}

		newMovie.ID = nextID
		nextID++
		movies = append(movies, newMovie)
		c.JSON(http.StatusCreated, newMovie)
	})

	// Update
	r.PUT("/movies/:id", func(c *gin.Context) {
		idStr := c.Param("id")
		id, err := strconv.Atoi(idStr)
		if err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid ID"})
			return
		}

		var updated Movie
		if err := c.ShouldBindJSON(&updated); err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			return
		}

		for i, m := range movies {
			if m.ID == id {
				updated.ID = id
				movies[i] = updated
				c.JSON(http.StatusOK, updated)
				return
			}
		}

		c.JSON(http.StatusNotFound, gin.H{"error": "Movie invalid"})
	})

	// Get
	r.GET("/movies", func(c *gin.Context) {
		pageStr := c.DefaultQuery("page", "1")
		limitStr := c.DefaultQuery("limit", "10")

		page, err1 := strconv.Atoi(pageStr)
		limit, err2 := strconv.Atoi(limitStr)
		if err1 != nil || err2 != nil || page <= 0 || limit <= 0 {
			c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid"})
			return
		}

		start := (page - 1) * limit
		end := start + limit

		if start > len(movies) {
			c.JSON(http.StatusOK, []Movie{})
			return
		}
		if end > len(movigo run main.go
			es) {
			end = len(movies)
		}

		c.JSON(http.StatusOK, movies[start:end])
	})

	// Searching
	r.GET("/movies/search", func(c *gin.Context) {
		query := strings.ToLower(c.Query("q"))
		if query == "" {
			c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid"})
			return
		}

		var results []Movie
		for _, m := range movies {
			if strings.Contains(strings.ToLower(m.Title), query) ||
				strings.Contains(strings.ToLower(m.Description), query) ||
				containsIgnoreCase(m.Artists, query) ||
				containsIgnoreCase(m.Genres, query) {
				results = append(results, m)
			}
		}
		c.JSON(http.StatusOK, results)
	})

	r.Run(":8080")
}

func containsIgnoreCase(arr []string, str string) bool {
	for _, a := range arr {
		if strings.Contains(strings.ToLower(a), str) {
			return true
		}
	}
	return false
}


