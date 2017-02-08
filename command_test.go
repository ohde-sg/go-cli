package cli

import (
	"strings"
	"testing"
)

var (
	expectedCommandName string = "main sub1 sub2"
	expectedArgs        []string
	resultCommandName   string
	resultArgs          []string
)

type MyCommand struct {
	MyName        string
	MySubCommands []Command
}

func (mc MyCommand) Name() string {
	return mc.MyName
}

func (mc MyCommand) SubCommands() []Command {
	return mc.MySubCommands
}

func (mc MyCommand) Execute(commandName string, args []string) {
	resultCommandName = commandName
	resultArgs = args
}

func IsCollect(t *testing.T, inp []string, commandName string, args []string) {
	t.Log("input:", strings.Join(inp, " "))
	t.Log("commandName:", resultCommandName)
	t.Log("arguments:", resultArgs)
	if commandName != expectedCommandName {
		t.Errorf("commandName: %s is expected, but %s", expectedCommandName, commandName)
	}
	for i, arg := range expectedArgs {
		tmpArg := args[i]
		if arg != tmpArg {
			t.Errorf("args: %s is expected, but %s", expectedArgs, args)
		}
	}
}

func TestCommand(t *testing.T) {
	cmd := MyCommand{MyName: "main"}
	cmd.MySubCommands = []Command{
		MyCommand{
			MyName: "sub1",
			MySubCommands: []Command{
				MyCommand{
					MyName: "sub2",
				},
			},
		},
	}
	OsArgs := []string{"./main", "sub1", "sub2", "-v", "-h"}
	expectedCommandName = "main sub1 sub2"
	expectedArgs = []string{"-v", "-h"}
	Execute(OsArgs, cmd)
	IsCollect(t, OsArgs, resultCommandName, resultArgs)
}

func TestCommand2(t *testing.T) {
	cmd := MyCommand{MyName: "main"}
	OsArgs := []string{"./main", "-v", "-h"}
	expectedCommandName = "main"
	expectedArgs = []string{"-v", "-h"}
	Execute(OsArgs, cmd)
	IsCollect(t, OsArgs, resultCommandName, resultArgs)
}

func TestCommand3(t *testing.T) {
	cmd := MyCommand{MyName: "main"}
	sub1_cmd := MyCommand{MyName: "sub1"}
	cmd.MySubCommands = []Command{sub1_cmd}
	OsArgs := []string{"./main", "sub1", "-v", "-h"}
	expectedCommandName = "main sub1"
	expectedArgs = []string{"-v", "-h"}
	Execute(OsArgs, cmd)
	IsCollect(t, OsArgs, resultCommandName, resultArgs)
}

func TestCommand4(t *testing.T) {
	cmd := MyCommand{MyName: "main"}
	sub1_cmd := MyCommand{MyName: "sub1"}
	cmd.MySubCommands = []Command{sub1_cmd}
	OsArgs := []string{"./main", "sub1"}
	expectedCommandName = "main sub1"
	expectedArgs = []string{}
	Execute(OsArgs, cmd)
	IsCollect(t, OsArgs, resultCommandName, resultArgs)
}

func TestCommand5(t *testing.T) {
	cmd := MyCommand{MyName: "main"}
	sub1_cmd := MyCommand{MyName: "sub1"}
	sub2_cmd := MyCommand{MyName: "sub2"}
	sub1_cmd.MySubCommands = []Command{sub2_cmd}
	cmd.MySubCommands = []Command{sub1_cmd}
	OsArgs := []string{"./main", "sub1", "sub2"}
	expectedCommandName = "main sub1 sub2"
	expectedArgs = []string{}
	Execute(OsArgs, cmd)
	IsCollect(t, OsArgs, resultCommandName, resultArgs)
}

func TestCommand6(t *testing.T) {
	cmd := MyCommand{MyName: "main"}
	sub1_cmd := MyCommand{MyName: "sub1"}
	sub2_cmd := MyCommand{MyName: "sub2"}
	cmd.MySubCommands = []Command{sub1_cmd, sub2_cmd}
	OsArgs := []string{"./main", "sub1", "-v", "-h"}
	expectedCommandName = "main sub1"
	expectedArgs = []string{"-v", "-h"}
	Execute(OsArgs, cmd)
	IsCollect(t, OsArgs, resultCommandName, resultArgs)
}

func TestCommand7(t *testing.T) {
	cmd := MyCommand{MyName: "main"}
	sub1_cmd := MyCommand{MyName: "sub1"}
	sub2_cmd := MyCommand{MyName: "sub2"}
	cmd.MySubCommands = []Command{sub1_cmd, sub2_cmd}
	OsArgs := []string{"./main", "sub2", "-v", "-h"}
	expectedCommandName = "main sub2"
	expectedArgs = []string{"-v", "-h"}
	Execute(OsArgs, cmd)
	IsCollect(t, OsArgs, resultCommandName, resultArgs)
}
