package contracts

type IPubSub interface {
	Publish(stream string, message any) error
	Subscribe(stream string, handler func(data any)) error
}
