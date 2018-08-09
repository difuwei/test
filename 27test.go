package main
import("fmt")

func main(){
	const world = "世界"
	fmt.Println(world)
    fmt.Println("hello,world!")
    sum := 0
    for i:=1;i<10;i++{
    	sum +=i
    }
    fmt.Println(sum)
}
