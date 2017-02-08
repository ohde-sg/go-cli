package cli

func Execute(args []string, cmd Command) {
	if len(args) < 2 {
		//オプション無しルートコマンド
		cmd.Execute(cmd.Name(), []string{})
		return
	}
	execute(cmd.Name(), args[1:], cmd)
}

func execute(commandName string, args []string, cmd Command) {
	arg := args[0] //argはコマンド名orオプション
	for _, subCmd := range cmd.SubCommands() {
		if subCmd.Name() == arg {
			if len(args) < 2 { // オプション無しサブコマンド
				subCmd.Execute(commandName+" "+subCmd.Name(), []string{})
				return
			}
			//さらにサブコマンドがある場合
			if len(subCmd.SubCommands()) > 0 {
				execute(commandName+" "+subCmd.Name(), args[1:], subCmd)
				return
			}
			//オプションありサブコマンド
			subCmd.Execute(commandName+" "+subCmd.Name(), args[1:])
			return
		}
	}
	//オプションありコマンド
	cmd.Execute(commandName, args)
}
