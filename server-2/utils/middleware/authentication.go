package middleware

import (
	"fmt"
	"restaurant-management/server-2/utils/jwt"
	"strings"

	"github.com/gofiber/fiber/v2"
)

// Auth is the authentication middleware
func Auth(c *fiber.Ctx) error {
	h := c.Get("Authorization")

	if h == "" {
		return fiber.ErrUnauthorized
	}

	// Split the header into chunks
	chunks := strings.Split(h, " ")
	test := strings.Split(chunks[1], " ")

	// If header signature is not like `Bearer <token>`, then throw
	// This is also required, otherwise chunks[1] will throw out of bound error
	if len(chunks) < 2 {
		return fiber.ErrUnauthorized
	}

	// Verify the token which is in the chunks
	user, err := jwt.Verify(chunks[1])

	if err != nil {
		return fiber.ErrUnauthorized
	}
	fmt.Println(test[0])

	c.Locals("USER", user.ID)

	return c.Next()
}
