package basicauth

import (
	"net/http/httptest"
	"testing"

	"github.com/gofiber/fiber/v2"
	"github.com/stretchr/testify/assert"
)

func Test_BasicAuth_Validation(t *testing.T) {
	app := fiber.New()
	app.Use(New(Config{Users: map[string]string{"user": "pass:word:123"}}))

	tests := []struct {
		name           string
		authHeader     string
		expectedStatus int
	}{
		{"Missing Prefix", "dXNlcjpwYXNz", 401},
		{"Invalid Base64", "Basic invalid_base64_chars_!!!", 401},
		{"Missing Colon", "Basic dXNlcg==", 401},
		{"Valid with Colons in Password", "Basic dXNlcjpwYXNzOndvcmQ6MTIz", 200},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			req := httptest.NewRequest("GET", "/", nil)
			req.Header.Set("Authorization", tt.authHeader)
			resp, _ := app.Test(req)
			assert.Equal(t, tt.expectedStatus, resp.StatusCode)
		})
	}
}