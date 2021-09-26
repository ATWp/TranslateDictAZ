# TranslateDictAZ #

Программа называется "az" не просто так. Основная идея идет из перечисления символов в ASCII кодировке с A->z. Из-за того, что
набирать название в консоли CapsLock or Shift + A (CapsLock) + z  и протеворечит методологии vim, первая буква была преобразована
в нижний регист: az
	
The program is called "az" for a reason. The basic idea comes from enumerating characters in ASCII encoding with A-> z. Because of
typing the name in the console CapsLock or Shift + A (CapsLock) + z and contradicts the vim methodology, the first letter was converted
to lowercase: az
	
## Русский ##
-------------------------- 
	
Данная программа "az" - это консольное приложение по переводу слов с любого языка на любой другой язык.
	
Основа программы - словарь(map) в котором храняться ключ = слово + значение = перевод слова.
Во время загрузки программа автоматически проверяет существование прямого и обратного значения слова -> его перевода
Т.к. эта программа часть обучения новым языкам, в ней не предусмотрен базовый словарь. Все слова и их переводы, вам
придется добавить самостоятельно.
	
Удачи в изучении новых языков!!
	
## English ##
-------------------------- 
	
This program "az" is a console application for translating words from any language to any other language.

The basis of the program is a dictionary (map) in which the key = word + value = word translation is stored.
During the boot program, it automatically checks the existence of the forward and reverse meanings of the word -> its translation
Because this program is part of learning new languages, it does not provide a basic vocabulary. All words and their translations, you
have to add it yourself.

Good luck learning new languages !!
	
## Stack(Программный стек) ##
-------------------------- 
	
1. **Golang** with libs
	a. os
	b. fmt
	c. encoding/json
	d. io/ioutil
		
## Install ##
-------------------------- 
Не забудьте добавить программу в Windows - "Переменные среды", чтобы можно было вызывать ее из любого положения.
	
1. Откройте проводник "win+e"
2. Нажмите правой кнопкой мыши на мой(этот) компьютер.
3. Выберите свойства
4. Нажмите "Дополнительные параметры системы"
5. Нажмите "Переменные среды"
6. Нажмите "Path"
7. Нажмите "Создать"
8. Пропишите путь к расположению "az.exe", который вы скомпилировали из "az.go"
9. Откройте консоль "win+r" -> cmd -> enter
10. Пропишите команду "az version" или "az ver" или "az v"
