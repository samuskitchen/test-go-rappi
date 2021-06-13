package main

import (
    "fmt"
    "math"
)

func main() {
    ////Test Cases
    //values := []int32{1, 2, 3, 4, 5} iterations := int32(14)
    //values := []int32{2, 2, 3, 1} iterations := int32(2)
    //values := []int32{2, 3} iterations := int32(1)
    values := number()
    iterations := int32(500)
    fmt.Println(minSum(values, iterations))


}

func number() []int32 {

    result := make([]int32, 0)
    for i := 0; i < 100; i++ {
        result = append(result, 1)
    }

    return result
}


func minSum(num []int32, k int32) int32 {

    limit := k

    if k > int32(len(num)){
        limit = int32(len(num))
    }

    calculateValue := int32(0)
    for i := int32(0); i < limit; i++ {

        result := make([]int32, len(num))
        copy(result, num)

        result[i] = 0
        value := calculate(num[i]) + sum(result)

        if calculateValue == 0 || calculateValue > value{
            calculateValue = value
        }
    }

    return calculateValue
}

func calculate(number int32) int32 {
    return int32(math.Ceil(float64(number)/2))
}

func sum(num []int32) int32 {

    result := int32(0)
    for _, nm := range num  {
        result += nm
    }

    return result
}