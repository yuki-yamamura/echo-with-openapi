package user

import (
	"net/http"

	"github.com/labstack/echo/v4"
	"github.com/yuki-yamamura/echo-with-openapi/internal/presentation/schema"
)

func CreateUser(c echo.Context) error {
	var user schema.PostUsersJSONBody
	if err := c.Bind(&user); err != nil {
		return c.JSON(http.StatusBadRequest, err)
	}
	var userId int64 = 1
	data := schema.User{
		Id:    userId,
		Name:  user.Name,
		Email: user.Email,
	}

	return c.JSON(http.StatusCreated, data)
}

func GetUsers(c echo.Context) error {
	return c.JSON(http.StatusOK, schema.GetUsersResponseDataJSONResponse{
		{
			Id:    1,
			Name:  "Alice",
			Email: "alice@example.com",
		},
		{
			Id:    2,
			Name:  "Bob",
			Email: "bob@example.com",
		},
	})
}

func GetUsersId(c echo.Context, id int64) error {
	return c.JSON(http.StatusOK, schema.GetUserResponseDataJSONResponse{
		Id:    id,
		Name:  "Alice",
		Email: "bob@example.com",
	})
}

func UpdateUser(c echo.Context, id int64) error {
	var user schema.PutUsersId200JSONResponse
	if err := c.Bind(&user); err != nil {
		return c.JSON(http.StatusBadRequest, err)
	}
	data := schema.User{
		Id:    id,
		Name:  user.Name,
		Email: user.Email,
	}
	return c.JSON(http.StatusOK, data)
}

func DeleteUser(c echo.Context, id int64) error {
	return c.NoContent(http.StatusNoContent)
}
