package handler

import (
	"github.com/go-chi/chi"
	"github.com/go-chi/render"
)

func (s *Server) routes() {
	s.router.Use(render.SetContentType(render.ContentTypeJSON))

	s.router.Route("/api/friendship", func(r chi.Router) {
		// r.Get("/", s.handleListFriends)
		// r.Post("/", s.handleCreateAlbum)
		// r.Route("/{id}", func(r chi.Router) {
		// 	r.Get("/", s.handleGetMovie)
		// 	r.Put("/", s.handleUpdateMovie)
		// 	r.Delete("/", s.handleDeleteMovie)
		// })
	})
}
