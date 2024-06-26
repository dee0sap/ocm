package vault

import (
	"sync"

	"github.com/mandelsoft/goutils/errors"

	"github.com/open-component-model/ocm/pkg/contexts/credentials/cpi"
	"github.com/open-component-model/ocm/pkg/contexts/datacontext"
)

const ATTR_REPOS = "github.com/open-component-model/ocm/pkg/contexts/credentials/repositories/vault"

type Repositories struct {
	lock  sync.Mutex
	repos map[cpi.ProviderIdentity]*Repository
}

func newRepositories(datacontext.Context) interface{} {
	return &Repositories{
		repos: map[cpi.ProviderIdentity]*Repository{},
	}
}

func (r *Repositories) GetRepository(ctx cpi.Context, spec *RepositorySpec) (*Repository, error) {
	var repo *Repository

	if spec.ServerURL == "" {
		return nil, errors.ErrInvalid("server url")
	}
	r.lock.Lock()
	defer r.lock.Unlock()

	var err error
	key := spec.GetKey()
	repo = r.repos[key]
	if repo == nil {
		repo, err = NewRepository(ctx, spec)
		if err == nil {
			r.repos[key] = repo
		}
	}
	return repo, err
}
