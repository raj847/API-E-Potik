package bussiness_test

import (
	"minpro_arya/features/company"
	bussiness "minpro_arya/features/company/bussiness"
	_companyMock "minpro_arya/features/company/mocks"
	"minpro_arya/helpers/encrypt"
	"minpro_arya/middleware"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
)

var (
	mockCompanyRepository _companyMock.Repository
	companyService        company.Service
	companyDomain         company.Domain
)

func TestMain(m *testing.M) {
	jwtAuth := &middleware.ConfigJWT{
		SecretJWT:       "minpro_arya",
		ExpiresDuration: 1,
	}
	companyService = bussiness.NewServiceCompany(&mockCompanyRepository, 1, jwtAuth)
	companyDomain = company.Domain{
		ID:           1,
		Username:     "kimiafarma",
		Email:        "kimiafarma@mail.com",
		Password:     "kimiafarma",
		Name:         "Kimia Farma",
		Phone_Number: "024-1111",
		Photo:        "kimiafarma.png",
	}
}

func TestRegister(t *testing.T) {
	t.Run("test case 1, valid register", func(t *testing.T) {
		password, _ := encrypt.HashingPassword("kimiafarma")
		outputDomain := company.Domain{
			Username:     "kimiafarma",
			Email:        "kimiafarma@mail.com",
			Password:     password,
			Name:         "Kimia Farma",
			Phone_Number: "024-1111",
			Photo:        "kimiafarma.png",
		}
		inputService := company.Domain{
			Username:     "kimiafarma",
			Email:        "kimiafarma@mail.com",
			Password:     "kimiafarma",
			Name:         "Kimia Farma",
			Phone_Number: "024-1111",
			Photo:        "kimiafarma.png",
		}
		mockCompanyRepository.On("Register", mock.Anything).Return(outputDomain, nil).Once()

		resp, err := companyService.Register(&inputService)
		assert.Nil(t, err)
		assert.Equal(t, inputService.Username, resp.Username)
	})

	t.Run("test case 2, invalid register", func(t *testing.T) {
		password, _ := encrypt.HashingPassword("kimiafarma")
		outputDomain := company.Domain{
			Username:     "kimiafarma",
			Email:        "kimiafarma@mail.com",
			Password:     password,
			Name:         "Kimia Farma",
			Phone_Number: "024-1111",
			Photo:        "kimiafarma.png",
		}
		inputService := company.Domain{
			Username:     "kimiafarma",
			Email:        "kimiafarma@mail.com",
			Password:     "kimiafarma",
			Name:         "Kimia Farma",
			Phone_Number: "024-1111",
			Photo:        "kimiafarma.png",
		}
		mockCompanyRepository.On("Register", mock.Anything).Return(outputDomain, bussiness.ErrInternalServer).Once()

		resp, err := companyService.Register(&inputService)
		assert.Empty(t, resp)
		assert.Equal(t, err, bussiness.ErrInternalServer)
	})
	t.Run("test case 3, invalid register duplicate", func(t *testing.T) {
		password, _ := encrypt.HashingPassword("kimiafarma")
		outputDomain := company.Domain{
			Username:     "kimiafarma",
			Email:        "kimiafarma@mail.com",
			Password:     password,
			Name:         "Kimia Farma",
			Phone_Number: "024-1111",
			Photo:        "kimiafarma.png",
		}
		inputService := company.Domain{
			Username:     "kimiafarma",
			Email:        "kimiafarma@mail.com",
			Password:     "kimiafarma",
			Name:         "Kimia Farma",
			Phone_Number: "024-1111",
			Photo:        "kimiafarma.png",
		}

		mockCompanyRepository.On("Register", mock.Anything).Return(company.Domain{}, bussiness.ErrDuplicateData).Once()

		resp, err := companyService.Register(&inputService)

		assert.NotNil(t, err)
		assert.NotEqual(t, outputDomain, resp)
	})
}

func TestLogin(t *testing.T) {
	t.Run("test case 1, valid login", func(t *testing.T) {
		password, _ := encrypt.HashingPassword("kimiafarma")
		outputDomain := company.Domain{
			Email:    "kimiafarma@mail.com",
			Password: password,
		}
		inputService := company.Domain{
			Email:    "kimiafarma@mail.com",
			Password: "kimiafarma",
		}
		mockCompanyRepository.On("Login", mock.AnythingOfType("string"), mock.AnythingOfType("string")).Return(outputDomain, nil).Once()

		resp, err := companyService.Login(inputService.Email, inputService.Password)
		assert.Nil(t, err)
		assert.NotEmpty(t, resp)
	})

	t.Run("test case 2, invalid login", func(t *testing.T) {
		password, _ := encrypt.HashingPassword("kimiafarma")
		outputDomain := company.Domain{
			Email:    "kimiafarma@mail.com",
			Password: password,
		}
		inputService := company.Domain{
			Email:    "kimiafarma@mail.com",
			Password: "fisikakarma",
		}
		mockCompanyRepository.On("Login", mock.AnythingOfType("string"), mock.AnythingOfType("string")).Return(outputDomain, bussiness.ErrEmailorPass).Once()

		resp, err := companyService.Login(inputService.Email, inputService.Password)
		assert.Empty(t, resp)
		assert.Equal(t, err, bussiness.ErrEmailorPass)
	})

	t.Run("test case 3, invalid login", func(t *testing.T) {
		password, _ := encrypt.HashingPassword("kimiafarma")
		outputDomain := company.Domain{
			Email:    "kimiafarma@mail.com",
			Password: password,
		}
		inputService := company.Domain{
			Email:    "fisikakarma@mail.com",
			Password: "kimiafarma",
		}
		mockCompanyRepository.On("Login", mock.AnythingOfType("string"), mock.AnythingOfType("string")).Return(outputDomain, bussiness.ErrEmailorPass).Once()

		resp, err := companyService.Login(inputService.Email, inputService.Password)
		assert.Empty(t, resp)
		assert.Equal(t, err, bussiness.ErrEmailorPass)
	})
}
