// Copyright 2021 Harran Ali <harran.m@gmail.com>. All rights reserved.
// Use of this source code is governed by MIT-style
// license that can be found in the LICENSE file.

package http

import (
	"learn-gocondor/http/handlers"

	"github.com/gocondor/core/routing"
)

// RegisterRoutes to register your routes
func RegisterRoutes() {
	router := routing.Resolve()

	//Define your routes here
	router.Get("/", handlers.HomeShow)

	// Users
	router.Get("/users", handlers.GetUsers)
	router.Post("/users", handlers.CreateUser)
	router.Get("/users/:id", handlers.GetUserById)
	router.Put("/users/:id", handlers.UpdateUserById)
	router.Delete("/users/:id", handlers.DeleteUserById)

	// Category
	router.Get("/categories", handlers.GetCategories)
	router.Get("/categories/:id", handlers.GetCategoryById)
	router.Post("/categories", handlers.CreateCategory)

	// Posts
	router.Get("/posts", handlers.GetAllPosts)
}
