package stack

type Stacker interface {
	Push(int)
	PushAll(...int)
	Pop() (int, error)
	// isEmpty() bool

	String() string
}
