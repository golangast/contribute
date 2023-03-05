
# contribute

run the following
go mod init "yourproject"
go mod tidy
go install github.com/spf13/cobra-cli@latest
export PATH="~/go/bin:$PATH"
cobra-cli init

to create command
cobra add your command name

to run it 
go run main.go your command
