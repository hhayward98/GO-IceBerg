package main

import (
	"bufio"
	"fmt"
	"os"
	"log"
	"strconv"
	"strings"

)

func Debugger(err error) {
	if err != nil{
		log.Fatal(err)
	}
}






/*


User Input
Title: is the Name of website
DomainName:  	Ex. http://localhost:8080/ 
NumTemplate: Generates n templates. adds makers for modifacation points
Styling: If True: 
	- Ask user for Primary and secondary color to use
	- accept Hex color Codes as well as CSS standard color strings.
Forms: If True: 
		- Ask user for Form Name
		- Ask User for HTML Page to assign the form to
		- Ask User for Form Parameters/Dimensions

Database: If True:
	- Ask user if Database Name
	- Ask user if they want to create a table based on Form Data if they have a Form
		- if True then Ask for table paramaters to be used with From that was generated
		- Ask user for table parameters
	- Ask user if 

*/




// check if the Database already exists before trying to create a newone
func InitDatabase(DBname string) {
	fmt.Println("Setting Up Database....")
	// create a database to use for web app 
	//full interface might be required

	fmt.Println("Database Creation complete!")

}

// ask if user needs any input forms
func AppendForms(FormName string, PageName string) {
	fmt.Println("Adding Forms....")
	// let the user specify form parameters (size, var_names, var_types, etc...) 


}



func AppendStyling(BK_color string, Text_color string, Second_color string) {
	fmt.Println("Adding Styling")
	// add styling to css file based on user input
}



// ask if user needs a database
func AppendDataBase(DBname string) {
	fmt.Println("Implementing Database....")
	// must create and setup database before trying to use in our program
	InitDatabase(DBname)
	// add code that connects to a database as part of a page function 
	// let the user choose what route/page will be modified

}
