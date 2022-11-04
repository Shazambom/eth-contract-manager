package listener

type EventHandlerService interface {
	Handle(key, val string, err error) error
	InitService() error
}
