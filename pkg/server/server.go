package server

import (
	"net"
	"time"

	"github.com/atpons/m2proxy/pkg/handler"
	"github.com/atpons/m2proxy/pkg/packet"
	"github.com/atpons/m2proxy/pkg/request"
	"github.com/atpons/m2proxy/pkg/response"
	"github.com/atpons/m2proxy/pkg/storage"
	"github.com/atpons/m2proxy/pkg/util"
)

type Server struct {
	Storage       *storage.Storage
	listenAddress *net.TCPAddr
}

func NewServer(lisAddr string, st *storage.Storage) *Server {
	addr, err := net.ResolveTCPAddr("tcp", lisAddr)
	if err != nil {
		panic(err)
	}

	server := &Server{Storage: st, listenAddress: addr}
	return server
}

func (s *Server) ListenAndServe() {
	listener, err := net.ListenTCP("tcp", s.listenAddress)
	if err != nil {
		panic(err)
	}
	_ = s.handleListener(listener)
}

func (s *Server) handleListener(l *net.TCPListener) error {
	for {
		conn, err := l.AcceptTCP()
		if err != nil {
			return err
		}
		go s.handleRequest(conn)
	}
}

func (s *Server) handleRequest(conn net.Conn) {
	defer conn.Close()
	for {
		conn.SetReadDeadline(time.Now().Add(10 * time.Second))

		buf := make([]byte, 1024)

		_, err := conn.Read(buf)

		if err != nil {
			return
		}

		req, _ := request.ParseRequest(buf)

		if util.Debug > 0 {
			req.Print()
		}

		res := response.Response{}
		switch req.Opcode {
		case packet.CmdVersion:
			res = handler.Version(*req)
		case packet.CmdSet, packet.CmdSetQ, packet.CmdReplace, packet.CmdReplaceQ, packet.CmdAdd, packet.CmdAddQ:
			res = handler.Set(*s.Storage, *req)
		case packet.CmdGet, packet.CmdGetQ, packet.CmdGetK, packet.CmdGetKQ:
			res = handler.Get(*s.Storage, *req)
		case packet.CmdDelete, packet.CmdDeleteQ:
			res = handler.Delete(*s.Storage, *req)
		case packet.CmdIncrement, packet.CmdIncrementQ, packet.CmdDecrement, packet.CmdDecrementQ:
			res = handler.Calc(*s.Storage, *req)
		case packet.CmdNoop:
			res = handler.Noop(*req)
		case packet.CmdQuit:
			res = handler.Quit(*req)
		case packet.CmdQuitQ:
		default:
			res = *response.BuildResponse(*req, req.Opcode, packet.StatusUnknownCommand, []byte{}, []byte{})
		}

		cmd := packet.Cmd(req.Opcode)

		if cmd.Quietly() || (cmd.Quietly() && res.Status == packet.StatusKeyNotFound) {
		} else {
			_, err = conn.Write(res.ToBytes())
		}

		if (req.Opcode == packet.CmdQuit) || (req.Opcode == packet.CmdQuitQ) {
			return
		}
	}
}
