package RankUp

/*Your task, is to create a NxN spiral with a given size.

For example, spiral with size 5 should look like this:

00000
....0
000.0
0...0
00000
and with the size 10:

0000000000
.........0
00000000.0
0......0.0
0.0000.0.0
0.0..0.0.0
0.0....0.0
0.000000.0
0........0
0000000000
Return value should contain array of arrays, of 0 and 1, with the first row being composed of 1s. For example for given size 5 result should be:

[[1,1,1,1,1],[0,0,0,0,1],[1,1,1,0,1],[1,0,0,0,1],[1,1,1,1,1]]
Because of the edge-cases for tiny spirals, the size will be at least 5.

General rule-of-a-thumb is, that the snake made with '1' cannot touch to itself.*/

func Spiralize(size int) [][]int {
	res := make([][]int, size)
	for i := 0; i < size; i++ {
		res[i] = make([]int, size)
	}
	start, end, v := 0, size-1, 1
	for start < end {
		// right
		for j := start; j < end; j++ {
			res[start][j] = v
		}

		// down
		for i := start; i < end; i++ {
			res[i][end] = v
		}

		// left
		for j := end; j > start; j-- {
			res[end][j] = v
		}

		// up
		for i := end; i > start+1; i-- {
			res[i][start] = v
		}

		v = 1 - v
		if start+1 < end {
			res[start+1][start] = v
		}

		start++
		end--

	}

	if size&1 == 1 {
		res[size/2][size/2] = v
	}

	return res
}
