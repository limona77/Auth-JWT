package controller

import (
	"auth/internal/model"
	"auth/internal/service"
	mock_service "auth/internal/service/mocks"
	"context"
	"fmt"
	"io"
	"net/http/httptest"
	"strings"
	"testing"

	"github.com/go-playground/assert/v2"
	"github.com/gofiber/fiber/v2"
	"github.com/golang/mock/gomock"
)

func TestRegister(t *testing.T) {
	type args struct {
		ctx    context.Context
		params service.AuthParams
	}
	type MockBehavior func(m *mock_service.MockAuth, args args)

	testTable := []struct {
		name            string
		inputBody       string
		args            args
		mockBehavior    MockBehavior
		wantStatus      int
		wantRequestBody string
	}{
		{
			name:      "OK",
			inputBody: `{"email":"test1@gmail.com","password":"12345"}`,
			args: args{
				ctx: context.Background(),
				params: service.AuthParams{
					Email:    "test1@gmail.com",
					Password: "12345",
				},
			},
			mockBehavior: func(m *mock_service.MockAuth, args args) {
				m.EXPECT().CreateUser(gomock.Any(), gomock.Any()).Return(model.User{
					ID:       1,
					Email:    "test1@gmail.com",
					Password: "",
				}, nil)
				m.EXPECT().GenerateTokens(gomock.Any(), gomock.Any()).Return(service.Tokens{
					AccessToken:  "token",
					RefreshToken: "token",
				}, nil)
			},
			wantStatus:      200,
			wantRequestBody: `{"accessToken":"token","refreshToken":"token","user":{"ID":1,"Email":"test1@gmail.com","Password":""}}`,
		},
		{
			name:            "field Email must have at least 8 characters",
			inputBody:       `{"email":"t4@g.c","password":"123455"}`,
			args:            args{},
			mockBehavior:    func(m *mock_service.MockAuth, args args) {},
			wantStatus:      400,
			wantRequestBody: `{"message":"field Email must have at least 8 characters"}`,
		},
		{
			name:            "field Password must have at least 5 characters",
			inputBody:       `{"email":"test1@gmail.com","password":"1"}`,
			args:            args{},
			mockBehavior:    func(m *mock_service.MockAuth, args args) {},
			wantStatus:      400,
			wantRequestBody: `{"message":"field Password must have at least 5 characters"}`,
		},
		{
			name:            "field Password is required",
			inputBody:       `{"email":"test1@gmail.com"}`,
			args:            args{},
			mockBehavior:    func(m *mock_service.MockAuth, args args) {},
			wantStatus:      400,
			wantRequestBody: `{"message":"field Password is required"}`,
		},
		{
			name:            "field Email is required",
			inputBody:       `{"password":"12345"}`,
			args:            args{},
			mockBehavior:    func(m *mock_service.MockAuth, args args) {},
			wantStatus:      400,
			wantRequestBody: `{"message":"field Email is required"}`,
		},
	}
	for _, tc := range testTable {
		t.Run(tc.name, func(t *testing.T) {
			c := gomock.NewController(t)
			defer c.Finish()

			// init service mock
			auth := mock_service.NewMockAuth(c)
			tc.mockBehavior(auth, tc.args)

			// init service
			services := service.Services{Auth: auth}

			// init test server
			f := fiber.New()
			g := f.Group("/auth")
			newAuthRoutes(g, services.Auth)

			// init request
			req := httptest.NewRequest("POST", "/auth/register", strings.NewReader(tc.inputBody))
			req.Header.Set("Content-Type", "application/json")
			resp, err := f.Test(req)
			if err != nil {
				fmt.Println(err)
			}

			// check response
			if resp != nil {
				body, err := io.ReadAll(resp.Body)
				if err != nil {
					fmt.Println("Error reading response body:", err)
				}

				assert.Equal(t, tc.wantStatus, resp.StatusCode)
				assert.Equal(t, tc.wantRequestBody, string(body))
			}
		})
	}
}
