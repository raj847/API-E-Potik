package bussiness

import (
	"minpro_arya/features/transactions"
	randomcode "minpro_arya/helpers/random-code"
)

type serviceTrans struct {
	transRepository transactions.Repository
}

func NewServiceTrans(repoTrans transactions.Repository) transactions.Service {
	return &serviceTrans{
		transRepository: repoTrans,
	}
}

func (serv *serviceTrans) Trans(userID int, domain *transactions.Domain) (transactions.Domain, error) {

	randCode, _ := randomcode.GenerateCode(8)

	domain.Transaction_code = randCode

	result, err := serv.transRepository.Trans(userID, domain)

	if err != nil {
		return transactions.Domain{}, ErrInternalServer
	}

	return result, nil
}

func (serv *serviceTrans) GetTransByID(id int) (transactions.Domain, error) {

	result, err := serv.transRepository.GetTransByID(id)

	if err != nil {
		return transactions.Domain{}, ErrNotFound
	}

	return result, nil

}

func (serv *serviceTrans) GetAllTrans() ([]transactions.Domain, error) {

	result, err := serv.transRepository.GetAllTrans()

	if err != nil {
		return []transactions.Domain{}, err
	}

	return result, err
}

func (serv *serviceTrans) GetAllUserTrans(userID int) ([]transactions.Domain, error) {

	result, err := serv.transRepository.GetAllUserTrans(userID)

	if err != nil {
		return []transactions.Domain{}, err
	}

	return result, err
}
