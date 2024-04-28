package controller

import (
	custom_errors "auth/internal/custom-errors"
	"auth/internal/model"
	"auth/internal/service"
	mock_service "auth/internal/service/mocks"
	"context"
	"fmt"
	"io"
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"

	"github.com/go-playground/assert/v2"
	"github.com/gofiber/fiber/v2"
	"github.com/golang/mock/gomock"
)

func TestRegister(t *testing.T) {
	type args struct {
		ctx        context.Context
		authParams service.AuthParams
		tokenModel model.Token
		tokens     service.Tokens
		user       model.User
	}
	type MockBehavior func(m *mock_service.MockIAuth, args args)

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
				authParams: service.AuthParams{
					Email:    "test1@gmail.com",
					Password: "12345",
				},
				tokenModel: model.Token{
					RefreshToken: "token",
					UserID:       1,
				},
				tokens: service.Tokens{
					AccessToken:  "token",
					RefreshToken: "token",
				},
				user: model.User{
					ID:       1,
					Email:    "test1@gmail.com",
					Password: "",
				},
			},
			mockBehavior: func(m *mock_service.MockIAuth, args args) {
				m.EXPECT().Register(gomock.Any(), gomock.Any()).Return(
					args.tokens,
					args.user,
					nil)
			},
			wantStatus:      200,
			wantRequestBody: `{"user":{"ID":1,"Email":"test1@gmail.com","Password":""},"refreshToken":"token","accessToken":"token"}`,
		},
		{
			name:            "field Email must have at least 8 characters",
			inputBody:       `{"email":"t4@g.c","password":"123455"}`,
			args:            args{},
			mockBehavior:    func(m *mock_service.MockIAuth, args args) {},
			wantStatus:      400,
			wantRequestBody: `{"message":"поле Email должно содержать как минимум 8 символов"}`,
		},
		{
			name:            "field Password must have at least 5 characters",
			inputBody:       `{"email":"test1@gmail.com","password":"1"}`,
			args:            args{},
			mockBehavior:    func(m *mock_service.MockIAuth, args args) {},
			wantStatus:      400,
			wantRequestBody: `{"message":"поле Password должно содержать как минимум 5 символов"}`,
		},
		{
			name:            "field Password is required",
			inputBody:       `{"email":"test1@gmail.com"}`,
			args:            args{},
			mockBehavior:    func(m *mock_service.MockIAuth, args args) {},
			wantStatus:      400,
			wantRequestBody: `{"message":"поле Password обязательно для заполнения"}`,
		},
		{
			name:            "field Email is required",
			inputBody:       `{"password":"12345"}`,
			args:            args{},
			mockBehavior:    func(m *mock_service.MockIAuth, args args) {},
			wantStatus:      400,
			wantRequestBody: `{"message":"поле Email обязательно для заполнения"}`,
		},
	}
	for _, tc := range testTable {
		t.Run(tc.name, func(t *testing.T) {
			c := gomock.NewController(t)
			defer c.Finish()

			// init service mock
			auth := mock_service.NewMockIAuth(c)
			tc.mockBehavior(auth, tc.args)

			// init service
			services := &service.Services{IAuth: auth}

			// init test server
			f := fiber.New()
			g := f.Group("/auth")
			newAuthRoutes(g, services.IAuth)

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

func TestLogin(t *testing.T) {
	type args struct {
		ctx        context.Context
		authParams service.AuthParams
		tokenModel model.Token
		tokens     service.Tokens
		user       model.User
	}

	type MockBehavior func(m *mock_service.MockIAuth, args args)

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
				authParams: service.AuthParams{
					Email:    "test1@gmail.com",
					Password: "12345",
				},
				tokenModel: model.Token{
					RefreshToken: "token",
					UserID:       1,
				},
				tokens: service.Tokens{
					AccessToken:  "token",
					RefreshToken: "token",
				},
				user: model.User{
					ID:       1,
					Email:    "test1@gmail.com",
					Password: "",
				},
			},
			mockBehavior: func(m *mock_service.MockIAuth, args args) {
				m.EXPECT().Login(gomock.Any(), gomock.Any()).Return(
					args.tokens,
					args.user, nil)
			},
			wantStatus:      200,
			wantRequestBody: `{"user":{"ID":1,"Email":"test1@gmail.com","Password":""},"refreshToken":"token","accessToken":"token"}`,
		},
		{
			name:      "wrong password",
			inputBody: `{"email":"test1@gmail.com","password":"4214214"}`,
			args: args{
				ctx: context.Background(),
				authParams: service.AuthParams{
					Email:    "test1@gmail.com",
					Password: "4214214",
				},
			},
			mockBehavior: func(m *mock_service.MockIAuth, args args) {
				m.EXPECT().Login(gomock.Any(), gomock.Any()).Return(
					args.tokens, model.User{}, custom_errors.ErrWrongCredetianls)
			},
			wantStatus:      400,
			wantRequestBody: `{"message":"неверная почта или пароль,попробуйте ещё раз"}`,
		},
		{
			name:      "user not found",
			inputBody: `{"email":"userNotFound@gmail.com","password":"12345"}`,
			args: args{
				ctx: context.Background(),
				authParams: service.AuthParams{
					Email:    "userNotFound@gmail.com",
					Password: "12345",
				},
			},
			mockBehavior: func(m *mock_service.MockIAuth, args args) {
				m.EXPECT().Login(gomock.Any(), gomock.Any()).Return(
					args.tokens,
					model.User{}, custom_errors.ErrUserNotFound)
			},
			wantStatus:      400,
			wantRequestBody: `{"message":"такого пользователья не существует"}`,
		},
		{
			name:            "field Email must have at least 8 characters",
			inputBody:       `{"email":"t4@g.c","password":"123455"}`,
			args:            args{},
			mockBehavior:    func(m *mock_service.MockIAuth, args args) {},
			wantStatus:      400,
			wantRequestBody: `{"message":"поле Email должно содержать как минимум 8 символов"}`,
		},
		{
			name:            "field Password must have at least 5 characters",
			inputBody:       `{"email":"test1@gmail.com","password":"1"}`,
			args:            args{},
			mockBehavior:    func(m *mock_service.MockIAuth, args args) {},
			wantStatus:      400,
			wantRequestBody: `{"message":"поле Password должно содержать как минимум 5 символов"}`,
		},
		{
			name:            "field Password is required",
			inputBody:       `{"email":"test1@gmail.com"}`,
			args:            args{},
			mockBehavior:    func(m *mock_service.MockIAuth, args args) {},
			wantStatus:      400,
			wantRequestBody: `{"message":"поле Password обязательно для заполнения"}`,
		},
		{
			name:            "field Email is required",
			inputBody:       `{"password":"12345"}`,
			args:            args{},
			mockBehavior:    func(m *mock_service.MockIAuth, args args) {},
			wantStatus:      400,
			wantRequestBody: `{"message":"поле Email обязательно для заполнения"}`,
		},
	}

	for _, tc := range testTable {
		t.Run(tc.name, func(t *testing.T) {
			c := gomock.NewController(t)
			defer c.Finish()

			// init service mock
			auth := mock_service.NewMockIAuth(c)
			tc.mockBehavior(auth, tc.args)

			// init service
			services := &service.Services{IAuth: auth}

			// init test server
			f := fiber.New()
			g := f.Group("/auth")
			newAuthRoutes(g, services.IAuth)

			// init request
			req := httptest.NewRequest("POST", "/auth/login", strings.NewReader(tc.inputBody))
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

func TestRefresh(t *testing.T) {
	type args struct {
		ctx        context.Context
		tokenModel model.Token
		tokens     service.Tokens
		user       model.User
	}
	type MockBehavior func(m *mock_service.MockIAuth, args args)

	testTable := []struct {
		name            string
		args            args
		mockBehavior    MockBehavior
		wantStatus      int
		wantRequestBody string
	}{
		{
			name: "OK",
			args: args{
				ctx: context.Background(),
				tokenModel: model.Token{
					RefreshToken: "token",
					UserID:       1,
				},
				tokens: service.Tokens{
					AccessToken:  "token",
					RefreshToken: "token",
				},
				user: model.User{
					ID:       1,
					Email:    "test1@gmail.com",
					Password: "",
				},
			},
			mockBehavior: func(m *mock_service.MockIAuth, args args) {
				m.EXPECT().Refresh(gomock.Any(), gomock.Any()).
					Return(
						args.tokens,
						args.user,
						nil)
			},
			wantStatus:      200,
			wantRequestBody: `{"user":{"ID":1,"Email":"test1@gmail.com","Password":""},"refreshToken":"token","accessToken":"token"}`,
		},
		{
			name:            "Empty refresh token",
			args:            args{},
			mockBehavior:    func(m *mock_service.MockIAuth, args args) {},
			wantStatus:      400,
			wantRequestBody: `{"message":"пользователь не авторизован"}`,
		},
		{
			name: "Unauthorized refresh token",
			args: args{},
			mockBehavior: func(m *mock_service.MockIAuth, args args) {
				m.EXPECT().Refresh(gomock.Any(), gomock.Any()).
					Return(
						args.tokens,
						args.user,
						custom_errors.ErrUserUnauthorized)
			},
			wantStatus:      401,
			wantRequestBody: `{"message":"пользователь не авторизован"}`,
		},
	}

	for _, tc := range testTable {
		t.Run(tc.name, func(t *testing.T) {
			c := gomock.NewController(t)
			defer c.Finish()

			// init service mock
			auth := mock_service.NewMockIAuth(c)
			tc.mockBehavior(auth, tc.args)

			// init service
			services := &service.Services{IAuth: auth}

			// init test server
			f := fiber.New()
			g := f.Group("/auth")
			newAuthRoutes(g, services.IAuth)

			// init request
			req := httptest.NewRequest("GET", "/auth/refresh", nil)
			req.Header.Set("Content-Type", "application/json")
			if tc.name != "Empty refresh token" {
				// set refresh token
				cookie := &http.Cookie{
					Name:     "refreshToken",
					Value:    "token",
					MaxAge:   30 * 24 * 60 * 60 * 1000,
					HttpOnly: true,
				}
				req.AddCookie(cookie)
			}

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
