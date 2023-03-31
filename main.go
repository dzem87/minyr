package main

import (
	"fmt"
        "os"
        "log"
	"bufio"
	yr "github.com/dzem87/minyr/yr"
)

func main(){

       var input string
       //creates a scanner that reads data from the standard input
       scanner := bufio.NewScanner(os.Stdin)
       
       /*the call to scanner.Scan will block the terminalwindow and let the user write commands,
       while the program is still running
       */
       for scanner.Scan () {
		input = scanner.Text()
		//exit the program if input equals q or exit 
		if input == "q" || input == "exit" {
			fmt.Println("exit")
			os.Exit(0)
		//convert file with celsius values to fahrenheit
		} else if  input == "convert" {
			fmt.Println("Konverterer alle malingene gitt i grader Celsius til grader Farhrenheit")
			//open file
			src, err := os.Open("kjevik-temp-celsius-20220318-20230318.csv")
        		if err != nil {
                		log.Fatal(err)	
        		}
			
			defer src.Close()

			//create a new file
			dst, err := os.OpenFile("kjevik-temp-fahr-20220318-20230318.csv", os.O_CREATE|os.O_TRUNC|os.O_RDWR, 0666)
                        if err!= nil {
                                log.Fatal(err)
                        }

                        defer dst.Close()
			
			//scanner for src
			scanner := bufio.NewScanner(src)
			//writer for dst
			writer := bufio.NewWriter(dst)
			
			//ignoring the first line if the Scan method returns true
			if scanner.Scan() {
			firstLine := (scanner.Text()) //retrives the line
			fmt.Fprintln(writer, firstLine) //writes a formated string to the specified writer
			}

			//read one line at a time at a time and give it to the variable
			for scanner.Scan() {
			line := scanner.Text()
			newLine, err := yr.CelsiusToFahrenheitLine(line)
			if err!= nil {
				log.Fatal(err)
			}

			//write data to the new file
			fmt.Fprintln(writer, newLine)

			}

			 //ensure that any remaining data in the buffer is written out 
                        writer.Flush()				
				
		//calculate the average temperature either in the cel og fahr file 
		} else if input == "average" {
            		fmt.Println("cel eller fahr")
            		var tempInput string
			
			//scanner that reads input from the teminal
            		if scanner.Scan() {
                		tempInput = scanner.Text()
            		} else {
                		log.Fatal("Error reading input:", scanner.Err())
            		}
			
			//calculate average celsius
            		if tempInput == "cel" {
                		celAvg := yr.AverageT("kjevik-temp-celsius-20220318-20230318.csv")
				fmt.Printf("gjennomsnittstemperatur i celsius: %2.0f\n", celAvg)
			//calculate average fahr	
            		} else if tempInput == "fahr" {
                		fahrAvg := yr.AverageT("kjevik-temp-fahr-20220318-20230318.csv")
				fmt.Printf("gjennomsnittstemperatur i fahrenheit: %2.0f\n", fahrAvg)
			}
		}
       	}
      
}
