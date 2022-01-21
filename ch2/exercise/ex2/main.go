// Write a general-purpose unit-conversion program analogous to cf that reads numbers from its command-line arguments or from the standard input if there are no arguments, and converts each number into units like temperature in Celsius and Fahrenheit, length in feets and meters, weight in pounds and kilograms, and the like

package main

import (
	"bufio"
	"flag"
	"fmt"
	"os"
	"strconv"
	"strings"
	"the_gopl/ch2/tempconv"
)

type Feet float64
type Meter float64
type Pound float64
type Kilogram float64

// set of the variables pre-defined are not necessary here, I'm just doing this
// to testout flag.*Var() functions
var d bool
var m bool
var t bool
var dist string
var mass string
var temp string
var sep string

func init() {
	flag.BoolVar(&d, "d", false, "Enter distance via std. input")
	flag.BoolVar(&m, "m", false, "Enter weight via std. input")
	flag.BoolVar(&t, "t", false, "Enter temp via std. input")
	flag.StringVar(&dist, "dist", "nil", "Provide distance as arg")
	flag.StringVar(&mass, "mass", "nil", "Provide mass as arg")
	flag.StringVar(&temp, "temp", "nil", "Provide temp as arg")
	flag.StringVar(&sep, "sep", " ", "Provide a seperator when using non-boolean parameter")
}

func (f Feet) String() string {
	return fmt.Sprintf("%.6gF", f)
}

func (m Meter) String() string {
	return fmt.Sprintf("%.6gM", m)
}

func (p Pound) String() string {
	return fmt.Sprintf("%.6gP", p)
}

func (k Kilogram) String() string {
	return fmt.Sprintf("%.6gKg", k)
}

func FToM(f Feet) Meter {
	return Meter(f * 0.305)
}

func MToF(m Meter) Feet {
	return Feet(m * 3.281)
}

func PToK(p Pound) Kilogram {
	return Kilogram(p * (9.0 / 20))
}

func KToP(k Kilogram) Pound {
	return Pound(k * (20 / 9))
}


func main() {
	flag.Parse()
	input := bufio.NewScanner(os.Stdin)
	//fmt.Println(flag.NFlag())
	switch {
	case d:
		callConversion("1", input)
	case m:
		callConversion("2", input)
	case t:
		callConversion("3", input)
	}
	if flag.NFlag() != 0 {
		//fmt.Println(dist, mass)
		if dist != "nil" {
			for i, sd := range strings.Split(dist, sep) {
				fd, _ := strconv.ParseFloat(sd, 64)
				ucD(i + 1, fd)
			}
		}
		if mass != "nil" {
			for i, sd := range strings.Split(mass, sep) {
				fd, _ := strconv.ParseFloat(sd, 64)
				ucM(i + 1, fd)
			}
		}
		if temp != "nil" {
			for i, sd := range strings.Split(temp, sep) {
				fd, _ := strconv.ParseFloat(sd, 64)
				ucT(i + 1, fd)
			}
		}
	} else {
		initStdinInput(input)
	}
}

func initStdinInput(input *bufio.Scanner) {
	for {
		fmt.Printf("What do you want to convert:\n")
		fmt.Println("1: Distance")
		fmt.Println("2: Mass")
		fmt.Println("3: Temp")
		fmt.Println("q: Quit")
		fmt.Print("Enter the choice[number]: ")
		input.Scan()
		in := input.Text()
		if in != "q" {
			callConversion(in, input)
		} else {
			fmt.Println("Exiting converter!!!")
			break
		}
	}
}

func callConversion(in string, input *bufio.Scanner) {
	fmt.Println("Enter conversion numbers:")
	for count :=1; input.Scan(); count++ {
		userInput := input.Text()
		if userInput == "q" {
			fmt.Println("Exiting!")
			return;
		} else {
			i, _ := strconv.ParseFloat(userInput, 64)
			switch in {
			case "1":
				ucD(count, i)
			case "2":
				ucM(count, i)
			case "3":
				ucT(count, i)
			default:
				fmt.Println("Didn't enter correct conversion choice. Select again")
				return;
			}
		}
	}
}

// unit conversion of distance
func ucD(i int, v float64) {
	fmt.Printf("%d: Distance conversion for: %g\n", i, v)
	f := Feet(v)
	m := Meter(v)
	fmt.Printf("%*s = %s, %s = %s\n",
		len(strconv.Itoa(i)) + len(f.String()) + 2, f, FToM(f), m, MToF(m))
}

// unit conversion of mass
func ucM(i int, v float64) {
	fmt.Printf("%d: Mass conversion for: %g\n", i, v)
	p := Pound(v)
	k := Kilogram(v)
	fmt.Printf("%*s = %s, %s = %s\n",
		len(strconv.Itoa(i)) + len(p.String()) + 2, p, PToK(p), k, KToP(k))
}

// unit conversion of temperature
func ucT(i int, v float64) {
	fmt.Printf("%d: Temp conversion for: %g\n", i, v)
	c := tempconv.Celsius(v)
	f := tempconv.Fahrenheit(v)
	k := tempconv.Kelvin(v)
	fmt.Printf("%*s = %s, %s = %s, %s = %s\n",
		len(strconv.Itoa(i)) + len(c.String()) + 1, c, tempconv.CToF(c), f, tempconv.FToK(f), k, tempconv.KToC(k))
}


/** Note: I suppose there could be a better way to write initStdinInput() and
* callConversion() without changing much logic, but anyways...
* I'll try to comeback to this problem someday, who know */
