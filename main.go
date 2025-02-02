package main

import (
	"github.com/gin-gonic/gin"
)

var numberList []int

func main() {
	r := gin.Default()
	r.POST("/add", handleInput) // API endpoint

	r.Run(":8081") // Running on 8080 port
}

func handleInput(c *gin.Context) {
	var input struct {
		Number int `json:"number"`
	}
	if err := c.ShouldBindJSON(&input); err != nil {
		c.JSON(400, gin.H{"error": "Invalid input"})
		return
	}

	numberList = processNumber(input.Number, numberList)
	c.JSON(200, gin.H{"updated_list": numberList})
}

func processNumber(num int, list []int) []int {
	if len(list) == 0 || (list[0] >= 0 && num >= 0) || (list[0] < 0 && num < 0) {
		list = append(list, num)
	} else {
		absNum := abs(num)
		for len(list) > 0 && absNum > 0 {
			if absNum >= abs(list[0]) {
				absNum -= abs(list[0])
				list = list[1:] 
			} else {
				list[0] += num
				break
			}
		}
	}
	return list
}

func abs(x int) int {
	if x < 0 {
		return -x
	}
	return x
}

