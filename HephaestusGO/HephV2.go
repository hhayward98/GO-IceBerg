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
ConfigSetting User Input Def

Title: is the Name of website
DomainName:  	Ex. http://localhost:8080/ 
NumTemplate: Generates n templates. adds makers for modifacation points
Styling: If True: 
	- Ask user for Primary and secondary color to use
	- accept Hex color Codes as well as CSS standard color strings.
*/

type ConfigSetting struct {
	Title string
	DomainName string
	NumTemplate int
	Styling Styles

}

type Styles struct {
	Pc string
	Sc string
}




func GoFRAME_L2(PrgSet struct){
	fmt.Println("Creating GOFRAME....")
	
	// makes main.go
	InitAppHead(PrgSet)

	// makes Dockerfile 
	InitDocker()


	// makes n amount of html files depending on userinput 
	InitTemplates(PrgSet)


	// makes css file
	InitCSS(PrgSet)


	fmt.Println("GoFRAME Complete!!!")

}



func InitAppHead(PrgSet struct) {
	fmt.Println("Creating webapp File....")

	// TODO: integrate PrgSet.Title into main as needed
	Tem := "package main\n\nimport (\n\t'fmt'\n\t'log'\n\t'net/http'\n\t'html/template'\n)\n\n\nvar tpl *template.Template\n\nfunc Home(w http.ResponseWriter, r *http.Request) {\n\n\tfmt.Println('Home')\n\n\ttpl.ExecuteTemplate(w, 'index.html', '')\n\n\treturn\n\n}\n\n\n//here\n\n\nfunc AppRouts() {\n\n\thttp.Handle('/static/', http.StripPrefix('/static/', http.FileServer(http.Dir('static'))))\n\thttp.HandleFunc('/', Home)\n\n\tlog.Fatal(http.ListenAndServe(':8080', nil))\n\n}\n\n\nfunc main() {\n\n\ttpl, _ = template.ParseGlob('./static/templates/*html')\n\n\tlog.Print('Listening....')\n\tAppRouts()\n\n}\n\n"
	New := strings.ReplaceAll(Tem, "'", `"`,)
	f, err := os.Create("./main.go")
	_, err2 := f.WriteString(New)
	Debugger(err)
	Debugger(err2)

	defer f.Close()
	fmt.Println("Done")

}

func InitDocker(){
	fmt.Println("Creating Docker File....")

	Tem := "FROM golang:1.18\n\nRUN mkdir /GoWeb\n\nADD . /GoWeb\n\nWORKDIR /GoWeb\n\nCOPY go.* ./\n\nRUN go mod download && go mod verify\n\nRUN go build -o app .\n\nEXPOSE 8080\n\nCMD ['/GoWeb/app']"
	New := strings.ReplaceAll(Tem, "'", `"`,)
	f, err := os.Create("./Dockerfile")
	_, err2 := f.WriteString(New)
	Debugger(err)
	Debugger(err2)

	defer f.Close()
	fmt.Println("Done")
}


func InitTemplates(PrgSet struct) {

	Num := PrgSet.NumTemplate

	fmt.Println("Creating Templates....")
	if err := os.MkdirAll("./static/templates/", os.ModePerm); err != nil {
		log.Fatal(err)
	}
	// if the user only needs 1 page, program makes just the index page
	if Num == 1 {
		WriteHTML("index.html", PrgSet)
	}else if Num > 1 {
		WriteHTML("index.html", PrgSet)
		// makes n amount of html files
		for i := 1; i < Num; i++ {
			strBuf := "Page"
			s1 := strconv.Itoa(i)
			strBuf += s1
			strBuf += ".html"
			fmt.Println(strBuf)
			// writes code to html files
			WriteHTML(strBuf, PrgSet)
		}
	}


	// opens app.go and adds routes to the html files that were created
	// takes Num so the number of routes is == to the number of html files created
	AppendRoutsMain(Num, PrgSet)

	fmt.Println("Done")
}


func WriteHTML(Fname string, PrgSet struct) {

	// integrate PrgSet as needed 
	DomName := PrgSet.DomainName

	Tem := "<!DOCTYPE html>\n<html>\n<head>\n\t<meta charset='utf-8'>\n\t<meta name='viewport' content='width=device-width, initial-scale=1'>\n\t<meta http-equiv='X-UA-Compatible' content='ie=edge' />\n\t<title>Home</title>\n</head>\n<body>\n\t<h3>{{.}}</h3>\n\t<p>Information</p>\n</body>\n</html>"
	New := strings.ReplaceAll(Tem, "'", `"`,)

	f, err := os.Create("./static/templates/"+ Fname)
	_, err2 := f.WriteString(New)
	Debugger(err)
	Debugger(err2)

	defer f.Close()

}


