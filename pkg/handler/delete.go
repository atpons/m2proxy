package handler

import (
	"github.com/atpons/m2proxy/pkg/packet"
	"github.com/atpons/m2proxy/pkg/request"
	"github.com/atpons/m2proxy/pkg/response"
	"github.com/atpons/m2proxy/pkg/storage"
)

func Delete(s storage.Storage, req request.Request) response.Response {
	err := s.Delete(req.Body[0:req.TotalBodyLength])

	switch err {
	case storage.ErrKeyNotFound:
		return *response.BuildResponse(req, req.Opcode, packet.StatusKeyNotFound, []byte{}, []byte{})
	case nil:
		return *response.BuildResponse(req, req.Opcode, packet.StatusNoError, []byte{}, []byte{})
	default:
		return *response.BuildResponse(req, req.Opcode, packet.StatusInternalError, []byte{}, []byte{})
	}
}
