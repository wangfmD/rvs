package remotev

import (
	"fmt"
	"golang.org/x/crypto/ssh"
	"log"
	"net"
	// "os"
	"strings"
	"time"
)

func connect(user, password, host string, port int) (*ssh.Session, error) {
	var (
		auth         []ssh.AuthMethod
		addr         string
		clientConfig *ssh.ClientConfig
		client       *ssh.Client
		session      *ssh.Session
		err          error
	)
	// get auth method
	auth = make([]ssh.AuthMethod, 0)
	auth = append(auth, ssh.Password(password))

	clientConfig = &ssh.ClientConfig{
		User:    user,
		Auth:    auth,
		Timeout: 30 * time.Second,
		HostKeyCallback: func(hostname string, remote net.Addr, key ssh.PublicKey) error {
			return nil
		},
	}

	// connet to ssh
	addr = fmt.Sprintf("%s:%d", host, port)

	if client, err = ssh.Dial("tcp", addr, clientConfig); err != nil {
		return nil, err
	}

	// create session
	if session, err = client.NewSession(); err != nil {
		return nil, err
	}

	return session, nil
}

func GetVersionMaps(name, pw, addr string) (map[string]string, error) {
	// session, err := connect("root", "xungejiaoyu2015", "10.1.41.60", 22)
	session, err := connect(name, pw, addr, 22)

	if err != nil {
		log.Fatal(err)
		return nil, err
	}
	defer session.Close()

	// session.Stderr = os.Stderr
	cmd := "docker ps | awk '{print $2}' | grep -v ID"
	res, err := session.Output(cmd)
	if err != nil {
		fmt.Println(err)
		return nil, err
	}
	s := string(res)
	// fmt.Println(strings.Split(s, "5000/"))
	version := make(map[string]string)
	for _, value := range strings.Split(s, "\n") {
		if value != "" {
			s := strings.Split(value, "5000/")
			// fmt.Println(strings.Split(value, "5000/"))
			s1 := strings.Split(s[1], "/")
			// fmt.Println(s1)
			s2 := strings.Split(s1[1], ":")
			// fmt.Println(s2)
			version[s2[0]] = s2[1]
		}
	}
	fmt.Println(version)
	return version, nil
}
