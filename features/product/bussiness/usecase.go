package bussiness

import (
	"minpro_arya/features/product"
)

type serviceProduct struct {
	productRepository product.Repository
}

func NewServiceProduct(repoProduct product.Repository) product.Service {
	return &serviceProduct{
		productRepository: repoProduct,
	}
}

func (serv *serviceProduct) AllProduct() ([]product.Domain, error) {

	result, err := serv.productRepository.AllProduct()

	if err != nil {
		return []product.Domain{}, err
	}

	return result, nil
}

func (serv *serviceProduct) Create(orgID int, domain *product.Domain) (product.Domain, error) {

	result, err := serv.productRepository.Create(orgID, domain)

	if err != nil {
		return product.Domain{}, err
	}

	return result, nil
}

func (serv *serviceProduct) Update(orgID int, evID int, domain *product.Domain) (product.Domain, error) {

	result, err := serv.productRepository.Update(orgID, evID, domain)

	if err != nil {
		return product.Domain{}, err
	}

	return result, nil
}

func (serv *serviceProduct) Delete(orgID int, id int) (string, error) {

	result, err := serv.productRepository.Delete(orgID, id)

	if err != nil {
		return "", ErrNotFound
	}

	return result, nil
}

func (serv *serviceProduct) MyProductByCompany(orgID int) ([]product.Domain, error) {

	result, err := serv.productRepository.MyProductByCompany(orgID)

	if err != nil {
		return []product.Domain{}, err
	}

	return result, nil
}

func (serv *serviceProduct) ProductByID(id int) (product.Domain, error) {

	result, err := serv.productRepository.ProductByID(id)

	if err != nil {
		return product.Domain{}, err
	}

	return result, nil
}
func (serv *serviceProduct) ProductByIdCompany(orgzID int) ([]product.Domain, error) {

	result, err := serv.productRepository.ProductByIdCompany(orgzID)

	if err != nil {
		return []product.Domain{}, err
	}

	return result, nil

}
