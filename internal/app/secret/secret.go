package secret

import (
	"github.com/labstack/echo"

	"github.com/trueone/beetest/internal/app/secret/presenter"
	"github.com/trueone/beetest/internal/app/secret/presenter/http/json"
	"github.com/trueone/beetest/internal/app/secret/repository"
	"github.com/trueone/beetest/internal/app/secret/usecase"
)

type registry struct {
	repository repository.Repository
	create     usecase.Create
	get        usecase.Get
}

// Registry constructor
func New() presenter.Registry {
	return &registry{
		repository: repository.NewRepositoryMap(),
	}
}

// Register json handlers
func (r *registry) RegisterJsonHandlers(e *echo.Echo) {
	json.NewHandler(e, r)
}

// Create usecase getter
func (r *registry) Create() usecase.Create {
	if r.create == nil {
		r.create = usecase.NewCreate(r.repository)
	}

	return r.create
}

// Get usecase getter
func (r *registry) Get() usecase.Get {
	if r.get == nil {
		r.get = usecase.NewGet(r.repository)
	}

	return r.get
}
