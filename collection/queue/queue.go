package queue

type Queuer interface {
	Push(int)
	PushAll(...int)
	Pop() (int, error)
	// isEmpty() bool

	String() string
}
