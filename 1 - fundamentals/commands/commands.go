package main

import "fmt"

func main() {
	fmt.Printf("Another program in %s\n")
	fmt.Printf("Another program in %s\n", "Go")
}

/*
> go help <topic/command_name>***: Shows the docummentation about a topic or command
> godoc -http 'localhost:6060'*** (Maybe you should to install go tools): Download offline docs and
	get it available at a disred address (localhost:6060 in this case), so you can access throgh browser by typing: http://localhost:6060
> go env: Displays go environment variables
> go doc cmd/vet: Displays the documentation of Vet command, which reports some code issues
> go vet <file_name>.g: Analisys the code in the file and report some code issues
> go build <filename>.go: Build the file and generate an exacutable at the same folder level
> go run <filename>.go: Build and runs the code in the file
> go get -u <source_repo/username/projectname>: Download the code from repositorie's project as an code depency,
	which is installed at 'src/repo.com/projectname'. e.g: go get github.com/go-sql-driver/mysql
*/
