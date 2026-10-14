
package main

import (

	"net/http"
	"log"
	"os"
	"io"
	"regexp"
	"strings"
	"strconv"
	"fmt"

)

var verbose_on bool = true

func verbose(message string){

	log.Print(message)

}

type Application struct{


}

func main(){

	var app Application

	app.run()


}

func (app Application) run(){

	getFile()

}

func getFile(){

	//Arguements 

	var arguements []string = os.Args[1:]

	if (len(arguements) >= 2){

		fatalm("Please input one argument only!!")

	}else if (len(arguements) <= 0){

		fatalm("There's no input!!")

	}

	if (verbose_on == true){

		verbose("Program Start!")
		verbose("Getting the Url....")
		

	}

	//Getting the url

	var response *http.Response
	var err error

	response, err = http.Get(arguements[0])

	if (err != nil){

		if (verbose_on == true){
			verbose("Failed")
		}

		fatal(err)

	}

	if (verbose_on == true){

		verbose("Success!!")
		verbose("Copying the url body(Copy the url Body)...")

	}

	//Copying the inside of the file
	
	var fileBody []byte 
	fileBody, err = io.ReadAll(response.Body)

	if (err != nil){

		fatal(err)

	}

	if (verbose_on == true){

		verbose("Success!!")
		verbose("Closing the file body")

	}

	response.Body.Close()

	if (verbose_on == true){

		verbose("file body Closed")
		verbose("Creating the file...")

	}

	//Creating the file the file(where the inside of the body will be placed)

	var filePath string = "yogurt.txt"
	var flags int = int(os.O_RDONLY | os.O_WRONLY | os.O_CREATE)
	var filePerm os.FileMode = 0777
	var pFile *os.File 

	pFile, err = os.OpenFile(filePath, flags, filePerm)

	if (err != nil){

		fatal(err)

	}

	if(verbose_on == true){

		verbose("Success!!")
		verbose("Converting the file body from byte-string")

	}

	var fileString string = string(fileBody)
	var sizeOfString int  

	if (verbose_on == true){

		verbose("Pasting the url body into the document.txt!!")

	}

	//size of string that is being write in the file

	sizeOfString, err = pFile.WriteString(fileString)

	var wholeFileBuffer []byte = make([]byte, sizeOfString)

	if (verbose_on == true){

		verbose("Reading the document.txt")

	}

	//inside of the file in []byte

	wholeFileBuffer, err = os.ReadFile(filePath)

	if (err != nil){

		fatal(err)

	}

	var first_targetString = [4]string { "<span class=\"c0\">", "<span class=\"c1\">", "<span class=\"c2\">" } 
	var second_targetString string = "</span></p>"
	var num_count int = 0

	for{

		if (fileStringChecker(first_targetString[num_count], wholeFileBuffer) == true){

			break
		}

		num_count++
	}

	log.Print(first_targetString[num_count])
	
	var rexp *regexp.Regexp

	rexp, err = regexp.Compile(first_targetString[num_count])

	if (err != nil){

		fatal(err)

	}

	if (verbose_on == true){

		verbose("Searching the data inside of the file......")
		log.Print(sizeOfString)

	}

	var endFile int = -1
	var grepIndexLocation [][]int = rexp.FindAllIndex(wholeFileBuffer, endFile)

	var rexp_2 *regexp.Regexp
	rexp_2, err = regexp.Compile(second_targetString)

	if (err != nil){

		fatal(err)

	}

	if (rexp_2.Match(wholeFileBuffer) != true){

		fatalm("No Match")

	}
	
	var grepIndexLocation_2 [][]int = rexp_2.FindAllIndex(wholeFileBuffer, endFile)	

	if (verbose_on == true){

		verbose("Match Found!!")

	}


	var inioffFormula int = len(grepIndexLocation)


	var read *strings.Reader = strings.NewReader(string(wholeFileBuffer))

	var a int
        var b int
	var formula int 
      	var num_buff int
	var buffer []byte 
	var x_coor []int
	var char []string

	for i := 5; i < inioffFormula; i++ {

		var tempHolder int = 0

		a = grepIndexLocation_2[i][0] 
		b = grepIndexLocation[i][1] 
		formula = a - b

		buffer = make([]byte, formula)
		num_buff, err = read.ReadAt(buffer, int64(grepIndexLocation[i][1]))

		tempHolder, err = strconv.Atoi(string(buffer))

		if (err != nil){

			char = append(char, string(buffer))

		}
		
		x_coor = append(x_coor, tempHolder)


	}

	fmt.Print("\x1b[2J")

	var count int = 0
	var count_2 int = 0

	for i := 0; i < len(char); i++ {

		count_2 = count_2 + 1

		fmt.Printf("\x1b[%d;%df", x_coor[count_2], x_coor[count])
		fmt.Printf("%s", char[i])

		count = count + 1

	}
	log.Print(num_buff)

	pFile.Close()

}

func fileStringChecker(compareString string, buffer []byte) bool{

	var check = false

	re, err := regexp.Compile(compareString)

	if (err != nil){

		log.Fatal(err)

	}

	check = re.Match(buffer)

	return check

}

func fatalm(message string){

	log.Fatal(message)

}

func fatal(message error){

	log.Fatal(message)

}
