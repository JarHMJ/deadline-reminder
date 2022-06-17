package receiver

type Receiver interface {
	Notify(msg string) error
}
