package matrix

/*
	matrix[row][col] => matrix[col][n-1-row]=> matrix[n-1-row][n-1-col] => matrix[n-1-col][row] => matrix[row][col]
	From the observation, we can know the essence of rotation.

	In addition, how do we figure out which point need to process like above?
	Because we have to change the position of four points, the times of rotate is 'n^2/4'
	- For odd number, (n^2 - 1)/4 = (n-1)/2 * (n+1)/2
	- For even number, we just need to iterate n^2/4 = n/2 * n/2
*/

func rotate(matrix [][]int) {
	n := len(matrix)
	for i := 0; i < n/2; i++ {
		for j := 0; j < (n+1)/2; j++ {
			matrix[i][j], matrix[j][n-1-i], matrix[n-1-i][n-1-j], matrix[n-1-j][i] = matrix[n-1-j][i], matrix[i][j], matrix[j][n-1-i], matrix[n-1-i][n-1-j]
		}
	}
}
