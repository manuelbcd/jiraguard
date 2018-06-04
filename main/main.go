package main

import(
	// "bufio"
	"fmt"
	// "os"
	"strings"
	// "syscall"

	jira "github.com/andygrunwald/go-jira"
	//"golang.org/x/crypto/ssh/terminal"
)

const url = "https://hosturl.com"
const user = "yourusername"
const pw = "your_password"

func main() {
	//r := bufio.NewReader(os.Stdin)
	//fmt.Print("Jira URL: ")
	//jiraURL, _ := r.ReadString('\n')
	//fmt.Print("Jira Username: ")
	//username, _ := r.ReadString('\n')
	//fmt.Print("Jira Password: ")
	//bytePassword, _ := terminal.ReadPassword(int(syscall.Stdin))
	//password := string(bytePassword)

	var username = user
	var password = pw
	var jiraURL = url

	tp := jira.BasicAuthTransport{
		Username: strings.TrimSpace(username),
		Password: strings.TrimSpace(password),
	}

	client, err := jira.NewClient(tp.Client(), strings.TrimSpace(jiraURL))
	if err != nil {
		fmt.Printf("\nerror: %v\n", err)
		return
	}

	u, _, err := client.Issue.Get("OCI-1111", nil)

	if err != nil {
		fmt.Printf("\nerror: %v\n", err)
		return
	}

	fmt.Printf("\nCaption: %v\nSuccess!\n", u.Fields.Summary)
}
