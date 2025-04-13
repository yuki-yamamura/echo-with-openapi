package router

import (
	"encoding/json"
	"net/http"

	"github.com/getkin/kin-openapi/openapi3"
	"github.com/go-playground/validator/v10"
	"github.com/labstack/echo/v4"
	mw "github.com/oapi-codegen/echo-middleware"
	"github.com/yuki-yamamura/echo-with-openapi/internal/presentation/handlers/user"
	"github.com/yuki-yamamura/echo-with-openapi/internal/presentation/schema"
)

type Server struct{}

func (s *Server) PostUsers(ctx echo.Context) error {
	return user.CreateUser(ctx)
}

func (s *Server) GetUsers(ctx echo.Context) error {
	return user.GetUsers(ctx)
}

func (s *Server) GetUsersId(ctx echo.Context, id int64) error {
	return user.GetUsersId(ctx, id)
}

func (s *Server) PutUsersId(ctx echo.Context, id int64) error {
	return user.UpdateUser(ctx, id)
}

func (s *Server) DeleteUsersId(ctx echo.Context, id int64) error {
	return user.DeleteUser(ctx, id)
}

type RequestValidator struct {
	validator *validator.Validate
}

func (rv *RequestValidator) Validate(i interface{}) error {
	if err := rv.validator.Struct(i); err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, err.Error())
	}
	return nil
}

func NewRequestValidator() *RequestValidator {
	return &RequestValidator{
		validator: validator.New(),
	}
}

func WithSwaggerJson(s *openapi3.T) func(c echo.Context) error {
	return func(c echo.Context) error {
		c.Response().Header().Set(echo.HeaderContentType, echo.MIMEApplicationJSON)
		c.Response().WriteHeader(http.StatusOK)
		encoder := json.NewEncoder(c.Response())

		return encoder.Encode(s)
	}
}

func NewRouter() *echo.Echo {
	e := echo.New()
	swagger, err := schema.GetSwagger()
	if err != nil {
		e.Logger.Fatal(err)
	}

	docsGroup := e.Group("/docs")
	docsGroup.GET("/swagger.json", WithSwaggerJson(swagger))

	rootGroup := e.Group("")
	rootGroup.Use(mw.OapiRequestValidator(swagger))
	schema.RegisterHandlers(rootGroup, &Server{})
	return e
}
