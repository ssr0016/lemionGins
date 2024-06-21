package main

import (
	"net/http"

	"gins/config"
	"gins/controller"
	"gins/helper"
	"gins/model"
	"gins/repository"
	"gins/service"

	"gins/router"

	"github.com/go-playground/validator/v10"
	"github.com/rs/zerolog/log"
)

func main() {

	log.Info().Msg("Starting server...")

	// Database
	db := config.DatabaseConnection()
	validate := validator.New()

	db.Table("tags").AutoMigrate(&model.Tags{})

	// Repository
	tagRespository := repository.NewTagsRepositoryImpl(db)

	// Service
	tagsService := service.NewTagsServiceImpl(&tagRespository, validate)

	// Controller
	tagController := controller.NewTagasController(tagsService)

	// Router

	router := router.NewRouter(tagController)

	server := &http.Server{
		Addr:    ":8080",
		Handler: router,
	}

	err := server.ListenAndServe()
	helper.ErrorPanic(err)
}
