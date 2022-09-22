package Test52


import (
	"fmt"
)

// each entry must be an address to a Company struct
var CompanyList = []Company



type Company struct {
	Cname string
	employees []employee
}

type employee struct {
	name string
	salary int
	position string
}


func Company1() {

	// creating employee Object
	Bob := employee{"Bob", 120000, "Back-end Developer"}
	Jake := employee{"Jake", 90000, "Front-end Developer"}
	Bill := employee{"Bill", 150000,"Cloud Engineer"}

	//creating the list of structs
	employees := []employee{Bob, Jake, Bill}

	// creating the company Name
	company := company{"DemonicLabs", employees}

	CompanyList = append(company)

}


func Company2() {


	Joe := employee{"Bob", 95000, "Jr Golang Developer"}
	Marry := employee{"Jake", 100000, "Database Engineer"}
	Kara := employee{"Bill", 180000,"Sr. Golang Developer"}

	employees := []employee{Joe, Marry, Kara}

	company := company{"AesirContructs", employees}


	CompanyList = append(company)

}

func main() {

	fmt.Println(CompanyList)
	fmt.Println("Test52: ")
	Company1()
	Company2()

	fmt.Println(CompanyList)

}