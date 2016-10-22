package main

import (
	"fmt"
	"io/ioutil"
	"net/http"
	"net/url"
	"os"
	"strings"
)

const serverURL = "http://upe21.cs.rpi.edu:3000/api/"

// Strategy defines approach to choosing the next move
type Strategy interface {
	execute()
}

// Player encapsulates the user's credentials
type Player struct {
	DevKey   string
	Username string
}

// NewPlayer reads the given credential file and returns a new Player
func NewPlayer(credentialFile *string) *Player {
	fp, err := os.Open(*credentialFile)
	if err != nil {
		panic(err)
	}
	defer func() {
		if err := fp.Close(); err != nil {
			panic(err)
		}
	}()

	buf := make([]byte, 1024)
	n, err := fp.Read(buf)
	if err != nil {
		panic(err)
	}
	if n == 0 {
		return nil
	}
	res := string(buf[:n])
	strres := string(res)
	bits := strings.Split(strres, "\n")
	p := &Player{bits[0], bits[1]}
	return p
}

// PracticeGame joins a practice game and begins making moves using the given strategy
func PracticeGame(player *Player, strategy *Strategy) {
	resp, err := http.PostForm(serverURL+"games/search",
		url.Values{"devKey": {player.DevKey}, "username": {player.Username}})
	if err != nil {
		panic(err)
	}
	defer resp.Body.Close()
	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		panic(err)
	}
	fmt.Println(string(body))
}

func main() {
	args := os.Args
	if len(args) < 1 {
		fmt.Fprintf(os.Stderr, "Usage: go run ai.go <credential_file>\n")
		os.Exit(1)
	}
	player := newPlayer(&args[1])
	fmt.Printf("Loaded credentials for %s\n", player.Username)

}
