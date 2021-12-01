package bussiness_test

import (
	"minpro_arya/features/product"
	bussiness "minpro_arya/features/product/bussiness"
	_productMock "minpro_arya/features/product/mocks"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
)

var (
	mockProductRepository _productMock.Repository
	productService        product.Service
	productDomain         product.Domain
)

func TestMain(m *testing.M) {
	productService = bussiness.NewServiceProduct(&mockProductRepository)
	productDomain = product.Domain{
		ID:          1,
		CompanyID:   2,
		Name:        "panadol",
		Description: "banyak dipakai untuk meredakan sakit kepala ringan akut, nyeri ringan hingga sedang, serta demam.",
		NetWeight:   500,
		Price:       150000,
	}
}

func TestAllProduct(t *testing.T) {
	t.Run("test case 1, valid all product", func(t *testing.T) {
		expectedReturn := []product.Domain{
			{
				ID:          1,
				CompanyID:   2,
				Name:        "panadol",
				Description: "banyak dipakai untuk meredakan sakit kepala ringan akut, nyeri ringan hingga sedang, serta demam.",
				NetWeight:   500,
				Price:       150000,
			},
			{
				ID:          5,
				CompanyID:   2,
				Name:        "farmadol",
				Description: "digunakan untuk mengatasi demam dan sebagai pereda nyeri seperti sakit kepala, sakit gigi dan nyeri ringan lainnya pada bayi dan anak-anak.",
				NetWeight:   200,
				Price:       10000,
			},
		}
		mockProductRepository.On("AllProduct").Return(expectedReturn, nil).Once()
		_, err := productService.AllProduct()
		assert.Nil(t, err)
	})
	t.Run("test case 2, invalid all product", func(t *testing.T) {
		expectedReturn := []product.Domain{}
		mockProductRepository.On("AllProduct").Return(expectedReturn, assert.AnError).Once()
		_, err := productService.AllProduct()
		assert.Equal(t, err, assert.AnError)

	})
}

func TestCreate(t *testing.T) {
	t.Run("test case 1, valid create product", func(t *testing.T) {
		outputDomain := product.Domain{
			ID:          1,
			CompanyID:   2,
			Name:        "panadol",
			Description: "banyak dipakai untuk meredakan sakit kepala ringan akut, nyeri ringan hingga sedang, serta demam.",
			NetWeight:   500,
			Price:       150000,
		}
		inputService := product.Domain{
			ID:          1,
			CompanyID:   2,
			Name:        "panadol",
			Description: "banyak dipakai untuk meredakan sakit kepala ringan akut, nyeri ringan hingga sedang, serta demam.",
			NetWeight:   500,
			Price:       150000,
		}
		mockProductRepository.On("Create", mock.Anything, mock.Anything).Return(outputDomain, nil).Once()
		resp, err := productService.Create(inputService.CompanyID, &inputService)
		assert.Nil(t, err)
		assert.Equal(t, inputService.Name, resp.Name)
	})

	t.Run("test case 2, invalid create product", func(t *testing.T) {
		outputDomain := product.Domain{}
		inputService := product.Domain{
			ID:          1,
			CompanyID:   2,
			Name:        "panadol",
			Description: "banyak dipakai untuk meredakan sakit kepala ringan akut, nyeri ringan hingga sedang, serta demam.",
			NetWeight:   500,
			Price:       150000,
		}
		mockProductRepository.On("Create", mock.Anything, mock.Anything).Return(outputDomain, assert.AnError).Once()

		resp, err := productService.Create(inputService.CompanyID, &inputService)
		assert.Empty(t, resp)
		assert.Equal(t, err, assert.AnError)
	})
}

func TestUpdate(t *testing.T) {
	t.Run("test case 1, valid update product", func(t *testing.T) {
		outputDomain := product.Domain{
			ID:          1,
			CompanyID:   2,
			Name:        "panadol",
			Description: "banyak dipakai untuk meredakan sakit kepala ringan akut, nyeri ringan hingga sedang, serta demam.",
			NetWeight:   500,
			Price:       150000,
		}
		inputService := product.Domain{
			ID:          1,
			CompanyID:   2,
			Name:        "panadol",
			Description: "banyak dipakai untuk meredakan sakit kepala ringan akut, nyeri ringan hingga sedang, serta demam.",
			NetWeight:   500,
			Price:       150000,
		}
		mockProductRepository.On("Update", mock.Anything, mock.Anything, mock.Anything).Return(outputDomain, nil).Once()
		resp, err := productService.Update(inputService.CompanyID, inputService.ID, &inputService)
		assert.Nil(t, err)
		assert.Equal(t, inputService.Name, resp.Name)
	})

	t.Run("test case 2, invalid update product", func(t *testing.T) {
		outputDomain := product.Domain{}
		inputService := product.Domain{
			ID:          1,
			CompanyID:   2,
			Name:        "panadol",
			Description: "banyak dipakai untuk meredakan sakit kepala ringan akut, nyeri ringan hingga sedang, serta demam.",
			NetWeight:   500,
			Price:       150000,
		}
		mockProductRepository.On("Update", mock.Anything, mock.Anything, mock.Anything).Return(outputDomain, assert.AnError).Once()
		resp, err := productService.Update(inputService.CompanyID, inputService.ID, &inputService)

		assert.Empty(t, resp)
		assert.Equal(t, err, assert.AnError)
	})
}

