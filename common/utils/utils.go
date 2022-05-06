package utils

import (
	"context"
	"strconv"
)

func ExtractOneFromStringMap(m map[string]string) (string, string, bool) {
	if len(m) != 1 {
		return "", "", false
	}
	for k, v := range m {
		return k, v, true
	}
	return "", "", false
}

func ExtractUserIdFromCtx(ctx context.Context) (uint64, bool) {
	userId, ok := ctx.Value("x-user-id").(uint64)
	return userId, ok
}

func StringUid(a uint64) string {
	return "0x" + strconv.FormatUint(a, 16)
}

func UintUid(a string) uint64 {
	x, err := strconv.ParseUint(a, 0, 64)
	if err != nil {
		return 0
	}

	return x
}
