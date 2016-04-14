package ssh

import (
	"io"
	"log"
	"os"

	"github.com/molsbee/clc-cli/api"

	"golang.org/x/crypto/ssh"
	"golang.org/x/crypto/ssh/terminal"
)

var terminalModes = ssh.TerminalModes{
	ssh.ECHO:          0,
	ssh.ECHOCTL:       0,
	ssh.TTY_OP_ISPEED: 14400,
	ssh.TTY_OP_OSPEED: 14400,
}

// Connect Create SSH Connection to Server Provided
func Connect(serverAlias string) {
	server := api.Server{
		Name: serverAlias,
	}

	serverCredentials := server.GetCredentials()
	ipAddress := server.Get().Details.IPAddresses[0].Internal

	config := &ssh.ClientConfig{
		User: serverCredentials.Username,
		Auth: []ssh.AuthMethod{ssh.Password(serverCredentials.Password)},
	}

	conn := establishConnection(ipAddress, config)
	session := createSession(conn)
	defer session.Close()

	fd := int(os.Stdin.Fd())
	oldState, err := terminal.MakeRaw(fd)
	if err != nil {
		log.Fatal("Unable to put terminal in Raw Mode", err)
	}

	width, height, _ := terminal.GetSize(fd)
	if err := session.RequestPty("xterm", height, width, terminalModes); err != nil {
		log.Fatalln("Request for Psuedo Terminal Failed: ", err)
	}

	redirectStreamToSession(session)

	if err := session.Shell(); err != nil {
		log.Fatalln("Failed to start shell: ", err)
	}

	session.Wait()
	terminal.Restore(fd, oldState)
}

func establishConnection(ipAddress string, config *ssh.ClientConfig) *ssh.Client {
	conn, err := ssh.Dial("tcp", ipAddress+":22", config)
	if err != nil {
		log.Fatalln("Unable to create tcp connection to server: "+ipAddress+":22", err)
	}

	return conn
}

func createSession(conn *ssh.Client) *ssh.Session {
	session, err := conn.NewSession()
	if err != nil {
		log.Fatalln("Unable to create session with server: ", err)
	}

	return session
}

func redirectStreamToSession(session *ssh.Session) {
	stdin, err := session.StdinPipe()
	if err != nil {
		log.Fatal("Unable to setup stdin for session", err)
	}
	go io.Copy(stdin, os.Stdin)

	stdout, err := session.StdoutPipe()
	if err != nil {
		log.Fatal("Unable to setup stdout for session", err)
	}
	go io.Copy(os.Stdout, stdout)

	stderr, err := session.StderrPipe()
	if err != nil {
		log.Fatal("Unable to setup stderr for session", err)
	}
	go io.Copy(os.Stderr, stderr)
}
