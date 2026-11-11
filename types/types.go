package types

type Request struct {
	Command string
	Data    []string
}

type ResponseKind int

const (
	KindUnknown      ResponseKind = iota
	KindSimpleString              // +OK
	KindError                     // -ERR
	KindInteger                   // :1
	KindBulkString                // $5 hello
	KindNull                      // $-1, ключа нет
)

type Response struct {
	Kind ResponseKind
	Str  string
	Int  int
}

func NewSimpleString(s string) Response {
	return Response{Kind: KindSimpleString, Str: s}
}

func NewError(msg string) Response {
	return Response{Kind: KindError, Str: msg}
}

func NewInteger(n int) Response {
	return Response{Kind: KindInteger, Int: n}
}

func NewBulkString(s string) Response {
	return Response{Kind: KindBulkString, Str: s}
}

func NewNull() Response {
	return Response{Kind: KindNull}
}

func NewUnknown() Response {
	return Response{Kind: KindUnknown}
}
