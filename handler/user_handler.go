package handler

import (
	"fmt"
	"golang-restapi/dto"
	"golang-restapi/models"
	"golang-restapi/repositories"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

type UserHandler struct {
	repo *repositories.UserRepository
}

func NewUserHandler(repo *repositories.UserRepository) *UserHandler {
	return &UserHandler{repo: repo}
}

// CreateUser godoc
// @Summary      Create new user
// @Description  Create a new user
// @Tags         User
// @Accept       json
// @Produce      json
// @Param        body  body      dto.CreateUser  true  "Create user payload"
// @Success      201   {object}  models.UserModel
// @Failure      400   {object}  map[string]string
// @Failure      409   {object}  map[string]string
// @Router       /v1/user [post]
func (h *UserHandler) CreateUser(c *gin.Context) {
	var req dto.CreateUser

	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSONP(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	exists, err := h.repo.ExistsByEmail(req.Email)

	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	if exists {
		c.JSON(http.StatusConflict, gin.H{"error": "Email already exists"})
		return
	}

	user := models.UserModel{
		Name:      req.Name,
		Email:     req.Email,
		Age:       req.Age,
		AvatarURL: req.AvatarURL,
	}

	if err := h.repo.CreatUser(&user); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
	}

	c.JSON(http.StatusCreated, user)

}

// UpdateUser godoc
// @Summary      Update user
// @Description  Update user by ID
// @Tags         User
// @Accept       json
// @Produce      json
// @Param        id    path      int                true  "User ID"
// @Param        body  body      dto.UpdateUserDTO  true  "Update user payload"
// @Success      200   {string}  string  "success"
// @Failure      400   {object}  map[string]string
// @Failure      409   {object}  map[string]string
// @Router       /v1/user/{id} [put]

func (h *UserHandler) UpdateUser(c *gin.Context) {
	var req dto.UpdateUserDTO

	if err := c.ShouldBindJSON(&req); err != nil {

		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	id, _ := strconv.Atoi(c.Param("id"))

	exists, err := h.repo.ExistsByEmail(req.Name)

	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	if exists {
		c.JSON(http.StatusConflict, gin.H{"error": "Email already exists"})
		return
	}

	if err := h.repo.UpdateUser(uint(id), map[string]interface{}{
		"name":      req.Name,
		"email":     req.Email,
		"age":       req.Age,
		"avatarURL": req.AvatarURL,
	}); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})

		return
	}
	c.JSON(http.StatusOK, gin.H{"message": "User updated"})
}

// DeleteUser godoc
// @Summary      Delete user
// @Description  Delete user by ID
// @Tags         User
// @Produce      json
// @Param        id   path      int  true  "User ID"
// @Success      200  {object}  models.UserModel
// @Failure      400  {object}  map[string]string
// @Router       /v1/user/{id} [delete]

func (h *UserHandler) DeleteUser(c *gin.Context) {

	id, err := strconv.Atoi(c.Param("id"))

	ID := uint(id)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid user id"})
		return
	}

	fmt.Println(id)
	exists, err := h.repo.ExistsById(ID)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": error.Error})
		return
	}

	fmt.Println("exists: ", exists)

	if !exists {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "User does not existed"})
		return
	}

	user := models.UserModel{
		Model: gorm.Model{
			ID: ID,
		},
	}

	if err := h.repo.DeleteUser(&user, ID); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "Delete user successfully"})
}

// GetAllUsers godoc
// @Summary      Get all users
// @Description  Get list of all users
// @Tags         User
// @Produce      json
// @Success      200  {array}  models.UserModel
// @Router       /v1/user [get]
func (h *UserHandler) GetAllUsers(c *gin.Context) {
	users, err := h.repo.GetAllUsers()

	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	// fmt.Println(users)
	c.JSON(http.StatusOK, users)

}
