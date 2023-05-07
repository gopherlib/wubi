package wubi

// codeSlice 对codes进行排序，简码在前，全码在后
type codeSlice []string

func (x codeSlice) Len() int           { return len(x) }
func (x codeSlice) Less(i, j int) bool { return len(x[i]) < len(x[j]) }
func (x codeSlice) Swap(i, j int)      { x[i], x[j] = x[j], x[i] }
