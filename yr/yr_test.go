package yr

import (
	"testing"
)


/*
This function test a specific unit of code (function). 
*/
func TestCountLines (t *testing.T) {

	//defines a type test which is a struct
	type test struct {
		input string
		want float64
	}
	
	//creates an array of test instances calles tests
	tests := []test{
		{input: "/home/dzem/minyr/kjevik-temp-celsius-20220318-20230318.csv", want: 16756},
		{input: "/home/dzem/minyr/kjevik-temp-fahr-20220318-20230318.csv", want: 16756},
	}

	//loop that iterates through each test case
	for _, tc := range tests {
		got := countLines(tc.input)
		//if the outpur does not match print the details of the failed test case
		if !(tc.want == got) {
			t.Errorf("expected %v, got %v", tc.want, got) //prints the values as they are
		}
	}
} 

func TestAverage (t *testing.T) {
	type test struct {
		input string
		want float64
	}

	 tests := []test{
                {input: "/home/dzem/minyr/kjevik-temp-celsius-20220318-20230318.csv", want: 8.56},
                {input: "/home/dzem/minyr/kjevik-temp-fahr-20220318-20230318.csv", want: 47.4},
        }

	 for _, tc := range tests {
                got := AverageT(tc.input)
                if !(tc.want == got) {
                        t.Errorf("expected %v, got %v", tc.want, got) 
                }
        }
}


func TestCelsiusToFahrenheitString(t *testing.T) {
	
     type test struct {
	input string
	want string
     }

     tests := []test{
	     {input: "6", want: "42.8"},
	     {input: "0", want: "32.0"},
     }

     for _, tc := range tests {
	     got, _ := CelsiusToFahrenheitString(tc.input)
	     if !(tc.want == got) {
		     t.Errorf("expected %s, got: %s", tc.want, got)
	     }
     }
}

// Forutsetter at vi kjenner strukturen i filen og denne implementasjon 
// er kun for filer som inneholder linjer hvor det fjerde element
// på linjen er verdien for temperatrmaaling i grader celsius
func TestCelsiusToFahrenheitLine(t *testing.T) {
     type test struct {
	input string
	want string
     }
     tests := []test{
	     {input: "Kjevik;SN39040;18.03.2022 01:50;6", want: "Kjevik;SN39040;18.03.2022 01:50;42.8"},
	     {input: "Kjevik;SN39040;18.03.2022 01:50;0", want: "Kjevik;SN39040;18.03.2022 01:50;32.0"},
	     {input: "Kjevik;SN39040;18.03.2022 01:50;-11", want: "Kjevik;SN39040;18.03.2022 01:50;12.2"},
	     {input: "Kjevik;SN39040;18.03.2022 01:50", want: ""},
	     {input: "Data er gyldig per 18.03.2023 (CC BY 4.0), Meteorologisk institutt (MET);;;",
	     want: "Data er basert paa gyldig data (per 18.03.2023) (CCBY 4.0) fra Meteorologisk institutt (MET);endringen er gjort av Dzemil Alendar"},

     }

     for _, tc := range tests {
	     got, _ := CelsiusToFahrenheitLine(tc.input)
	     if !(tc.want == got) {
		     t.Errorf("expected %s, got: %s", tc.want, got)
	     }
     }

	
}
