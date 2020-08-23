package db

import "bear/db/req"

var Module = NewDBModule()

const (
	GET_LOGIN_DATA = "getLoginData"
	CREATE_ROLE    = "createRole"
)

type ModuleReq struct {
	msgtype string
	handler func(*req.Request)
}

/**
*** All db module requests register here
**/
var requests = []*ModuleReq{
	{msgtype: GET_LOGIN_DATA, handler: req.GetLoginData},
	{msgtype: CREATE_ROLE, handler: req.CreateRole},
}
