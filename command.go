package cli

type Command interface {
	Name() string                              // return command name
	SubCommands() []Command                    // return registered slice of sub command
	Execute(commandName string, args []string) // execute with commandName and args
}
