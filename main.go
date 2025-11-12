package main

import "fmt"

/* temporary testing until the rework */
func main() {
	todos := Todos{}
	todos.Add("This is ID 1")
	todos.Add("This is ID 2")
	todos.Add("Should be ID 3")
	fmt.Printf("%+v\n\n", todos)

	fmt.Println("Deleting ID1")
	todos.Delete(1)
	fmt.Printf("%+v\n\n", todos)

	todos.Delete(2)
	fmt.Println("Now Deleting ID2")
	fmt.Printf("%+v\n\n", todos)

	todos.Toggle(3)
	fmt.Println("Toggling ID3")
	fmt.Printf("%+v\n", todos)

	fmt.Println("Adding the last task.")
	todos.Add("should be ID 4")
	fmt.Printf("%+v", todos)
}
