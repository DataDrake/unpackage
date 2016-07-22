package environment

import (
	"fmt"
	"io"
	"os"
)

func AppendPath( stream io.Writer, paths map[string][]string ) (n int, err error){
	n = 0
	for name,list := range paths {
		variable := os.Getenv(name)
		empty := len(variable) == 0
		for _,p := range list {
			if empty {
				variable = p
				empty = false
			} else {
				variable = variable + ":" + p
			}
		}
		m, err := fmt.Fprintf(stream, "export %s=%s\n", name, variable)
		n += m
		if err != nil {
			return
		}
	}
	return
}

func PrependPath( stream io.Writer, paths map[string][]string ) (n int, err error){
	n = 0
	for name,list := range paths {
		variable := os.Getenv(name)
		empty := len(variable) == 0
		for _,p := range list {
			if empty {
				variable = p
				empty = false
			} else {
				variable = p + ":" + variable
			}
		}
		m, err := fmt.Fprintf(stream, "export %s=%s\n", name, variable)
		n += m
		if err != nil {
			return
		}
	}
	return
}

func Setenv( stream io.Writer, vars map[string]string ) (n int, err error) {
	n = 0
	for k,v := range vars {
		m,err := fmt.Fprintf(stream, "export %s=%s\n", k , v)
		n += m
		if err != nil {
			return
		}
	}
	return
}

func Unsetenv( stream io.Writer, vars []string ) (n int, err error) {
	n = 0
	for _,v := range vars {
		m,err := fmt.Fprintf(stream, "unset %s\n", v)
		n += m
		if err != nil {
			return
		}
	}
	return
}