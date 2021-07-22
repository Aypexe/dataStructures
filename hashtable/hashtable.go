package hashtable

type HashTable interface {
	Insert(string)
	Lookup(string) int
	Delete(string)

	String() string
}
