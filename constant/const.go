package constant

const (
	ServiceName          = `foo`
	Host                 = `localhost`
	GrpcPort             = `9082`
	HttpPort             = `8082`
	CircuitBreakerTimout = 1000 * 30 // change the second operand, this means 30 second timeout
	ZipkinHostPort       = `:0`
)
