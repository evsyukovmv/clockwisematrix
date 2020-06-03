package clockwisematrix

// Flatten - flattens square matrix in clockwise order
func Flatten(m [][]int) ([]int, error) {
	vm := &vectorMatrix{matrix: m}

	if err := vm.initSquareMatrix(); err != nil {
		return []int{}, err
	}

	for i := 0; len(vm.vector) < cap(vm.vector); i++ {
		vm.clockwiseStep(i)
	}
	return vm.vector, nil
}
