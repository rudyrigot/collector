package collector

const (
	transformOp = iota
	filterOp    = iota
	eachOp      = iota
)

type operation struct {
	op  int
	lbd func(...string) string
}

type operationSet struct {
	alreadyPerformed bool
	initialArray     []string
	ops              []*operation
}

// Transform (which may called "map" in other languages) allows to apply a function on each
// element of an array, separately.
// The variadic parameter in the function should only expect one string.
func (opSet *operationSet) Transform(f func(...string) string) *operationSet {
	op := operation{op: transformOp, lbd: f}
	opSet.ops = append(opSet.ops, &op)
	return opSet
}

func (opSet *operationSet) Into(s *[]string) error {
	// Duplicating the slice
	buffer := make([]string, len(opSet.initialArray))
	copy(buffer, opSet.initialArray)

	// Performing the operations
	for i, elem := range buffer {
		for _, op := range opSet.ops {
			if op.op == transformOp {
				buffer[i] = op.lbd(elem)
			}
		}
	}

	// Returning the results
	s = &buffer
	return nil
}
