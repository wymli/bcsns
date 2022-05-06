package resolver

import (
	"strings"

	"google.golang.org/grpc/resolver"
)

func init() {
	resolver.Register(&directResolverBuilder{})
}

const ENDPOINT_SEP = ","

type directResolverBuilder struct{}

func (b *directResolverBuilder) Build(target resolver.Target, cc resolver.ClientConn, opts resolver.BuildOptions) (resolver.Resolver, error) {
	urls := strings.Split(target.URL.Path, ENDPOINT_SEP)

	addrs := make([]resolver.Address, len(urls))
	for i, url := range urls {
		addrs[i] = resolver.Address{Addr: b.removeScheme(url)}
	}

	if err := cc.UpdateState(resolver.State{
		Addresses: addrs,
	}); err != nil {
		return nil, err
	}

	return &nopResolver{cc: cc}, nil
}

func (b *directResolverBuilder) Scheme() string {
	return "direct"
}

func (b *directResolverBuilder) removeScheme(url string) string {
	return strings.TrimLeft(url, b.Scheme()+":///")
}
