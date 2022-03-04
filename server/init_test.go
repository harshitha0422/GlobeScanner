package main

import (
	"database/sql"
	"net/http"
	"strconv"
	"testing"

	"github.com/go-chi/chi"
	"github.com/qkgo/yin"
)

// // func TestuserRegister(t *testing.T) {
// //     name := "Gladys"
// //     want := regexp.MustCompile(`\b`+name+`\b`)
// // 	msg, err := userLogin(c *gin.Context)
// //     msg, err := Hello("Gladys")
// //     if !want.MatchString(msg) || err != nil {
// //         t.Fatalf(`Hello("Gladys") = %q, %v, want match for %#q, nil`, msg, err, want)
// //     }
// // }

func TestAddMovies(t *testing.T) {

	db, _ := sql.Open("sqlite3", "./moviedatabaseV3.db")
	feed := moviefeed.NewFeed(db)

	r := chi.NewRouter()
	r.Use(yin.SimpleLogger)

	r.Post("/addMovie", func(w http.ResponseWriter, r *http.Request) {
		res, req := yin.Event(w, r)
		body := map[string]string{}
		req.BindBody(&body)
		movie := moviefeed.Movie{
			Name:   body["name"],
			Desc:   body["description"],
			Review: body["review"],
			Rating: body["rating"],
			Genre:  body["genre"],
		}
		feed.Add(movie)
		res.SendStatus(204)
		assert.Equal(t, 204, r.Response.StatusCode)
	})

}

func TestDeleteMovie(t *testing.T) {

	db, _ := sql.Open("sqlite3", "./moviedatabaseV3.db")
	feed := moviefeed.NewFeed(db)

	r := chi.NewRouter()
	r.Use(yin.SimpleLogger)

	r.Post("/deleteMovieByID", func(w http.ResponseWriter, r *http.Request) {
		res, req := yin.Event(w, r)
		body := map[string]string{}
		req.BindBody(&body)
		id, _ := strconv.Atoi(body["movieid"])
		movie := moviefeed.Movie{
			ID: id,
		}
		feed.DeleteMovieByID(movie)
		res.SendStatus(204)
	})

}

func TestUpdateMovies(t *testing.T) {

	db, _ := sql.Open("sqlite3", "./moviedatabaseV3.db")
	feed := moviefeed.NewFeed(db)

	r := chi.NewRouter()
	r.Use(yin.SimpleLogger)

	r.Put("/updateMovieByID", func(w http.ResponseWriter, r *http.Request) {
		res, req := yin.Event(w, r)
		body := map[string]string{}
		req.BindBody(&body)
		id, _ := strconv.Atoi(body["movieid"])
		movie := moviefeed.Movie{
			ID:     id,
			Rating: body["rating"],
		}
		feed.UpdateMovieByID(movie)
		res.SendStatus(204)
	})

}
