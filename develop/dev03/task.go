package main

import (
	"bufio"
	"flag"
	"fmt"
	"log"
	"os"
	"regexp"
	"sort"
	"strconv"
	"strings"
	"unicode"
)

//Отсортировать строки в файле по аналогии с консольной утилитой sort (man sort — смотрим описание и основные параметры): на входе подается файл из несортированными строками, на выходе — файл с отсортированными.
//
//Реализовать поддержку утилитой следующих ключей:
//
//-k — указание колонки для сортировки (слова в строке могут выступать в качестве колонок, по умолчанию разделитель — пробел)
//-n — сортировать по числовому значению
//-r — сортировать в обратном порядке
//-u — не выводить повторяющиеся строки
//
//Дополнительно
//
//Реализовать поддержку утилитой следующих ключей:
//
//-M — сортировать по названию месяца
//-b — игнорировать хвостовые пробелы
//-c — проверять отсортированы ли данные
//-h — сортировать по числовому значению с учетом суффиксов

// Примечание: не реализована возможность сортировать наименования месяцев с суффиксами, поэтому сортировку по колонке и по месяцем нагляднее делать в разных файлах примерах(example.txt и example_month.txt). При этом сортировка по месяцам в example.txt отработает верно за исключением названия месяца с суфиксом.
// run config example:
// go run task.go -M example_month.txt
// go run task.go -r example.txt
// go run task.go -h example.txt
// go run task.go -c example.txt

type input struct {
	data []string
	k    int
}

func newInput(data []string, k int) *input {
	return &input{
		data: data,
		k:    k,
	}
}

func (i *input) sort() { //метод сортировки строк без параметров.
	sort.Strings(i.data)
}

func (i *input) getUniqueStr() { //метод возврата уникальных строк.
	unicalMap := make(map[string]bool, 0) //инициализируем карту.
	for _, val := range i.data {          //циклом проходим по слайсу строк.
		_, ok := unicalMap[val] //проверяем по ключу есть ли в карте данная строка.
		if !ok {
			unicalMap[val] = true //если нет, записываем строку как ключ.
		}
	}
	result := make([]string, 0) //инициализируем слайс для записи только уникальных строк.
	for k := range unicalMap {  //итерируемся по карте.
		result = append(result, k) //добавляем ключи из карты в слайс.
	}
	i.data = result
}

func (i *input) sortNum() { //метод сортировки строк из файла по числовому значению.
	sort.Slice(i.data, func(x, j int) bool {
		vx, _ := strconv.Atoi(strings.Fields(i.data[x])[0]) //в качестве аргумента функции используется число без суффикса
		vj, _ := strconv.Atoi(strings.Fields(i.data[j])[0]) //то есть сортировка будет проведена только по первому числу до пробела без учета чисел с префиксами и суффиксами.
		return vx < vj
	})
}

func (i *input) sortByColumn() { //метод сортировки строк из файла по колонке.
	table := make([][]string, 0) //инициализируем слайс слайсов строк.
	for _, val := range i.data { //итерируемся по строкам из файла.
		table = append(table, strings.Split(val, " ")) //каждую строку файла разбиваем по пробелу и получившийся слайс строк кладем в table.
	}
	sort.Slice(table, func(x, j int) bool { //сортируем слайс слайсов по k-элементу.
		return table[x][i.k-1] < table[j][i.k-1]
	})

	for idx, val := range table { //итерируемся по полученному слайсу слайсов.
		var str string
		for _, value := range val { //итерируемся слайсу (строке файла).
			str += value + " " //собираем слайс в одну строку.
		}
		str = strings.TrimRight(str, " ") //обрезаем пробел справа.

		i.data[idx] = str
	}
}

func (i *input) reversSort() { //метод реверса данных.
	reverse := make([]string, 0)            //инициализируем слайс.
	for y := len(i.data) - 1; y >= 0; y-- { //пишем в новый слайс все данные в обратном порядке.
		reverse = append(reverse, i.data[y])
	}
	i.data = reverse
}

