package implement

import (
	"idev-cms-service/config"
	"idev-cms-service/service/tokens"
	"idev-cms-service/service/users"
	wrp "idev-cms-service/service/users/wrapper"
	"idev-cms-service/service/util"
	"idev-cms-service/service/util/logs"
	"idev-cms-service/service/validator"
)

const (
	PREFIX_AUTH_TOKEN           = "auth-token"
	EmailActivateExpiredAtError = "activation links that have been sent can still be used"
	ActivateTokenHasExpired     = "activate token has expired"
)

var monthList = map[string]interface{}{
	"01": "มกราคม",
	"02": "กุมภาพันธ์",
	"03": "มีนาคม",
	"04": "เมษายน",
	"05": "พฤษภาคม",
	"06": "มิถุนายน",
	"07": "กรกฎาคม",
	"08": "สิงหาคม",
	"09": "กันยายน",
	"10": "ตุลาคม",
	"11": "พฤศจิกายน",
	"12": "ธันวาคม",
}

type implementation struct {
	*UserServiceConfig
}

type UserServiceConfig struct {
	Validator     validator.Validator
	RepoUsers     util.Repository
	RepoTokens    util.Repository
	RepoRoles     util.Repository
	Config        *config.Config
	ServiceTokens tokens.Service
	UUID          util.UUID
	GenCode       util.GenCode
	DateTime      util.DateTime
	FilterString  util.FilterString
	Encrypt       util.Encrypt
	Log           logs.Log
}

func New(config *UserServiceConfig) (service users.Service) {
	return &wrp.Wrapper{
		Service: &implementation{config},
	}
}
