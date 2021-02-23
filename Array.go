package main
import "fmt"


func main(){
//var array[10]string
//array :=[...]int{50,60,70,88}

/*array[0]="my Name Azizul hoq"
array[1]="I Liked"
array[2]="Programming"
array[3]="Specially"
array[4]="Golang"
array[5]="C"
array[6]="etc"
*/
//fmt.Println(array)


//fmt.Println(array,len(array))

//slice
//p:=[12]string{"Azizul Hoq","Very Lezy","in This ","Futur Plan"}



//var Q []string=p[0:4]

//fmt.Println(Q,len(Q))
//var orenge []float32

//orenge = append(orenge,30,36,40)
//fmt.Println(orenge,len(orenge))
//fmt.Printf("%T",orenge)//type

//a := reflect.TypeOf(orenge).Kind().String()
//a := reflect.TypeOf(orenge).Kind().String(array)
//fmt.Println(a, orenge,len(orenge))

//map

z :=make(map[int]int)
x :=make(map[string]string)
z[2]=2020
z[3]=2021

x["Name"]="Azizul"
x["Height"]="Height: 5.8"
x["From"]="From: Bangladesh"
x["since"]="Year"


//fmt.Println(z[2],z[3])
fmt.Println(x["Name"],x["Height"],x["From"],x["since"],z[2],z[3])
//fmt.Println(z,x)




}