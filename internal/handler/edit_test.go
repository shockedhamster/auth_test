package handler

import (
	"bytes"
	"errors"
	"net/http/httptest"
	"testing"

	"github.com/auth_test/internal/service"
	mock_service "github.com/auth_test/internal/service/mocks"
	"github.com/gin-gonic/gin"
	"github.com/golang/mock/gomock"
	"github.com/magiconair/properties/assert"
)

func TestHandler_deleteUser(t *testing.T) {
	type mockBehavior func(s *mock_service.MockEdit, user string)

	testTable := []struct {
		name                string
		inputBody           string
		inputUser           string
		mockBehavior        mockBehavior
		expectedStatusCode  int
		expectedRequestBody string
	}{
		{
			name:      "OK",
			inputBody: `{"username": "Test"}`,
			inputUser: "Test",
			mockBehavior: func(s *mock_service.MockEdit, user string) {
				s.EXPECT().DeleteUser(user)
			},
			expectedStatusCode:  200,
			expectedRequestBody: `{"status":"user is deleted"}`,
		},
		{
			name:                "EMPTY FIELDS",
			inputBody:           `{"userna": "Test"}`,
			mockBehavior:        func(s *mock_service.MockEdit, user string) {},
			expectedStatusCode:  400,
			expectedRequestBody: `{"message":"Invalid input body"}`,
		},
		{
			name:      "SERVICE FAILURE",
			inputBody: `{"username": "Test"}`,
			inputUser: "Test",
			mockBehavior: func(s *mock_service.MockEdit, user string) {
				s.EXPECT().DeleteUser(user).Return(errors.New("service failure"))
			},
			expectedStatusCode:  500,
			expectedRequestBody: `{"message":"service failure"}`,
		},
	}

	for _, testCase := range testTable {
		t.Run(testCase.name, func(t *testing.T) {
			c := gomock.NewController(t)
			defer c.Finish()

			edit := mock_service.NewMockEdit(c)
			testCase.mockBehavior(edit, testCase.inputUser)

			services := &service.Service{Edit: edit}
			handler := NewHandler(services)

			r := gin.New()

			r.DELETE("/delete-user", handler.deleteUser)

			w := httptest.NewRecorder()
			req := httptest.NewRequest("DELETE", "/delete-user",
				bytes.NewBufferString(testCase.inputBody))

			r.ServeHTTP(w, req)

			assert.Equal(t, testCase.expectedRequestBody, w.Body.String())
			assert.Equal(t, testCase.expectedStatusCode, w.Code)

		})
	}

}

func TestHandler_updateUser(t *testing.T) {
	type mockBehavior func(s *mock_service.MockEdit, username, newUsername string)

	testTable := []struct {
		name                string
		inputBody           string
		inputUsername       string
		inputNewUsername    string
		mockBehavior        mockBehavior
		expectedStatusCode  int
		expectedRequestBody string
	}{
		{
			name:             "OK",
			inputBody:        `{"username": "Test", newusername: "newTest"}`,
			inputUsername:    "Test",
			inputNewUsername: "newTest",
			mockBehavior: func(s *mock_service.MockEdit, username, newUsername string) {
				s.EXPECT().UpdateUsername(username, newUsername)
			},
			expectedStatusCode:  200,
			expectedRequestBody: `{"status": "New username is set","id": 1,}`,
		},
		// {
		// 	name:                "EMPTY FIELDS",
		// 	inputBody:           `{"userna": "Test"}`,
		// 	mockBehavior:        func(s *mock_service.MockEdit, username, newUsername string) {},
		// 	expectedStatusCode:  400,
		// 	expectedRequestBody: `{"message":"Invalid input body"}`,
		// },
		// {
		// 	name:      "SERVICE FAILURE",
		// 	inputBody: `{"username": "Test"}`,
		// 	inputUser: "Test",
		// 	mockBehavior: func(s *mock_service.MockEdit, username, newUsername string) {
		// 		s.EXPECT().UpdateUsername(username, newUsername)
		// 	},
		// 	expectedStatusCode:  500,
		// 	expectedRequestBody: `{"message":"service failure"}`,
		// },
	}

	for _, testCase := range testTable {
		t.Run(testCase.name, func(t *testing.T) {
			c := gomock.NewController(t)
			defer c.Finish()

			edit := mock_service.NewMockEdit(c)
			testCase.mockBehavior(edit, testCase.inputUsername, testCase.inputNewUsername)

			services := &service.Service{Edit: edit}
			handler := NewHandler(services)

			r := gin.New()

			r.PATCH("/update-user", handler.updateUsername)

			w := httptest.NewRecorder()
			req := httptest.NewRequest("PATCH", "/update-user",
				bytes.NewBufferString(testCase.inputBody))

			r.ServeHTTP(w, req)

			assert.Equal(t, testCase.expectedRequestBody, w.Body.String())
			assert.Equal(t, testCase.expectedStatusCode, w.Code)

		})
	}

}
