package main

import (
	"fmt"
	"unsafe"
)
// Пример композиции

func NewAction(human Human, description string) *Action {
	return &Action{humInfo: human, SomeActionsDescription: description}
}

type Human struct { // Родительский класс
	Parents []string
	Age     int
	Sex     string
}

type Action struct {
	humInfo                Human // Объект родительского класса. Это поле недоступно для изменения извне(просто для примера)
	SomeActionsDescription string
}

type Group struct {
	PeopleInvolved []Human
}

type A struct {
	People Group
}

type B struct {
	People *Group
}

func (action *Action) GetHumanInfo() Human { // Получение копии родительского объекта
	return action.humInfo
}

func main() {
	action := NewAction(Human{Age: 18}, "NOTHING") // Создание объекта action

	fmt.Println(action)

	count := 20000
	peopleInvolved := make([]Human, count) // Создаем 20000 человек
	for i := range peopleInvolved {
		peopleInvolved[i] = Human{
			Parents: nil,
			Age:     i,
			Sex:     "male",
		}
	}

	group := Group{
		PeopleInvolved: peopleInvolved,
	}

	fmt.Println(unsafe.Sizeof(A{People: group}))
	fmt.Println(unsafe.Sizeof(B{People: &group})) // Размер ссылки в GO

}

//func main() {
//	init := []byte("package main\n\nfunc main() {}")
//	for i := 2; i < 35; i++ {
//		fileName := "./task_"+strconv.Itoa(i)+".go"
//		f, err := os.Create(fileName)
//		if err != nil {
//			log.Fatalf(err.Error())
//		}
//		err = ioutil.WriteFile(fileName, init, 0644)
//		if err != nil {
//			log.Fatalf(err.Error())
//		}
//		f.Close()
//	}
//}
