package bussiness

import (
	"minpro_arya/features/customer"
	"minpro_arya/helpers/encrypt"
	"minpro_arya/middleware"
	"time"
)

type serviceCustomer struct {
	customerRepository customer.Repository
	contextTimeout     time.Duration
	jwtAuth            *middleware.ConfigJWT
}

func NewServiceCustomer(repoCustomer customer.Repository, timeout time.Duration, jwtauth *middleware.ConfigJWT) customer.Service {
	return &serviceCustomer{
		customerRepository: repoCustomer,
		contextTimeout:     timeout,
		jwtAuth:            jwtauth,
	}
}

func (serv *serviceCustomer) Register(domain *customer.Domain) (customer.Domain, error) {

	hashedPassword, err := encrypt.HashingPassword(domain.Password)

	if err != nil {
		return customer.Domain{}, ErrInternalServer
	}

	domain.Password = hashedPassword

	result, err := serv.customerRepository.Register(domain)

	if err != nil {
		return customer.Domain{}, ErrInternalServer
	}
	return result, nil
}

func (serv *serviceCustomer) Login(email, password string) (customer.Domain, error) {

	result, err := serv.customerRepository.Login(email, password)

	if err != nil {
		return customer.Domain{}, ErrEmailorPass
	}

	checkPass := encrypt.CheckPasswordHash(password, result.Password)

	if !checkPass {
		return customer.Domain{}, ErrEmailorPass
	}

	result.Token = serv.jwtAuth.GenerateToken(result.ID, "customer")

	return result, nil
}
