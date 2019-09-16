package handler

import (
	"encoding/binary"
	"fmt"
	"os"

	"github.com/atpons/m2proxy/pkg/util"

	"github.com/atpons/m2proxy/pkg/packet"
	"github.com/atpons/m2proxy/pkg/request"
	"github.com/atpons/m2proxy/pkg/response"
	"github.com/atpons/m2proxy/pkg/storage"
)

type RequestSet struct {
	Flags      uint32
	Expiration uint32
	Key        []byte
	Value      []byte
}

func (r *RequestSet) Print() {
	fmt.Fprintf(os.Stderr, "Flags: %v, Expiration: %v, Key: %s, Value: %s\n", r.Flags, r.Expiration, util.StringBytes(r.Key), util.StringBytes(r.Value))
}

func Set(s storage.Storage, req request.Request) response.Response {
	setBody := ExtractSetBody(req)

	if util.Debug > 2 {
		setBody.Print()
	}

	prev, err := s.Get(setBody.Key)

	if err == storage.ErrKeyNotFound && (req.Opcode == packet.CmdReplace || req.Opcode == packet.CmdReplaceQ) {
		return *response.BuildResponse(req, req.Opcode, packet.StatusKeyNotFound, []byte{}, []byte("Key Not Found."))
	}

	if prev != nil && (req.Opcode == packet.CmdAdd || req.Opcode == packet.CmdAddQ) {
		return *response.BuildResponse(req, req.Opcode, packet.StatusKeyExists, []byte{}, []byte("Data Exists for Key."))
	}

	rec := storage.Record{
		Key:   string(setBody.Key),
		Value: setBody.Value,
		Flag:  setBody.Flags,
		Exp:   setBody.Expiration,
		CAS:   req.Cas,
	}

	cas, err := s.Set(rec)

	if err != nil {
		return *response.BuildResponse(req, req.Opcode, packet.StatusInternalError, []byte{}, []byte{})
	}
	res := response.BuildResponse(req, req.Opcode, packet.StatusNoError, []byte{}, []byte{})
	res.Cas = cas

	return *res
}

func ExtractSetBody(r request.Request) RequestSet {
	return RequestSet{
		Flags:      binary.BigEndian.Uint32(r.Body[0:4]),
		Expiration: binary.BigEndian.Uint32(r.Body[4:8]),
		Key:        r.Body[8 : 8+r.KeyLength],
		Value:      r.Body[8+r.KeyLength : int(8+r.KeyLength)+(int(r.TotalBodyLength)-8-int(r.KeyLength))],
	}
}
