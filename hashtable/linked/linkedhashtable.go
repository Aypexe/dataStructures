package hashtable

const LENGTH int = 10

type HashTable struct {
	buckets [LENGTH]*bucket
}

type bucket struct {
	head *bucketNode
}

type bucketNode struct {
	next  *bucketNode
	key   string
	value string
}

func New() *HashTable {
	h := &HashTable{}
	for i := range h.buckets {
		h.buckets[i] = &bucket{}
	}
	return h
}

func hash(s string) int {
	sum := 0
	for _, c := range s {
		sum += int(c)
	}

	return sum % LENGTH
}

//TODO: add Insert, Lookup, and Delete
