package collector

const (
	mapOp    = iota
	selectOp = iota
	eachOp   = iota
)

type operation struct {
	operation int
	lbd       func(...string) string
}

type operationsOrError struct {
	alreadyPerformed bool
	initialArray     []interface{}
	ops              []operation
	err              error
}
