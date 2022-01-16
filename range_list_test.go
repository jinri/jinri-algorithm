/*
 * @Author: jinri
 * @Date: 2022-01-16 11:12:50
 * @Description: RangeList Imp Test File
 * @Email: chunfengbishu@hotmail.com
 */

package main

import(
    "testing"
	_"fmt"
	"math/rand"
	"time"
)

var rl = RangeList{}

func BenchmarkAdd(b *testing.B) {
	rand.Seed(time.Now().UnixNano())
	for i := 0; i < b.N; i++ {
		left := rand.Intn(b.N)
		right := rand.Intn(b.N)
		//fmt.Printf("b.n=%d,left=%d.right=%d\n",b.N,left,right)
		rl.Add([2]int{left, right})
	}
}

func BenchmarkRemove(b *testing.B) {

	rand.Seed(time.Now().UnixNano())
	for i := 0; i < b.N; i++ {
		left := rand.Intn(b.N)
		right := rand.Intn(b.N)
		//fmt.Printf("b.n=%d,left=%d.right=%d\n",b.N,left,right)
		rl.Remove([2]int{left, right})
	}
}

func BenchmarkQuery(b *testing.B) {
	rand.Seed(time.Now().UnixNano())
	for i := 0; i < b.N; i++ {
		left := rand.Intn(b.N)
		right := rand.Intn(b.N)
		//fmt.Printf("b.n=%d,left=%d.right=%d\n",b.N,left,right)
		rl.Query([2]int{left, right})
	}
}

func BenchmarkPrint(b *testing.B) {
	//no meaning
	/*
	rand.Seed(time.Now().UnixNano())
	for i := 0; i < b.N; i++ {
		rl.Print()
	}
	*/
}