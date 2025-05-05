package guide

import (
	"profkom/internal/models"
	"strconv"

	"github.com/gofiber/fiber/v2"
)

type Handler struct {
	service service
}

func New(service service) *Handler {
	return &Handler{
		service: service,
	}
}

func (h *Handler) GetGuide(c *fiber.Ctx) error {
	resp, err := h.service.GetGuide(c.Context())
	if err != nil {
		return err
	}

	return c.JSON(resp)
}

func (h *Handler) CreateGuide(c *fiber.Ctx) error {
	var guides []models.Guide

	if err := c.BodyParser(&guides); err != nil {
		return err
	}

	guidesType := c.Query("type")

	err := h.service.InsertGuide(c.Context(), guidesType, guides)
	if err != nil {
		return err
	}

	return c.SendStatus(fiber.StatusCreated)
}

func (h *Handler) DeleteGuide(c *fiber.Ctx) error {
	id, err := strconv.Atoi(c.Params("guide_id"))
	if err != nil {
		return err
	}

	err = h.service.DeleteGuide(c.Context(), id)
	if err != nil {
		return err
	}

	return err
}

func (h *Handler) DeleteTheme(c *fiber.Ctx) error {
	id, err := strconv.Atoi(c.Params("theme_id"))
	if err != nil {
		return err
	}

	err = h.service.DeleteTheme(c.Context(), id)
	if err != nil {
		return err
	}

	return err
}
