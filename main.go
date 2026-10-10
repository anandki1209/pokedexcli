package main

func main() {
	cfg := config{
		supportedCommands: getCommands(),
	}
	startRepl(&cfg)

}
