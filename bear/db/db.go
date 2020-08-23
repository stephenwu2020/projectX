package db

import "bear/db/req"

var Module = NewDBModule()

/**
*** All db module requests register here
**/
var requests = []*req.ModuleReq{
	{MsgType: req.GET_LOGIN_DATA, Handler: req.GetLoginData},
	{MsgType: req.CREATE_ROLE, Handler: req.CreateRole},
}
