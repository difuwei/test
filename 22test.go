package main
import(
"fmt"
"math"
// "runtime"
"time"
)
func main(){
	// fmt.Println("hello,world")
	sum:=1
	for sum < 200 {
		sum +=sum
	}
	fmt.Println(sum)
	fmt.Println(sqrt(2),sqrt(-8))
	fmt.Println(1,2,3)
	fmt.Println(pow(3,2,10))

	fmt.Println(time.Now().Weekday()+1)
	fmt.Println(time.Saturday)

	// defer fmt.Println("world")
	// fmt.Println("hello")

	// for i:= 0;i<10;i++{
	// 	defer fmt.Println(i)
	// }


	// i,j := 42,27112
	// p := &i
	// fmt.Println(*p)
	// fmt.Println(p)
	// *p = 10
	// fmt.Println(i)
	// p = &j
	// *p = *p / 2
	// fmt.Println(j)

	// v :=struc{101,2}
	// fmt.Println(v.x)

	v := struc{1, 2}
	p := &v
	p.x = 1e9
	fmt.Println(v)

	 a:= [3]int{133,2}
	fmt.Println(a[0])

	p2 := []int{12,75,2,1,4,4,2,5}
	fmt.Println(p2[1:3])

	aa := make([]int ,5)
	fmt.Println(aa)

}

	type struc struct {
		x int
		y int
	}

func sqrt(x float64) string{
	if x <0 {
		return sqrt(-x) + "i"
	}

	return fmt.Sprint(math.Sqrt(x))
}

func pow(x,n,lim float64) float64{
	if v := math.Pow(x,n);v < lim {
		return v
	}
	return lim
}

// func testSwitch(x int){
// 	switch x;x {
// 	case "1":
// 		fmt.Println("hello1")
// 	case "2":
// 		fmt.Println("hello2")
// 	default：
// 	fmt.Println("hello3")
// 	}
// }
