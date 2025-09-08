package config

import "flag"

type Config struct {
	KeyColumn    int  // sort by colunmd
	Numeric      bool // sort by column in numeric order
	Reverse      bool // sort in reverse way
	Unique       bool // output only unique values
	MonthSort    bool // sort by date format
	IgnoreBlanks bool // ingore trailing blanks
	CheckSorted  bool // is data sorted?
	HumanNumeric bool // sort by numeric values for human representation
}

func ParseConfig() *Config {
	conf := &Config{}

	flag.IntVar(&conf.KeyColumn, "k", 1, "sort by column number")
	flag.BoolVar(&conf.Numeric, "n", false, "sort by column in numeric order")
	flag.BoolVar(&conf.Reverse, "r", false, "sort in reverse way")
	flag.BoolVar(&conf.Unique, "u", false, "output only unique values")
	flag.BoolVar(&conf.MonthSort, "d", false, "sort by date format")
	flag.BoolVar(&conf.IgnoreBlanks, "b", false, "ignoring trailing blanks")
	flag.BoolVar(&conf.CheckSorted, "c", false, "check is sorted")
	flag.BoolVar(&conf.HumanNumeric, "h", false, "human numeric read")

	flag.Parse()
	return conf
}
