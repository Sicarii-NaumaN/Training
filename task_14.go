package main

import (
	"fmt"
	"reflect"
)

// Инициализация из имеющегося среза
func NewSet(elems []string) *Set {
	check := make(map[interface{}]bool)
	setInit := make([]interface{}, 0)
	for i := range elems {
		if !check[elems[i]] {
			setInit = append(setInit, elems[i])
			check[elems[i]] = true
		}
	}
	return &Set{SetElements: setInit, isUniq: check}
}

// Само множество и таблица состояния
type Set struct {
	SetElements []interface{}
	isUniq      map[interface{}]bool
}

// Добавление в таблицу нового элемента
func (s *Set) Insert(elem interface{}) {
	if !s.isUniq[elem] {
		s.SetElements = append(s.SetElements, elem)
	}
}

// Печатаем только само множество
func (s *Set) Println() {
	fmt.Println(reflect.ValueOf(s).Elem().Field(0))
}

func main() {
	str := []string{"cat", "cat", "dog", "cat", "tree"}

	set := NewSet(str)
	set.Insert("cat")
	set.Insert("mouse")

	set.Println()
}
