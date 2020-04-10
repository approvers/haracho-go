package client

func ExecuteCommand(parser CommandParser, collection *CommandCollection) {
	arg, ctx := parser.Parse()
	collection.execute(arg, ctx)
}
