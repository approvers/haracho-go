package client

type Arg struct {
	CommandName string
	Args        []string
}

func (a *Arg) IsSame(arg string) bool {
	return a.CommandName == arg
}

func (a *Arg) Next() *Arg {
	return &Arg{CommandName: a.Args[0], Args: a.Args[1:]}
}
