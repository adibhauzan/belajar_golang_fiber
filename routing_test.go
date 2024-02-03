package main

import (
	"fmt"
	"github.com/gofiber/fiber/v2"
	"github.com/stretchr/testify/assert"
	"io"
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"
)

func TestRouting(t *testing.T) {
	app := fiber.New()

	app.Get("/", func(ctx *fiber.Ctx) error {
		return ctx.SendString("Hello world")
	})

	request := httptest.NewRequest(fiber.MethodGet, "/", nil)
	response, err := app.Test(request)
	assert.Nil(t, err)
	assert.Equal(t, 200, response.StatusCode)

	bytes, err := io.ReadAll(response.Body)
	assert.Nil(t, err)
	assert.Equal(t, "Hello world", string(bytes))
}

func TestCTX(t *testing.T) {
	app := fiber.New()

	app.Get("/hello", func(ctx *fiber.Ctx) error {
		name := ctx.Query("name", "guest")
		return ctx.SendString("Hello " + name)
	})

	request := httptest.NewRequest(fiber.MethodGet, "/hello?name=Adib", nil)
	response, err := app.Test(request)
	assert.Nil(t, err)
	assert.Equal(t, 200, response.StatusCode)

	bytes, err := io.ReadAll(response.Body)
	assert.Nil(t, err)
	assert.Equal(t, "Hello Adib", string(bytes))

	fmt.Println(string(bytes))
}

func TestHttpRequest(t *testing.T) {
	app := fiber.New()

	app.Get("/request", func(ctx *fiber.Ctx) error {
		first := ctx.Get("firstname")
		last := ctx.Cookies("lastname")
		return ctx.SendString("Hello " + first + " " + last)
	})

	request := httptest.NewRequest(fiber.MethodGet, "/request", nil)
	request.Header.Set("firstname", "Adib")
	request.AddCookie(&http.Cookie{Name: "lastname", Value: "Hauzan"})
	response, err := app.Test(request)
	assert.Nil(t, err)
	assert.Equal(t, 200, response.StatusCode)

	bytes, err := io.ReadAll(response.Body)
	assert.Nil(t, err)
	assert.Equal(t, "Hello Adib Hauzan", string(bytes))

	fmt.Println(string(bytes))
}

func TestRouteParameter(t *testing.T) {
	app := fiber.New()

	app.Get("/users/:userId/orders/:orderId", func(ctx *fiber.Ctx) error {
		userIdParams := ctx.Params("userId")
		orderIdParams := ctx.Params("orderId")
		return ctx.SendString("Get Order " + orderIdParams + " From User " + userIdParams)
	})

	request := httptest.NewRequest(fiber.MethodGet, "/users/1/orders/2", nil)
	response, err := app.Test(request)
	assert.Nil(t, err)
	assert.Equal(t, 200, response.StatusCode)

	bytes, err := io.ReadAll(response.Body)
	assert.Nil(t, err)
	assert.Equal(t, "Get Order 2 From User 1", string(bytes))
	fmt.Println(string(bytes))
}

func TestFormRequest(t *testing.T) {
	app := fiber.New()

	app.Post("/hello", func(ctx *fiber.Ctx) error {
		name := ctx.FormValue("name")
		return ctx.SendString("Hello " + name)
	})

	body := strings.NewReader("name=Adib")
	request := httptest.NewRequest(fiber.MethodPost, "/hello", body)
	request.Header.Set("Content-Type", "application/x-www-form-urlencoded")
	response, err := app.Test(request)
	assert.Nil(t, err)
	assert.Equal(t, 200, response.StatusCode)

	bytes, err := io.ReadAll(response.Body)
	assert.Nil(t, err)
	assert.Equal(t, "Hello Adib", string(bytes))
	fmt.Println(string(bytes))
}
