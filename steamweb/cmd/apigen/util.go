package main

import (
	"strconv"
	"strings"

	"github.com/13k/valve.go/steamlib"
)

func stringJoinInt(ii []int, sep string) string {
	ss := make([]string, len(ii))

	for i, n := range ii {
		ss[i] = strconv.FormatInt(int64(n), 10)
	}

	return strings.Join(ss, sep)
}

func stringJoinUint32(ii []uint32, sep string) string {
	ss := make([]string, len(ii))

	for i, n := range ii {
		ss[i] = strconv.FormatUint(uint64(n), 10)
	}

	return strings.Join(ss, sep)
}

func stringJoinAppIDs(ii []steamlib.AppID, sep string) string {
	uints := make([]uint32, len(ii))

	for i, appID := range ii {
		uints[i] = uint32(appID)
	}

	return stringJoinUint32(uints, sep)
}
