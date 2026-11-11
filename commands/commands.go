package commands

import (
	"errors"
	"strings"

	"github.com/westside-jpg/Rediz/storage"
	"github.com/westside-jpg/Rediz/types"
)

func CommandResolve(request types.Request, storage *storage.Storage) (types.Response, error) {
	request.Command = strings.ToLower(request.Command)
	command := request.Command

	switch command {
	case "ping":
		response, err := commandPING(request)
		if err != nil {
			return response, err
		}
		return response, nil
	case "echo":
		response, err := commandECHO(request)
		if err != nil {
			return response, err
		}
		return response, nil
	case "get":
		response, err := commandGET(request, storage)
		if err != nil {
			return response, err
		}
		return response, nil
	case "set":
		response, err := commandSET(request, storage)
		if err != nil {
			return response, err
		}
		return response, nil
	}

	return types.NewUnknown(), errors.New("Unknown command")
}

func commandPING(request types.Request) (types.Response, error) {
	if len(request.Data) > 1 {
		return types.NewError("Command PING expects 0 or 1 arguments"), errors.New("Command PING expects 0 or 1 arguments")
	} else if len(request.Data) == 1 {
		return types.NewBulkString(request.Data[0]), nil
	}

	return types.NewSimpleString("PONG"), nil
}

func commandECHO(request types.Request) (types.Response, error) {
	if len(request.Data) != 1 {
		return types.NewError("Command ECHO expects 1 arguments"), errors.New("Command ECHO expects 1 arguments")
	}

	return types.NewBulkString(request.Data[0]), nil
}

func commandGET(request types.Request, storage *storage.Storage) (types.Response, error) {
	if len(request.Data) != 1 {
		return types.NewError("(GET) invalid request data"), errors.New("(GET) invalid request data")
	}

	data := storage.Get(request.Data[0])
	return types.NewBulkString(data), nil
}

func commandSET(request types.Request, storage *storage.Storage) (types.Response, error) {
	if len(request.Data) != 2 {
		return types.NewError("(SET) invalid request data"), errors.New("(SET) invalid request data")
	}

	storage.Set(request.Data[0], request.Data[1])
	return types.NewSimpleString("OK"), nil
}
