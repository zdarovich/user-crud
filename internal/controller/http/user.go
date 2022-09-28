package http

import (
	"binance-order-matcher/internal/model"
	"binance-order-matcher/internal/service"
	"binance-order-matcher/internal/service/util"
	"binance-order-matcher/pkg/logger"
	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
	"net/http"
	"strconv"
	"time"
)

const (
	defaultPage  = "0"
	defaultLimit = "100"
)

type userRoutes struct {
	u service.UserRepo
	l logger.Interface
}

func newUserRoutes(handler *gin.RouterGroup, t service.UserRepo, l logger.Interface) {
	r := &userRoutes{t, l}

	h := handler.Group("/users")
	{
		h.GET("", r.getUsers)
		h.POST("", r.postUser)
		h.PUT("", r.putUser)
		h.DELETE("", r.deleteUser)
	}
}

type usersResponse struct {
	Users []*model.User `json:"users"`
}

func (r *userRoutes) getUsers(c *gin.Context) {
	country := c.Query("country")
	page := c.DefaultQuery("page", defaultPage)
	limit := c.DefaultQuery("limit", defaultLimit)
	filter := model.User{}
	if len(country) != 0 {
		filter.Country = country
	}
	pageInt, err := strconv.Atoi(page)
	if err != nil {
		r.l.Error(err, "'page' query is invalid")
		errorResponse(c, http.StatusBadRequest, "'page' query is invalid")
		return
	}
	limitInt, err := strconv.Atoi(limit)
	if err != nil {
		r.l.Error(err, "'limit' query is invalid")
		errorResponse(c, http.StatusBadRequest, "'limit' query is invalid")
		return
	}
	users, err := r.u.Get(c.Request.Context(), pageInt, limitInt, filter)
	if err != nil {
		r.l.Error(err, "http - v1 - getUsers")
		errorResponse(c, http.StatusInternalServerError, "database problems")
		return
	}
	if len(users) == 0 {
		r.l.Error(err, "http - v1 - getUsers - notFound")
		errorResponse(c, http.StatusNotFound, "users not found")
		return
	}

	c.JSON(http.StatusOK, usersResponse{users})
}

type postUserRequest struct {
	Id        string `json:"id"`
	FirstName string `json:"first_name"`
	LastName  string `json:"last_name"`
	Nickname  string `json:"nickname"`
	Password  string `json:"password"`
	Email     string `json:"email"`
	Country   string `json:"country"`
}

func (r *userRoutes) postUser(c *gin.Context) {
	var request postUserRequest
	if err := c.ShouldBindJSON(&request); err != nil {
		r.l.Error(err, "http - v1 - postUser")
		errorResponse(c, http.StatusBadRequest, "invalid request body")
		return
	}
	id := uuid.New()
	newUser := &model.User{
		Id:        id.String(),
		FirstName: request.FirstName,
		LastName:  request.LastName,
		Nickname:  request.Nickname,
		Password:  request.Password,
		Email:     request.Email,
		Country:   request.Country,
		CreatedAt: time.Now(),
	}
	err := util.Validate(newUser)
	if err != nil {
		r.l.Error(err, "http - v1 - postUser")
		errorResponse(c, http.StatusBadRequest, err.Error())
		return
	}
	err = r.u.Save(
		c.Request.Context(),
		newUser,
	)
	if err != nil {
		r.l.Error(err, "http - v1 - postUser")
		errorResponse(c, http.StatusInternalServerError, "user service problems")
		return
	}

	c.JSON(http.StatusOK, newUser)
}

type putUserRequest struct {
	Id        string `json:"id"`
	FirstName string `json:"first_name"`
	LastName  string `json:"last_name"`
	Nickname  string `json:"nickname"`
	Password  string `json:"password"`
	Email     string `json:"email"`
	Country   string `json:"country"`
}

func (r *userRoutes) putUser(c *gin.Context) {
	var request putUserRequest
	if err := c.ShouldBindJSON(&request); err != nil {
		r.l.Error(err, "http - v1 - putUser")
		errorResponse(c, http.StatusBadRequest, "invalid request body")

		return
	}
	id := uuid.New()
	newUser := &model.User{
		Id:        id.String(),
		FirstName: request.FirstName,
		LastName:  request.LastName,
		Nickname:  request.Nickname,
		Password:  request.Password,
		Email:     request.Email,
		Country:   request.Country,
		UpdatedAt: time.Now(),
	}
	err := r.u.Update(
		c.Request.Context(),
		newUser,
	)
	if err != nil {
		r.l.Error(err, "http - v1 - putUser")
		errorResponse(c, http.StatusInternalServerError, "user service problems")

		return
	}

	c.JSON(http.StatusOK, newUser)
}

type deleteUserRequest struct {
	Id string `json:"id"`
}

func (r *userRoutes) deleteUser(c *gin.Context) {
	var request deleteUserRequest
	if err := c.ShouldBindJSON(&request); err != nil {
		r.l.Error(err, "http - v1 - deleteUser")
		errorResponse(c, http.StatusBadRequest, "invalid request body")
		return
	}
	id := uuid.New()
	userToDelete := &model.User{
		Id: id.String(),
	}
	err := r.u.Delete(
		c.Request.Context(),
		userToDelete,
	)
	if err != nil {
		r.l.Error(err, "http - v1 - deleteUser")
		errorResponse(c, http.StatusInternalServerError, "user service problems")
		return
	}

	c.JSON(http.StatusOK, userToDelete)
}
