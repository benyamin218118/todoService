package contracts

type IPubSub interface {
	Publish(stream string, message map[string]any) error
}
