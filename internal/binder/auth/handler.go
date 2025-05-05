package auth

import (
	"profkom/internal/models"
	"profkom/internal/service/auth"

	"github.com/gofiber/fiber/v2"
)

type Handler struct {
	service *auth.Service
}

func New(service *auth.Service) *Handler {
	return &Handler{
		service: service,
	}
}

func (h *Handler) SignUp(c *fiber.Ctx) error {
	var request models.SignUpRequest

	if err := c.BodyParser(&request); err != nil {
		return err
	}

	resp, err := h.service.AdminSingUp(c.Context(), request)
	if err != nil {
		return err
	}

	return c.JSON(resp)
}

func (h *Handler) PostInviteToken(c *fiber.Ctx) error {
	var request models.PostInviteTokenRequest

	if err := c.BodyParser(&request); err != nil {
		return err
	}

	resp, err := h.service.CreateInviteToken(c.Context(), request)
	if err != nil {
		return err
	}

	return c.JSON(resp)
}

func (h *Handler) SignIn(c *fiber.Ctx) error {
	var request models.AdminSignInRequest

	if err := c.BodyParser(&request); err != nil {
		return err
	}

	resp, err := h.service.AdminSignIn(c.Context(), request)
	if err != nil {
		return err
	}

	return c.JSON(resp)
}
