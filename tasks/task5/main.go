package main

type customError struct {
	msg string
}

func (e *customError) Error() string {
	return e.msg
}

func test() *customError {
	// ... do something
	// для примера можно задавать не только тип, но и указывать значение полю msg
	// return &customError{msg : "some custom error"}
	// и после вызова ф-и test() в мейне мы сможем получить более точную информацию о ошибке сделав вызов err.Error()
	return nil
}

func main() {
	var err error // Объявляем переменную интерфейсного типа error, чтобы он стал таковым типом ему
	// необходимо имплементировать предписанные интерфейсом методы(Error() string)
	err = test()    //здесь упаковываем значение интерфейсного типа error к кастомному customError, поэтому теперь у интерфейса тип customError а значение nil
	if err != nil { // чтобы переменная интерфейсного типа являлась nil, нужно чтобы и тип и значение были равны nil, однако в нашем случае type = customError value = nil
		println("error") //Поэтому вывод эту строку и сделает return
		return
	}
	println("ok") // Выводилась бы эта строка если бы мы не кастили тип, вызывая функцию test()
}
