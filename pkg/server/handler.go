package server

import (
	"Run_Hse_Run/pkg/service"
	"github.com/go-chi/chi"
	"github.com/go-chi/chi/middleware"
)

type Handler struct {
	services *service.Service
}

func NewHandler(services *service.Service) *Handler {
	return &Handler{services: services}
}

func (h *Handler) InitRoutes() chi.Router {
	router := chi.NewRouter()
	router.Use(middleware.Logger)

	router.Route("/api", func(router chi.Router) {
		router.Route("/game", func(router chi.Router) {
			router.Get("/get-rooms-by-code", h.getRoomByCode)
			router.Put("/put-in-queue", h.putInQueue)
			router.Delete("/delete-from-queue", h.deleteFromQueue)
			router.Put("/add-call", h.addCall)
			router.Delete("/delete-call", h.deleteCall)
			router.Put("/send-time", h.sendTime)
		})

	})

	return router
}
