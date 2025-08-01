// package main

// import "fmt"

// func main() {
// 	// Ввод из консоли имени
// 	var name string
// 	fmt.Println("Введи своё имя:")
// 	fmt.Scanln(&name)

// 	// Ввод из консоли возраста
// 	var age int
// 	fmt.Println("Введи свой возраст:")
// 	fmt.Scanln(&age)

// 	// Ввод из консоли среднего балла по любимому предмету
// 	var mark float64
// 	fmt.Println("Введи средний балл по своему любимому предмету:")
// 	fmt.Scanln(&mark)

// 	// Константа с секретным номером
// 	const secretNumber = 54423617

// 	// Формирование вывода и сам вывод
// 	fmt.Println(
// 		fmt.Sprintf(
// 			"Привет, я робот-предсказатель! Вижу, что тебе сейчас %d, а зовут тебя %s. Уникальное число твоего предсказания %d. А средняя оценка по твоему любимому предмету %f. Как я догадался? Пусть это останется секретом моего мастерства!",
// 			age,
// 			name,
// 			secretNumber,
// 			mark,
//         ),
//     )
// }

// package main
// import "fmt"

// func main() {
//     var inp string
//     fmt.Scanln(&inp)
//     fmt.Println(
//         fmt.Sprintf("А ещё я люблю %s!", inp),
//     )
// }

// package main

// import (
// 	"fmt"
// 	"math"
// )

// func main() {
// 	fmt.Println(math.Pi)    // Число пи: 3.141592653589793
// 	fmt.Println(math.E)     // Число e: 2.718281828459045
// 	fmt.Println(math.Sqrt(4)) // Значение корня из двух: 1.4142135623730951
// }


// package main

// import (
// 	"fmt"
// 	"time"
// 	"math"
// )
 
// func main() {
// 	var date, name, surname, otch string
// 	var s1, s2, s3 float64
// 	fmt.Scanln(&date)
// 	t, err := time.Parse("02.01.2006", date)
// 	if err != nil {
// 		return
// 	}
// 	t2 := t.AddDate(0, 1, 15)
// 	fmt.Scanln(&name)
// 	fmt.Scanln(&surname)
// 	fmt.Scanln(&otch)
// 	fmt.Scanln(&s1)
// 	fmt.Scanln(&s2)
// 	fmt.Scanln(&s3)

// 	summ := s1 + s2 + s3
// 	summ_kop := (summ - math.Floor(summ)) * 100

// 	fmt.Printf("Уважаемый, %s %s %s, доводим до вашего сведения, что бухгалтерия сформировала документы по факту выполненной вами работы.\nДата подписания договора: ", name, surname, otch)
// 	fmt.Print(t2.Format("02.01.2006"))
// 	fmt.Printf(". Просим вас подойти в офис в любое удобное для вас время в этот день.\nОбщая сумма выплат составит %.0f руб. %.0f коп.\n\nС уважением,Гл. бух. Иванов А.Е.", math.Floor(summ), math.Round(summ_kop))

// }

// package main

// import (
// 	"fmt"
// 	"strings"
// )

// func main() {
// var str, str2, word string
// var count, s int
// fmt.Scanln(&count)
// fmt.Scanln(&str)
// str2 = strings.ToLower(str)
// for i := 0; i < count; i++ {
// 	fmt.Scanln(&word)
// 	if word == str || word == str2{
// 		s++
// 	}

// }
// fmt.Println(s)
// }


// package main

// import (
// 	"fmt"

// )

// func main() {
// var str string
// var count int
// fmt.Scan(&str)
// fmt.Scan(&count) 
// fmt.Println(count)


// }

// package main

// import (
// 	"fmt"
// )

// func main() {
// var str, str1, str2, str3, str4, str5 string
// var place, count int

// for {
//     fmt.Scanf("%s %d", &str, &place)
// 	if (str == "конец") {
// 		if str1 == "" {
// 			fmt.Println("1. -")
// 		} else {
// 			fmt.Printf("1. %s\n", str1)
// 		}

// 		if str2 == "" {
// 			fmt.Println("2. -")
// 		} else {
// 			fmt.Printf("2. %s\n", str2)
// 		}

// 		if str3 == "" {
// 			fmt.Println("3. -")
// 		} else {
// 			fmt.Printf("3. %s\n", str3)
// 		}

// 		if str4 == "" {
// 			fmt.Println("4. -")
// 		} else {
// 			fmt.Printf("4. %s\n", str4)
// 		}

// 		if str5 == "" {
// 			fmt.Println("5. -")
// 		} else {
// 			fmt.Printf("5. %s\n", str5)
// 		}
// 		return
// 	} else if str == "количество" {
// 		fmt.Printf("Осталось свободных мест: %d\nВсего человек в очереди: %d\n", 5 - count, count)
//         str = ""
// 		place = 0
        
