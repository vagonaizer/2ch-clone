package http

import (
	"github.com/vladimirfedunov/2chan-clone/internal/usecase"
)

type Handler struct {
	BoardService  usecase.BoardService
	ThreadService usecase.ThreadService
	PostService   usecase.PostService
	AdminService  *usecase.AdminService
}

func NewHandler(board usecase.BoardService, thread usecase.ThreadService, post usecase.PostService, admin *usecase.AdminService) *Handler {
	return &Handler{
		BoardService:  board,
		ThreadService: thread,
		PostService:   post,
		AdminService:  admin,
	}
}
