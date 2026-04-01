package main

import(
	"fmt"
)

func main(){
	var students []StudentInfo
	var choice int

	for {
		fmt.Println("\n___Student Manager___")
		fmt.Println("1. Add Student")
		fmt.Println("2. View student")
		fmt.Println("3. Exit")
		fmt.Println("Select a choice: ")

		fmt.Scanln(&choice)

		switch choice{
		case 1:
			addStudent(students)
		case 2:
			viewStudent(students)
		case 3:
			fmt.Println("Goodbye!")
			return
		default:
			fmt.Println("Invalid choice")		
		}
	}



	

}
func addStudent(students []StudentInfo) []StudentInfo{
		var name string
		var age int
		var contact string
		var school string

		fmt.Println("Enter your name: ")
		fmt.Scanln(&name)

		fmt.Println("Enter your age: ")
		fmt.Scanln(&age)

		fmt.Println("Enter your contact number: ")
		fmt.Scanln(&contact)

		fmt.Println("Enter the name of your school: ")
		fmt.Scanln(&school)


		newStudent := StudentInfo{
			Name: name,
			Age: age,
			Contact: contact,
			School: school,
		}

		students = append(students, newStudent)
		fmt.Println("Student added successfully")

		return students
}

func viewStudent(students []StudentInfo){
	if len(students)== 0{
		fmt.Println("No student found")
		return
	}
fmt.Println("\n Student list:")
for i, student := range students{
	fmt.Println("%d. %s (%d years old)\n", i+1, student.Name, student.Age)
}

}

type StudentInfo struct{
		Name string
		Age int
		Contact string
		School string
	}
