# Package flattens square matrix in clockwise order

For example

```go
matrix := [][]int{
    {1, 2, 3},
    {4, 5, 6},
    {7, 8, 9},
}

clockwisematrix.Flatten(matrix) // => []int{1, 2, 3, 6, 9, 8, 7, 4, 5}
```

## Code Status

[![CircleCI](https://circleci.com/gh/evsyukovmv/clockwisematrix.svg?style=svg)](https://circleci.com/gh/evsyukovmv/clockwisematrix)
