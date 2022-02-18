package main

import (
	"bufio"
	"encoding/json"
	"fmt"
	"io"
	"io/ioutil"
	"math"
	"math/rand"
	"net/http"
	"os"
	"sort"
	"strconv"
	"strings"
	"time"
)

const aConst int = 68

const url = "http://services.explorecalifornia.org/json/tours.php"

func main2() {
	resp, err := http.Get(url)
	checkError(err)

	// fmt.Printf("Responce type %T\n", resp)

	defer resp.Body.Close()

	bytes, err := ioutil.ReadAll(resp.Body)
	checkError(err)

	content := string(bytes)
	// fmt.Println(content)

	tours := toursFromJson(content)

	for _, tour := range tours {
		fmt.Println(tour.Name)
	}

	fmt.Println("Network request")
}

func toursFromJson(content string) []Tour {
	tours := make([]Tour, 0, 20)
	decoder := json.NewDecoder(strings.NewReader(content))
	_, err := decoder.Token()
	checkError(err)
	var tour Tour
	for decoder.More() {
		err := decoder.Decode(&tour)
		checkError(err)
		tours = append(tours, tour)
	}

	return tours
}

type Tour struct {
	tourId       string
	packageId    string
	packageTitle string
	Name         string
	blurb        string
	description  string
	Price        string
	difficulty   string
}

func main0601() {
	fmt.Println("Files")
	content := "Hello from Go!"
	file, err := os.Open("./fromString.txt")

	checkError(err)
	length, err := io.WriteString(file, content)
	checkError(err)
	fmt.Printf("Wrote a file with %v char\n", length)
	defer file.Close()
	defer readFile(file)

}

func readFile(file *os.File) {
	data, err := ioutil.ReadFile(file.Name())
	checkError(err)
	fmt.Println("Text is:", string(data))

}

func checkError(err error) {
	if err != nil {
		panic(err)
	}
}

func main0503() {
	dog := newDog()
	fmt.Println(dog)
	fmt.Printf("%+v\n", dog)
	dog.Weight = 9
	fmt.Println(dog)
	dog.Speak()
	dog.Sound = "Arf"
	dog.Speak()

}

// Speak is how the dog speaks
func (d Dog) Speak() {
	fmt.Println(d.Sound)
}

func newDog() Dog {
	poodle := Dog{"Poodle", 10, "Woof"}
	return poodle
}

type Dog struct {
	Breed  string
	Weight int
	Sound  string
}

func main0501() {
	doSomething()
	sum := addValue(5, 8)
	fmt.Println(sum)
	sum = addAllValues(1, 2, 3, 4, 5)
	fmt.Println(sum)
}

func doSomething() {
	fmt.Println("Doing something")
}

func addValue(v1, v2 int) int {
	return v1 + v2
}

func addAllValues(values ...int) int {
	total := 0
	for _, v := range values {
		total += v
	}

	return total
}

func main0403() {
	colors := []string{"R", "C", "B"}
	fmt.Println(colors)

	for i := 0; i < len(colors); i++ {
		fmt.Println(colors[i])
	}

	for i := range colors {
		fmt.Println(colors[i])
	}

	for _, color := range colors {
		fmt.Println(color)
	}

	value := 1
	for value < 10 {
		fmt.Println(value)
		value++
	}
	rand.Seed(time.Now().Unix())
	dow := rand.Intn(7) + 1
	fmt.Println(dow)
	ans := 42
	var result string
	if ans > 0 {
		result = "Greater"
	}
	fmt.Println(result)
}

func main0306() {
	poodle := newDog()
	fmt.Println(poodle)
	fmt.Printf("%+v\n", poodle)
	poodle.Weight = 9
	fmt.Println(poodle)

}

