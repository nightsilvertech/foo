package endpoint

import (
	"github.com/go-kit/kit/endpoint"
	"github.com/nightsilvertech/foo/constant"
	"github.com/nightsilvertech/foo/middleware"
	_interface "github.com/nightsilvertech/foo/service/interface"
)

type FooEndpoint struct {
	AddFooEndpoint       endpoint.Endpoint
	EditFooEndpoint      endpoint.Endpoint
	DeleteFooEndpoint    endpoint.Endpoint
	GetAllFooEndpoint    endpoint.Endpoint
	GetDetailFooEndpoint endpoint.Endpoint
}

func NewFooEndpoint(svc _interface.FooService) FooEndpoint {
	var addFooEp endpoint.Endpoint
	{
		const name = `AddFoo`
		addFooEp = makeAddFooEndpoint(svc)
		addFooEp = middleware.CircuitBreakerMiddleware(constant.ServiceName)(addFooEp)
	}

	var editFooEp endpoint.Endpoint
	{
		const name = `EditFoo`
		editFooEp = makeEditFooEndpoint(svc)
		editFooEp = middleware.CircuitBreakerMiddleware(constant.ServiceName)(editFooEp)
	}

	var deleteFooEp endpoint.Endpoint
	{
		const name = `DeleteFoo`
		deleteFooEp = makeDeleteFooEndpoint(svc)
		deleteFooEp = middleware.CircuitBreakerMiddleware(constant.ServiceName)(deleteFooEp)
	}

	var getAllFooEp endpoint.Endpoint
	{
		const name = `GetAllFoo`
		getAllFooEp = makeGetAllFooEndpoint(svc)
		getAllFooEp = middleware.CircuitBreakerMiddleware(constant.ServiceName)(getAllFooEp)
	}

	var getDetailFooEp endpoint.Endpoint
	{
		const name = `GetDetailFoo`
		getDetailFooEp = makeGetDetailFooEndpoint(svc)
		getDetailFooEp = middleware.CircuitBreakerMiddleware(constant.ServiceName)(getDetailFooEp)
	}

	return FooEndpoint{
		AddFooEndpoint:       addFooEp,
		EditFooEndpoint:      editFooEp,
		DeleteFooEndpoint:    deleteFooEp,
		GetAllFooEndpoint:    getAllFooEp,
		GetDetailFooEndpoint: getDetailFooEp,
	}
}
