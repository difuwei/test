package main 
import(
"fmt"
"math"
"time"
// "runtime"
)


func main(){
fmt.Println("11")
fmt.Println(pow(3,2,10))
fmt.Println(math.Pow(2,3))
fmt.Println(time.Now().Hour())

a := 12
p := &a
fmt.Println(p)
fmt.Println(Vertex{2,4}.x)
m := []int{1,3,4,2,7}
fmt.Println(m[1:2])

b := make([]int,5)
fmt.Println(b)

var c []int
fmt.Println(c,cap(c))
c = append(c,1111,2222,3333)
fmt.Println(c)  
var pow = []int{1, 2, 4, 8, 16, 32, 64, 128}
for i, v := range pow {
		// fmt.Printf("2**%d = %d\n", i, v)
	fmt.Println(i,v)
	}  


}


func pow(x,n,lim float64) float64 {

	if v:= math.Pow(x,n);v<lim{
		return v
	}
	return lim
}

type Vertex struct{
	x int
	y int
}


