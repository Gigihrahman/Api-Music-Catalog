package memberships

import (
	"api-music/internal/configs"
	"api-music/internal/models/memberships"
	"fmt"
	"testing"

	"github.com/stretchr/testify/assert"
	"go.uber.org/mock/gomock"
	"gorm.io/gorm"
)

func Test_service_Login(t *testing.T) {
	ctrlMock := gomock.NewController(t)
	defer ctrlMock.Finish()
	mockRepo := NewMockrepository(ctrlMock)
	type args struct {
		request memberships.LoginRequest
	}
	tests := []struct {
		name    string
		s       *service
		args    args
		want    string
		wantErr bool
		mockFn  func(args)
	}{
		{
			name: "success",
			args: args{
				request: memberships.LoginRequest{
					Email:    "test@gmail.com",
					Password: "password",
				},
			},
			wantErr: false,
			mockFn: func(args args) {
				mockRepo.EXPECT().GetUser(args.request.Email, "", uint(0)).Return(&memberships.User{
					Model: gorm.Model{
						ID: 1,
					},
					Email:    "test@gmail.com",
					Password: "$2a$12$wtKHxdS16rsg2o4J5EJCeeocBzDGx0R6OQd/wyW9fWP2VOym/9I7K",
					Username: "test",
				}, nil)
			},
		},
		{
			name: "fail when get data",
			args: args{
				request: memberships.LoginRequest{
					Email:    "test@gmail.com",
					Password: "password",
				},
			},
			wantErr: true,
			mockFn: func(args args) {
				mockRepo.EXPECT().GetUser(args.request.Email, "", uint(0)).Return(nil, assert.AnError)
			},
		},
		{
			name: "fail user and pass not match",
			args: args{
				request: memberships.LoginRequest{
					Email:    "test@gmail.com",
					Password: "password",
				},
			},
			wantErr: true,
			mockFn: func(args args) {
				mockRepo.EXPECT().GetUser(args.request.Email, "", uint(0)).Return(&memberships.User{
					Model: gorm.Model{
						ID: 1,
					},
					Email:    "test@gmail.com",
					Password: "wrongpasword",
					Username: "test",
				}, nil)
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {

			s := &service{
				cfg: &configs.Config{
					Service: configs.Service{
						SecretJWT: "abc",
					},
				},
				repository: mockRepo,
			}
			tt.mockFn(tt.args)
			got, err := s.Login(tt.args.request)
			if (err != nil) != tt.wantErr {
				t.Errorf("service.Login() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !tt.wantErr {
				fmt.Printf("test case:%+v", tt.name)
				assert.NotEmpty(t, got)
			} else {
				fmt.Printf("test case:%s\n", tt.name)
				assert.Empty(t, got)
			}
		})
	}
}
