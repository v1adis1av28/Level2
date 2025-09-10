package config

import (
	"flag"
	"fmt"
	"regexp"
	"strings"
)

type Config struct {
	Pattern string // это либо строка или регулярка
	Flags   *Flags

	LineNum       int
	MatchesCount  int
	PostContext   int
	LastMatchLine int
	Buffer        []Line
}

type Line struct {
	Num   int
	Text  string
	Match bool
}

type Flags struct {
	IgnoreFlag       bool // -i игнорировать регистр
	InvertFlag       bool // -v инвертирующий фильтр
	StrictStringFlag bool // -F шаблон как фиксированная строка
	LineNumberFlag   bool // -n выводить номер строки пере каждой строкой
	CountOfLineFlag  bool // -c выводить только то кол-во строк совпадающих с паттерном

	AdditionalLineFlag int // -A N после каждой найденной строки дополнительно вывести N строк после неё (контекст).
	PreviousLineFlag   int //	-B N — вывести N строк до каждой найденной строки.
	AroundLineFlag     int // -C N — вывести N строк контекста вокруг найденной строки (включает и до, и после; эквивалентно -A N -B N).
}

func ParseConfig() (*Config, error) {
	flags := &Flags{}

	flag.BoolVar(&flags.IgnoreFlag, "i", false, "ignore register")
	flag.BoolVar(&flags.InvertFlag, "v", false, "invert order")
	flag.BoolVar(&flags.StrictStringFlag, "F", false, "setting fix string")
	flag.BoolVar(&flags.LineNumberFlag, "n", false, "show numbers of string per each")
	flag.BoolVar(&flags.CountOfLineFlag, "c", false, "show count of matching strings with pattern")

	flag.IntVar(&flags.AdditionalLineFlag, "A", 0, "after each find string additional show n string after")
	flag.IntVar(&flags.PreviousLineFlag, "B", 0, "before each find string show previous n strings")
	flag.IntVar(&flags.AroundLineFlag, "C", 0, "show before and after n strings")

	flag.Parse()
	args := flag.Args()
	if len(args) == 0 {
		return nil, fmt.Errorf("You should write regexp or pattern to find strings")
	}

	pattern := args[0]
	return &Config{Pattern: pattern, Flags: flags, Buffer: make([]Line, 0)}, nil
}

func (c *Config) Matcher() (func(string) bool, error) {
	if !c.Flags.StrictStringFlag {
		pattern := c.Pattern
		if c.Flags.IgnoreFlag {
			pattern = "(?i)" + pattern
		}
		re, err := regexp.Compile(pattern)
		if err != nil {
			return nil, err
		}
		return func(s string) bool {
			match := re.MatchString(s)
			if c.Flags.InvertFlag {
				return !match
			}
			return match
		}, nil
	} else {
		str := c.Pattern
		if c.Flags.IgnoreFlag {
			str = strings.ToLower(str)
		}
		return func(s string) bool {
			tmp := s
			if c.Flags.IgnoreFlag {
				tmp = strings.ToLower(tmp)
			}
			match := strings.Contains(tmp, str)
			if c.Flags.InvertFlag {
				return !match
			}
			return match
		}, nil
	}
}
