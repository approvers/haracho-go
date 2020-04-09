package client

type Arg struct {
	CommandName string
	Args        []string
}

func (a *Arg) IsSameCommand(arg string) bool {
	return a.CommandName == arg
}

func (a *Arg) Next() *Arg {
	var args []string
	if len(a.Args) > 1 {
		args = a.Args[1:]
	} else {
		args = []string{}
	}
	return &Arg{CommandName: a.Args[0], Args: args}
}
