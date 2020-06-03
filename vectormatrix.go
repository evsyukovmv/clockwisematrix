package clockwisematrix

import "fmt"

type vectorMatrix struct {
	matrix [][]int
	vector []int

	maxIndex    int
	middleIndex int
}

func (vm *vectorMatrix) initSquareMatrix() error {
	matrixLen := len(vm.matrix)
	for i := range vm.matrix {
		if matrixLen != len(vm.matrix[i]) {
			return fmt.Errorf("not valid, square matrix required")
		}
	}
	vm.maxIndex = len(vm.matrix) - 1
	vm.middleIndex = vm.maxIndex / 2
	vm.vector = make([]int, 0, len(vm.matrix)*len(vm.matrix))
	return nil
}

func (vm *vectorMatrix) clockwiseStep(step int) {
	if step == vm.middleIndex && vm.maxIndex%2 == 0 {
		vm.vector = append(vm.vector, vm.matrix[vm.middleIndex][vm.middleIndex])
		return
	}

	vm.directWalk(step)
	vm.reverseWalk(step)
}

func (vm *vectorMatrix) directWalk(step int) {
	vm.horizontalWalk(step, step, vm.maxIndex-step-1, true)
	vm.verticalWalk(vm.maxIndex-step, step, vm.maxIndex-step-1, true)
}

func (vm *vectorMatrix) reverseWalk(step int) {
	vm.horizontalWalk(vm.maxIndex-step, step+1, vm.maxIndex-step, false)
	vm.verticalWalk(step, step+1, vm.maxIndex-step, false)
}

func (vm *vectorMatrix) horizontalWalk(row int, from int, to int, direct bool) {
	if direct {
		for i := from; i <= to; i++ {
			vm.vector = append(vm.vector, vm.matrix[row][i])
		}
	} else {
		for i := to; i >= from; i-- {
			vm.vector = append(vm.vector, vm.matrix[row][i])
		}
	}
}

func (vm *vectorMatrix) verticalWalk(column int, from int, to int, direct bool) {
	if direct {
		for i := from; i <= to; i++ {
			vm.vector = append(vm.vector, vm.matrix[i][column])
		}
	} else {
		for i := to; i >= from; i-- {
			vm.vector = append(vm.vector, vm.matrix[i][column])
		}
	}
}
