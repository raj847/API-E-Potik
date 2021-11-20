package admin

import (
	"minpro_arya/bussiness"
)

type adminServices struct {
	userRepository Repository
	// jwtAuth        *middleware.ConfigJwt
}

func NewadminService(userRepo Repository) Service {
	return &adminServices{
		userRepository: userRepo,
	}
}

func (service *adminServices) Register(userData *Domain) (Domain, error) {

	// hashedPassword, _ := enkrips.Hash(userData.Password)
	// userData.Password = string(hashedPassword)
	res, err := service.userRepository.Register(userData)
	if res == (Domain{}) {
		return Domain{}, bussiness.ErrDuplicateData
	}
	if err != nil {
		return Domain{}, err
	}
	return res, nil
}

func (service *adminServices) Login(username, password string) error {
	_, err := service.userRepository.GetByUsername(username)
	if err != nil {
		return bussiness.ErrInvalidLoginInfo
	}
	return nil

	// if !enkrips.ValidateHash(password, userDomain.Password) {
	// 	return "", bussiness.ErrInvalidLoginInfo
	// }

	// token := service.jwtAuth.GenerateToken(int(userDomain.Id))
	// return token, nil
}

// func (service *adminServices) GetByID(id int) (error) {
// 	_, err := service.userRepository.GetByID(id)
// 	if err != nil {
// 		return err
// 	}
// 	return  nil
// }
