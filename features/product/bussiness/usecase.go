// package bussiness

// import (
// 	"minpro_arya/features/product"

// )

// type serviceProduct struct {
// 	productRepository product.Repository
// }

// func NewServiceEvent(repoEvent Repository) Service {
// 	return &serviceEvent{
// 		eventRepository: repoEvent,
// 	}
// }

// func (serv *serviceEvent) AllEvent() ([]Domain, error) {

// 	result, err := serv.eventRepository.AllEvent()

// 	if err != nil {
// 		return []Domain{}, err
// 	}

// 	return result, nil
// }

// func (serv *serviceEvent) Create(orgID int, domain *Domain) (Domain, error) {

// 	result, err := serv.eventRepository.Create(orgID, domain)

// 	if err != nil {
// 		return Domain{}, err
// 	}

// 	return result, nil
// }

// func (serv *serviceEvent) Update(orgID int, evID int, domain *Domain) (Domain, error) {

// 	result, err := serv.eventRepository.Update(orgID, evID, domain)

// 	if err != nil {
// 		return Domain{}, err
// 	}

// 	return result, nil
// }

// func (serv *serviceEvent) Delete(orgID int, id int) (string, error) {

// 	result, err := serv.eventRepository.Delete(orgID, id)

// 	if err != nil {
// 		return "", business.ErrNotFound
// 	}

// 	return result, nil
// }

// func (serv *serviceEvent) MyEventByOrganizer(orgID int) ([]Domain, error) {

// 	result, err := serv.eventRepository.MyEventByOrganizer(orgID)

// 	if err != nil {
// 		return []Domain{}, err
// 	}

// 	return result, nil
// }

// func (serv *serviceEvent) EventByID(id int) (Domain, error) {

// 	result, err := serv.eventRepository.EventByID(id)

// 	if err != nil {
// 		return Domain{}, err
// 	}

// 	return result, nil
// }
// func (serv *serviceEvent) EventByIdOrganizer(orgzID int) ([]Domain, error) {

// 	result, err := serv.eventRepository.EventByIdOrganizer(orgzID)

// 	if err != nil {
// 		return []Domain{}, err
// 	}

// 	return result, nil

// }