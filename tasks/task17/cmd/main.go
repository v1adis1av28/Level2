package main

import (
	"fmt"
	"log"
	"os"
	"os/signal"
	"sync"

	"github.com/v1adis1av28/level2/tasks/task17/internal/config"
	"github.com/v1adis1av28/level2/tasks/task17/internal/tcp"
)

//Реализовать простой telnet-клиент с возможностью соединяться к TCP-серверу и взаимодействовать с ним:

// Программа должна принимать параметры: хост, порт и опционально таймаут соединения
// (через флаг --timeout, по умолчанию 10 секунд).
// После запуска, telnet-клиент устанавливает TCP-соединение с указанным host:port.
// Все, что пользователь вводит в STDIN, должно отправляться в сокет;
// все, что приходит из сокета — печататься в STDOUT.
// При нажатии комбинации клавиш Ctrl+D клиент должен закрыть соединение и завершиться.
// Если сервер закрыл соединение, клиент тоже завершается.
// В случае, если попытка подключения не удалась (например, сервер недоступен)
// — программа завершается через заданный timeout с соответствующим сообщением об ошибке.
// Проверить программу можно, например, подключившись к какому-нибудь публичному echo-серверу или SMTP (порт: 25)
// и вручную отправляя команды.

// Обратите внимание на обработку буферов: желательно запускать чтение/запись в отдельных горутинах
//  (для конкурентного ввода/вывода). Код должен быть без гонок. Реализация данной утилиты подразумевает
//  использование пакета net (тип net.Conn), и возможно bufio для удобства чтения/записи.

func main() {
	cfg, err := config.NewConfig()
	if err != nil {
		log.Fatalf("conifg error: %v", err)
	}

	conn, err := tcp.SetConnection(cfg)
	if err != nil {
		log.Fatalf("error: %v", err)
	}
	defer conn.Close()

	fmt.Println("connection established. You can typing commands.")
	fmt.Println("Ctrl+D to exit.")

	exit := make(chan os.Signal, 1)
	signal.Notify(exit, os.Interrupt)
	var wg sync.WaitGroup
	wg.Add(2)
	go tcp.Writer(conn, exit, &wg)
	go tcp.Reader(conn, exit, &wg)
	<-exit
	conn.Close()
	wg.Wait()

}
