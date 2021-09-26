package main

import (
	"fmt"
	"os"
	"encoding/json"
	"io/ioutil"
)


func Save_map(file_name string, data map[string]string) {
	/*
		Сохранение данных из словаря в файл
	*/
	data_out, err := json.Marshal(data)
	if err != nil {
		fmt.Println(err)
	}
		
	err = ioutil.WriteFile(file_name, data_out, 0644)
	if err != nil {
		fmt.Println(err)
	}
}

func Load_map(file_name string) map[string]string {
	/*
		Загрузка данных из файла в словарь
	*/
	var data map[string]string
	
	file, err := ioutil.ReadFile(file_name)
	if err != nil {
		fmt.Println(err)
	}
		
	err = json.Unmarshal(file, &data)
	if err != nil {
        fmt.Println(err)
    }
	
	data = CheckFullWords(data)

	return data
}


func Create(data *map[string]string, key_word string, translate_word string) {
	/*
		Создаем новый перевод слова и создаем также обратный перевод
	*/
	(*data)[key_word] = translate_word
	
	_, ok := (*data)[translate_word]
	if !ok{
		(*data)[translate_word] = key_word
	}	
}

func Update(data *map[string]string, key_word string, new_translate_word string) {
	/*
		Записываем новый перевод слова, обратный перевод слова и удаляем старый перевод
	*/
	val, ok := (*data)[key_word]
	(*data)[key_word] = new_translate_word
	
	if ok{
		_, ok := (*data)[val]
		if ok {
			delete((*data), val)
		}
	}
	
	
	_, ok = (*data)[new_translate_word]
	if !ok{
		(*data)[new_translate_word] = key_word
	}
}

func Delete(data *map[string]string, key_word string) {
	/*
		Удаляем перевод слова и обратный перевод
	*/
	val, ok := (*data)[key_word]
	
	if ok{
		delete((*data), key_word)
	}
	
	_, ok = (*data)[val]
	if ok{
		delete((*data), val)
	}
}

func Translate(data map[string]string, key_word string) {
	/*
		Переводим слово, которое ввел пользователь
	*/
	_, ok := data[key_word]
	if ok{
		fmt.Println("Перевод слова: ", key_word, " -> ",  data[key_word])
	} else {
		fmt.Println("Данное слово отсутствует в словаре!")
	}
}


func CheckFullWords(data map[string]string) map[string]string {
	/*
	Проверка полной связанности перевода слов на разных языках
	Проверка того, что перевод слова и обратный перевод есть в словаре
	*/
	var data_out map[string]string
	data_out = make(map[string]string)
	
	for key, value := range data{
		data_out[value] = key
		data_out[key] = value
	}
	
	return data_out
}

func main(){
	
	/*
	Русский
	
	Данная программа "az" - это консольное приложение по переводу слов с любого языка на любой другой язык.
	
	Основа программы - словарь(map) в котором храняться ключ = слово + значение = перевод слова.
	Во время загрузки программа автоматически проверяет существование прямого и обратного значения слова -> его перевода
	Т.к. эта программа часть обучения новым языкам, в ней не предусмотрен базовый словарь. Все слова и их переводы, вам
	придется добавить самостоятельно.
	
	Удачи в изучении новых языков!!
	
	English
	
	This program "az" is a console application for translating words from any language to any other language.

	The basis of the program is a dictionary (map) in which the key = word + value = word translation is stored.
	During the boot program, it automatically checks the existence of the forward and reverse meanings of the word -> its translation
	Because this program is part of learning new languages, it does not provide a basic vocabulary. All words and their translations, you
	have to add it yourself.

	Good luck learning new languages !!
	
	Git: https://github.com/ATWp/TranslateDictAZ
	*/
	
	
	//Версия программы
	Version := "1.0"

	file := "dict.json"
	data := Load_map(file)
	
	arg := os.Args
	ln_arg := len(arg)
	
	if ln_arg >= 3 {
		command := arg[1]
		key_word := arg[2]
		//fmt.Println("command: ", command)
		fmt.Println("key_word: ", key_word)
		
		switch command{
		case "create":
			if ln_arg != 4 {
				fmt.Println("Недостаточно параметров для создания нового перевода")
			}else{
				translate_word := arg[3]
				fmt.Println("translate_word: ", translate_word)
				Create(&data, key_word, translate_word)
			}
			fmt.Println("command: create, doing: create")
			
		case "update":
			if ln_arg != 4 {
				fmt.Println("Недостаточно параметров для обновления перевода")
			}else{
				translate_word := arg[3]
				fmt.Println("translate_word: ", translate_word)
				Update(&data, key_word, translate_word)
			}
			fmt.Println("command: update, doing: update")
			
		case "translate":
			fmt.Println("command: translate, doing: translate")
			Translate(data, key_word)
			
		case "cr":
			if ln_arg != 4 {
				fmt.Println("Недостаточно параметров для создания нового перевода")
			}else{
				translate_word := arg[3]
				fmt.Println("translate_word: ", translate_word)
				Create(&data, key_word, translate_word)
			}
			fmt.Println("command: cr, doing: create")
			
		case "upd":
			if ln_arg != 4 {
				fmt.Println("Недостаточно параметров для обновления перевода")
			}else{
				translate_word := arg[3]
				fmt.Println("translate_word: ", translate_word)
				Update(&data, key_word, translate_word)
			}
			fmt.Println("command: upd, doing: update")
			
		case "tr":
			fmt.Println("command: tr, doing: translate")
			Translate(data, key_word)
			
		case "look":
			fmt.Println("command: look, doing: translate")
			Translate(data, key_word)
			
		case "delete":
			fmt.Println("command: delete, doing: delete")	
			Delete(&data, key_word)
			
		case "del":
			fmt.Println("command: del, doing: delete")
			Delete(&data, key_word)
			
		case "dl":
			fmt.Println("command: dl, doing: delete")		
			Delete(&data, key_word)
			
		case "version":
			fmt.Println("command: delete, doing: version")	
			fmt.Println("version: ", Version)	
			
		case "ver":
			fmt.Println("command: del, doing: version")
			fmt.Println("version: ", Version)
			
		case "v":
			fmt.Println("command: v, doing: version")		
			fmt.Println("version: ", Version)
			
		default:
			fmt.Println("Error: unknown command!")
			fmt.Println("Ошибка: Неопознанная команда!")
		
		}
	}else{
		if ln_arg == 2{
			command := arg[1]
			if command == "help"{
			
				fmt.Print("Все команды:\n commands:[\n \tcreate,\n \tcr,\n]\n doings: create\n --------\n")
				fmt.Print("commands:[\n \tupdate,\n \tupd,\n]\n doings: update\n --------\n")
				fmt.Print("commands:[\n \ttranslate,\n \ttr,\n \tlook,\n]\n doings: translate\n --------\n")
				fmt.Print("commands:[\n \tdelete,\n \tdel,\n \tdl,\n]\n doings: delete\n --------\n")
				fmt.Print("commands:[\n \tversion,\n \tver,\n \tv,\n]\n doings: version\n --------\n")
				
			}else{
				fmt.Println("Error: unknown command!")
				fmt.Println("Ошибка: Неопознанная команда!")
			}			
		}
	}

	Save_map(file, data)
}