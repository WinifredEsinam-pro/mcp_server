package main
import ("fmt")

func main() {
  var student string = "Winifred"
var student1 = 23;
var x int
var y string
var z bool
var i, j string= "Hello", "World"
var arr = [5]int{2,4,6,7,8}
mySlice := []int{}
mySlice1 := []string{"Go", "Slices", "Are", "Powerful"}
a := 334
  fmt.Println(student)
  fmt.Println(student1)
  fmt.Println(x)
  fmt.Println(y)
  fmt.Println(z)
  fmt.Println(a)
  fmt.Print(i, "\n", j, "\n")
  fmt.Printf("%b\n", student1)
  fmt.Println(len(arr))
  fmt.Println(mySlice)
  fmt.Println(len(mySlice))
  fmt.Println(cap(mySlice))
  fmt.Println(mySlice1)
  fmt.Println(len(mySlice1))
  fmt.Println(cap(mySlice1))

  myFunction()
  mySecondFunction("Winifred")
  fmt.Println(myNextFunction(2, 5))


  var per1 struct1
  var per2 struct1

  //for person 1
  per1.name = "Ama"
  per1.age = 19
  per1.job = "Student"
  per1.salary = 3000000

  //for person 2
  per2.name = "Kofi"
  per2.age = 22
  per2.job = "Backend developer"
  per2.salary = 4000000

  //for person 1
  fmt.Println("Name:", per1.name)
  fmt.Println("Name:", per1.age)
  fmt.Println("Name:", per1.job)
  fmt.Println("Name:", per1.salary)
   
  //for person 2
    fmt.Println("Name:", per2.name)
      fmt.Println("Name:", per2.age)
        fmt.Println("Name:", per2.job)
          fmt.Println("Name:", per2.salary)
}

func myFunction(){
fmt.Println("This is my first created function")
}

func mySecondFunction(fname string){
fmt.Println("Hello", fname, "Esinam!")
}

func myNextFunction(x int, y int) (result int){
  result = x + y
  return result
}

type struct1 struct{
  name string
  age int
  job string
  salary int
}
