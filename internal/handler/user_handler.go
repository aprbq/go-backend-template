package handler

import (
	"errors"
	"net/http"
	"strconv"

	"go-backend-template/internal/dto"
	"go-backend-template/internal/errs"
	"go-backend-template/internal/service"
	"go-backend-template/internal/utils"

	"github.com/gin-gonic/gin"
)

type userHandler struct {
	userSrv service.UserService
}

func NewUserHandler(userSrv service.UserService) userHandler {
	return userHandler{userSrv: userSrv}
}

// GetProfile returns the profile of the currently authenticated user.
func (h userHandler) GetProfile(c *gin.Context) {
	userID, err := currentUserID(c)
	if err != nil {
		utils.Error(c, http.StatusUnauthorized, err.Error())
		return
	}

	user, err := h.userSrv.GetProfile(userID)
	if err != nil {
		handleError(c, err)
		return
	}

	utils.Success(c, http.StatusOK, user)
}

// UpdateProfile updates the profile of the currently authenticated user.
func (h userHandler) UpdateProfile(c *gin.Context) {
	userID, err := currentUserID(c)
	if err != nil {
		utils.Error(c, http.StatusUnauthorized, err.Error())
		return
	}

	var input dto.UserRequest
	if err := c.ShouldBindJSON(&input); err != nil {
		utils.Error(c, http.StatusBadRequest, err.Error())
		return
	}

	user, err := h.userSrv.UpdateProfile(userID, input)
	if err != nil {
		handleError(c, err)
		return
	}

	utils.Success(c, http.StatusOK, user)
}

// DeleteUser deletes the user identified by the :id path parameter.
func (h userHandler) DeleteUser(c *gin.Context) {
	userID, err := parseUserID(c.Param("id"))
	if err != nil {
		utils.Error(c, http.StatusBadRequest, "invalid user id")
		return
	}

	if err := h.userSrv.DeleteUser(userID); err != nil {
		handleError(c, err)
		return
	}

	utils.Success(c, http.StatusOK, gin.H{"id": userID})
}

// GetAllUsers returns every user.
func (h userHandler) GetAllUsers(c *gin.Context) {
	users, err := h.userSrv.GetAllUsers()
	if err != nil {
		handleError(c, err)
		return
	}

	utils.Success(c, http.StatusOK, users)
}

// currentUserID reads the authenticated user id set by AuthMiddleware.
func currentUserID(c *gin.Context) (uint, error) {
	raw, ok := c.Get("userId")
	if !ok {
		return 0, errors.New("unauthorized")
	}

	id, ok := raw.(string)
	if !ok {
		return 0, errors.New("unauthorized")
	}

	return parseUserID(id)
}

func parseUserID(s string) (uint, error) {
	id, err := strconv.ParseUint(s, 10, 64)
	if err != nil {
		return 0, err
	}
	return uint(id), nil
}

// handleError maps AppError to its status code, falling back to 500.
func handleError(c *gin.Context, err error) {
	var appErr errs.AppError
	if errors.As(err, &appErr) {
		utils.Error(c, appErr.Code, appErr.Message)
		return
	}
	utils.Error(c, http.StatusInternalServerError, err.Error())
}
