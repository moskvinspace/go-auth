package handlers

import (
	"github.com/gin-gonic/gin"
	"github.com/moskvinspace/simple-web-app/pkg/models"
	"net/http"
	"regexp"
	"strings"
)

type signUpRequest struct {
	FirstName string `json:"first_name"`
	LastName  string `json:"last_name"`
	Email     string `json:"email"`
	Password1 string `json:"password_1"`
	Password2 string `json:"password_2"`
}

type signInRequest struct {
	Email    string `json:"email"`
	Password string `json:"password"`
}

func SignUp(c *gin.Context) {
	var req signUpRequest

	if err := c.BindJSON(&req); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	regexpName := regexp.MustCompile("[a-zA-Z]")
	if !regexpName.MatchString(req.FirstName) || !regexpName.MatchString(req.LastName) {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid symbol! "})
		return
	}

	regexpEmail := regexp.MustCompile("^[a-zA-Z0-9.!#$%&'*+/=?^_`{|}~-]+@[a-zA-Z0-9](?:[a-zA-Z0-9-]{0,61}[a-zA-Z0-9])?(?:\\.[a-zA-Z0-9](?:[a-zA-Z0-9-]{0,61}[a-zA-Z0-9])?)*$")
	if !regexpEmail.MatchString(req.Email) {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Oops! It doesn't look like correct email! "})
		return
	}

	if models.IsEmailExist(strings.ToLower(req.Email)) {
		c.JSON(400, gin.H{"error": "Email is already in use. "})
		return
	}

	if !isPasswordValid(req.Password1) {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Your password must be 8+ characters and contain at least one uppercase, one symbol and one number digit "})
		return
	}

	if req.Password1 != req.Password2 {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Passwords doesn't match "})
		return
	}

	hash, err := hashPassword(req.Password1)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	user := models.User{
		FirstName: req.FirstName,
		LastName:  req.LastName,
		Email:     strings.ToLower(req.Email),
		Password:  hash,
	}

	if err = user.Create(); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusCreated, user)
}

func SignIn(c *gin.Context) {
	var req signInRequest

	if err := c.BindJSON(req); err != nil {
		c.JSON(http.StatusInternalServerError, err.Error())
		return
	}

	if !models.IsEmailExist(strings.ToLower(req.Email)) {
		c.JSON(http.StatusUnauthorized, gin.H{"error": "Email doesn't exist! "})
		return
	}

	user, err := models.GetUser(strings.ToLower(req.Email))
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	if !checkPasswordHash(req.Password, user.Password) {
		c.JSON(http.StatusUnauthorized, gin.H{"error": "Password is wrong! "})
		return
	}

	c.JSON(http.StatusOK, gin.H{"token": ""})
}
