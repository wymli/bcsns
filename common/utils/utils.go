package utils

import (
	"context"
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
	userId, ok := ctx.Value("UserId").(uint64)
	return userId, ok
}

func CtxWithUserId(ctx context.Context, userId uint64) context.Context {
	return context.WithValue(ctx, "UserId", userId)
}
