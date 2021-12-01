package bussiness_test

import (
	"minpro_arya/features/customer"
	bussiness "minpro_arya/features/customer/bussiness"
	_customerMock "minpro_arya/features/customer/mocks"
	"minpro_arya/helpers/encrypt"
	"minpro_arya/middleware"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
)

var (
	mockCustomerRepository _customerMock.Repository
	customerService        customer.Service
	customerDomain         customer.Domain
)

func TestMain(m *testing.M) {
	jwtAuth := &middleware.ConfigJWT{
		SecretJWT:       "minpro_arya",
		ExpiresDuration: 1,
	}
	customerService = bussiness.NewServiceCustomer(&mockCustomerRepository, 1, jwtAuth)
	customerDomain = customer.Domain{
		ID:           1,
		Username:     "customer1",
		Email:        "customer1@mail.com",
		Password:     "customer1",
		Name:         "customer1",
		Dob:          "3-3-2001",
		Phone_Number: "0877777",
		Photo:        "customer1.jpg",
	}
}

func TestLogin(t *testing.T) {
	t.Run("test case 1 | valid test", func(t *testing.T) {
		password, _ := encrypt.HashingPassword("customer1")

		outputDomain := customer.Domain{
			Email:    "customer1@mail.com",
			Password: password,
		}

		inputService := customer.Domain{
			Email:    "customer1@mail.com",
			Password: "customer1",
		}

		mockCustomerRepository.On("Login", mock.AnythingOfType("string"), mock.AnythingOfType("string")).Return(outputDomain, nil).Once()

		resp, err := customerService.Login(inputService.Email, inputService.Password)

		assert.Nil(t, err)
		assert.NotEmpty(t, resp)
	})

	t.Run("test case 2 | wrong password test", func(t *testing.T) {
		password, _ := encrypt.HashingPassword("customer1")
		outputDomain := customer.Domain{
			Email:    "customer1@mail.com",
			Password: password,
		}

		inputService := customer.Domain{
			Email:    "customer1@mail.com",
			Password: "curut1",
		}

		mockCustomerRepository.On("Login", mock.AnythingOfType("string"), mock.AnythingOfType("string")).Return(outputDomain, bussiness.ErrEmailorPass).Once()

		resp, err := customerService.Login(inputService.Email, inputService.Password)
		assert.Empty(t, resp)
		assert.Equal(t, err, bussiness.ErrEmailorPass)
	})

	t.Run("test case 3 | no email and password test", func(t *testing.T) {
		password, _ := encrypt.HashingPassword("customer1")
		outputDomain := customer.Domain{
			Email:    "customer1@mail.com",
			Password: password,
		}

		inputService := customer.Domain{
			Email:    "curut1@mail.com",
			Password: "curut1",
		}

		mockCustomerRepository.On("Login", mock.AnythingOfType("string"), mock.AnythingOfType("string")).Return(outputDomain, bussiness.ErrEmailorPass).Once()

		resp, err := customerService.Login(inputService.Email, inputService.Password)
		assert.Equal(t, err, bussiness.ErrEmailorPass)
		assert.Empty(t, resp)
	})
}

func TestRegister(t *testing.T) {

	t.Run("test case 1, valid register", func(t *testing.T) {
		password, _ := encrypt.HashingPassword("customer1")
		outputDomain := customer.Domain{
			Username:     "customer1",
			Email:        "customer1@mail.com",
			Password:     password,
			Name:         "customer1",
			Dob:          "3-3-2001",
			Phone_Number: "0877777",
			Photo:        "customer1.jpg",
		}
		inputService := customer.Domain{
			Username:     "customer1",
			Email:        "customer1@mail.com",
			Password:     "customer1",
			Name:         "customer1",
			Dob:          "3-3-2001",
			Phone_Number: "0877777",
			Photo:        "customer1.jpg",
		}
		mockCustomerRepository.On("Register", mock.Anything).Return(outputDomain, nil).Once()

		resp, err := customerService.Register(&inputService)
		assert.Nil(t, err)
		assert.Equal(t, inputService.Username, resp.Username)
	})
	t.Run("test case 2, fail registration", func(t *testing.T) {
		outputDomain := customer.Domain{}
		inputService := customer.Domain{
			Username:     "customer1",
			Email:        "customer1@mail.com",
			Password:     "customer1",
			Name:         "customer1",
			Dob:          "3-3-2001",
			Phone_Number: "0877777",
			Photo:        "customer1.jpg",
		}
		mockCustomerRepository.On("Register", mock.Anything).Return(outputDomain, bussiness.ErrInternalServer).Once()

		resp, err := customerService.Register(&inputService)
		assert.Empty(t, resp)
		assert.Equal(t, err, bussiness.ErrInternalServer)
	})

	t.Run("test case 3, fail hashed", func(t *testing.T) {
		password, _ := encrypt.HashingPassword("sssssss")
		outputDomain := customer.Domain{}
		inputService := customer.Domain{
			Username:     "customer1",
			Email:        "customer1@mail.com",
			Password:     password,
			Name:         "customer1",
			Dob:          "3-3-2001",
			Phone_Number: "0877777",
			Photo:        "customer1.jpg",
		}
		mockCustomerRepository.On("Register", mock.Anything).Return(outputDomain, bussiness.ErrInternalServer).Once()

		resp, err := customerService.Register(&inputService)
		assert.Empty(t, resp)
		assert.Equal(t, err, bussiness.ErrInternalServer)
	})
}