// 	} else if str == "очередь" {
// 		if str1 == "" {
// 			fmt.Println("1. -")
// 		} else {
// 			fmt.Printf("1. %s\n", str1)
// 		}

// 		if str2 == "" {
// 			fmt.Println("2. -")
// 		} else {
// 			fmt.Printf("2. %s\n", str2)
// 		}

// 		if str3 == "" {
// 			fmt.Println("3. -")
// 		} else {
// 			fmt.Printf("3. %s\n", str3)
// 		}

// 		if str4 == "" {
// 			fmt.Println("4. -")
// 		} else {
// 			fmt.Printf("4. %s\n", str4)
// 		}

// 		if str5 == "" {
// 			fmt.Println("5. -")
// 		} else {
// 			fmt.Printf("5. %s\n", str5)
// 		}
//         str = ""
// 		place = 0
// 	} else if str == "" {
//         continue
//     } else {
// 		// fmt.Scan(place)
//         // fmt.Println(place)
// 		if place <= 0 || place > 5 {
// 			fmt.Printf("Запись на место номер %d невозможна: некорректный ввод\n", place)
// 			continue
// 		}
// 		if count == 5 {
// 			fmt.Printf("Запись на место номер %d невозможна: очередь переполнена\n", place)
// 			continue
// 		}
// 		switch place {
// 		case 1: 
// 			if str1 != "" {
// 				fmt.Printf("Запись на место номер %d невозможна: место уже занято\n", place)
// 				continue
// 			}
// 			str1 = str
// 			count++
// 		case 2: 
// 			if str2 != "" {
// 				fmt.Printf("Запись на место номер %d невозможна: место уже занято\n", place)
// 				continue
// 			}
// 			str2 = str
// 			count++

// 		case 3: 
// 			if str3 != "" {
// 				fmt.Printf("Запись на место номер %d невозможна: место уже занято\n", place)
// 				continue
// 			}
// 			str3 = str
// 			count++

// 		case 4: 
// 			if str4 != "" {
// 				fmt.Printf("Запись на место номер %d невозможна: место уже занято\n", place)
// 				continue
// 			}
// 			str4 = str
// 			count++

// 		case 5: 
// 			if str5 != "" {
// 				fmt.Printf("Запись на место номер %d невозможна: место уже занято\n", place)
// 				continue
// 			}
// 			str5 = str
// 			count++
//         default: 
//             // fmt.Println(str)
//             // fmt.Println(place)


// 		}
// 		str = ""
// 		place = 0
// 	}
	

// }

// }


// package main

// import (
// 	"fmt"
// 	"strings"
// )

// func main() {
// var str, str2, word string
// var count, s int
// fmt.Scanln(&count)
// fmt.Scanln(&str)
// str2 = strings.ToLower(str)
// for i := 0; i < count; i++ {
// 	fmt.Scan(&word)
// 	if strings.ToLower(word) == str2{
// 		s++
// 	}

// }
// fmt.Println(s)
// }


// package main

// import (
// 	"match"
// )


// type Share interface {
// 	Area() float64
// }

// type Circle struct {
// 	radius float64
// }

// func (c Circle) Area() float64 {
// 	return math.Pi * c.radius
// }

// type Rectangle {
// 	height float64
// 	width float64
// }


// func (r Rectangle) Area() float64 {
// 	return r.width * r.height
// }


// package main

// import (
// 	"fmt"
// )


// type Logger interface {
// 	Log(message string)
// }

// type  LogLevel struct {
// 	str string
// }

// const Error LogLevel{str: "Error"}
// const Info LogLevel{str: "Info"}

// type Log struct {
// 	Level LogLevel
// }

// func (l Log) Log(st string) {
// 	if l.Level == Error {
// 		fmt.Printf("ERROR: %s\n", st)
// 	} else {
// 		fmt.Printf("INFO:: %s\n", st)
// 	}
// }


// package main 

// import (
// 	"errors"
// )

// type User struct {
// 	ID int
// 	Name string
// 	Email string
// 	Age int
// }

// type  Report struct {
// 	User
// 	ReportID int
// 	Date string
// }

// func CreateReport(user User, reportDate string) Report {
// 	return Report{user, ReportID: user.ID+100, Date: reportDate}
// }

// func GenerateUserReports(users []User, reportDate string) []Report {
// 	rep := []Report{}
// 	for i := 0; i < len(users); i++ {
// 		rep = append(rep, CreateReport(users[i], reportDate))

// 	}
// 	return rep
// }

// package main 

// import (
// 	"strings"
// 	"fmt"
// )

// type Writer interface {
// 	Write(p []byte) int
// }

// type Reader interface {
// 	Read() []byte
// }

// type ReaderWriter interface {
// 	Writer
// 	Reader
// }

// type UpperReaderWriter struct {
// 	UpperString string
// }

// func (u *UpperReaderWriter) Write(p []byte) int {
// 	for i :=0; i < len(p); i++ {
// 		u.UpperString += string(p[i])
// 	}
// 	u.UpperString = strings.ToUpper(u.UpperString)
// 	return len(p)
// }


