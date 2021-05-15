package sample

import "log"

// удаляем элемент по индексу
func RemoveAtIndexSlice(source []int, index int) []int {
	log.Println("удаляем элемент с индексом:")
	lastIndex := len(source) - 1
	//меняем последнее значение и значение, которое хотим удалить, местами
	log.Printf("Было так: %d", source) //
	source[index], source[lastIndex] = source[lastIndex], source[index]
	// обрезаем последнее значение
	source = source[:lastIndex]
	log.Printf("удалили элемент с индексом: %d получилось так: %d", index, source) // [1 2 5 4]
	return source
}
