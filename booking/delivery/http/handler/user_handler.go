package handler

import (
	"net/http"
	"strconv"
	"booking/domain/entity"
	"booking/usecase/user"
	
	"github.com/gin-gonic/gin"
)

// UserHandler handles HTTP requests for user operations
type UserHandler struct {
	userUseCase user.UserUseCase
}

// NewUserHandler creates a new user handler
func NewUserHandler(userUseCase user.UserUseCase) *UserHandler {
	return &UserHandler{
		userUseCase: userUseCase,
	}
}

// CreateUserRequest represents the request body for creating a user
type CreateUserRequest struct {
	Email    string `json:"email" binding:"required"`
	Username string `json:"username" binding:"required"`
	Password string `json:"password" binding:"required"`
	FullName string `json:"full_name"`
	Phone    string `json:"phone"`
}

// UpdateUserRequest represents the request body for updating a user
type UpdateUserRequest struct {
	Email    string `json:"email"`
	Username string `json:"username"`
	Password string `json:"password"`
	FullName string `json:"full_name"`
	Phone    string `json:"phone"`
	IsActive *bool  `json:"is_active"`
}

// CreateUser handles POST /users
func (h *UserHandler) CreateUser(c *gin.Context) {
	var req CreateUserRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	
	user := &entity.User{
		Email:    req.Email,
		Username: req.Username,
		Password: req.Password,
		FullName: req.FullName,
		Phone:    req.Phone,
		IsActive: true,
	}
	
	if err := h.userUseCase.CreateUser(c.Request.Context(), user); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	
	c.JSON(http.StatusCreated, gin.H{
		"message": "User created successfully",
		"data":    user,
	})
}

// GetUser handles GET /users/:id
func (h *UserHandler) GetUser(c *gin.Context) {
	id, err := strconv.ParseUint(c.Param("id"), 10, 32)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid user ID"})
		return
	}
	
	user, err := h.userUseCase.GetUserByID(c.Request.Context(), uint(id))
	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "User not found"})
		return
	}
	
	c.JSON(http.StatusOK, gin.H{"data": user})
}

// ListUsers handles GET /users
func (h *UserHandler) ListUsers(c *gin.Context) {
	filter := &entity.UserFilter{}
	
	// Parse query parameters
	if email := c.Query("email"); email != "" {
		filter.Email = &email
	}
	if username := c.Query("username"); username != "" {
		filter.Username = &username
	}
	if isActive := c.Query("is_active"); isActive != "" {
		active := isActive == "true"
		filter.IsActive = &active
	}
	if limit := c.Query("limit"); limit != "" {
		if l, err := strconv.Atoi(limit); err == nil {
			filter.Limit = l
		}
	}
	if offset := c.Query("offset"); offset != "" {
		if o, err := strconv.Atoi(offset); err == nil {
			filter.Offset = o
		}
	}
	
	users, err := h.userUseCase.ListUsers(c.Request.Context(), filter)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	
	count, _ := h.userUseCase.CountUsers(c.Request.Context(), filter)
	
	c.JSON(http.StatusOK, gin.H{
		"data":  users,
		"total": count,
	})
}

// UpdateUser handles PUT /users/:id
func (h *UserHandler) UpdateUser(c *gin.Context) {
	id, err := strconv.ParseUint(c.Param("id"), 10, 32)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid user ID"})
		return
	}
	
	var req UpdateUserRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	
	user := &entity.User{
		ID:       uint(id),
		Email:    req.Email,
		Username: req.Username,
		Password: req.Password,
		FullName: req.FullName,
		Phone:    req.Phone,
	}
	
	if req.IsActive != nil {
		user.IsActive = *req.IsActive
	}
	
	if err := h.userUseCase.UpdateUser(c.Request.Context(), user); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	
	c.JSON(http.StatusOK, gin.H{
		"message": "User updated successfully",
		"data":    user,
	})
}

// DeleteUser handles DELETE /users/:id
func (h *UserHandler) DeleteUser(c *gin.Context) {
	id, err := strconv.ParseUint(c.Param("id"), 10, 32)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid user ID"})
		return
	}
	
	if err := h.userUseCase.DeleteUser(c.Request.Context(), uint(id)); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	
	c.JSON(http.StatusOK, gin.H{"message": "User deleted successfully"})
}

