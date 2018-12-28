package behavioral

import "testing"

func TestCommand(t *testing.T) {
	var tv TV
	openCommand := OpenCommand{tv}
	invoker := Invoker{openCommand}
	invoker.Do()

	invoker.SetCommand(CloseCommand{tv})
	invoker.Do()

	var openClose *OpenCloseCommand = NewOpenCloseCommand()
	openClose.AddCommand(OpenCommand{tv})
	openClose.AddCommand(CloseCommand{tv})
	invoker.SetCommand(openClose)
	invoker.Do()
}
