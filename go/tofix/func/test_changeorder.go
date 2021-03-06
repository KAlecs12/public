package main

import (
	"strconv"

	"./correct"
	"github.com/01-edu/public/go/lib"
)

type stuNode = NodeAddL
type solNode = correct.NodeAddL

func stuPushFront(node *stuNode, num int) *stuNode {
	tmp := &stuNode{Num: num}
	tmp.Next = node
	return tmp
}

func stuNumToList(num int) *stuNode {
	var res *stuNode
	for num > 0 {
		res = stuPushFront(res, num%10)
		num /= 10
	}
	return res
}

func stuListToNum(node *stuNode) int {
	var n int

	for tmp := node; tmp != nil; tmp = tmp.Next {
		n = n*10 + tmp.Num
	}
	return n
}

func solPushFront(node *solNode, num int) *solNode {
	tmp := &solNode{Num: num}
	tmp.Next = node
	return tmp
}

func printSol(node *solNode) string {
	var result string
	for tmp := node; tmp != nil; tmp = tmp.Next {
		result += strconv.Itoa(tmp.Num)
		if tmp.Next != nil {
			result += "->"
		}
	}
	return result
}

func solNumToList(num int) *solNode {
	var res *solNode
	for num > 0 {
		res = solPushFront(res, num%10)
		num /= 10
	}
	return res
}

func solListToNum(node *solNode) int {
	var n int

	for tmp := node; tmp != nil; tmp = tmp.Next {
		n = n*10 + tmp.Num
	}
	return n
}

func compareNodes(stuResult *stuNode, solResult *solNode, num1 int) {
	solList := printSol(solNumToList(num1))
	if stuResult == nil && solResult == nil {
		return
	}
	if stuResult != nil && solResult == nil {
		stuNum := stuListToNum(stuResult)
		lib.Fatalf("\nChangeorder(%s) == %v instead of %v\n\n",
			solList, stuNum, "")
	}
	if stuResult == nil && solResult != nil {
		solNum := solListToNum(solResult)
		lib.Fatalf("\nChangeorder(%s) == %v instead of %v\n\n",
			solList, "", solNum)
	}
	stuNum := stuListToNum(stuResult)
	solNum := solListToNum(solResult)
	if stuNum != solNum {
		lib.Fatalf("\nChangeorder(%s) == %v instead of %v\n\n",
			solList, stuNum, solNum)
	}
}

func main() {
	table := []int{1234567}
	table = append(table, lib.MultRandIntBetween(0, 1000000000)...)

	for _, arg := range table {
		stuResult := Changeorder(stuNumToList(arg))
		solResult := correct.Changeorder(solNumToList(arg))

		compareNodes(stuResult, solResult, arg)
	}
}
