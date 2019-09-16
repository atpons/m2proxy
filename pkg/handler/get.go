package handler

import (
	"encoding/binary"
	"fmt"

	"github.com/atpons/m2proxy/pkg/packet"
	"github.com/atpons/m2proxy/pkg/request"
	"github.com/atpons/m2proxy/pkg/response"
	"github.com/atpons/m2proxy/pkg/storage"
	"github.com/atpons/m2proxy/pkg/util"
)

const (
	RespExtraLength = 4
)

type RequestGet struct {
	Key []byte
}

func (r *RequestGet) Print() {
	fmt.Printf("Key: %s\n", util.StringBytes(r.Key))
}

func Get(s storage.Storage, req request.Request) response.Response {
	getBody := ExtractGetBody(req)
	if util.Debug > 2 {
		getBody.Print()
	}

	value, err := s.Get(getBody.Key)

	if err == storage.ErrKeyNotFound {
		return *response.BuildResponse(req, req.Opcode, packet.StatusKeyNotFound, []byte{}, []byte("Not found"))
	}

	if err != nil {
		return *response.BuildResponse(req, req.Opcode, packet.StatusInternalError, []byte{}, []byte{})
	}

	respGet := ResponseGet{Flags: value.Flag, Value: value.Value}

	var res *response.Response
	if req.Opcode == packet.CmdGetQ {
		res = response.BuildResponse(req, req.Opcode, packet.StatusNoError, []byte{}, respGet.toBytes())
	} else {
		res = response.BuildResponse(req, req.Opcode, packet.StatusNoError, []byte{}, respGet.toBytes())
	}

	res.Cas = value.CAS
	res.ExtraLength = RespExtraLength
	return *res
}

func ExtractGetBody(r request.Request) RequestGet {
	return RequestGet{
		Key: r.Body[0:r.TotalBodyLength],
	}
}

type ResponseGet struct {
	Flags uint32
	Value []byte
}

func (r *ResponseGet) toBytes() []byte {
	resp := make([]byte, 4)
	binary.BigEndian.PutUint32(resp[0:], r.Flags)
	resp = append(resp, r.Value...)
	return resp
}
