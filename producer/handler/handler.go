package handler

import (
	"github.com/gofiber/fiber/v2"
	"github.com/prongbang/kafka/producer/command"
	"net/http"
)

type ChatHandler interface {
	Post(c *fiber.Ctx) error
	Reply(c *fiber.Ctx) error
}

type chatHandler struct {
	chatCommander command.ChatCommander
}

func (c2 *chatHandler) Post(c *fiber.Ctx) error {
	cmd := command.PostCommand{}
	_ = c.BodyParser(&cmd)

	if err := c2.chatCommander.Post(cmd); err != nil {
		return fiber.NewError(http.StatusBadRequest, err.Error())
	}
	return c.Status(http.StatusOK).JSON(fiber.Map{"message": "success"})
}

func (c2 *chatHandler) Reply(c *fiber.Ctx) error {
	cmd := command.ReplyCommand{}
	_ = c.BodyParser(&cmd)

	if err := c2.chatCommander.Replay(cmd); err != nil {
		return fiber.NewError(http.StatusBadRequest, err.Error())
	}
	return c.Status(http.StatusOK).JSON(fiber.Map{"message": "success"})
}

func NewChatHandler(chatCommander command.ChatCommander) ChatHandler {
	return &chatHandler{
		chatCommander,
	}
}
