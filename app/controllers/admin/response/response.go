package response

import (
	"minpro_arya/bussiness/admin"
)

type Admins struct {
	ID       uint   `json:"id"`
	Username string `json:"username"`
	Password string `json:"password"`
	RoleID   uint   `json:"role_id"`
}

func FromDomain(domain admin.Domain) Admins {
	return Admins{
		ID:       domain.Id,
		Username: domain.Username,
		Password: domain.Password,
		RoleID:   domain.RoleID,
	}
}
