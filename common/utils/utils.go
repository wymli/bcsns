package utils

import (
	"context"
	"fmt"
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

func ExtractUserIdFromCtx(ctx context.Context) (uint64, error) {
	userId, ok := ctx.Value("UserId").(string)
	if !ok {
		return 0, fmt.Errorf("failed to get userId from ctx as string, get:%v", ctx.Value("UserId"))
	}

	res, err := strconv.ParseUint(userId, 10, 64)
	if err != nil {
		return 0, fmt.Errorf("failed to convert userId of string to uint64, get: %v", userId)
	}

	return res, nil
}
