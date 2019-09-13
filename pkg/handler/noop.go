package handler

import (
	"github.com/atpons/m2proxy/pkg/packet"
	"github.com/atpons/m2proxy/pkg/request"
	"github.com/atpons/m2proxy/pkg/response"
)

func Noop(req request.Request) response.Response {
	return *response.BuildResponse(req, packet.CmdNoop, packet.StatusNoError, []byte{}, []byte{})
}
