package bussiness

import (
	"minpro_arya/features/company"
	"minpro_arya/helpers/encrypt"
	"minpro_arya/middleware"
	"time"
)

type serviceCompany struct {
	companyRepository company.Repository
	contextTimeout    time.Duration
	jwtAuth           *middleware.ConfigJWT
}

func NewServiceCompany(repoCompany company.Repository, timeout time.Duration, jwtauth *middleware.ConfigJWT) company.Service {
	return &serviceCompany{
		companyRepository: repoCompany,
		contextTimeout:    timeout,
		jwtAuth:           jwtauth,
	}
}

func (serv *serviceCompany) Register(domain *company.Domain) (company.Domain, error) {

	hashedPassword, err := encrypt.HashingPassword(domain.Password)

	if err != nil {
		return company.Domain{}, ErrInternalServer
	}

	domain.Password = hashedPassword

	result, err := serv.companyRepository.Register(domain)

	if result == (company.Domain{}) {
		return company.Domain{}, ErrDuplicateData
	}

	if err != nil {
		return company.Domain{}, ErrInternalServer
	}
	return result, nil
}

func (serv *serviceCompany) Login(email, password string) (company.Domain, error) {

	result, err := serv.companyRepository.Login(email, password)

	if err != nil {
		return company.Domain{}, ErrEmailorPass
	}

	checkPass := encrypt.CheckPasswordHash(password, result.Password)

	if !checkPass {
		return company.Domain{}, ErrEmailorPass
	}

	result.Token = serv.jwtAuth.GenerateToken(result.ID, "company")

	return result, nil
}
