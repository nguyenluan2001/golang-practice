// You can edit this code!
// Click here and start typing.
package main

import "fmt"

type TrackingItem struct {
	index int
	value int
}

func nextGreaterElements(nums []int) []int {
	result := []int{}
	stack := []TrackingItem{}
	for i := 0; i < len(nums); i++ {
		result = append(result, -1)
	}
	for i := 0; i < len(nums); i++ {
		num := nums[i]
		if len(stack) == 0 {
			stack = append(stack, TrackingItem{
				index: i,
				value: num,
			})
		} else {
			for num > stack[len(stack)-1].value {
				// fmt.Println("num", num)
				// fmt.Println("stack", stack)
				lastStack := stack[len(stack)-1]
				result[lastStack.index] = num
				stack = stack[0 : len(stack)-1]
				if len(stack) == 0 {
					break
				}
			}
			stack = append(stack, TrackingItem{
				index: i,
				value: num,
			})
		}
	}
	if len(stack) != 0 {
		// for i := len(stack) - 1; i >= 0; i-- {
		// 	item := stack[i]
		// 	if stack[0].value > item.value {
		// 		result[item.index] = stack[0].value
		// 	} else {
		// 		result[item.index] = -1
		// 	}
		// }
		count := 0
		for i := len(stack) - 1; i >= 0; i-- {
			for count <= stack[i].index {
				if nums[count] > stack[i].value {
					result[stack[i].index] = nums[count]
					break
				} else {
					count++
				}
			}
			if count == i {
				result[stack[i].index] = -1
			}
			count = 0
		}
	}
	return result
}
func main() {
	// nums := []int{1, 5, 3, 4, 3}
	// nums := []int{1, 2, 1}
	// nums := []int{1, 2, 3, 4, 3}
	// nums := []int{1, 5, 3, 6, 8}
	nums := []int{1, 2, 3, 2, 1}
	result := nextGreaterElements(nums)
	fmt.Println(result)
}
