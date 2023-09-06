package dev04

import (
	"sort"
	"strings"
)

func SearchAnagram(input []string) map[string][]string {
	anagram := make(map[string][]string, 0) //инициализируем карту анаграмм.
	for idx, val := range input {           //итерируемся по получаемому массиву.
		word := strings.ToLower(val) //переводим все  значения массива в нижний регистр.
		if idx == 0 {                //первый элемент массива сразу пишем в карту.
			anagram[word] = append(anagram[word], word)
			continue
		}

		runeWord := []rune(word)                   //инициализируем слайс рун слов из массива.
		sort.Slice(runeWord, func(i, j int) bool { //сортируем слайс рун.
			return runeWord[i] < runeWord[j]
		})

		fl := false                   //флаг для определения того было ли слово записано в карту как анаграмма.
		for key, _ := range anagram { //итерируемся по карте.
			runeKey := []rune(key)                    //инициализируем слайс рун ключей карты.
			sort.Slice(runeKey, func(i, j int) bool { //сортируем слайс рун
				return runeKey[i] < runeKey[j]
			})
			if equal(runeWord, runeKey) { //сравниваем отсортированные слайсы слова вводного массива и ключей из карты.
				anagram[key] = append(anagram[key], word) //при их равенстве считаем слово анаграммой и пишем в карту.
				fl = true
				break
			}
		}
		if fl == false {
			anagram[word] = append(anagram[word], word)
		}
	}
	del := make([]string, 0)          //инициализируем массив, куда запишем ключи карты, по которым отсутствуют множества.
	for key, value := range anagram { //итерируемся по карте анаграмм.
		if len(value) == 1 { //при отсутствии множества по ключу пишем ключ в массив.
			del = append(del, key)
		}
		sort.Strings(value) //сортируем все слова в множествах в карте.
		anagram[key] = value
	}
	for _, val := range del { //итерируемся по массиву.
		delete(anagram, val) //удаляем из карты элементы, где отсутсвуют множества анаграмм.
	}

	return anagram
}

func equal(a, b []rune) bool { //функция сравнения массивов.
	if len(a) != len(b) { //проверяем равна ли длина массивов.
		return false
	}
	for i, v := range a { //итерируемся по первому массиву.
		if v != b[i] { //сравниваем каждый i-тый элемент первого массива с каждым i-тым элементом второго массива.
			return false
		}
	}
	return true
}
