package handler

import (
	"github.com/atpons/m2proxy/pkg/packet"
	"github.com/atpons/m2proxy/pkg/request"
	"github.com/atpons/m2proxy/pkg/response"
)

const (
	MemcachedVersion = "0.0.1"
)

func Version(request request.Request) response.Response {
	return *response.BuildResponse(request, packet.CmdVersion, packet.StatusNoError, []byte{}, []byte(MemcachedVersion))
}
