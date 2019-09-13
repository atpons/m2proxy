package handler

import (
	"encoding/binary"

	"github.com/atpons/m2proxy/pkg/packet"
	"github.com/atpons/m2proxy/pkg/request"
	"github.com/atpons/m2proxy/pkg/response"
	"github.com/atpons/m2proxy/pkg/storage"
)

func Calc(s storage.Storage, req request.Request) response.Response {
	getBody := ExtractCalcBody(req)

	value, err := s.Get(getBody.Body)
	if err == storage.ErrKeyNotFound && getBody.Expiration == 0xffffffff {
		initialValue := make([]byte, 8)
		binary.LittleEndian.PutUint64(initialValue, getBody.InitialValue)

		err := s.Set(
			storage.Record{
				Key:   string(getBody.Body),
				Value: initialValue,
				Exp:   getBody.Expiration,
				CAS:   req.Cas,
			},
		)

		if err != nil {
			return *response.BuildResponse(req, req.Opcode, packet.StatusInternalError, []byte{}, []byte{})
		}

		return *response.BuildResponse(req, req.Opcode, packet.StatusNoError, []byte{}, initialValue)
	}

	if err != nil {
		return *response.BuildResponse(req, req.Opcode, packet.StatusInternalError, []byte{}, []byte{})
	}

	new := make([]byte, 8)
	if req.Opcode == packet.CmdIncrement || req.Opcode == packet.CmdIncrementQ {
		binary.BigEndian.PutUint64(new, binary.BigEndian.Uint64(value.Value)+1)
	}

	if req.Opcode == packet.CmdDecrement || req.Opcode == packet.CmdDecrementQ {
		binary.BigEndian.PutUint64(new, binary.BigEndian.Uint64(value.Value)-1)
	}

	value.Value = new

	err = s.Set(*value)

	if err != nil {
		return *response.BuildResponse(req, req.Opcode, packet.StatusInternalError, []byte{}, []byte{})
	}

	return *response.BuildResponse(req, req.Opcode, packet.StatusNoError, []byte{}, value.Value)

}

func ExtractCalcBody(r request.Request) RequestCalc {
	return RequestCalc{
		Amount:       binary.BigEndian.Uint64(r.Body[0:8]),
		InitialValue: binary.BigEndian.Uint64(r.Body[8:16]),
		Expiration:   binary.BigEndian.Uint32(r.Body[16:20]),
		Body:         r.Body[20 : 20+int(r.TotalBodyLength)-int(r.ExtraLength)-1],
	}
}

type RequestCalc struct {
	Amount       uint64
	InitialValue uint64
	Expiration   uint32
	Body         []byte
}
