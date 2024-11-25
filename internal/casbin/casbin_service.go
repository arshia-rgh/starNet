package casbin

import (
	"github.com/casbin/casbin/v2"
)

type CasbinService interface {
	Enforce(sub, obj, act string) (bool, error)
}

type casbinService struct {
	enforcer *casbin.Enforcer
}

func NewCasbinService() (CasbinService, error) {
	enforcer, err := casbin.NewEnforcer("internal/casbin/model.conf", "internal/casbin/policy.csv")
	if err != nil {
		return nil, err
	}
	return &casbinService{enforcer: enforcer}, nil
}

func (c *casbinService) Enforce(sub, obj, act string) (bool, error) {
	return c.enforcer.Enforce(sub, obj, act)
}
