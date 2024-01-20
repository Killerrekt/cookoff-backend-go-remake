package controller

import (
	"database/sql"
	"errors"
	"fmt"
	"net/http"
	"time"

	"github.com/CodeChefVIT/cookoff-backend/internal/models"
	"github.com/CodeChefVIT/cookoff-backend/internal/service"
	"github.com/CodeChefVIT/cookoff-backend/internal/utils"
	"github.com/labstack/echo/v4"
	"golang.org/x/crypto/bcrypt"
)

func Login(c echo.Context) error {
	var req struct {
		Email    string `json:"email" validate:"required"`
		Password string `json:"password" validate:"required"`
	}

	var res struct {
		AccessToken  string `json:"access_token"`
		RefreshToken string `json:"refresh_token"`
	}

	if err := c.Bind(&req); err != nil {
		return c.JSON(http.StatusBadRequest, models.Response{
			Message: "Failed to parse the request",
			Status:  false,
			Data:    res,
		})
	}

	if err := c.Validate(req); err != nil {
		return c.JSON(http.StatusBadRequest, models.Response{
			Message: "Didn't pass all the fields",
			Status:  false,
			Data:    res,
		})
	}

	user, err := service.FindUserByEmail(req.Email)
	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			return c.JSON(http.StatusConflict, models.Response{
				Message: "User with given email doesn't exists",
				Status:  false,
				Data:    res,
			})
		}
		return c.JSON(http.StatusInternalServerError, models.Response{
			Message: "Something went wrong : " + err.Error(),
			Status:  false,
			Data:    res,
		})
	}

	if err := bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(req.Password)); err != nil {
		return c.JSON(http.StatusConflict, models.Response{
			Message: "Invalid Password please try again",
			Status:  false,
			Data:    res,
		})
	}

	user.TokenVersion += 1
	res.AccessToken, err = utils.CreateAccessToken(user, time.Hour)
	if err != nil {
		return c.JSON(http.StatusFailedDependency, models.Response{
			Message: "Failed to Create the Accesstoken",
			Status:  false,
			Data:    res,
		})
	}

	res.RefreshToken, err = utils.CreateRefreshToken(user, time.Hour*24)
	if err != nil {
		return c.JSON(http.StatusFailedDependency, models.Response{
			Message: "Failed to Create the Refreshtoken",
			Status:  false,
			Data:    res,
		})
	}

	user.RefreshToken = res.RefreshToken
	err = service.UpdateUserTokenDetails(user)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, models.Response{
			Message: "Failed to save the token details(DB error : " + err.Error() + ")",
			Status:  false,
			Data:    res,
		})
	}

	return c.JSON(http.StatusAccepted, models.Response{
		Message: "User Successfully Logged in",
		Status:  true,
		Data:    res,
	})
}

func SignUp(c echo.Context) error {
	var req struct {
		Name     string `json:"name" validate:"required"`
		Email    string `json:"email" validate:"required"`
		RegNo    string `json:"reg_no" validate:"required"`
		Password string `json:"password" validate:"required"`
	}

	if err := c.Bind(&req); err != nil {
		return c.JSON(http.StatusBadRequest, models.Response{
			Message: "Failed to parse the data",
			Status:  false,
		})
	}

	fmt.Println(req)

	if err := c.Validate(&req); err != nil {
		return c.JSON(http.StatusBadRequest, models.Response{
			Message: "Didn't pass all the required fields",
			Status:  false,
		})
	}

	password, err := bcrypt.GenerateFromPassword([]byte(req.Password), 10)
	if err != nil {
		return c.JSON(http.StatusFailedDependency, models.Response{
			Message: "Failed to hash the password",
			Status:  false,
		})
	}

	err = service.CreateUser(req.Name, req.Email, string(password), req.RegNo)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, models.Response{
			Message: "DB error : " + err.Error(),
			Status:  false,
		})
	}

	return c.JSON(http.StatusAccepted, models.Response{
		Message: "User has been created successfully",
		Status:  true,
	})

}
