package handler

import (
	"fmt"
	"os"

	"github.com/atpons/m2proxy/pkg/packet"
	"github.com/atpons/m2proxy/pkg/request"
	"github.com/atpons/m2proxy/pkg/response"
	"github.com/atpons/m2proxy/pkg/storage"
	"github.com/atpons/m2proxy/pkg/util"
)

func Flush(s storage.Storage, req request.Request) response.Response {
	if util.Debug > 1 {
		fmt.Fprintf(os.Stderr, "Flush is to ALL delete\n")
	}
	_ = s.Flush()
	return *response.BuildResponse(req, req.Opcode, packet.StatusNoError, []byte{}, []byte{})
}
