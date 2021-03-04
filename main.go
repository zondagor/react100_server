package main

import (
	"fmt"
	"github.com/go-chi/chi"
	"github.com/go-chi/cors"
	"log"
	"net/http"
	"react100_server/controllers/auth"
	"react100_server/controllers/profile"
	"react100_server/controllers/user"
)

func main() {
	r := chi.NewRouter()

	r.Use(cors.Handler(cors.Options{
		AllowedOrigins:   []string{"*"},
		AllowedMethods:   []string{"GET", "POST", "PUT", "DELETE", "OPTIONS"},
		AllowedHeaders:   []string{"Accept", "Authorization", "Content-Type", "X-CSRF-Token"},
		AllowCredentials: false,
		MaxAge:           300, // Maximum value not ignored by any of major browsers
	}))

	r.Route("/api/1.0/", func(r chi.Router) {
		r.Get("/auth/me", auth.AuthHandler)


		r.Route("/profile", func(r chi.Router) {
			r.Get("/{id}", profile.GetProfile)
			r.Get("/status/{id}", profile.GetUserStatus)
			r.Put("/status/{id}", profile.UpdateUserStatus)
		})

		r.Route("/users", func(r chi.Router) {
		  r.Get("/", user.GetUsers)
		  r.Post("/follow/{id}", user.FollowUser)
		  r.Delete("/follow/{id}", user.UnFollowUser)
		})
	})

	port := ":8081"
	fmt.Println("started server on port", port)
	log.Fatal(http.ListenAndServe(port, r))
}

// http://localhost:8081/api/1.0/users?page=1&count=5

