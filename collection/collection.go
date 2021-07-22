package collection

type Collectioner interface {
	Push(int)
	PushAll(...int)
	Pop() (int, error)
	// isEmpty() bool

	String() string
}
