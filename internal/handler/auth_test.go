package handler

import (
	"bytes"
	"errors"
	"net/http/httptest"
	"testing"

	"github.com/auth_test/internal/entity"
	"github.com/auth_test/internal/service"
	mock_service "github.com/auth_test/internal/service/mocks"
	"github.com/gin-gonic/gin"
	"github.com/golang/mock/gomock"
	"github.com/magiconair/properties/assert"
)

func TestHandler_signUp(t *testing.T) {
	type mockBehavior func(s *mock_service.MockAuthorization, user entity.User)

	testTable := []struct {
		name                string
		inputBody           string
		inputUser           entity.User
		mockBehavior        mockBehavior
		expectedStatusCode  int
		expectedRequestBody string
	}{
		{
			name:      "OK",
			inputBody: `{"username": "Test", "password": "Test"}`,
			inputUser: entity.User{
				Username: "Test",
				Password: "Test",
			},
			mockBehavior: func(s *mock_service.MockAuthorization, user entity.User) {
				s.EXPECT().CreateUser(user).Return(1, nil)
			},
			expectedStatusCode:  200,
			expectedRequestBody: `{"id":1}`,
		},
		{
			name:                "EMPTY FIELDS",
			inputBody:           `{"username": "Test"}`,
			mockBehavior:        func(s *mock_service.MockAuthorization, user entity.User) {},
			expectedStatusCode:  400,
			expectedRequestBody: `{"message":"Invalid input body"}`,
		},
		{
			name:      "SERVICE FAILURE",
			inputBody: `{"username": "Test", "password": "Test"}`,
			inputUser: entity.User{
				Username: "Test",
				Password: "Test",
			},
			mockBehavior: func(s *mock_service.MockAuthorization, user entity.User) {
				s.EXPECT().CreateUser(user).Return(1, errors.New("service failure"))
			},
			expectedStatusCode:  500,
			expectedRequestBody: `{"message":"service failure"}`,
		},
	}

	for _, testCase := range testTable {
		t.Run(testCase.name, func(t *testing.T) {
			c := gomock.NewController(t)
			defer c.Finish()

			auth := mock_service.NewMockAuthorization(c)
			testCase.mockBehavior(auth, testCase.inputUser)

			services := &service.Service{Authorization: auth}
			handler := NewHandler(services)

			r := gin.New()
			r.POST("/sign-up", handler.signUp)

			w := httptest.NewRecorder()
			req := httptest.NewRequest("POST", "/sign-up",
				bytes.NewBufferString(testCase.inputBody))

			r.ServeHTTP(w, req)

			assert.Equal(t, testCase.expectedRequestBody, w.Body.String())
			assert.Equal(t, testCase.expectedStatusCode, w.Code)
		})
	}
}
