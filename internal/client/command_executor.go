package client

func ExecuteCommand(req CommandParser, res CommandResponse) {
	command := req.parse()
	result := command.execute()
	res.respond(result)
}
