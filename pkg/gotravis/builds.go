package gotravis

import (
	"context"
	"net/http"

	"github.com/shuheiktgw/go-travis"
)

type Builds interface {
	ListByRepoSlug(ctx context.Context, repoSlug string, opt *travis.BuildsByRepoOption) ([]*travis.Build, *http.Response, error)
}