func main0305() {
	states := make(map[string]string)
	fmt.Println(states)
	states["WA"] = "was"
	states["OR"] = "Ora"
	states["CA"] = "Cal"
	fmt.Println(states)

	cali := states["CA"]
	fmt.Println(cali)

	delete(states, "OR")
	fmt.Println(states)

	states["NY"] = "New Y"
	fmt.Println(states)
	for k, v := range states {
		fmt.Printf("%v : %v\n", k, v)
	}

	keys := make([]string, len(states))
	i := 0
	for k := range states {
		keys[i] = k
		i++
	}

	fmt.Println(keys)

	var colors2 = []string{"R", "G", "B"}
	colors2 = append(colors2, "P")
	fmt.Println(colors2)
	colors2 = append(colors2[0 : len(colors2)-1])
	fmt.Println(colors2)

	number2 := make([]int, 5, 5)
	number2[0] = 124
	number2[1] = 12
	number2[2] = 14
	fmt.Println(number2)
	sort.Ints(number2)
	fmt.Println(number2)

	var colors [3]string
	colors[0] = "R"
	colors[1] = "G"
	colors[2] = "B"

	fmt.Println(colors)
	fmt.Println(colors[0])

	var numbers = [5]int{5, 3, 1, 2, 4}
	fmt.Println(numbers)

	fmt.Println("number of colors:", len(colors))
	fmt.Println("number of numbers:", len(numbers))

	m := make(map[string]int)
	m["the"] = 42
	fmt.Println(m)

	anInt := 42
	var p *int = &anInt
	fmt.Println("the value p: ", *p)

	v1 := 42.23
	p1 := &v1
	fmt.Println("v1 is:", *p1)

}

func challenge02() {
	reader := bufio.NewReader(os.Stdin)
	fmt.Print("Value1: ")
	value1str, _ := reader.ReadString('\n')
	aFloat1, err := strconv.ParseFloat(strings.TrimSpace(value1str), 64)
	if err != nil {
		panic(err)
	}

	fmt.Print("Value2: ")
	value2str, _ := reader.ReadString('\n')
	aFloat2, err := strconv.ParseFloat(strings.TrimSpace(value2str), 64)
	if err != nil {
		panic(err)
	}

	// sum := (math.Round(aFloat1+aFloat2) * 100) / 100
	sum := ((aFloat1 + aFloat2) * 100) / 100
	// fmt.Printf("The sum of %v and %v is %v", aFloat1, aFloat2, sum)
	fmt.Printf("The sum of %v and %v is %v", aFloat1, aFloat2, sum)

}

func main0206() {

	i1, i2, i3 := 12, 45, 68
	intSum := i1 + i2 + i3
	fmt.Println("Integer sum:", intSum)

	f1, f2, f3 := 23.5, 65.1, 76.3
	floatSum := f1 + f2 + f3
	fmt.Println("Float sum:", floatSum)

	floatSum = math.Round(floatSum*100) / 100
	fmt.Println("The sum is now", floatSum)

	circleRadius := 15.5
	circumference := circleRadius * 2 * math.Pi
	fmt.Printf("Circumference: %.2f\n", circumference)

}

func main0207() {

	n := time.Now()
	fmt.Println("I recorded this video at ", n)

	t := time.Date(2009, time.November, 10, 23, 0, 0, 0, time.UTC)
	fmt.Println("Go launched at ", t)
	fmt.Println(t.Format(time.ANSIC))

	parsedTime, _ := time.Parse(time.ANSIC, "Tue Nov 10 23:00:00 2009")
	fmt.Printf("The type of parsedTime is %T\n", parsedTime)

}

func mai0204() {

	reader := bufio.NewReader(os.Stdin)
	fmt.Print("Enter text: ")
	input, _ := reader.ReadString('\n')
	fmt.Println("You entered:", input)

	fmt.Print("Enter a number: ")
	numInput, _ := reader.ReadString('\n')
	aFloat, err := strconv.ParseFloat(strings.TrimSpace(numInput), 64)
	if err != nil {
		fmt.Println(err)
	} else {
		fmt.Println("Value of number:", aFloat)
	}

}

func main0203() {

	reader := bufio.NewReader(os.Stdin)
	fmt.Print("Enter text: ")
	input, _ := reader.ReadString('\n')
	fmt.Println("You entered:", input)

}

func main0202() {

	var aString string = "This is Go!"
	fmt.Println(aString)
	fmt.Printf("The variable's type is %T\n", aString)

	var anInteger int = 42
	fmt.Println(anInteger)

	var defaultInt int
	fmt.Println(defaultInt)

	var anotherString = "This is another string"
	fmt.Println(anotherString)
	fmt.Printf("The variable's type is %T\n", anotherString)

	var anotherInt = 53
	fmt.Println(anotherInt)
	fmt.Printf("The variable's type is %T\n", anotherInt)

	myString := "This is also a string"
	fmt.Println(myString)
	fmt.Printf("The variable's type is %T\n", myString)

	fmt.Println(aConst)
	fmt.Printf("The variable's type is %T\n", aConst)

}
