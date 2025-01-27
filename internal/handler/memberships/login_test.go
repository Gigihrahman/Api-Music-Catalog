package memberships

import (
	"api-music/internal/models/memberships"
	"bytes"
	"encoding/json"
	"fmt"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/gin-gonic/gin"
	"github.com/stretchr/testify/assert"
	"go.uber.org/mock/gomock"
)

func TestHandler_Login(t *testing.T) {
	ctrlMock := gomock.NewController(t)

	defer ctrlMock.Finish()
	mockSvc := NewMockservice(ctrlMock)

	tests := []struct {
		name               string
		mockFn             func()
		expectedStatusCode int
		expedtedBody       memberships.LoginResponse
		wantErr            bool
	}{
		{
			name: "succes",
			mockFn: func() {
				mockSvc.EXPECT().Login(memberships.LoginRequest{
					Email:    "test@gmail.com",
					Password: "password",
				}).Return("accesToken", nil)
			},
			expectedStatusCode: 200,
			expedtedBody: memberships.LoginResponse{
				AccesToken: "accesToken",
			},
			wantErr: false,
		}, {
			name: "fail",
			mockFn: func() {
				mockSvc.EXPECT().Login(memberships.LoginRequest{
					Email:    "test@gmail.com",
					Password: "password",
				}).Return("", assert.AnError)
			},
			expectedStatusCode: 400,
			expedtedBody:       memberships.LoginResponse{},
			wantErr:            true,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			tt.mockFn()
			api := gin.New()
			h := &Handler{
				Engine:  api,
				service: mockSvc,
			}
			h.RegisterRoute()
			w := httptest.NewRecorder()
			endpoint := `/memberships/login`
			model := memberships.LoginRequest{
				Email: "test@gmail.com",

				Password: "password",
			}
			val, err := json.Marshal(model)
			assert.NoError(t, err)
			body := bytes.NewReader(val)
			req, err := http.NewRequest(http.MethodPost, endpoint, body)
			assert.NoError(t, err)
			h.ServeHTTP(w, req)
			assert.Equal(t, tt.expectedStatusCode, w.Code)
			if !tt.wantErr {
				res := w.Result()
				defer res.Body.Close()
				response := memberships.LoginResponse{}
				err = json.Unmarshal(w.Body.Bytes(), &response)
				assert.NoError(t, err)
				fmt.Println("succes nih bang")
				assert.Equal(t, tt.expedtedBody, response)
			}
		})
	}
}