func (i *input) sortMonth() { //метод сортировки по месяцам
	date := make(map[string]string)   //инициализируем карту для дат и месяцев(ключ изначальная строка, значение номер или имя месяца).
	slDateString := make([]string, 0) //инициализируем слайс строк с датами и месяцами.
	slString := make([]string, 0)     //инициализируем слайс строк без дат и месяцев.
	monthS := [12]string{"jan", "feb", "mar", "apr", "may", "jun", "jul", "aug", "sep", "oct", "nov", "dec"}
	nameMonth := [12]string{"january", "february", "march", "april", "may", "june", "july", "august", "september", "october", "november", "december"}
	for _, val := range i.data {
		reDate, _ := regexp.Compile(`(\d{4})[./-](\d{2})[./-](\d{2})`) //регистрируем регулярное выражение.
		resDate := reDate.FindAllStringSubmatch(val, -1)               //ищем регулярное выражение в строке.
		reName, _ := regexp.Compile(`(\d{4})[./-](\w{3})[./-](\d{2})`) //регистрируем регулярное выражение.
		resName := reName.FindAllStringSubmatch(val, -1)               //ищем регулярное выражение в строке.
		for m := 0; m < 12; m++ {
			if resDate == nil && resName == nil && strings.ToLower(val) != monthS[m] && strings.ToLower(val) != nameMonth[m] {
				slString = append(slString, val) //строки без дат и месяцев пишем в отдельный массив.
				break
			}
		}
		for _, v := range resDate { //итерируемся по найденным значениям.
			date[val] = v[2] //пишем в карту ключ - изначальная строка, значение - значение с индексом месяца.
		}
		for _, v := range resName { //итерируемся по найденным значениям.
			date[val] = v[2] //пишем в карту ключ - изначальная строка, значение - значение с индексом месяца.
		}
		for m := 0; m < 12; m++ { //итерируемся 12 раз по количеству месяцев.
			if strings.ToLower(val) == monthS[m] || strings.ToLower(val) == nameMonth[m] { //проверяем, что значение является месяцем.
				date[val] = val //если да, пишем в карту.
			}
		}
	}
	for j := 1; j <= 12; j++ { //итерируемся 12 раз, это позволит отсоритровать по месяцам.
		for key, val := range date { //итерируемся по карте.
			var name string
			num, err := strconv.Atoi(val) //преобразуем значение из карты в число.
			if err != nil {
				name = strings.ToLower(val) //если это не число то записываем его в переменную name в нижнем регистре.
			}
			if num == j || name == monthS[j-1] { //если значение из карты равно итерации
				slDateString = append(slDateString, key) //то пишем ключ из карты в слайс строк с датами
			}
			if strings.ToLower(key) == monthS[j-1] || strings.ToLower(key) == nameMonth[j-1] { //если ключ является месяцем, то пишем его в слайс строк с датами.
				slDateString = append(slDateString, key)
			}
		}
	}
	i.data = append(slDateString, slString...) //объединяем слайс с датами и без них.

}

func (i *input) ignoreSpace() { //метод игнорирования хвостовых пробелов.
	table := make([][]string, 0) //инициализируем слайс слайсов("таблицу").
	for _, val := range i.data { //итерируемся по строкам из файла.
		table = append(table, strings.Fields(val)) //добавляем в "таблицу" слайс из строк каждой строки файла, используя функцию strings.Fields().
	} //функция strings.Fields() послужит средством удаления любых хвостовых пробелов.

	for idx, val := range table { //итерируемся по "таблице".
		var str string
		for _, value := range val { //собираем строку из одного слайса таблицы.
			str += value + " "
		}
		i.data[idx] = str
	}
}

func (i *input) checkSort() error { //метод проверки сортировки файла.
	for x := 0; x < len(i.data)-1; x++ {
		if i.data[x+1] < i.data[x] {
			return fmt.Errorf("неправильный порядок: %s", i.data[x+1])
		} //построчно проверяем сортировку строк в файле
	}
	return nil
}

