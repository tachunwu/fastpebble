package api

import (
	"net/http"

	"github.com/gofiber/fiber/v2"
	"github.com/tachunwu/fastpebble/pkg/api/presenter"
	"github.com/tachunwu/fastpebble/pkg/entity"
	"github.com/tachunwu/fastpebble/pkg/fastpebble"
)

func Get(service fastpebble.Service) fiber.Handler {
	return func(c *fiber.Ctx) error {
		k := c.Params("key")

		result, err := service.Get(k)
		if err != nil {
			c.Status(http.StatusBadRequest)
			return c.JSON(presenter.ErrorResponse(err))
		}
		return c.JSON(presenter.SuccessResponse(&[]entity.KeyValue{*result}))
	}
}

func Set(service fastpebble.Service) fiber.Handler {
	return func(c *fiber.Ctx) error {
		k := c.Params("key")
		v := &entity.Value{}
		err := c.BodyParser(v)
		if err != nil {
			c.Status(http.StatusBadRequest)
			return c.JSON(presenter.ErrorResponse(err))
		}

		kv := &entity.KeyValue{
			Key:   k,
			Value: v.Value,
		}

		result, err := service.Set(kv)
		if err != nil {
			c.Status(http.StatusBadRequest)
			return c.JSON(presenter.ErrorResponse(err))
		}
		return c.JSON(presenter.SuccessResponse(&[]entity.KeyValue{*result}))
	}
}

func Delete(service fastpebble.Service) fiber.Handler {
	return func(c *fiber.Ctx) error {

		k := c.Params("key")
		kv := &entity.KeyValue{
			Key: k,
		}
		err := service.Delete(k)
		if err != nil {
			c.Status(http.StatusBadRequest)
			return c.JSON(presenter.ErrorResponse(err))
		}
		return c.JSON(presenter.SuccessResponse(&[]entity.KeyValue{*kv}))
	}
}

func Scan(service fastpebble.Service) fiber.Handler {
	return func(c *fiber.Ctx) error {

		sr := &entity.ScanRequest{}
		err := c.BodyParser(sr)
		if err != nil {
			return c.JSON(presenter.ErrorResponse(err))
		}

		results, err := service.Scan(*sr)
		if err != nil {
			c.Status(http.StatusBadRequest)
			return c.JSON(presenter.ErrorResponse(err))
		}
		return c.JSON(presenter.SuccessResponse(results))
	}
}

func Batch(service fastpebble.Service) fiber.Handler {
	return func(c *fiber.Ctx) error {
		batch := &entity.BatchRequest{}
		if err := c.BodyParser(batch); err != nil {
			return c.JSON(presenter.ErrorResponse(err))
		}

		results, err := service.Batch(batch)

		if err != nil {
			c.Status(http.StatusBadRequest)
			return c.JSON(presenter.ErrorResponse(err))
		}
		return c.JSON(
			presenter.SuccessResponse(results))
	}

}
