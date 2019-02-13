package main

import (
	"bufio"
	"crypto/md5"
	"encoding/hex"
	"encoding/json"
	"fmt"
	"os"
	"strconv"

	"github.com/sethvargo/go-password/password"
)

// {
//   "username" : "combox-25",
//   "password" : "5f360d9a62441e2a6e148d06ccb3a05a",
//   "salt" : "",
//   "is_superuser" : false
// }

// User ...
type User struct {
	Username    string `json:"username"`
	Password    string `json:"password"`
	Salt        string `json:"salt"`
	IsSuperuser bool   `json:"is_superuser"`
}

// {
//   "username" : "combox-25",
//   "publish" : [
//       "asset/telemetry/#",
//       "asset/commandresp/#"
//   ],
//   "subscribe" : [
//       "asset/commands/#",
//       "asset/cmdsimul/#"
//   ]
// }

// ACL ...
type ACL struct {
	Username  string   `json:"username"`
	Publish   []string `json:"publish"`
	Subscribe []string `json:"subscribe"`
}

func check(e error) {
	if e != nil {
		panic(e)
	}
}

func md5Hash(text string) string {
	hasher := md5.New()
	hasher.Write([]byte(text))
	return hex.EncodeToString(hasher.Sum(nil))
}

func main() {
	if len(os.Args) < 4 {
		fmt.Println(
			"usage: ./credentials-generator <nb-credentials> <username-prefix> <counter-start>\n" +
				"usage example: ./credentials-generator 25 combox- 26")
		os.Exit(1)
	}

	nbCredentials, err := strconv.Atoi(os.Args[1]) // e.g. 25
	check(err)
	userPrefix := os.Args[2]                 // e.g. "combox-"
	counter, err := strconv.Atoi(os.Args[3]) // e.g. 26
	check(err)

	gen, err := password.NewGenerator(&password.GeneratorInput{
		Symbols: "",
	})
	check(err)

	publish := []string{
		"asset/telemetry/#",
		"asset/commandresp/#",
	}
	subscribe := []string{
		"asset/commands/#",
		"asset/cmdsimul/#",
	}

	f, err := os.Create("./credentials.json")
	check(err)
	defer f.Close()

	w := bufio.NewWriter(f)

	for i := 0; i < nbCredentials; i++ {
		username := fmt.Sprintf("%s%d", userPrefix, counter)
		password, err := gen.Generate(15, 5, 0, false, false)
		check(err)
		passwordHash := md5Hash(password)
		user, err := json.MarshalIndent(User{username, passwordHash, "", false}, "", "  ")
		check(err)
		acl, err := json.MarshalIndent(ACL{username, publish, subscribe}, "", "  ")
		check(err)
		counter++

		_, err = w.WriteString(fmt.Sprintf(
			"%s  %s\n%s\n%s\n\n", username, password, string(user), string(acl)))
		check(err)
	}

	w.Flush()
}
