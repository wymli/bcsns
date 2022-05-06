package resolver

import (
	"google.golang.org/grpc/resolver"
)

type kubeResolverBuilder struct{}

func (b *kubeResolverBuilder) Build(target resolver.Target, cc resolver.ClientConn, opts resolver.BuildOptions) (resolver.Resolver, error) {
	// todo

	return &nopResolver{cc: cc}, nil
}

func (b *kubeResolverBuilder) Scheme() string {
	return "k8s"
}
