package main

import (
	"io/ioutil"
	"log"
	"os"
	"strconv"
)

func main() {
	err := realMain()

	if err != nil {
		log.Fatal(err)
	}

	os.Exit(0)
}

func realMain() error {
	if len(os.Args) != 6 {
		// Ensure the empty input case is handled correctly
		return nil
	}

	// savetofile <filepath> <data> <uid> <gid>
	path := os.Args[1]
	data := os.Args[2]
	uid_arg := os.Args[3]
	gid_arg := os.Args[4]

	uid, err := strconv.Atoi(uid_arg)
	if err != nil {
		return err
	}

	gid, err := strconv.Atoi(gid_arg)
	if err != nil {
		return err
	}

	err = ioutil.WriteFile(path, []byte(data), 0700)
	if err != nil {
		return err
	}

	err = os.Chown(path, uid, gid)
	if err != nil {
		return err
	}

	err = os.Chmod(path, 0640)
	if err != nil {
		return err
	}

	return nil
}
