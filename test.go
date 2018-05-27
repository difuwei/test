package main
import "fmt"
func main(){

    var num = [...]int{1,2,3,4,5,6,6,7,8,8,4,42}
    slice := num[2:5:6]
    slice1 := append(slice,100,200,300)
    fmt.Println(slice1)

    mm :=map[int]string{1:"a",2:"b"}
    mm[4] = "ccc"
    fmt.Println(mm)
    e := mm[5]
    fmt.Println(e)

    var m string = "difuwe"
    fmt.Println(m)

    for i := 0; i < 5; i++ {
        fmt.Println(i);
    }
    ret := add(1,2222);
    fmt.Println(ret);

    p :=Person{"jodan",20,"nan"}
    fmt.Println(p.Name); 
}

    type Person struct {
        Name string
        Age int
        Sex string
    }

func add(num1 int,num2 int) int{

    return num1+num2;
}

