package arg

import "strings"

type Arg struct {
	commandName string
	args        []string
}

func NewArg(message string) *Arg {
	split := strings.Split(message, " ")
	return &Arg{
		commandName: split[0],
		args:        split[1:],
	}
}

func (a *Arg) IsSameCommand(arg string) bool {
	return a.commandName == arg
}

func (a *Arg) StartWith(prefix string) bool {
	return strings.HasPrefix(a.commandName, prefix)
}

func (a *Arg) Next() *Arg {
	var args []string
	if len(a.args) > 1 {
		args = a.args[1:]
	} else {
		args = []string{}
	}
	return &Arg{commandName: a.args[0], args: args}
}
