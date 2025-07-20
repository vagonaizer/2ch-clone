package http

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func (h *Handler) GetBoards(c *gin.Context) {
	boards, err := h.BoardService.ListBoards(c.Request.Context())
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, boards)
}

func (h *Handler) GetBoard(c *gin.Context) {
	slug := c.Param("slug")
	board, err := h.BoardService.GetBoard(c.Request.Context(), slug)
	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "board not found"})
		return
	}
	c.JSON(http.StatusOK, board)
}

func (h *Handler) IndexPage(c *gin.Context) {
	boards, err := h.BoardService.ListBoards(c.Request.Context())
	if err != nil {
		c.String(500, "Ошибка: %v", err)
		return
	}
	c.HTML(200, "index.html", gin.H{
		"Boards": boards,
	})
}

func (h *Handler) BoardPage(c *gin.Context) {
	boardSlug := c.Param("slug")
	previews, err := h.ThreadService.ListThreadPreviews(c.Request.Context(), boardSlug)
	if err != nil {
		c.HTML(500, "board.html", gin.H{"Error": err.Error()})
		return
	}
	board, err := h.BoardService.GetBoard(c.Request.Context(), boardSlug)
	if err != nil {
		c.String(404, "Доска не найдена")
		return
	}
	c.HTML(200, "board.html", gin.H{"Board": board, "ThreadPreviews": previews})
}
