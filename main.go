package main

import (
	"io/ioutil"
	"log"
	"os"
	"os/exec"
	"os/signal"
	"syscall"
	"time"

	"github.com/google/uuid"
)

// StartMainRunner use run process backgroud
func StartMainRunner() bool {
	cmd := exec.Command("/usr/bin/gitlab-runner", "run", "--user=gitlab-runner", "--working-directory=/home/gitlab-runner")
	cmd.Stdout = os.Stdout
	err := cmd.Start()
	if err != nil {
		log.Println(err)
		return false
	}
	return true
}

// StartRunner use init runner
func StartRunner() bool {

	Token := os.Getenv("TOKEN")
	URL := os.Getenv("URL")
	CMD := "gitlab-runner"
	Name := "--name=" + GetName()
	RegisterToken := "--registration-token=" + Token
	RegisterURL := "--url=" + URL
	Volume := "--docker-volumes=/var/run/docker.sock:/var/run/docker.sock"
	Excuter := "--executor=docker"
	Image := "--docker-image=docker"
	IsLock := "--locked=false"
	OutPut, err := exec.Command(CMD, "register", "-n", Name, RegisterURL, RegisterToken, Excuter, Image, IsLock, Volume).Output()
	if len(Token) < 1 {
		log.Println("Please declear environment variable is TOKEN and assign value.")
	}
	if len(URL) < 1 {
		log.Println("Please declear environment variable is  URL and assign value.")
	}
	if err != nil {
		log.Println("Please valide TOKEN and URL")
		return false
	}
	log.Println(string(OutPut), URL, Token)
	return true
}

// Clean runner is not usege
func Clean() bool {
	OutPut, err := exec.Command("gitlab-runner", "verify", "--delete", GetName()).Output()
	if err != nil {
		log.Println(err)
		return false
	}
	log.Println(OutPut)
	return true
}

// StopRunner use stop runner
func StopRunner() bool {
	OutPut, err := exec.Command("gitlab-runner", "unregister", "--name", GetName()).Output()
	if err != nil {
		log.Println(err)
		return false
	}
	log.Println(OutPut)
	return true

}

// IsStart use Control runner start
var (
	IsStart bool
	IsMain  bool
)

// GetName Use Name Runner
func GetName() string {
	if _, err := os.Stat("./config.txt"); os.IsNotExist(err) {
		ID := uuid.New()
		Name := "runner-" + ID.String()
		Text := []byte(Name)
		err := ioutil.WriteFile("./config.txt", Text, 777)
		if err != nil {
			log.Println(err)
		}
		return Name
	}
	File, err := ioutil.ReadFile("./config.txt")
	if err != nil {
		log.Println(err)
	}
	return string(File)
}

func init() {
	IsMain = StartMainRunner()
}

func main() {
	if !IsMain {
		log.Println("Main Process error")
		os.Exit(0)
	}

	IsStart = StartRunner()
	ticker := time.NewTicker(5 * time.Second)
	quit := make(chan struct{})
	Control := make(chan os.Signal, 2)
	signal.Notify(Control, syscall.SIGTERM, os.Interrupt)
	// handel exit process
	go func() {
		<-Control
		Clean()
		StopRunner()
		os.Exit(0)
	}()

	go func() {
		for {
			select {
			case <-ticker.C:
				if !IsStart {
					Clean()
					IsStart = StartRunner()
				}
			case <-quit:
				ticker.Stop()
				return

			}

		}

	}()

	<-quit
}
