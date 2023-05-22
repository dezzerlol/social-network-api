package auth

import (
	"context"
	"net/http"
	"social-network-api/internal/db/models"
	"social-network-api/internal/services/users"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/jackc/pgx/v5/pgxpool"
	"go.uber.org/zap"
)

type Handler interface {
}

type handler struct {
	logger      *zap.SugaredLogger
	userService users.Service
}

func New(logger *zap.SugaredLogger, db *pgxpool.Pool) Handler {
	return &handler{
		logger:      logger,
		userService: users.New(db),
	}
}

func (h *handler) Signup(c *gin.Context) {
	var input struct {
		Email     string `json:"email" binding:"required"`
		Password  string `json:"password" binding:"required"`
		Username  string `json:"username" binding:"required"`
		Firstname string `json:"firstname" binding:"required"`
		Lastname  string `json:"lastname" binding:"required"`
	}

	if err := c.ShouldBindJSON(input); err != nil {
		c.AbortWithError(http.StatusBadRequest, err)
		return
	}

	ctx, cancel := context.WithTimeout(c.Request.Context(), 5*time.Second)
	defer cancel()

	user := &models.User{
		Email:     input.Email,
		Password:  input.Password,
		Username:  input.Username,
		Firstname: input.Firstname,
		Lastname:  input.Lastname,
	}

	if err := h.userService.Create(ctx, user); err != nil {
		c.AbortWithError(http.StatusInternalServerError, err)
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"message": "User created successfully.",
		"user_id": user.Id,
	})

}

func (h *handler) Login(c *gin.Context) {
	var input struct {
		Email    string `json:"email" binding:"required"`
		Password string `json:"password" binding:"required"`
	}

	if err := c.ShouldBindJSON(input); err != nil {
		c.AbortWithError(http.StatusBadRequest, err)
		return
	}

	ctx, cancel := context.WithTimeout(c.Request.Context(), 5*time.Second)
	defer cancel()

	user, err := h.userService.FindByEmail(ctx, input.Email)

	if err != nil {
		c.AbortWithError(http.StatusInternalServerError, err)
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"token":   "TOKEN.",
		"user_id": user.Id,
	})
}
