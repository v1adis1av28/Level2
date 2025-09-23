package tcp

import (
	"bufio"
	"fmt"
	"io"
	"log"
	"net"
	"os"
	"sync"

	"github.com/v1adis1av28/level2/tasks/task17/internal/config"
)

func SetConnection(cfg *config.Config) (net.Conn, error) {
	address := fmt.Sprintf("%s:%d", cfg.Host, cfg.Port)

	conn, err := net.DialTimeout("tcp", address, cfg.TimeOut)
	if err != nil {
		return nil, fmt.Errorf("connection error: %w", err)
	}

	fmt.Printf("connected to %s\n", address)
	return conn, nil
}

func Writer(conn net.Conn, exit chan os.Signal, wg *sync.WaitGroup) {
	defer wg.Done()
	defer func() {
		fmt.Println("Writer stopped")
	}()

	scanner := bufio.NewScanner(os.Stdin)

	for scanner.Scan() {
		text := scanner.Text()
		text += "\n"

		_, err := conn.Write([]byte(text))
		if err != nil {
			log.Printf("error: %v", err)
			exit <- os.Interrupt
			return
		}
	}

	if err := scanner.Err(); err != nil { // in this case error acused while writing
		log.Printf("scanner error: %v", err)
	} else { // user typed Ctrl + d
		fmt.Println("\nclosing connection")
	}
	exit <- os.Interrupt
}

func Reader(conn net.Conn, exit chan os.Signal, wg *sync.WaitGroup) {
	defer wg.Done()
	defer func() {
		fmt.Println("Reader stopped")
	}()

	reader := bufio.NewReader(conn)

	for {
		select {
		case <-exit: // in case we can read from exit channel that means programm is stoping
			return
		default: // in other cases reading strings
			line, err := reader.ReadString('\n')
			if err != nil {
				if err == io.EOF {
					fmt.Println("\nServer closed the connection")
					exit <- os.Interrupt
					return
				}
				log.Printf("Read error: %v", err)
				exit <- os.Interrupt
				return
			}
			fmt.Print(line)
		}
	}
}
