package main

import (
    "fmt"
)

/*
Description

You are given an array prices where prices[i] is the price of a given stock on the ith day.

You want to maximize your profit by choosing a single day to buy one stock and choosing a different day in the future to sell that stock.

Return the maximum profit you can achieve from this transaction. If you cannot achieve any profit, return 0.

*/


func main(){
    fmt.Println("BUY AND SELL STOCK")
    prices := []int{7,1,5,3,6,4}
    fmt.Println(prices)
    fmt.Println(brute(prices))


    fmt.Println(minMax(prices))
}
func minMax(prices []int) int{
    // Sliding window
    return 0
}

func brute(prices []int) int{
    // Simple brute force solution,
    // Should be a better way
    var profit int

    for i:= 0; i < len(prices); i++ {
        for j:= i; j < len(prices); j++ {
            if prices[j] - prices[i] > profit {
                profit = prices[j] - prices[i]
            }
        }
    }

    return profit
}