func AppendRoutsMain(Num int, PrgSet struct) {


	// TODO: Integrate PrgSet as needed

	
	var strbuffer string
	f, err := os.Open("app.go")
	Debugger(err)

	defer f.Close()

	scanner := bufio.NewScanner(f)

	for scanner.Scan() {

		Line := scanner.Text()

		// adds functions for routs
		// allow users to choose methods for template serving
			// parsing Glob of html vs parsing individual files  

		if strings.Contains(Line, "//here") {

			for j := 1; j < Num; j++ {
				StJ := strconv.Itoa(j)
				Rfunc := "func Page"+StJ+"(w http.ResponseWriter, r *http.Request) {\n\n\ttmpl := template.Must(template.ParseFiles('static/templates/Page"+StJ+".html'))\n\ttmpl.Execute(w, 'Page"+StJ+"')\n\treturn\n\n}\n\n"
				Rinject := strings.ReplaceAll(Rfunc, "'", `"`,)
				//append to string Buffer code for injection
				strbuffer += Rinject 

			}

		}

		strbuffer += Line +"\n"

		if strings.Contains(Line, "main()"){
			for j := 1; j < Num; j++ {
				StJ := strconv.Itoa(j)
				strinj := "\thttp.HandleFunc('/Page"+StJ+"', Page"+StJ+")\n"
				Rinject := strings.ReplaceAll(strinj, "'", `"`,)
				//append to string Buffer code for injection
				strbuffer += Rinject 

			}
		}

	}

	h, err := os.Create("./app.go")
	_, err2 := h.WriteString(strbuffer)
	Debugger(err)
	Debugger(err2)

	defer h.Close()

	if err := scanner.Err(); err != nil {
		log.Fatal(err)
	}

}


func AppendStyling(BK_color string, Text_color string, Second_color string) {
	fmt.Println("Adding Styling")
	// add styling to css file based on user input
}


func InitCSS(PrgSet struct) {
	fmt.Println("Creating CSS....")

	// TODO: Integrate PrgSet as Needed

	if err := os.Mkdir("./static/css/", os.ModePerm); err != nil {
		log.Fatal(err)
	}


	f, err := os.Create("./static/css/main.css")
	_, err2 := f.WriteString("html {\n\t\n\t\n}\n\nhead {\n\t\n\t\n}\n\nbody {\n\t\n\t\n}\n\nfooter {\n\t\n\t\n}\n")
	Debugger(err)
	Debugger(err2)

	defer f.Close()

}









func InitUnitTest() {
	fmt.Println("Creating Test....")
	//add a unit test file that will create test for everything created and test to see if the creation was successful
}

func ParseUserInput(Uput string, Uvalue int) {

	// 0 == return Uput as string

	// 1 == return Uput as Int

	// 2 == return Uput as bool

	fmt.Println("Parsing Input....")

}









func main() {
	fmt.Println("Starting up Hephaestus....")

	fmt.Println("Welcome!!!!")
	fmt.Println("\nBuilding golang webapp skeleton\n")

	var WebName string
	var domainName string
	var NumTemp int
	var style bool

	fmt.Println("What is the Name of your WebApp?")
	fmt.Scanln(&WebName)
	ParseUserInput(WebName,0)
	fmt.Println("What is the DomainName for your WebApp?\n(enter http://localhost:8080/ for Defualt)")
	fmt.Scanln(&domainName)
	ParseUserInput(domainName,0)
	fmt.Println("Enter the number html Templates you want (as an int ex: 1)")
	fmt.Scanln(&NumTemp)
	ParseUserInput(NumTemp, 1)
	fmt.Println("Do you want Light Css styling?")
	fmt.Scanln(&style)
	ParseUserInput(style,2)

	if style == true {
		var PrimaryC string
		var SecondaryC string
		fmt.Println("What color do you want for Primary color?\n(Hex codes or css standard)")
		fmt.Scanln(&PrimaryC)
		ParseUserInput(PrimaryC,0)
		fmt.Println("What color do you want for Secondary color?\n(Hex codes or css standard)")
		fmt.Scanln(&SecondaryC)
		ParseUserInput(SecondaryC,0)

		appStyle := Styles{
			Pc: PrimaryC,
			Sc: SecondaryC,
		}

	}else {

		appStyle := Styles{
			Pc: "",
			Sc: "",
		}
	}

	// parse user Input


	ProgramSettings := ConfigSetting{
		Title: WebName,
		DomainName: domainName,
		NumTemplate: NumTemp,
		Styling: appStyle,
	}

	GoFRAME_L2(ProgramSettings)


	fmt.Println("Exiting Hephaestus!!")

}