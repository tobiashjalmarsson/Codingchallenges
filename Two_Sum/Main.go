package main

import (
    "fmt"
)


/*
Question Description:

Given an array of integers nums and an integer target, return indices of the two numbers such that they add up to target.

You may assume that each input would have exactly one solution, and you may not use the same element twice.

You can return the answer in any order.


*/




func main(){

    // Implement a hashmap with complement since lookup will run in constant time
    var target int = 9
    nums := []int{2,7,11,15}

    result := TwoSum2(nums, target)
    fmt.Println("First Resut: ",result)

    target = 6
    nums2 := []int{3,2,4}
    fmt.Println("Second result: ", TwoSum2(nums2, target))

    nums3 := []int{3,3}
    fmt.Println("Third result: ", TwoSum2(nums3, target))
}

func TwoSum2(nums []int, target int) []int {
    // A solution always exists
    // Runtime: n + n = 2n => O(n)
    result := []int{-1, -1}

    hash := make(map[int]int)
    for index,elem := range nums {
        complement := target-elem
        hash[complement] = index
    }

    // Now loop over the original slice and look for the element in the slice
    // if the element is in the slice it means that its complement exists and a solution is there
    for index, elem := range nums {
        // Check if complement exists
        // value : index of complement
        // index : value of value being checked for complement
        if value, ok := hash[elem]; ok {
            // We also have to make sure that they don't use the same value twice
            if value != index {
                result[0] = index
                result[1] = value
                return result
            }
        }
    }

    return result
}