func (i *input) sortNumWithS() { //метод сортировки строк файла с учетом суффиксов.
	numValues := make([][]string, 0) //инициализируем слайс слайсов, послужит аналогом карты, где ключи могут повторяться("карта").
	arrString := make([]string, 0)   //инициализируем слайс строк.
	arrInt := make([]string, 0)      //инициализируем слайс строк с числами.
	arrKey := make([]int, 0)         //инициализируем слайс чисел.
	for _, val := range i.data {     //итерируемся по строкам из файла.
		for index, el := range []rune(val) { //итерируемся по руническому представлению элементов строки.
			if unicode.IsDigit(el) { //если элемент является десятичным числом, приступаем к следующей итерации.
				if index == len([]rune(val))-1 { //если элемент является последним в строке, то пишем строку в аналог "карты" и по ключу и по значению.
					numValues = append(numValues, []string{val, val})
					break
				}
				continue
			} else if index != 0 { //если индекс не нулевой, то однозначно, что строка, либо целое число, либо число с суффиксом.
				numValues = append(numValues, []string{val[:index], val}) //пишем строку в значение "карты", число строки в ключ "карты".
				break
			} else {
				arrString = append(arrString, val) //если строка не числовое значение, пишем ее в слайс строк.
				break
			}
		}
	}

	for _, val := range numValues { //итерируемся по "карте"
		key, err := strconv.Atoi(val[0]) //преобразуем ключи "карты" в числа.
		if err != nil {
			log.Println(err)
		}
		arrKey = append(arrKey, key) //и пишем их в слайс ключей.
	}
	sort.Ints(arrKey)            //сортируем слайс ключей.
	for _, val := range arrKey { //итерируемся по слайсу ключей.
		for idx, sl := range numValues { //итерируемся по "карте".
			if strconv.Itoa(val) == sl[0] { //проверяем условие равенства ключа с ключом из "карты".
				arrInt = append(arrInt, sl[1])                            //если условие верно пишем значение из "карты" в слайс чисел.
				numValues = append(numValues[:idx], numValues[idx+1:]...) //удаляем данную пару ключ-значение из "карты" для избежания повторов.
				break
			}
		}
	}
	i.data = append(arrInt, arrString...) //объединяем в результате слайсы строк и чисел.

}

func main() {
	//инициализируем поддерживаемые флаги.
	k := flag.Int("k", 1, "колонка для сортировки")
	n := flag.Bool("n", false, "сортировка по числовому значению")
	r := flag.Bool("r", false, "сортировка в обратном порядке")
	u := flag.Bool("u", false, "не выводить повторяющиеся строки")
	m := flag.Bool("M", false, "сортировать по названию месяца")
	b := flag.Bool("b", false, "игнорировать хвостовые пробелы")
	c := flag.Bool("c", false, "проверять отсортированы ли данные")
	h := flag.Bool("h", false, "сортировать по числовому значению с учетом суффиксов")

	flag.Parse()

	file := flag.Arg(0) //считываем аргумент командной строки, после обработки флагов.
	if file == "" {
		log.Fatalln("не указано имя файла")
	}
	data := readData(file)
	input := newInput(data, *k)

	if *c {
		err := input.checkSort()
		if err != nil {
			fmt.Println(err)
			return
		}
	}
	if *b {
		input.ignoreSpace()
		input.sort()
	}
	if *u {
		input.getUniqueStr()
		input.sort()
	}
	if *k != 0 && *k != 1 {
		input.sortByColumn()
	}
	if *n {
		input.sortNum()
	}
	if *m {
		input.sortMonth()
	}
	if *h {
		input.sortNumWithS()
	}
	if !*b && !*u && *k == 1 && !*n && !*m && !*h {
		input.sort()
	}
	if *r {
		input.reversSort()
	}
	for _, val := range input.data {
		fmt.Println(val)
	}
}

func readData(file string) []string {
	dataFile := make([]string, 0)
	f, err := os.Open(file)
	if err != nil {
		log.Fatalln(err)
	}
	defer f.Close()

	scan := bufio.NewScanner(f)

	for scan.Scan() {
		dataFile = append(dataFile, scan.Text())
	}

	return dataFile
}
