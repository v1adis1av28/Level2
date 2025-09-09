// config/config.go
package config

import (
	"fmt"
	"os"
	"strconv"
	"strings"
)

type Config struct {
	KeyColumn int  // sort by column number
	Numeric   bool // sort by column in numeric order
	Reverse   bool // sort in reverse way
	Unique    bool // output only unique values
}

func ParseConfig() *Config {
	conf := &Config{
		KeyColumn: 1,
	}

	args := os.Args[1:]
	var i int
	for i < len(args) {
		arg := args[i]
		if strings.HasPrefix(arg, "-") && len(arg) > 1 && arg[1] != '-' {
			for j := 1; j < len(arg); j++ {
				switch arg[j] {
				case 'r':
					conf.Reverse = true
				case 'n':
					conf.Numeric = true
				case 'u':
					conf.Unique = true
				case 'k':
					if j+1 < len(arg) {
						keyStr := arg[j+1:]
						key, err := strconv.Atoi(keyStr)
						if err != nil || key < 1 {
							fmt.Fprintf(os.Stderr, "Invalid key specification: %s\n", arg[j:])
							return nil
						}
						conf.KeyColumn = key
						j = len(arg) - 1
					} else {
						if i+1 >= len(args) {
							fmt.Fprintln(os.Stderr, "Missing argument for -k")
							return nil
						}
						i++
						key, err := strconv.Atoi(args[i])
						if err != nil || key < 1 {
							fmt.Fprintf(os.Stderr, "Invalid key: %s\n", args[i])
							return nil
						}
						conf.KeyColumn = key
						j = len(arg) - 1
					}
				default:
					fmt.Fprintf(os.Stderr, "Unknown flag: -%c\n", arg[j])
					return nil
				}
			}
			i++
		} else if strings.HasPrefix(arg, "--") {
			fmt.Fprintf(os.Stderr, "Long flags not supported: %s\n", arg)
			return nil
		} else {
			i++
		}
	}

	if !checkFlagCompatibility(conf) {
		return nil
	}

	return conf
}

func checkFlagCompatibility(conf *Config) bool {
	if conf.KeyColumn < 1 {
		fmt.Fprintln(os.Stderr, "Error: column number must be >= 1")
		return false
	}
	return true
}
