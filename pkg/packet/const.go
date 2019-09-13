package packet

const (
	Request  = 0x80
	Response = 0x81

	StatusNoError                    = 0x0000
	StatusKeyNotFound                = 0x0001
	StatusKeyExists                  = 0x0002
	StatusValueTooLarge              = 0x0003
	StatusInvalidArguments           = 0x0004
	StatusItemNotStored              = 0x0005
	StatusIncrDecrOnNumericValue     = 0x0006
	StatusTheVbucketBelongsToAnother = 0x0007
	StatusAuthenticationError        = 0x0008
	StatusAuthenticationContinue     = 0x0009
	StatusUnknownCommand             = 0x0081
	StatusOutOfMemory                = 0x0082
	StatusNotSupported               = 0x0083
	StatusInternalError              = 0x0084
	StatusBusy                       = 0x0085
	StatusTemporaryFailure           = 0x0086

	CmdGet                = 0x00
	CmdSet                = 0x01
	CmdAdd                = 0x02
	CmdReplace            = 0x03
	CmdDelete             = 0x04
	CmdIncrement          = 0x05
	CmdDecrement          = 0x06
	CmdQuit               = 0x07
	CmdFlush              = 0x08
	CmdGetQ               = 0x09
	CmdNoop               = 0x0a
	CmdVersion            = 0x0b
	CmdGetK               = 0x0c
	CmdGetKQ              = 0x0d
	CmdAppend             = 0x0e
	CmdPrepend            = 0x0f
	CmdStat               = 0x10
	CmdSetQ               = 0x11
	CmdAddQ               = 0x12
	CmdReplaceQ           = 0x13
	CmdDeleteQ            = 0x14
	CmdIncrementQ         = 0x15
	CmdDecrementQ         = 0x16
	CmdQuitQ              = 0x17
	CmdFlushQ             = 0x18
	CmdAppendQ            = 0x19
	CmdVerbosity          = 0x1b
	CmdTouch              = 0x1c
	CmdGAT                = 0x1d
	CmdGATQ               = 0x1e
	CmdSASLListMechs      = 0x20
	CmdSASLAuth           = 0x21
	CmdSASLStep           = 0x22
	CmdRGet               = 0x30
	CmdRSet               = 0x31
	CmdRSetQ              = 0x32
	CmdRAppend            = 0x33
	CmdRAppendQ           = 0x34
	CmdRPrepend           = 0x35
	CmdRPrependQ          = 0x36
	CmdRDelete            = 0x37
	CmdRDeleteQ           = 0x38
	CmdRIncr              = 0x39
	CmdRIncrQ             = 0x3a
	CmdRDecr              = 0x3b
	CmdRDecrQ             = 0x3c
	CmdSetVBucket         = 0x3d
	CmdGetVBucket         = 0x3e
	CmdDelVBucket         = 0x3f
	CmdTAPConnect         = 0x40
	CmdTAPMutation        = 0x41
	CmdTAPDelete          = 0x42
	CmdTAPFlush           = 0x43
	CmdTAPOpaque          = 0x44
	CmdTAPVBucketSet      = 0x45
	CmdTAPCheckpointStart = 0x46
	CmdTAPCheckpointEnd   = 0x47

	DataTypeRaw = 0x00
)

var (
	StatusString = map[uint16]string{
		0x0000: "NoError",
		0x0001: "KeyNotFound",
		0x0002: "KeyExists",
		0x0003: "ValueTooLarge",
		0x0004: "InvalidArguments",
		0x0005: "ItemNotStored",
		0x0006: "IncrDecrOnNumericValue",
		0x0007: "TheVbucketBelongsToAnother",
		0x0008: "AuthenticationError",
		0x0009: "AuthenticationContinue",
		0x0081: "UnknownCommand",
		0x0082: "OutOfMemory",
		0x0083: "NotSupported",
		0x0084: "InternalError",
		0x0085: "Busy",
		0x0086: "TemporaryFailure",
	}

	CmdString = map[byte]string{
		0x00: "Get",
		0x01: "Set",
		0x02: "Add",
		0x03: "Replace",
		0x04: "Delete",
		0x05: "Increment",
		0x06: "Decrement",
		0x07: "Quit",
		0x08: "Flush",
		0x09: "GetQ",
		0x0a: "Noop",
		0x0b: "Version",
		0x0c: "GetK",
		0x0d: "GetKQ",
		0x0e: "Append",
		0x0f: "Prepend",
		0x10: "Stat",
		0x11: "SetQ",
		0x12: "AddQ",
		0x13: "ReplaceQ",
		0x14: "DeleteQ",
		0x15: "IncrementQ",
		0x16: "DecrementQ",
		0x17: "QuitQ",
		0x18: "FlushQ",
		0x19: "AppendQ",
		0x1b: "Verbosity",
		0x1c: "Touch",
		0x1d: "GAT",
		0x1e: "GATQ",
		0x20: "SASLListMechs",
		0x21: "SASLAuth",
		0x22: "SASLStep",
		0x30: "RGet",
		0x31: "RSet",
		0x32: "RSetQ",
		0x33: "RAppend",
		0x34: "RAppendQ",
		0x35: "RPrepend",
		0x36: "RPrependQ",
		0x37: "RDelete",
		0x38: "RDeleteQ",
		0x39: "RIncr",
		0x3a: "RIncrQ",
		0x3b: "RDecr",
		0x3c: "RDecrQ",
		0x3d: "SetVBucket",
		0x3e: "GetVBucket",
		0x3f: "DelVBucket",
		0x40: "TAPConnect",
		0x41: "TAPMutation",
		0x42: "TAPDelete",
		0x43: "TAPFlush",
		0x44: "TAPOpaque",
		0x45: "TAPVBucketSet",
		0x46: "TAPCheckpointStart",
		0x47: "TAPCheckpointEnd",
	}
)
