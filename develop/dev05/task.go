package main

import (
	"bufio"
	"errors"
	"flag"
	"fmt"
	"log"
	"os"
	"regexp"
	"strconv"
	"strings"
)

//Реализовать утилиту фильтрации по аналогии с консольной утилитой (man grep — смотрим описание и основные параметры).
//
//Реализовать поддержку утилитой следующих ключей:
//-A - "after" печатать +N строк после совпадения
//-B - "before" печатать +N строк до совпадения
//-C - "context" (A+B) печатать ±N строк вокруг совпадения
//-c - "count" (количество строк)
//-i - "ignore-case" (игнорировать регистр)
//-v - "invert" (вместо совпадения, исключать)
//-F - "fixed", точное совпадение со строкой, не паттерн
//-n - "line num", напечатать номер строки

type flags struct {
	A int
	B int
	C int
	c bool
	i bool
	v bool
	F bool
	n bool
}

func NewFlags() *flags {
	return &flags{
		A: *flag.Int("A", 0, "печатать +N строк после совпадения"),
		B: *flag.Int("B", 0, "печатать +N строк до совпадения"),
		C: *flag.Int("C", 0, "(A+B) печатать ±N строк вокруг совпадения"),
		c: *flag.Bool("c", false, "(количество строк)"),
		i: *flag.Bool("i", false, "(игнорировать регистр)"),
		v: *flag.Bool("v", false, "(вместо совпадения, исключать)"),
		F: *flag.Bool("F", false, "точное совпадение со строкой, не паттерн"),
		n: *flag.Bool("n", false, "напечатать номер строки"),
	}
}

func grep(F flags) {
	patt, data := getPatternAndSource(F) //получаем образец и данные из файла.

	a, b, intRes, lineRes, err := checkFlagsAndDo(F, patt, data)
	if err != nil {
		return
	} else {
		sum := 0
		for _, val := range intRes { //итерируемся по слайсу индексов строк с совпадениями.
			sum = min(b, val) //определяем меньшее из значение(количество до совпадение строк или индекс строки).
			for sum != 0 {
				lineRes = append(lineRes, data[val-sum]) //добавляем в слайс строк на вывод строки до совпадения строк.
				sum--
			}
			lineRes = append(lineRes, data[val]) //добавляем в слайс строк на вывод строки с совпадениями.
			sum = 0
			for sum != min(a, len(data)-val-1) {
				lineRes = append(lineRes, data[val+sum+1]) //добавляем в слайс строк на вывод строки после совпадения строк.
				sum++
			}
		}
		fmt.Println(strings.Join(lineRes, "\n")) //пишем строки в стандартный поток вывода.
	}
}

func getPatternAndSource(F flags) (string, []string) {
	patt := flag.Arg(0)            //получаем первый аргумент(образец) командной строки, после обработки флагов.
	fileName := flag.Arg(1)        //получаем второй аргумент(имя файла) командной строки, после обработки флагов.
	data := make([]string, 0)      //инициализируем слайс, куда будем писать данные из файла.
	file, err := os.Open(fileName) //открываем полученный файл.
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()
	sc := bufio.NewScanner(file) //создаем новый сканнер.
	for sc.Scan() {              //построчно сканируем данные из файлаи пишем их в слайс строк.
		if F.i {
			data = append(data, strings.ToLower(sc.Text()))
		} else {
			data = append(data, sc.Text())
		}
	}
	return patt, data
}

func checkFlagsAndDo(F flags, patt string, data []string) (int, int, []int, []string, error) {
	if F.F { //описываем регулярное выражение(образец) в зависимости от полученных флагов.
		patt = `\Q` + patt + `\E`
	}
	if F.i {
		patt = `(?i)` + patt
	}
	reg := regexp.MustCompile(patt) //инициализируем регулярное выражение
	a := max(F.A, F.C)              //определяем сколько строк писать после совпадения строк.
	b := max(F.B, F.C)              //определяем сколько строк писать до совпадения строк.
	intRes := make([]int, 0)        //инициализируем слайс для записи индексов строк с совпадением.
	for i, val := range data {      //итерируемся по слайсу строк из файла.
		if reg.Match([]byte(val)) { //ищем совпадения значения с регулярным выражением.
			intRes = append(intRes, i) //пишем индексы строк с совпадениями в слайс.
		}
	}

	if F.n { //приполучении флага добавляем в вывод номер строки.
		for i, val := range data { //итерируемся по слайсу строк из файла.
			val = strconv.Itoa(i) + " " + val //добавляем номер строк.
		}
	}
	lineRes := make([]string, 0) //инициализируем слайс строк для записи строк на вывод.
	if F.v {                     //при получении флага удаляем из результата строки, содержащие образец.
		for i, val := range data { //итерируемся по слайсу строк из файла.
			if findElem(intRes, i) { //сравниваем индексы строк с образцами и пропускаем их в случае совпадения.
				continue
			}
			lineRes = append(lineRes, val) //добавляем в слайс строк на вывод строку из файла.
		}
		fmt.Println(strings.Join(lineRes, "\n")) //пишем в поток стандартного вывода строки.
		return 0, 0, intRes, lineRes, errors.New("end")
	}
	if F.c { //при получении флага также выводим количество выведенных строк.
		fmt.Println(len(intRes))
	}
	return a, b, intRes, lineRes, nil
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}

func min(a, b int) int {
	if a < b {
		return a
	}
	return b
}

func findElem(str []int, elem int) bool {
	for _, index := range str {
		if index == elem {
			return true
		}
	}
	return false
}

func main() {
	F := NewFlags()
	flag.Parse()
	grep(*F)
}
