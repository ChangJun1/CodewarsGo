package Fundamental

/*Write a function that accepts a square matrix (N x N 2D array) and returns the determinant of the matrix.

How to take the determinant of a matrix -- it is simplest to start with the smallest cases:

A 1x1 matrix |a| has determinant a.

A 2x2 matrix [ [a, b], [c, d] ] or

|a  b|
|c  d|
has determinant: a*d - b*c.

The determinant of an n x n sized matrix is calculated by reducing the problem to the calculation of the determinants of n matrices ofn-1 x n-1 size.

For the 3x3 case, [ [a, b, c], [d, e, f], [g, h, i] ] or

|a b c|
|d e f|
|g h i|
the determinant is: a * det(a_minor) - b * det(b_minor) + c * det(c_minor) where det(a_minor) refers to taking the determinant of the 2x2 matrix created by crossing out the row and column in which the element a occurs:

|- - -|
|- e f|
|- h i|
Note the alternation of signs.

The determinant of larger matrices are calculated analogously, e.g. if M is a 4x4 matrix with first row [a, b, c, d], then:

det(M) = a * det(a_minor) - b * det(b_minor) + c * det(c_minor) - d * det(d_minor)*/

func Determinant(matrix [][]int) int {
	if len(matrix) == 0 {
		return 0
	}

	if len(matrix) == 1 {
		return matrix[0][0]
	}

	if len(matrix) == 2 {
		return matrix[0][0]*matrix[1][1] - matrix[0][1]*matrix[1][0]
	}

	res := 0
	sign := 1
	for j := 0; j < len(matrix); j++ {
		res += sign * matrix[0][j] * Determinant(minorMatrix(matrix, 0, j))
		sign = -sign
	}
	return res
}

func minorMatrix(A [][]int, i, j int) [][]int {
	n := len(A)
	minor := make([][]int, n-1)
	for k := range minor {
		minor[k] = make([]int, n-1)
	}
	for p := 0; p < n; p++ {
		if p == i {
			continue
		}
		for q := 0; q < n; q++ {
			if q == j {
				continue
			}
			x, y := p, q
			if p > i {
				x--
			}
			if q > j {
				y--
			}
			minor[x][y] = A[p][q]
		}
	}
	return minor
}

// the most popular
/*func Determinant(matrix [][]int) int {
	if len(matrix) == 1 {
		return matrix[0][0]
	}

	sliced := matrix[1:]
	res := 0

	for n := 0; n < len(matrix); n++ {
		arr := [][]int{}

		for _, v := range sliced {
			arr = append(arr, append(append([]int{}, v[:n]...), v[n+1:]...))
		}

		sign := -1

		if n%2 == 0 {
			sign = 1
		}

		res += sign * matrix[0][n] * Determinant(arr)
	}

	return res
}*/
