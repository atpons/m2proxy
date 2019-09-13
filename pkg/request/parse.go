package request

import (
	"encoding/binary"
	"fmt"
	"os"

	"github.com/atpons/m2proxy/pkg/util"

	"github.com/atpons/m2proxy/pkg/packet"
)

type Request struct {
	Magic           byte
	Opcode          packet.Cmd
	KeyLength       uint16
	ExtraLength     byte
	DataType        byte
	VbucketId       uint16
	TotalBodyLength uint32
	Opaque          uint32
	Cas             uint32
	Body            []byte
}

func ParseRequest(req []byte) (*Request, error) {
	request := Request{}
	request.Magic = req[0]
	request.Opcode = packet.Cmd(req[1])
	request.KeyLength = binary.BigEndian.Uint16(req[2:4])
	request.ExtraLength = req[4]
	request.DataType = req[5]
	request.VbucketId = binary.BigEndian.Uint16(req[6:8])
	request.TotalBodyLength = binary.BigEndian.Uint32(req[8:12])
	request.Opaque = binary.BigEndian.Uint32(req[12:16])
	request.Cas = binary.BigEndian.Uint32(req[16:24])
	request.Body = req[24:]
	return &request, nil
}

func (r *Request) Print() {
	fmt.Fprintf(os.Stderr, "Magic: %v, Opcode: %s, Keylength: %v, ExtraLength: %v, DataType: %v, VbucketId: %v, TotalBodyLength: %v, Opaque: %v, Cas: %v\n",
		r.Magic,
		r.Opcode.String(),
		r.KeyLength,
		r.ExtraLength,
		r.DataType,
		r.VbucketId,
		r.TotalBodyLength,
		r.Opaque,
		r.Cas,
	)

	if util.Debug > 1 {
		fmt.Fprintf(os.Stderr, "Body: %v\n",
			r.Body,
		)
	}
}
