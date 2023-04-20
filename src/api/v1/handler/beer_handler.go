package handler

import (
	"errors"
	"fmt"
	"github.com/gofiber/fiber/v2"
	"karhub.backend.developer.test/src/api/v1/errors/apierr"
	"karhub.backend.developer.test/src/api/v1/handler/request"
	"karhub.backend.developer.test/src/api/v1/service"
)

type BeerHandler struct {
	beerService     service.BeerService
	playlistService service.PlaylistService
}

func NewBeerHandler(beerService service.BeerService, playlistService service.PlaylistService) BeerHandler {
	return BeerHandler{beerService: beerService, playlistService: playlistService}
}

func (h BeerHandler) HandleGetAll(c *fiber.Ctx) error {
	beers, err := h.beerService.GetAll()
	if err != nil {
		return fiber.NewError(fiber.StatusUnprocessableEntity, fmt.Errorf("failed to get all beers: %w", err).Error())
	}

	return c.JSON(beers)
}

func (h BeerHandler) HandleCreate(c *fiber.Ctx) error {
	var req request.CreateBeerRequest
	if err := c.BodyParser(&req); err != nil {
		return fiber.NewError(fiber.StatusBadRequest, fmt.Errorf("failed to parse body: %w", err).Error())
	}

	if err := req.Validate(); err != nil {
		return fiber.NewError(fiber.StatusBadRequest, fmt.Errorf("failed to validate body: %w", err).Error())
	}

	beer, err := h.beerService.Create(req.ToDomain())
	if err != nil {
		return fiber.NewError(fiber.StatusUnprocessableEntity, fmt.Errorf("failed to create beer: %w", err).Error())
	}

	return c.Status(fiber.StatusCreated).JSON(beer)
}

func (h BeerHandler) HandleUpdate(c *fiber.Ctx) error {
	id, err := c.ParamsInt("id")
	if err != nil {
		return fiber.NewError(fiber.StatusBadRequest, "invalid beer id")
	}

	var req request.UpdateBeerRequest
	err = c.BodyParser(&req)
	if err != nil {
		return fiber.NewError(fiber.StatusBadRequest, fmt.Errorf("failed to parse body: %w", err).Error())
	}

	err = req.Validate()
	if err != nil {
		return fiber.NewError(fiber.StatusBadRequest, fmt.Errorf("failed to validate body: %w", err).Error())
	}

	beer, err := h.beerService.Update(id, req.ToDomain())
	if err != nil {
		if errors.Is(err, apierr.ErrBeerNotFound) {
			return fiber.NewError(fiber.StatusNotFound, err.Error())
		}

		return fiber.NewError(fiber.StatusUnprocessableEntity, fmt.Errorf("failed to update beer: %w", err).Error())
	}

	return c.JSON(beer)
}

func (h BeerHandler) HandleDelete(c *fiber.Ctx) error {
	id, err := c.ParamsInt("id")
	if err != nil {
		return fiber.NewError(fiber.StatusBadRequest, "invalid beer id")
	}

	err = h.beerService.Delete(id)
	if err != nil {
		if errors.Is(err, apierr.ErrBeerNotFound) {
			return fiber.NewError(fiber.StatusNotFound, err.Error())
		}

		return fiber.NewError(fiber.StatusUnprocessableEntity, fmt.Errorf("failed to delete beer: %w", err).Error())
	}

	return c.SendStatus(fiber.StatusNoContent)
}

func (h BeerHandler) HandleGetClosestBeerStyles(c *fiber.Ctx) error {
	var req request.GetClosestBeerStylesRequest
	err := c.BodyParser(&req)
	if err != nil {
		return fiber.NewError(fiber.StatusBadRequest, fmt.Errorf("failed to parse body: %w", err).Error())
	}

	err = req.Validate()
	if err != nil {
		return fiber.NewError(fiber.StatusBadRequest, fmt.Errorf("failed to validate body: %w", err).Error())
	}

	beers, err := h.beerService.GetClosestBeerStyles(req.Temperature)
	if err != nil {
		return fiber.NewError(fiber.StatusUnprocessableEntity, fmt.Errorf("failed to get beer styles: %w", err).Error())
	}

	spotifyToken := c.Locals("token").(string)

	playlists, err := h.playlistService.SearchPlaylists(beers, spotifyToken)
	if err != nil {
		return fiber.NewError(fiber.StatusUnprocessableEntity, fmt.Errorf("failed to get spotify playlists: %w", err).Error())
	}

	return c.JSON(playlists)
}
