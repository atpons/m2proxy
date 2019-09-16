package handler

import (
	"encoding/binary"
	"fmt"
	"os"
	"strconv"

	"github.com/atpons/m2proxy/pkg/util"

	"github.com/atpons/m2proxy/pkg/packet"
	"github.com/atpons/m2proxy/pkg/request"
	"github.com/atpons/m2proxy/pkg/response"
	"github.com/atpons/m2proxy/pkg/storage"
)

// will be replaced to work properly
func IncrDecr(s storage.Storage, req request.Request) response.Response {
	getBody := ExtractCalcBody(req)

	value, err := s.Get([]byte(string(getBody.Body)))

	if err == storage.ErrKeyNotFound {
		initialValue := make([]byte, 8)
		binary.BigEndian.PutUint64(initialValue, getBody.InitialValue)

		_, err := s.Set(
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

	existNum, err := strconv.Atoi(string(value.Value))
	if err != nil {
		return *response.BuildResponse(req, req.Opcode, packet.StatusIncrDecrOnNumericValue, []byte{}, []byte("Value: Non-numeric server-side value for incr or decr"))

	}

	if ((req.Opcode == packet.CmdDecrement || req.Opcode == packet.CmdDecrementQ) && (existNum-1 < 0)) || ((req.Opcode == packet.CmdIncrement || req.Opcode == packet.CmdIncrementQ) && (4294967295 < existNum+1)) {
		return *response.BuildResponse(req, req.Opcode, packet.StatusIncrDecrOnNumericValue, []byte{}, []byte("Value: Non-numeric server-side value for incr or decr"))
	}

	if util.Debug > 1 {
		fmt.Fprintf(os.Stderr, "incr/decr: Value Size: %d, Prev: %d, Incr=%d, Decr=%d\n", len(value.Value), existNum, existNum+1, existNum-1)
	}

	resp := make([]byte, 8)
	var v []byte
	if req.Opcode == packet.CmdIncrement || req.Opcode == packet.CmdIncrementQ {
		v = []byte(fmt.Sprint(uint32(existNum + 1)))
		binary.BigEndian.PutUint64(resp, uint64(existNum+1))
	}

	if req.Opcode == packet.CmdDecrement || req.Opcode == packet.CmdDecrementQ {
		v = []byte(fmt.Sprint(uint32(existNum - 1)))
		binary.BigEndian.PutUint64(resp, uint64(existNum-1))
	}

	value.Value = v

	_, err = s.Set(*value)

	if err != nil {
		return *response.BuildResponse(req, req.Opcode, packet.StatusInternalError, []byte{}, []byte{})
	}

	return *response.BuildResponse(req, req.Opcode, packet.StatusNoError, []byte{}, resp)

}

func ExtractCalcBody(r request.Request) RequestCalc {
	return RequestCalc{
		Amount:       binary.BigEndian.Uint64(r.Body[0:8]),
		InitialValue: binary.BigEndian.Uint64(r.Body[8:16]),
		Expiration:   binary.BigEndian.Uint32(r.Body[16:20]),
		Body:         r.Body[20:int(r.TotalBodyLength)],
	}
}

type RequestCalc struct {
	Amount       uint64
	InitialValue uint64
	Expiration   uint32
	Body         []byte
}
