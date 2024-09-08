package errorx

const (
	Success = 1000
	Error   = 2000

	ParseRequsetError     = 3000
	GenerateIdError       = 3001
	InsertToDbError       = 3002
	OriginalUrlDuplicated = 3003
	ShortUrlDuplicated    = 3004
	ShortUrlNotFound      = 3005
)

var MsgFlags = map[uint]string{
	Success: "Operation success",
	Error:   "Operation failed",

	ParseRequsetError:     "Failed to parse request, invalid params",
	GenerateIdError:       "Generate uuid failed",
	InsertToDbError:       "Failed to insert data to database",
	OriginalUrlDuplicated: "Original url duplicated",
	ShortUrlDuplicated:    "Short url duplicated",
	ShortUrlNotFound:      "Short url not found",
}

func GetMsg(code uint) string {
	msg, ok := MsgFlags[code]
	if ok {
		return msg
	}
	return MsgFlags[Error]
}
