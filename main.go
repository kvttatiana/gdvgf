package main

import (
	"fmt"
	"math/rand"
	"time"
)

func main() {

	for {
		var gen string
		var symbols string
		var specials string
		var big string
		fmt.Println("Хотите сгенерировать пароль?")
		fmt.Println("1. Да")
		fmt.Println("2. Нет")
		fmt.Scan(&gen)
		if gen == "1" {
			fmt.Println("Хотите добавить символы в пароль? ")
			fmt.Println("1. Да")
			fmt.Println("2. Нет")
			fmt.Scan(&symbols)
			fmt.Println("Хотите добавить спец. символы в пароль?")
			fmt.Println("1. Да")
			fmt.Println("2. Нет")
			fmt.Scan(&specials)
			fmt.Println("Хотите использовать большие буквы?")
			fmt.Println("1. Да")
			fmt.Println("2. Нет")
			fmt.Scan(&big)

		}
		if symbols == "1" && specials == "1" && big == "1" {
			rand.Seed(time.Now().UnixNano())
			//символы
			numbers := "0123456789"
			symbol := "!@#$%^&*"
			all := numbers + symbol + "QWERTYUIOPASDFGHJKLZXCVBNM" + "qwertyuiopasdfghjklzxcvbnm" + "ЙЦУКЕНГШЩЗФЫВАПРОЛДЖЭЯЧСМИТЬ" + "йцукенгшщзхъфывапролджэячсмитьбю"
			var length int
			fmt.Println("Введите длину желаемого пароля:")
			fmt.Scan(&length)
			//массив с символами
			password := make([]byte, length)
			//Первый член - рандомная цифра
			password[0] = numbers[rand.Intn(len(numbers))]
			//Второй член - рандомный символ
			password[1] = symbol[rand.Intn(len(symbol))]
			//генерируем хуйню не менее чем из восьми символов, используя рандом из всего
			for i := 1; i < length; i++ {
				password[i] = all[rand.Intn(len(all))]

			}
			//Перемешиваем элементы среза
			rand.Shuffle(len(password), func(i int, j int) {
				password[i], password[j] = password[j], password[i]
			})
			str := string(password)
			fmt.Println(str)
			break
		}
		if symbols == "1" && specials == "2" && big == "1" {
			rand.Seed(time.Now().UnixNano())
			//символы
			numbers := "0123456789"

			all := numbers + "QWERTYUIOPASDFGHJKLZXCVBNM" + "qwertyuiopasdfghjklzxcvbnm" + "ЙЦУКЕНГШЩЗФЫВАПРОЛДЖЭЯЧСМИТЬ" + "йцукенгшщзхъфывапролджэячсмитьбю"
			var length int
			fmt.Println("Введите длину желаемого пароля:")
			fmt.Scan(&length)
			//массив с символами
			password := make([]byte, length)
			//Первый член - рандомная цифра
			password[0] = numbers[rand.Intn(len(numbers))]
			//генерируем хуйню не менее чем из восьми символов, используя рандом из всего
			for i := 1; i < length; i++ {
				password[i] = all[rand.Intn(len(all))]

			}
			//Перемешиваем элементы среза
			rand.Shuffle(len(password), func(i int, j int) {
				password[i], password[j] = password[j], password[i]
			})
			str := string(password)
			fmt.Println(str)
			break
		}
		if symbols == "1" && specials == "2" && big == "2" {
			rand.Seed(time.Now().UnixNano())
			//символы
			numbers := "0123456789"
			all := numbers + "qwertyuiopasdfghjklzxcvbnm" + "йцукенгшщзхъфывапролджэячсмитьбю"
			var length int
			fmt.Println("Введите длину желаемого пароля:")
			fmt.Scan(&length)
			//массив с символами
			password := make([]byte, length)
			//Первый член - рандомная цифра
			password[0] = numbers[rand.Intn(len(numbers))]
			//генерируем хуйню не менее чем из восьми символов, используя рандом из всего
			for i := 1; i < length; i++ {
				password[i] = all[rand.Intn(len(all))]

			}
			//Перемешиваем элементы среза
			rand.Shuffle(len(password), func(i int, j int) {
				password[i], password[j] = password[j], password[i]
			})
			str := string(password)
			fmt.Println(str)
			break
		}
		if symbols == "1" && specials == "1" && big == "2" {
			rand.Seed(time.Now().UnixNano())
			//символы
			numbers := "0123456789"
			symbol := "!@#$%^&*"
			all := numbers + symbol + "qwertyuiopasdfghjklzxcvbnm" + "йцукенгшщзхъфывапролджэячсмитьбю"
			var length int
			fmt.Println("Введите длину желаемого пароля:")
			fmt.Scan(&length)
			//массив с символами
			password := make([]byte, length)
			//Первый член - рандомная цифра
			password[0] = numbers[rand.Intn(len(numbers))]
			//Второй член - рандомный символ
			password[1] = symbol[rand.Intn(len(symbol))]
			//генерируем хуйню не менее чем из восьми символов, используя рандом из всего
			for i := 1; i < length; i++ {
				password[i] = all[rand.Intn(len(all))]

			}
			//Перемешиваем элементы среза
			rand.Shuffle(len(password), func(i int, j int) {
				password[i], password[j] = password[j], password[i]
			})
			str := string(password)
			fmt.Println(str)
			break
		}
		if symbols == "2" && specials == "2" && big == "2" {
			rand.Seed(time.Now().UnixNano())
			//символы
			all := "qwertyuiopasdfghjklzxcvbnm" + "йцукенгшщзхъфывапролджэячсмитьбю"
			var length int
			fmt.Println("Введите длину желаемого пароля:")
			fmt.Scan(&length)
			//массив с символами
			password := make([]byte, length)
			//генерируем хуйню не менее чем из восьми символов, используя рандом из всего
			for i := 1; i < length; i++ {
				password[i] = all[rand.Intn(len(all))]

			}
			//Перемешиваем элементы среза
			rand.Shuffle(len(password), func(i int, j int) {
				password[i], password[j] = password[j], password[i]
			})
			str := string(password)
			fmt.Println(str)
			fmt.Println("Пароль ненадёжный")
			break
		}
		if symbols == "2" && specials == "1" && big == "1" {
			rand.Seed(time.Now().UnixNano())
			//символы
			symbol := "!@#$%^&*"
			all := symbol + "QWERTYUIOPASDFGHJKLZXCVBNM" + "qwertyuiopasdfghjklzxcvbnm" + "ЙЦУКЕНГШЩЗФЫВАПРОЛДЖЭЯЧСМИТЬ" + "йцукенгшщзхъфывапролджэячсмитьбю"
			var length int
			fmt.Println("Введите длину желаемого пароля:")
			fmt.Scan(&length)
			//массив с символами
			password := make([]byte, length)
			//Второй член - рандомный символ
			password[1] = symbol[rand.Intn(len(symbol))]
			//генерируем хуйню не менее чем из восьми символов, используя рандом из всего
			for i := 1; i < length; i++ {
				password[i] = all[rand.Intn(len(all))]

			}
			//Перемешиваем элементы среза
			rand.Shuffle(len(password), func(i int, j int) {
				password[i], password[j] = password[j], password[i]
			})
			str := string(password)
			fmt.Println(str)
			break
		}
		if symbols == "2" && specials == "2" && big == "1" {
			rand.Seed(time.Now().UnixNano())
			//символы
			all := "QWERTYUIOPASDFGHJKLZXCVBNM" + "qwertyuiopasdfghjklzxcvbnm" + "ЙЦУКЕНГШЩЗФЫВАПРОЛДЖЭЯЧСМИТЬ" + "йцукенгшщзхъфывапролджэячсмитьбю"
			var length int
			fmt.Println("Введите длину желаемого пароля:")
			fmt.Scan(&length)
			//массив с символами
			password := make([]byte, length)
			//генерируем хуйню не менее чем из восьми символов, используя рандом из всего
			for i := 1; i < length; i++ {
				password[i] = all[rand.Intn(len(all))]

			}
			//Перемешиваем элементы среза
			rand.Shuffle(len(password), func(i int, j int) {
				password[i], password[j] = password[j], password[i]
			})
			str := string(password)
			fmt.Println(str)
			break
		}
		if symbols == "2" && specials == "1" && big == "2" {
			rand.Seed(time.Now().UnixNano())
			//символы
			symbol := "!@#$%^&*"
			all := symbol + "qwertyuiopasdfghjklzxcvbnm" + "йцукенгшщзхъфывапролджэячсмитьбю"
			var length int
			fmt.Println("Введите длину желаемого пароля:")
			fmt.Scan(&length)
			//массив с символами
			password := make([]byte, length)
			//Второй член - рандомный символ
			password[1] = symbol[rand.Intn(len(symbol))]
			//генерируем хуйню не менее чем из восьми символов, используя рандом из всего
			for i := 1; i < length; i++ {
				password[i] = all[rand.Intn(len(all))]

			}
			//Перемешиваем элементы среза
			rand.Shuffle(len(password), func(i int, j int) {
				password[i], password[j] = password[j], password[i]
			})
			str := string(password)
			fmt.Println(str)
			break
		}
		if symbols == "1" && specials == "1" && big == "2" {
			rand.Seed(time.Now().UnixNano())
			//символы
			numbers := "0123456789"
			symbol := "!@#$%^&*"
			all := numbers + symbol + "qwertyuiopasdfghjklzxcvbnm" + "йцукенгшщзхъфывапролджэячсмитьбю"
			var length int
			fmt.Println("Введите длину желаемого пароля:")
			fmt.Scan(&length)
			//массив с символами
			password := make([]byte, length)
			//Первый член - рандомная цифра
			password[0] = numbers[rand.Intn(len(numbers))]
			//Второй член - рандомный символ
			password[1] = symbol[rand.Intn(len(symbol))]
			//генерируем хуйню не менее чем из восьми символов, используя рандом из всего
			for i := 1; i < length; i++ {
				password[i] = all[rand.Intn(len(all))]

			}
			//Перемешиваем элементы среза
			rand.Shuffle(len(password), func(i int, j int) {
				password[i], password[j] = password[j], password[i]
			})
			str := string(password)
			fmt.Println(str)
			break
		}
		if symbols == "2" && specials == "1" && big == "2" {
			rand.Seed(time.Now().UnixNano())
			//символы
			symbol := "!@#$%^&*"
			all := symbol + "QWERTYUIOPASDFGHJKLZXCVBNM" + "qwertyuiopasdfghjklzxcvbnm" + "ЙЦУКЕНГШЩЗФЫВАПРОЛДЖЭЯЧСМИТЬ" + "йцукенгшщзхъфывапролджэячсмитьбю"
			var length int
			fmt.Println("Введите длину желаемого пароля:")
			fmt.Scan(&length)
			//массив с символами
			password := make([]byte, length)
			//Второй член - рандомный символ
			password[1] = symbol[rand.Intn(len(symbol))]
			//генерируем хуйню не менее чем из восьми символов, используя рандом из всего
			for i := 1; i < length; i++ {
				password[i] = all[rand.Intn(len(all))]

			}
			//Перемешиваем элементы среза
			rand.Shuffle(len(password), func(i int, j int) {
				password[i], password[j] = password[j], password[i]
			})
			str := string(password)
			fmt.Println(str)
			break
		}
	}
}
