package main

import (
	"fmt"
	"time"
	"math/rand"
	"log"
	
	"github.com/eiannone/keyboard"
)
// (1) Інформація
// (2) Чекаємо 3 секунди, виводимо букву і чекаємо знову 5 секунд
// (3) Ранжомізація

var alphabet = []string{
	"A", "B", "C", "D", "E", "F", "G", "H", "I", "J", "K", "L", "M", "N", "O", "P", "Q", "R", "S", "T", "U", "V", "W", "X", "Y", "Z",
	"0", "1", "2", "3", "4", "5", "6", "7", "8", "9",
}



func main(){
	fmt.Println("Вітаю вас у цій програмі. Мета програми — відстежити швидкість вашої реакції.")  
	fmt.Println("Я задам один символ і буду генерувати та виводити випадкові символи.")  
	fmt.Println("Ваше завдання — коли побачите символ, який я задав, натиснути клавішу любу зручну для вас клавішу.")

	fmt.Println("Оберіть швидкість:")
	fmt.Println("1 - Повільно")
	fmt.Println("2 - Середньо")
	fmt.Println("3 - Дуже швидко")
	fmt.Println("4 - Супер швидко")
	answer := "1";
	fmt.Scanln(&answer)

	rand.Seed(time.Now().UnixNano())
	letter := alphabet[rand.Intn(len(alphabet))]
	fmt.Println("Я загадав символ:", letter)
	fmt.Println("Подивіться на нього уважно, і через 4 секунди я почну генерувати символи.")
	time.Sleep(4 * time.Second)
	if err := keyboard.Open(); err != nil { // відкриваємо клавіатуру
		log.Fatal(err)
	}
	defer keyboard.Close()

	isWas := false
	isPress := false
	start := time.Now()
	go func(){
		for {
			_, _, err := keyboard.GetKey()
			if err != nil {
				log.Fatal("Програма завершила свою роботу!")
			}
			end := time.Now()
			if isWas {
				fmt.Println("Швидкість реакції:", end.Sub(start))
			} else {
				fmt.Println("Такого символа ще не було!")
			}
			time.Sleep(4 * time.Second)
			isPress = true
		}
	}()
	delay := 1 * time.Second;
	if answer == "2" {
		delay = 500 * time.Millisecond
	} else if answer == "3" {
		delay = 250 * time.Millisecond
	} else if answer == "4" {
		delay = 100 * time.Millisecond
	}
	for !isPress{
		// Виводимо символ або спеціальний код клавіші
		time.Sleep(delay)
		rand.Seed(time.Now().UnixNano())
		randLetter := alphabet[rand.Intn(len(alphabet))]
		if randLetter == letter {
			start = time.Now()
			isWas = true
		}
		if !isPress {
			fmt.Println(randLetter)
		}
	}
}