func TestDelete(t *testing.T) {
	t.Run("test case 1, valid delete product", func(t *testing.T) {
		mockProductRepository.On("Delete", mock.Anything, mock.Anything).Return("Product has been delete", nil).Once()
		resp, err := productService.Delete(1, 1)
		assert.Nil(t, err)
		assert.Equal(t, "Product has been delete", resp)
	})

	t.Run("test case 2, invalid delete product", func(t *testing.T) {
		mockProductRepository.On("Delete", mock.Anything, mock.Anything).Return("", bussiness.ErrNotFound).Once()
		resp, err := productService.Delete(2, 2)
		assert.NotNil(t, err)
		assert.Equal(t, "", resp)
	})
}

func TestMyProductByCompany(t *testing.T) {
	t.Run("test case 1, valid all product", func(t *testing.T) {
		expectedReturn := []product.Domain{
			{
				ID:          1,
				CompanyID:   2,
				Name:        "panadol",
				Description: "banyak dipakai untuk meredakan sakit kepala ringan akut, nyeri ringan hingga sedang, serta demam.",
				NetWeight:   500,
				Price:       150000,
			},
			{
				ID:          5,
				CompanyID:   2,
				Name:        "farmadol",
				Description: "digunakan untuk mengatasi demam dan sebagai pereda nyeri seperti sakit kepala, sakit gigi dan nyeri ringan lainnya pada bayi dan anak-anak.",
				NetWeight:   200,
				Price:       10000,
			},
		}
		mockProductRepository.On("MyProductByCompany", mock.Anything).Return(expectedReturn, nil).Once()
		_, err := productService.MyProductByCompany(2)
		assert.Nil(t, err)
	})
	t.Run("test case 2, invalid all product", func(t *testing.T) {
		expectedReturn := []product.Domain{}
		mockProductRepository.On("MyProductByCompany", mock.Anything).Return(expectedReturn, assert.AnError).Once()
		_, err := productService.MyProductByCompany(2)
		assert.Equal(t, err, assert.AnError)

	})
}

func TestProductByID(t *testing.T) {
	t.Run("test case 1, valid all product", func(t *testing.T) {
		expectedReturn := product.Domain{
			ID:          1,
			CompanyID:   2,
			Name:        "panadol",
			Description: "banyak dipakai untuk meredakan sakit kepala ringan akut, nyeri ringan hingga sedang, serta demam.",
			NetWeight:   500,
			Price:       150000,
		}
		mockProductRepository.On("ProductByID", mock.Anything).Return(expectedReturn, nil).Once()
		_, err := productService.ProductByID(1)
		assert.Nil(t, err)
	})
	t.Run("test case 2, invalid all Product", func(t *testing.T) {
		expectedReturn := product.Domain{}
		mockProductRepository.On("ProductByID", mock.Anything).Return(expectedReturn, assert.AnError).Once()
		_, err := productService.ProductByID(1)
		assert.Equal(t, err, assert.AnError)

	})
}

func TestProductByIdCompany(t *testing.T) {
	t.Run("test case 1, valid all product", func(t *testing.T) {
		expectedReturn := []product.Domain{
			{
				ID:          1,
				CompanyID:   2,
				Name:        "panadol",
				Description: "banyak dipakai untuk meredakan sakit kepala ringan akut, nyeri ringan hingga sedang, serta demam.",
				NetWeight:   500,
				Price:       150000,
			},
			{
				ID:          5,
				CompanyID:   2,
				Name:        "farmadol",
				Description: "digunakan untuk mengatasi demam dan sebagai pereda nyeri seperti sakit kepala, sakit gigi dan nyeri ringan lainnya pada bayi dan anak-anak.",
				NetWeight:   200,
				Price:       10000,
			},
		}
		mockProductRepository.On("ProductByIdCompany", mock.Anything).Return(expectedReturn, nil).Once()
		_, err := productService.ProductByIdCompany(2)
		assert.Nil(t, err)
	})
	t.Run("test case 2, invalid all Product", func(t *testing.T) {
		expectedReturn := []product.Domain{}
		mockProductRepository.On("ProductByIdCompany", mock.Anything).Return(expectedReturn, assert.AnError).Once()
		_, err := productService.ProductByIdCompany(2)
		assert.Equal(t, err, assert.AnError)

	})
}
