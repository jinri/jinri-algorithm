/*
 * @Author: jinri
 * @Date: 2022-01-14 18:43:43
 * @Description: RangeList Imp File
 * @Email: chunfengbishu@hotmail.com
 */

package main

import (
    "sort"
    "fmt"
)

type Range struct {
	Left, Right int
}

type RangeList struct {
	SortedRanges []Range
}

/**
 * @description:  Add Range[] is in the list
 * @param {[2]int} arr
 * @return {*}
 */

//Time complexity  O(n)
//Space complexity O(n)
func (this *RangeList) Add(arr [2]int) {

	var left = arr[0]
	var right = arr[1]
	if left > right{
		return
	}
	if this.SortedRanges == nil{
		this.SortedRanges = make([]Range,0)
	} 
	i := 0
	newRange := []Range{}
	for i < len(this.SortedRanges) {
		if this.SortedRanges[i].Right < left {
			//temp sorted [left,right] list
			newRange = append(newRange, this.SortedRanges[i])
			//position need add
			i++
		} else {
			// use min left merge
			left = min(this.SortedRanges[i].Left, left)
			break
		}
	}
	
	//init merge position
	j := i

	for j < len(this.SortedRanges) {
		if this.SortedRanges[j].Left > right {
			// left > right need not deal
			break
		} else {
			// use max right merge
			right = max(right, this.SortedRanges[j].Right)
			j++
		}
	}
	//before i list + new i
	newRange = append(newRange, Range{left, right})
	//after i list
	for i := j; i < len(this.SortedRanges); i++ {
		newRange = append(newRange, this.SortedRanges[i])
	}
	this.SortedRanges = newRange
}

/**
 * @description: Query Range[] is in the list
 * @param {[2]int} arr
 * @return bool
 */

 //Time complexity  O(log2 n)
func (this *RangeList) Query(arr [2]int) bool {
	var left = arr[0]
	var right = arr[1]
	
    //binary-search, return index
	index := sort.Search(len(this.SortedRanges), func(i int) bool {
		return this.SortedRanges[i].Left > left
	})
	if index == 0 {
		return false
	}

	target := this.SortedRanges[index-1]
	if target.Left <= left && target.Right >= right {
		return true
	} else {
		return false
	}
}

/**
 * @description: print odered range 
 * @param {}
 * @return {}
 */
func (this *RangeList) Print()  {
	for i := 0; i < len(this.SortedRanges); i++ {
		fmt.Printf("[%d, %d) ",this.SortedRanges[i].Left,this.SortedRanges[i].Right)
	}
	fmt.Printf("\n")
}

/**
 * @description: Remove Range[] is in the list
 * @param {[2]int} arr
 * @return {*}
 */

//Time complexity  O(n)
//Space complexity O(n)
func (this *RangeList) Remove(arr [2]int) {

	var left = arr[0]
	var right = arr[1]
	if left > right{
		return
	}
	
	newRange := []Range{}
	i := 0
	for i < len(this.SortedRanges) {
		if this.SortedRanges[i].Right < left {
			newRange = append(newRange, this.SortedRanges[i])
			//position need delete
			i++
		} else {
			break
		}
	}
	j := i
	for j < len(this.SortedRanges) {
		if this.SortedRanges[j].Left > right {
			break
		} else {
			//j-i,need delete list index 
			j++
		}
	}
	for p := i; p < j; p++ {
		curRange := this.SortedRanges[p]
		if left <= curRange.Left && right >= curRange.Right {
			continue
		}
		if curRange.Left < left {
			//use left left merge
			newRange = append(newRange, Range{curRange.Left, left})
			if curRange.Right > right {
				// split two
				newRange = append(newRange, Range{right, curRange.Right})
			}
		}else {
			//use right right merge
			newRange = append(newRange, Range{right, curRange.Right})
		}
	}
	for i := j; i < len(this.SortedRanges); i++ {
		newRange = append(newRange, this.SortedRanges[i])
	}
	this.SortedRanges = newRange
}

func min(x, y int) int {
	if x < y {
		return x
	}
	return y
}
func max(x, y int) int {
	if x > y {
		return x
	}
	return y
}

func main(){
	rl := RangeList{}
	
	rl.Add([2]int{1, 5})
    rl.Print()
	// Should display: [1, 5)
	rl.Add([2]int{10, 20})
	rl.Print()
	// Should display: [1, 5) [10, 20)
	rl.Add([2]int{20, 20})
	rl.Print()
	// Should display: [1, 5) [10, 20)
	rl.Add([2]int{20, 21})
	rl.Print()
	// Should display: [1, 5) [10, 21)
	rl.Add([2]int{2, 4})
	rl.Print()
	// Should display: [1, 5) [10, 21)
	rl.Add([2]int{3, 8})
	rl.Print()
	// Should display: [1, 8) [10, 21)
	rl.Remove([2]int{10, 10})
	rl.Print()
	// Should display: [1, 8) [10, 21)
	rl.Remove([2]int{10, 11})
	rl.Print()
	// Should display: [1, 8) [11, 21)
	rl.Remove([2]int{15, 17})
	rl.Print()
	// Should display: [1, 8) [11, 15) [17, 21)
	rl.Remove([2]int{3, 19})
	rl.Print()
	// Should display: [1, 3) [19, 21)
	
}


