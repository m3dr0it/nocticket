package constant

import "noctiket/model/entity"

const (
	RoleAdmin           entity.Role = "admin"
	RoleUser            entity.Role = "user"
	RoleNocEngineer     entity.Role = "noc-engineer"
	RoleFieldEngineer   entity.Role = "field-engineer"
	RoleNetworkEngineer entity.Role = "network-engineer"
)

var (
	Engineers entity.RoleGroup = []entity.Role{RoleFieldEngineer, RoleNocEngineer, RoleNetworkEngineer}
)
