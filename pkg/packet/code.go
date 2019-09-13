package packet

type Status uint16

type Cmd uint8

func (c *Cmd) String() string {
	return CmdString[byte(*c)]
}

func (c *Cmd) Quietly() bool {
	switch *c {
	case CmdGetQ:
	case CmdGetKQ:
	case CmdSetQ:
	case CmdAddQ:
	case CmdReplaceQ:
	case CmdDeleteQ:
	case CmdIncrementQ:
	case CmdDecrementQ:
	case CmdQuitQ:
	case CmdFlushQ:
	case CmdAppendQ:
	case CmdGATQ:
	case CmdRSetQ:
	case CmdRAppendQ:
	case CmdRPrependQ:
	case CmdRDeleteQ:
	case CmdRIncrQ:
	case CmdRDecrQ:
	default:
		return false
	}
	return true
}

func (s *Status) String() string {
	return StatusString[uint16(*s)]
}
