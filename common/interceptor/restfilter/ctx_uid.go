package restfilter

import (
	"context"
	"net/http"
	"strconv"
)

func ExtractUserIdToCtx(next http.HandlerFunc) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		userId, _ := strconv.ParseInt(r.Header.Get("X-UserId"), 10, 64)
		ctx := r.Context()
		ctx = context.WithValue(ctx, "UserId", userId)

		next(w, r.WithContext(ctx))
	}
}
