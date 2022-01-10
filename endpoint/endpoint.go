package endpoint

import (
	"github.com/go-kit/kit/endpoint"
	kitoc "github.com/go-kit/kit/tracing/opencensus"
	"github.com/nightsilvertech/foo/constant"
	_interface "github.com/nightsilvertech/foo/service/interface"
	"github.com/nightsilvertech/utl/middlewares"
)

type FooEndpoint struct {
	AddFooEndpoint       endpoint.Endpoint
	EditFooEndpoint      endpoint.Endpoint
	DeleteFooEndpoint    endpoint.Endpoint
	GetAllFooEndpoint    endpoint.Endpoint
	GetDetailFooEndpoint endpoint.Endpoint
}

func NewFooEndpoint(svc _interface.FooService) FooEndpoint {
	const (
		username    = `mr-dummy`
		phoneNumber = `081299019285`
		secret      = `295817ALKXQM`
	)

	var addFooEp endpoint.Endpoint
	{
		const name = `AddFoo`
		addFooEp = makeAddFooEndpoint(svc)
		addFooEp = middlewares.CircuitBreakerMiddleware(constant.ServiceName)(addFooEp)
		addFooEp = kitoc.TraceEndpoint(name)(addFooEp)
		addFooEp = middlewares.JwtTestMiddleware(username, phoneNumber, secret)(addFooEp)
	}

	var editFooEp endpoint.Endpoint
	{
		const name = `EditFoo`
		editFooEp = makeEditFooEndpoint(svc)
		editFooEp = middlewares.CircuitBreakerMiddleware(constant.ServiceName)(editFooEp)
		editFooEp = kitoc.TraceEndpoint(name)(editFooEp)
		editFooEp = middlewares.JwtTestMiddleware(username, phoneNumber, secret)(editFooEp)
	}

	var deleteFooEp endpoint.Endpoint
	{
		const name = `DeleteFoo`
		deleteFooEp = makeDeleteFooEndpoint(svc)
		deleteFooEp = middlewares.CircuitBreakerMiddleware(constant.ServiceName)(deleteFooEp)
		deleteFooEp = kitoc.TraceEndpoint(name)(deleteFooEp)
		deleteFooEp = middlewares.JwtTestMiddleware(username, phoneNumber, secret)(deleteFooEp)
	}

	var getAllFooEp endpoint.Endpoint
	{
		const name = `GetAllFoo`
		getAllFooEp = makeGetAllFooEndpoint(svc)
		getAllFooEp = middlewares.CircuitBreakerMiddleware(constant.ServiceName)(getAllFooEp)
		getAllFooEp = kitoc.TraceEndpoint(name)(getAllFooEp)
		getAllFooEp = middlewares.JwtTestMiddleware(username, phoneNumber, secret)(getAllFooEp)
	}

	var getDetailFooEp endpoint.Endpoint
	{
		const name = `GetDetailFoo`
		getDetailFooEp = makeGetDetailFooEndpoint(svc)
		getDetailFooEp = middlewares.CircuitBreakerMiddleware(constant.ServiceName)(getDetailFooEp)
		getDetailFooEp = kitoc.TraceEndpoint(name)(getDetailFooEp)
		getDetailFooEp = middlewares.JwtTestMiddleware(username, phoneNumber, secret)(getDetailFooEp)
	}

	return FooEndpoint{
		AddFooEndpoint:       addFooEp,
		EditFooEndpoint:      editFooEp,
		DeleteFooEndpoint:    deleteFooEp,
		GetAllFooEndpoint:    getAllFooEp,
		GetDetailFooEndpoint: getDetailFooEp,
	}
}
