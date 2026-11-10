package parser

import (
	"bufio"
	"errors"
	"strconv"
	"strings"

	"github.com/westside-jpg/Rediz/types"
)

func RequestParser(reader *bufio.Reader) (types.Request, error) {

	firstByte, err := reader.ReadByte()
	if err != nil {
		return types.Request{}, err
	} else if firstByte != '*' {
		return types.Request{}, errors.New("invalid request")
	}

	var req types.Request

	elementsCountRaw, err := reader.ReadString('\n')
	elementsCount, err := strconv.Atoi(strings.TrimSuffix(elementsCountRaw, "\r\n"))

	for i := range elementsCount {

		tokenLengthRaw, err := reader.ReadString('\n')
		if err != nil {
			return types.Request{}, err
		}

		tokenLengthRaw = strings.TrimSuffix(tokenLengthRaw, "\r\n")
		tokenLengthRaw = strings.TrimPrefix(tokenLengthRaw, "$")
		tokenLength, err := strconv.Atoi(tokenLengthRaw)
		if err != nil {
			return types.Request{}, err
		}

		token, err := reader.ReadString('\n')
		if err != nil {
			return types.Request{}, err
		}

		token = strings.TrimSuffix(token, "\r\n")

		if len(token) != tokenLength {
			return types.Request{}, errors.New("invalid request")
		}

		if i == 0 {
			req.Command = token
		} else {
			req.Data = append(req.Data, token)
		}

	}

	return req, nil
}
