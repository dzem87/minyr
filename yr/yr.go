package yr

import (
	"fmt"
	"strconv"
	"strings"
	"errors"
	"log"
	"bufio"
	"os"
	"github.com/dzem87/funtemps/conv"
)
/*
This function counts the lines of a specified file and returns the number of lines
*/
func countLines (filename string) float64{
	file, err := os.Open(filename) //open the specified file
	if err!= nil { //check if there was an error opening the file
		log.Fatal(err)
	}

	defer file.Close()

	scanner := bufio.NewScanner(file) //create a new scanner to read the file 
	count := 0.0
	
	for scanner.Scan() { //the for loop reads each line of the file using the scanner 
		count++ //increment the counter for each line read
	}

	if err := scanner.Err(); err !=nil { //check if there was an error scanning the file
		log.Fatal(err)
	}

	return count //return the final count
}



/*
The function reads a file that contains weather data, 
and returns the average temperature  
*/
func AverageT (filename string) float64 {
	file, err := os.Open(filename)
        if err!= nil {
                log.Fatal(err)
        }

        scanner := bufio.NewScanner(file)
        count := 0.0
	tempSum := 0.0
	avg := 0.0

        for scanner.Scan() {
		line := scanner.Text() //obtaining the recent line that was read
		
		//skip the to the next line without performing any furhter operations on the specified line
		if strings.Contains(line, "temperatur") {
		continue  
		} 
		if strings.Contains(line, "institutt") {
		continue
		}

		//split up the line and get a specifif element
		elementsInLine := strings.Split(line, ";")
		//convert string to float64 value
		temp, err := strconv.ParseFloat(elementsInLine[3], 64)
		if err!= nil {
			log.Fatal(err)
		}

		tempSum += temp //increment the temp counter
                count++ //increment the line counter
        }

        if err := scanner.Err(); err !=nil { //check if there was an error scanning the file
                log.Fatal(err)
        }

	//calculate the averga
	avg = (tempSum/count)

        return avg
}


/*
Takes a celsius string as input, converts it to fahrenheit using the conv package,
and retursn the Fahr temperature in string format 
*/
func CelsiusToFahrenheitString(celsius string) (string, error) {
	var fahrFloat float64
	var err error
	/*converts the celsius string to a float64 value, and assigns it to the celsiusFloat variable
	and also assings the nil value to error. If there is any error, the function will skip the next step
	*/
	if celsiusFloat, err := strconv.ParseFloat(celsius, 64); err == nil {
		fahrFloat = conv.CelsiusToFahrenheit(celsiusFloat)
	}
	//converts fahrFloat to a string format
	fahrString := fmt.Sprintf("%.1f", fahrFloat)
	return fahrString, err
}

/* 
The function takes a string as input and assumes that it contains four elements seperated by ";",
it converts the fourth element using the function declared earlier. It retruns the modified input string
with the fourth element converted to fahrenheit
*/
func CelsiusToFahrenheitLine(line string) (string, error) {
	//splits the input into a slice of elements
	elementsInLine := strings.Split(line, ";")
	var err error

	/*
	checks if the number of elements in the slice is 4 and proceeds to the next step
	if not it returns an error message
	*/
	if (len(elementsInLine) == 4) {	

		//if the line contains contains a specified string return a new string and an error 
		if strings.HasPrefix(elementsInLine[0], "Data er gyldig") {
                return "Data er basert paa gyldig data (per 18.03.2023) (CCBY 4.0) fra Meteorologisk institutt (MET);endringen er gjort av Dzemil Alendar", err
		}

		//converts celsius to fahrenheit and replaces the old value with the new
		elementsInLine[3], err = CelsiusToFahrenheitString(elementsInLine[3])
		if err != nil {
			return "", err
		}
		} else {
			return "", errors.New("linje har ikke forventet format")
		}

	//joins the elements in the slice back into a semicolon seperated string
	return strings.Join(elementsInLine, ";"), nil


}




