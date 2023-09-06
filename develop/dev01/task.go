package main

import (
	"fmt"
	"github.com/beevik/ntp"
	"log"
	"os"
)

/*
=== Базовая задача ===

Создать программу печатающую точное время с использованием NTP библиотеки.Инициализировать как go module.
Использовать библиотеку https://github.com/beevik/ntp.
Написать программу печатающую текущее время / точное время с использованием этой библиотеки.

Программа должна быть оформлена с использованием как go module.
Программа должна корректно обрабатывать ошибки библиотеки: распечатывать их в STDERR и возвращать ненулевой код выхода в OS.
Программа должна проходить проверки go vet и golint.
*/

func main() {
	response, err := ntp.Time("0.beevik-ntp.pool.ntp.org") //делаем запрос к ntp серверу.
	if err != nil {
		_, errWr := fmt.Fprintln(os.Stderr, err) //пишем ошибку в stderr.
		if errWr != nil {
			log.Fatalln(errWr) //при ошибке записи в файл обрабатываем ошибку уже встроенным методом, log.Fatal() пишет ошибку в логи уровня ERROR(stderr) и выходит с кодом 1.
		}
		os.Exit(1) //выходим из программы с кодом 1.
	}

	fmt.Println(response)
}

//TODO make module
