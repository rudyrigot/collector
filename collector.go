package collector

// CString is the entry point of collector.
// Call it on your array, and then you can do things like C(arr).map....
func CString(array []string) *operationSet {
	return &operationSet{initialArray: array}
}
