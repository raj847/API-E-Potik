package request

import "minpro_arya/bussiness/admin"

type Admins struct {
	Username string `json:"username"`
	Password string `json:"password"`
	RoleID   uint   `json:"role_id"`
}

type AdminsLogin struct {
	Username string `json:"username"`
	Password string `json:"password"`
	RoleID   uint   `json:"role_id"`
}

func (req *Admins) ToDomain() *admin.Domain {
	return &admin.Domain{
		Username: req.Username,
		Password: req.Password,
		RoleID:   req.RoleID,
	}
}
