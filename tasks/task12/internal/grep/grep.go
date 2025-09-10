package grep

import (
	"bufio"
	"fmt"
	"os"

	"github.com/v1adis1av28/level2/tasks/task12/internal/config"
)

func Grep(c *config.Config) error {
	matcher, err := c.Matcher() //ф-я матчера для определения соотвветсвия строки паттерну
	if err != nil {
		return err
	}

	scanner := bufio.NewScanner(os.Stdin)
	firstMatch := true

	for scanner.Scan() {
		c.LineNum++
		line := scanner.Text()
		match := matcher(line) // проверяем соответствие
		//добавляем в буффер текущую строку отмечая ее соответсвтие
		c.Buffer = append(c.Buffer, config.Line{Num: c.LineNum, Text: line, Match: match})

		//если размер бефера превышает n строк до найденной то изменяем размер буффера
		if len(c.Buffer) > c.Flags.PreviousLineFlag+1 {
			c.Buffer = c.Buffer[1:]
		}

		//если мэтч увеличиваем соответсвие коунтов на +1 для флага с мэтчами
		if match {
			c.MatchesCount++

			if c.Flags.CountOfLineFlag {
				continue
			}
			//разделитель соответсвий
			if !firstMatch && c.LastMatchLine+c.Flags.AdditionalLineFlag+1 < c.LineNum {
				fmt.Println("--")
			}
			firstMatch = false
			c.LastMatchLine = c.LineNum

			for i := 0; i < len(c.Buffer)-1; i++ {
				printLine(c, c.Buffer[i], true)
			}
			printLine(c, c.Buffer[len(c.Buffer)-1], false)
			c.PostContext = c.Flags.AdditionalLineFlag

		} else if c.PostContext > 0 {
			printLine(c, config.Line{Num: c.LineNum, Text: line, Match: false}, true)
			c.PostContext--
		}
	}

	if err := scanner.Err(); err != nil {
		return err
	}

	if c.Flags.CountOfLineFlag {
		fmt.Println(c.MatchesCount)
	}

	return nil
}

func printLine(c *config.Config, line config.Line, isContext bool) {
	if c.Flags.CountOfLineFlag {
		return
	}

	prefix := ":"
	if isContext {
		prefix = "-"
	}

	if c.Flags.LineNumberFlag {
		fmt.Printf("%d%s%s\n", line.Num, prefix, line.Text)
	} else {
		fmt.Println(line.Text)
	}
}
