package metadata

import (
	"context"

	"github.com/wymli/bcsns/common/utils"
	"google.golang.org/grpc/metadata"
)

const XUSERID = "x-user-id"

func ExtractUserIdFromGRPC(ctx context.Context) uint64 {
	md, ok := metadata.FromIncomingContext(ctx)
	if !ok {
		return 0
	}

	return utils.UintUid(md[XUSERID][0])
}

func CtxWithUserId(ctx context.Context, userId uint64) context.Context {
	return context.WithValue(ctx, XUSERID, userId)
}

func CtxGetUserId(ctx context.Context) uint64 {
	return ctx.Value(XUSERID).(uint64)
}
