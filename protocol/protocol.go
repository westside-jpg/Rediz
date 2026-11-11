package protocol

import (
	"strconv"

	"github.com/westside-jpg/Rediz/types"
)

func Encode(response types.Response) []byte {
	switch response.Kind {
	case types.KindSimpleString:
		return []byte("+" + response.Str + "\r\n")
	case types.KindError:
		return []byte("-" + response.Str + "\r\n")
	case types.KindInteger:
		return []byte(":" + strconv.FormatInt(int64(response.Int), 10) + "\r\n")
	case types.KindBulkString:
		return []byte("$" + strconv.Itoa(len(response.Str)) + "\r\n" + response.Str + "\r\n")
	case types.KindNull:
		return []byte("$-1\r\n")
	case types.KindUnknown:
		return []byte("-ERR unknown command\r\n")
	default:
		return []byte("-ERR internal error\r\n")
	}
}
