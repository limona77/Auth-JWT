package controller

import (
	"auth/internal/model"
	"auth/internal/service"
	mock_service "auth/internal/service/mocks"
	"context"
	"github.com/go-playground/assert/v2"
	"github.com/gofiber/fiber/v2"
	"github.com/golang/mock/gomock"
	"net/http/httptest"
	"strings"
	"testing"
)

func TestSignUp(t *testing.T) {
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
				m.EXPECT().CreateUser(args.ctx, args.params).Return(model.User{
					ID:       1,
					Email:    "test1@gmail.com",
					Password: "",
				}, nil)
			},
			wantStatus:      200,
			wantRequestBody: `{"user": {"ID": 1,"Email": "test1@gmail.com","Password": ""}}`,
		},
	}

	for _, tc := range testTable {
		t.Run(tc.name, func(t *testing.T) {
			c := gomock.NewController(t)
			defer c.Finish()

			//init service mock
			auth := mock_service.NewMockAuth(c)
			tc.mockBehavior(auth, tc.args)

			//init service
			services := service.Services{Auth: auth}

			//init test server
			f := fiber.New()
			g := f.Group("/auth")
			newAuthRoutes(g, services.Auth)

			//init request
			req := httptest.NewRequest("POST", "http://127.0.0.1:8000/auth/register", strings.NewReader(tc.inputBody))
			req.Header.Set("Content-Type", "application/json")
			resp, _ := f.Test(req, 1)

			// check response
			if resp != nil {
				assert.Equal(t, tc.wantStatus, resp.Status)
				assert.Equal(t, tc.wantRequestBody, resp.Body)
			}
			return
		})
	}
}