// func (u *UpperReaderWriter) Read() []byte {
// 	res := []byte{}
// 	fmt.Println(len(u.UpperString))
// 	for i := 0; i < len(u.UpperString); i++ {
// 		res = append(res, u.UpperString[i])
// 	}
// 	return res
// }

// func main() {
// 	urw := &UpperReaderWriter{}

// 	testData := []byte("hello, world!")
// 	bytesWritten := urw.Write(testData)
// 	fmt.Printf("Записано %d байт: %s\n", bytesWritten, urw.UpperString)

// 	readData := urw.Read()
// 	fmt.Printf("Прочитано: %s\n", readData)

// 	if urw.UpperString != strings.ToUpper(string(testData)) {
// 		fmt.Println("Ошибка: строка не преобразована в верхний регистр")
// 	}

// 	var _ Reader = urw
// 	var _ Writer = urw
// 	var _ ReaderWriter = urw
// }

// package main

// import (
//     "fmt"
//     "Errors"
// )

// type Account struct {
// 	balance float64
// 	owner string
// }

// func  NewAccount(owner string) *Account {
// 	return &Account{balance: 0, owner: owner}
// }

// func (a *Account) SetBalance(value float64) error {
// 	if value < 0 {
// 		return errors.New("неправильный баланс")
// 	}

// 	a.balance = value
// 	return nil
// }

// func (a *Account) GetBalance() float64 {
// 	return a.balance
// }

// func (a *Account) Deposit(value float64) error {
// 	if value < 0 {
// 		return errors.New("неправильный баланс")
// 	}

// 	a.balance += value
// 	return nil

// }

// func (a *Account) Withdraw(value float64) error {
// 	if value < 0 && a.balance > value {
// 		return errors.New("неправильный баланс")
// 	}

// 	a.balance -= value
// 	return nil

// }


// package config

// type Config struct {
// 	DefaultBalance float64
// }


// func NewConfig() *Config {
// 	return &Config(DefaultBalance: 1000)
// }


// package main 

// import (
// 	"fmt"
// )

// type Animal interface {
// 	MakeSound() string
// 	GetName() string
// 	GetInfo() string
// }

// type animal struct {
// 	name string
// 	species string
// 	age int
// 	sound string
// }

// func (a animal) MakeSound() string {
// 	return a.sound
// }

// func (a animal) GetName() string {
// 	return a.name
// }

// func (a animal) GetInfo() string {
// 	return fmt.Sprintf("Имя: %s, Вид: %s, Возраст: %d", a.name, a.species, a.age)
// }

// func NewAnimal(name, species string, age int, sound string) Animal {
// 	return Animal{name: name, species: species, age: age, sound: sound}
// }

// func ZooShow(animals []Animal) {
// 	for i := range animals {
// 		fmt.Println(animals[i].GetInfo())
// 		fmt.Println(animals[i].MakeSound())
// 	}
// }

// type ZooKeeper struct {

// }

// func (z ZooKeeper) Feed(animal Animal) {
// 	fmt.Printf("Смотритель зоопарка кормит %s. %s!\n", animal.GetName(), animal.MakeSound())
// }


// func TestPrintHello(t *testing.T) {
//     got := PrintHello("Igor")
//     expected := "Hello, Igor!"

//     if got != expected {
//     t.Fatalf(`PrintHello("Igor") = %q, want %q`, got, expected)
//     }
// }


// package main

// import "testing"

// func TestSum(t *testing.T) {
// 	if Sum(1, 2) != 3 {
// 		t.Error("Do not work")
// 	}
// }

// package main

// import "testing"

// var data = map[int]string{
// 	0: "zero",
// 	-2: "negative",
// 	7: "short",
// 	35: "long",
// 	190: "very long",
// }

// func TestLength(t *testing.T) {
// 	for k, v := range data {
//         if Length(k) != v {
// 			t.Errorf("для %d ожидали %s получили %s", k, v, Length(k))
// 		}
//     }

	
// }

// package main

// import "testing"

// func TestMultiply(t *testing.T) {
// 	if Multiply(1, 2) != 2 {
// 		t.Error("Do not work")
// 	} 
// 	if Multiply(1, 0) != 0 {
// 		t.Error("Do not work")
// 	} 
// 	if Multiply(-1, 2) != -2 {
// 		t.Error("Do not work")
// 	} 
// 	if Multiply(-1, -2) != 2 {
// 		t.Error("Do not work")
// 	} 
// }


package main

import "testing"

func TestDeleteVowels(t *testing.T) {
	cases := map[string]string{
		"apple" : "pple",
		"ye" : "e",
		"oau" : "",
		"лол" : "лол",
		"psd" : "psd",
	}

	for k, v := range cases {
		if DeleteVowels(k) != v {
			t.Errorf("для %s ожидали %s получили %s", k, v, DeleteVowels(k))

		}
	}


}