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
		os.Exit(1)
	}

	os.Exit(0)
}

func realMain() error {
	if len(os.Args) != 6 {
		return nil
	}

	// savetofile <mode> <filepath> <uid> <gid> <data>
	mode := os.Args[1]
	path := os.Args[2]
	uid_arg := os.Args[3]
	gid_arg := os.Args[4]
	data := os.Args[5]

	uid, err := strconv.Atoi(uid_arg)
	if err != nil {
		return err
	}

	gid, err := strconv.Atoi(gid_arg)
	if err != nil {
		return err
	}

	switch mode {
	case "append":
		f, err := os.OpenFile(path, os.O_APPEND|os.O_WRONLY, 0700)
		if err != nil {
			return err
		}

		defer f.Close()

		if _, err = f.WriteString(data); err != nil {
			return err
		}
	case "append-nl":
		f, err := os.OpenFile(path, os.O_APPEND|os.O_WRONLY, 0700)
		if err != nil {
			return err
		}

		defer f.Close()

		if _, err = f.WriteString(data); err != nil {
			return err
		}

		if _, err = f.WriteString("\n"); err != nil {
			return err
		}
	case "create-nl":
		err := ioutil.WriteFile(path, append([]byte(data), []byte("\n")...), 0700)
		if err != nil {
			return err
		}
	default: // "create"
		err := ioutil.WriteFile(path, []byte(data), 0700)
		if err != nil {
			return err
		}
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
