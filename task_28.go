package main

type Сlient struct {  // Допустим купили телефон с алиэкспресса, а там китайская вилка
}

type Socket interface {  // Интерфейс розетки (eu), ему удовлетворяет только EU вилка
	Connect()
}

func (c *Сlient) ChargePhone(socket Socket) {
	socket.Connect()
}

type euCharger struct {
}

func (euCharger *euCharger) Connect() {
}

type chineseCharger struct{}

func (chCharger *chineseCharger) ChinesePlug() {
}

//  Создание структуры адаптера, позволяющей преобразовать нашу китайскую вилку к еврепейской
type chineseToEuAdapter struct {
	adapter *chineseCharger
}

// Метод, благодаря которому chineseToEuAdapter удовлетворяющий интерфейсу розетки
func (chToEu *chineseToEuAdapter) Connect() {
	chToEu.adapter.ChinesePlug()
}

func main() {

	client := &Сlient{}

	eu := &euCharger{}    // В данном случае никаких дополнительных вмешательств не нужно
	client.ChargePhone(eu)

	chinese := &chineseCharger{}
	chToEuAdapter := &chineseToEuAdapter{ // Теперь можем со спойной душой пихать в розетку
		adapter: chinese,
	}
	client.ChargePhone(chToEuAdapter)
}
