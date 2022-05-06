package logic

import (
	"context"
	"encoding/json"
	"fmt"

	"github.com/dgraph-io/dgo/v210/protos/api"
	"github.com/pkg/errors"
	pbauth "github.com/wymli/bcsns/app/auth_rpc/pb"
	"github.com/wymli/bcsns/app/user_center/internal/model"
	"github.com/wymli/bcsns/app/user_center/internal/svc"
	"github.com/wymli/bcsns/app/user_center/pb"
	"go.opentelemetry.io/otel/attribute"
	"go.opentelemetry.io/otel/trace"

	"github.com/wymli/bcsns/common/errx"
	"github.com/wymli/bcsns/common/logx"
	"github.com/wymli/bcsns/common/tracer"
	"github.com/wymli/bcsns/common/utils"
)

type RegisterUserLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewRegisterUserLogic(ctx context.Context, svcCtx *svc.ServiceContext) *RegisterUserLogic {
	return &RegisterUserLogic{
		Logger: logx.WithTraceCtx(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *RegisterUserLogic) RegisterUser(in *pb.RegisterUserReq) (*pb.RegisterUserResp, error) {
	user := model.User{
		Phone:     int(in.Phone),
		Nickname:  in.Nickname,
		Sex:       int(in.Sex),
		Age:       int(in.Age),
		Avatar:    in.Avatar,
		Address:   in.Address,
		PublicKey: in.PublicKey,
		Password:  in.Password,
	}

	userJson, err := json.Marshal(user)
	if err != nil {
		return nil, errors.Wrapf(errx.ERROR_SERVER_COMMON, "failed to marshall json, err:%v", err)
	}

	qFormat := `
  query {
      user as var(func: eq(phone, "%d"))
  }`
	query := fmt.Sprintf(qFormat, in.Phone)

	mu := &api.Mutation{
		Cond: `@if(eq(len(user), 0))`,
		// SetNquads: []byte(`uid(user) <email> "correct_email@dgraph.io" .`),
		SetJson: userJson,
	}

	apiReq := &api.Request{
		Query:     query,
		Mutations: []*api.Mutation{mu},
		CommitNow: true,
	}

	// span start
	_, span := tracer.Tracer().Start(l.ctx, "dgraph-mutate", trace.WithAttributes(attribute.String("user", fmt.Sprintf("%#v", user))))

	res, err := l.svcCtx.DgraphClient.NewTxn().Do(context.Background(), apiReq)
	if err != nil {
		// span end
		span.End()
		return nil, errx.Wrapf(errx.ERROR_DB, "failed to insert user in dgraph, err:%v", err)
	}

	// span end
	span.End()

	if len(res.Uids) == 0 {
		return nil, errx.Wrapf(errx.ERROR_USER_DUPLICATE, "duplicate register of user phone: %v", in.Phone)
	}

	_, v, ok := utils.ExtractOneFromStringMap(res.Uids)
	if !ok {
		return nil, errx.Wrapf(errx.ERROR_SERVER_COMMON, "failed to extract uid from dgraph response, want 1 uid, get %d uid", len(res.Uids))
	}

	uid := utils.UintUid(v)

	authRsp, err := l.svcCtx.AuthClient.GenerateToken(l.ctx, &pbauth.GenerateTokenReq{
		UserId: uid,
	})
	if err != nil {
		return nil, errx.Wrapf(err, "failed to generate token with userid:%d", uid)
	}

	return &pb.RegisterUserResp{
		UserId:       uid,
		AccessToken:  authRsp.AccessToken,
		AccessExpire: authRsp.AccessExpire,
		RefreshAfter: authRsp.RefreshAfter,
	}, nil
}
