package clockwisematrix

// Flatten - flattens square matrix in clockwise order
func Flatten(m [][]int) ([]int, error) {
	vm, err := newVectorMatrix(m)
	if err != nil {
		return []int{}, err
	}

	for i := 0; len(vm.vector) < cap(vm.vector); i++ {
		vm.clockwiseStep(i)
	}
	return vm.vector, nil
}

func newVectorMatrix(m [][]int) (*vectorMatrix, error) {
	vm := vectorMatrix{matrix: m}
	err := vm.initSquareMatrix()
	return &vm, err
}
