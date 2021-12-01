package bussiness_test

import (
	"minpro_arya/features/transactions"
	"minpro_arya/features/transactions/bussiness"
	_transMock "minpro_arya/features/transactions/mocks"
	randomcode "minpro_arya/helpers/random-code"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
)

var (
	mockTransRepository _transMock.Repository
	transService        transactions.Service
	transDomain         transactions.Domain
)

func TestMain(m *testing.M) {
	transService = bussiness.NewServiceTrans(&mockTransRepository)
	transDomain = transactions.Domain{
		ID:         1,
		CustomerID: 1,
		ProductID:  5,
	}
	m.Run()
}

func TestTrans(t *testing.T) {
	t.Run("test case 1, valid trans", func(t *testing.T) {
		uniq, _ := randomcode.GenerateCode(8)
		outputDomain := transactions.Domain{
			ID:               1,
			CustomerID:       1,
			ProductID:        5,
			Transaction_code: uniq,
		}
		inputService := transactions.Domain{
			ID:               1,
			CustomerID:       1,
			ProductID:        5,
			Transaction_code: uniq,
		}
		mockTransRepository.On("Trans", mock.Anything, mock.Anything).Return(outputDomain, nil).Once()

		resp, err := transService.Trans(inputService.CustomerID, &inputService)
		assert.Nil(t, err)
		assert.Equal(t, inputService.ID, resp.ID)
	})

	t.Run("test case 1, invalid trans", func(t *testing.T) {
		uniq, _ := randomcode.GenerateCode(8)
		outputDomain := transactions.Domain{
			ID:               1,
			CustomerID:       2,
			ProductID:        5,
			Transaction_code: uniq,
		}
		inputService := transactions.Domain{
			ID:               1,
			CustomerID:       1,
			ProductID:        5,
			Transaction_code: uniq,
		}
		mockTransRepository.On("Trans", mock.Anything, mock.Anything).Return(transactions.Domain{}, bussiness.ErrInternalServer).Once()

		resp, err := transService.Trans(inputService.CustomerID, &inputService)
		assert.NotNil(t, err)
		assert.NotEqual(t, outputDomain, resp)
	})
}

func TestGetTransByID(t *testing.T) {
	t.Run("test case 1, valid get trans by id", func(t *testing.T) {
		uniq, _ := randomcode.GenerateCode(8)
		expectedReturn := transactions.Domain{
			ID:               1,
			CustomerID:       1,
			ProductID:        5,
			Transaction_code: uniq,
		}
		mockTransRepository.On("GetTransByID", mock.Anything).Return(expectedReturn, nil).Once()
		_, err := transService.GetTransByID(1)
		assert.Nil(t, err)
	})
	t.Run("test case 2, invalid all events", func(t *testing.T) {
		expectedReturn := transactions.Domain{}
		mockTransRepository.On("GetTransByID", mock.Anything).Return(expectedReturn, bussiness.ErrNotFound).Once()
		_, err := transService.GetTransByID(1)
		assert.Equal(t, err, bussiness.ErrNotFound)

	})
}

func TestGetAllTrans(t *testing.T) {
	t.Run("test case 1, valid all trans", func(t *testing.T) {
		uniq, _ := randomcode.GenerateCode(8)
		expectedReturn := []transactions.Domain{
			{
				ID:               1,
				CustomerID:       1,
				ProductID:        5,
				Transaction_code: uniq,
			},
			{
				ID:               2,
				CustomerID:       1,
				ProductID:        5,
				Transaction_code: uniq,
			},
		}
		mockTransRepository.On("GetAllTrans").Return(expectedReturn, nil).Once()
		_, err := transService.GetAllTrans()
		assert.Nil(t, err)
	})
	t.Run("test case 2, invalid all trans", func(t *testing.T) {
		expectedReturn := []transactions.Domain{}
		mockTransRepository.On("GetAllTrans").Return(expectedReturn, assert.AnError).Once()
		_, err := transService.GetAllTrans()
		assert.Equal(t, err, assert.AnError)

	})
}

func TestGetAllCustomerTrans(t *testing.T) {
	t.Run("test case 1, valid all events", func(t *testing.T) {
		uniq, _ := randomcode.GenerateCode(8)
		expectedReturn := []transactions.Domain{
			{
				ID:               1,
				CustomerID:       1,
				ProductID:        5,
				Transaction_code: uniq,
			},
			{
				ID:               2,
				CustomerID:       1,
				ProductID:        5,
				Transaction_code: uniq,
			},
		}
		mockTransRepository.On("GetAllCustomerTrans", mock.Anything).Return(expectedReturn, nil).Once()
		_, err := transService.GetAllCustomerTrans(2)
		assert.Nil(t, err)
	})
	t.Run("test case 2, invalid all events", func(t *testing.T) {
		expectedReturn := []transactions.Domain{}
		mockTransRepository.On("GetAllCustomerTrans", mock.Anything).Return(expectedReturn, assert.AnError).Once()
		_, err := transService.GetAllCustomerTrans(2)
		assert.Equal(t, err, assert.AnError)

	})
}
