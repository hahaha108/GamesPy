package leetcode

//给定一个 n × n 的二维矩阵表示一个图像。
//
//将图像顺时针旋转 90 度。
//
//说明：
//
//你必须在原地旋转图像，这意味着你需要直接修改输入的二维矩阵。请不要使用另一个矩阵来旋转图像。
//
//来源：力扣（LeetCode）
//链接：https://leetcode-cn.com/problems/rotate-image
//著作权归领扣网络所有。商业转载请联系官方授权，非商业转载请注明出处。

//执行用时：0 ms
func Rotate(matrix [][]int) {
	x := len(matrix)
	for i := x - 1; i >= 0; i-- {
		for j := 0; j < x; j++ {
			matrix[j] = append(matrix[j], matrix[i][j])
		}
		matrix[i] = matrix[i][x:]
	}
}
