package response

import (
	"encoding/binary"

	"github.com/atpons/m2proxy/pkg/packet"
	"github.com/atpons/m2proxy/pkg/request"
)

type Response struct {
	Magic           byte
	Opcode          packet.Cmd
	KeyLength       uint16
	ExtraLength     byte
	DataType        byte
	Status          packet.Status
	TotalBodyLength uint32
	Opaque          uint32
	Cas             uint32
	Body            []byte
}

type Option func(Response)

func BuildResponse(req request.Request, opcode packet.Cmd, status packet.Status, key, value []byte) *Response {
	response := Response{}
	response.Magic = packet.Response
	response.Opcode = opcode
	response.KeyLength = uint16(len(key))
	response.ExtraLength = 0
	response.DataType = packet.DataTypeRaw
	response.Status = status
	response.TotalBodyLength = uint32(len(value))
	response.Opaque = req.Opaque
	response.Cas = 0
	response.Body = value
	return &response
}

func (r *Response) SetBody(b []byte) {
	r.Body = b
	r.TotalBodyLength = uint32(len(b))
}

func (r *Response) ToBytes() []byte {
	header := make([]byte, 24)
	header[0] = r.Magic
	header[1] = byte(r.Opcode)
	binary.BigEndian.PutUint16(header[2:4], r.KeyLength)
	header[4] = r.ExtraLength
	header[5] = r.DataType
	binary.BigEndian.PutUint16(header[6:8], uint16(r.Status))
	binary.BigEndian.PutUint32(header[8:12], r.TotalBodyLength)
	binary.BigEndian.PutUint32(header[12:16], r.Opaque)
	binary.BigEndian.PutUint32(header[16:20], r.Cas)
	header = append(header, r.Body...)
	return header
}
