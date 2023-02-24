package middleware

import (
	"github.com/gofiber/fiber/v2"
	"github.com/golang-jwt/jwt/v4"
	"github.com/pavanish/go-react-jwt/database"
	"github.com/pavanish/go-react-jwt/models"
)

const SecreteKey = "secrete"

func Authenticate() fiber.Handler {
	return func(c *fiber.Ctx) error {
		cookie := c.Cookies("jwt") // returns cookie Value by key
		token, err := jwt.ParseWithClaims(cookie, &jwt.StandardClaims{}, func(token *jwt.Token) (interface{}, error) {
			return []byte(SecreteKey), nil
		})
		if err != nil {
			c.Status(fiber.StatusUnauthorized)
			return c.JSON(fiber.Map{
				"message": "Unauthrised",
			})
		}

		claims := token.Claims.(*jwt.StandardClaims)
		if err := c.Next(); err != nil {
			return err
		}
		var user models.User
		database.DB.Where("id = ?", claims.Issuer).First(&user)
		//returnc.JSON(user)
		return c.SendString(user.Email)

		//log.Printf("authenticated log")
		//return nil
	}
}
