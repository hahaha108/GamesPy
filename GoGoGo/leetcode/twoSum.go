package leetcode

//给定一个整数数组 nums 和一个目标值 target，请你在该数组中找出和为目标值的那 两个 整数，并返回他们的数组下标。
//
//你可以假设每种输入只会对应一个答案。但是，数组中同一个元素不能使用两遍。
//
//来源：力扣（LeetCode）
//链接：https://leetcode-cn.com/problems/two-sum
//著作权归领扣网络所有。商业转载请联系官方授权，非商业转载请注明出处。


//运行时间 48 ms
func TwoSum(nums []int, target int) []int {
	for i:=0;i< len(nums);i++ {
		for j:=i+1;j<len(nums);j++{
			if nums[i] + nums[j] == target {
				return []int{i,j}
			}
		}
	}
	return nil
}