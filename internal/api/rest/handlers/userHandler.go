package handlers

import (
	"github.com/gofiber/fiber/v2"
	"github.com/karokojnr/duka/internal/api/rest"
	"github.com/karokojnr/duka/internal/dto"
	"github.com/karokojnr/duka/internal/repository"
	"github.com/karokojnr/duka/internal/service"
	"net/http"
)

type UserHandler struct {
	svc service.UserService
}

func SetupUserRoutes(restHandler *rest.Handler) {
	app := restHandler.App

	svc := service.UserService{
		UsrRepo: repository.NewUserRepository(restHandler.DB),
		Auth:    restHandler.Auth,
	}
	handler := UserHandler{
		svc,
	}

	publicRouteGrp := app.Group("/users")

	// public endpoints
	publicRouteGrp.Post("/register", handler.Register)
	publicRouteGrp.Post("/login", handler.Login)

	privateRoutesGrp := publicRouteGrp.Group("/", restHandler.Auth.Authorize)

	// private endpoints
	privateRoutesGrp.Get("/verify", handler.GetVerificationCode)
	privateRoutesGrp.Post("/verify", handler.Verify)

	privateRoutesGrp.Get("/profile", handler.GetProfile)
	privateRoutesGrp.Post("/profile", handler.CreateProfile)

	privateRoutesGrp.Get("/cart", handler.GetCart)
	privateRoutesGrp.Post("/cart", handler.AddToCart)

	privateRoutesGrp.Get("/order", handler.GetOrders)
	privateRoutesGrp.Get("/order/:id", handler.GetOrder)

	privateRoutesGrp.Post("/role", handler.AddRole)

}

func (hndlr *UserHandler) Register(ctx *fiber.Ctx) error {
	user := dto.UserRegisterDto{}
	err := ctx.BodyParser(&user)
	if err != nil {
		return ctx.Status(http.StatusBadRequest).JSON(&fiber.Map{
			"message": "invalid inputs",
		})
	}

	token, err := hndlr.svc.Register(user)
	if err != nil {
		return ctx.Status(http.StatusInternalServerError).JSON(&fiber.Map{
			"message": "error signing up",
		})
	}

	return ctx.Status(http.StatusOK).JSON(&fiber.Map{
		"message": "Login successful",
		"data":    token,
	})
}

func (hndlr *UserHandler) Login(ctx *fiber.Ctx) error {
	loginInput := dto.UserLoginDto{}
	err := ctx.BodyParser(&loginInput)
	if err != nil {
		return ctx.Status(http.StatusBadRequest).JSON(&fiber.Map{
			"message": "invalid inputs",
		})
	}

	token, err := hndlr.svc.Login(loginInput)
	if err != nil {
		return ctx.Status(http.StatusInternalServerError).JSON(&fiber.Map{
			"message": "wrong email or password",
		})
	}
	return ctx.Status(http.StatusOK).JSON(&fiber.Map{
		"message": "Login successful",
		"data":    token,
	})
}

func (hndlr *UserHandler) GetVerificationCode(ctx *fiber.Ctx) error {
	return ctx.Status(http.StatusOK).JSON(&fiber.Map{
		"message": "get-verification-code",
	})
}

func (hndlr *UserHandler) Verify(ctx *fiber.Ctx) error {
	return ctx.Status(http.StatusOK).JSON(&fiber.Map{
		"message": "verify",
	})
}

func (hndlr *UserHandler) GetProfile(ctx *fiber.Ctx) error {
	user := hndlr.svc.Auth.GetCurrentUser(ctx)
	return ctx.Status(http.StatusOK).JSON(&fiber.Map{
		"message": "get-profile",
		"data":    user,
	})
}
func (hndlr *UserHandler) CreateProfile(ctx *fiber.Ctx) error {
	return ctx.Status(http.StatusOK).JSON(&fiber.Map{
		"message": "create-profile",
	})
}

func (hndlr *UserHandler) GetCart(ctx *fiber.Ctx) error {
	return ctx.Status(http.StatusOK).JSON(&fiber.Map{
		"message": "get-cart",
	})
}

func (hndlr *UserHandler) AddToCart(ctx *fiber.Ctx) error {
	return ctx.Status(http.StatusOK).JSON(&fiber.Map{
		"message": "add-to-cart",
	})
}

func (hndlr *UserHandler) GetOrders(ctx *fiber.Ctx) error {
	return ctx.Status(http.StatusOK).JSON(&fiber.Map{
		"message": "get-orders",
	})
}

func (hndlr *UserHandler) GetOrder(ctx *fiber.Ctx) error {
	return ctx.Status(http.StatusOK).JSON(&fiber.Map{
		"message": "get-order-by-id",
	})
}

func (hndlr *UserHandler) AddRole(ctx *fiber.Ctx) error {
	return ctx.Status(http.StatusOK).JSON(&fiber.Map{
		"message": "add-role",
	})
}
