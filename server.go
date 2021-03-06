package main

import (
	"encoding/json"
	"fmt"
	"github.com/joho/godotenv"
	"github.com/sevlyar/go-daemon"
	"log"
	"net/http"
	"os"
	"os/exec"
	"path/filepath"
)

type DockerPayload struct {
	Callback_Url string
	Repository   struct {
		Is_Trusted bool
		Is_Private bool
		Repo_Url   string
		Repo_Name  string
	}
}

func dockerPull(r string) {
	dir, err := filepath.Abs(filepath.Dir(os.Args[0]))
	if err != nil {
		log.Fatal(err)
	}

	// point command to dockerpull script
	cmdPath := string(dir) + "/dockerPull.sh"
	cmd := exec.Command(cmdPath, r)
	err = cmd.Run()
	if err != nil {
		os.Exit(1)
	}

	fmt.Println("Docker pull script run successfully")
}

func handler(w http.ResponseWriter, req *http.Request) {
	var p DockerPayload
	if req.Body == nil {
		http.Error(w, "Please provide a request body", 400)
	}

	// attempt to parse body
	err := json.NewDecoder(req.Body).Decode(&p)
	if err != nil {
		http.Error(w, err.Error(), 400)
		return
	}

	// send response header
	w.Header().Set("Content-Type", "text/plain; charset=utf-8")
	w.WriteHeader(http.StatusOK)
	w.Write([]byte("Payload accepted. Pulling " + p.Repository.Repo_Name + "\n"))

	fmt.Println(p.Callback_Url)
	dockerPull(p.Repository.Repo_Name)
}

func main() {

	cntxt := &daemon.Context{
		PidFileName: "pid",
		PidFilePerm: 0644,
		LogFileName: "log",
		LogFilePerm: 0640,
		WorkDir:     "./",
		Umask:       027,
		Args:        []string{"[go-daemon process]"},
	}

	d, err := cntxt.Reborn()
	if err != nil {
		log.Fatal("Unable to run crane daemon: ", err)
	}
	if d != nil {
		return
	}
	defer cntxt.Release()

	log.Print("----------------")
	log.Print("crane daemon started")

	// read environment variables
	err = godotenv.Load()
	if err != nil {
		log.Fatal("No .env file found! Please make one using env.example as a mock")
	}

	// set certification and key
	certFileAndPath := os.Getenv("CERT_FILE_AND_PATH")
	keyFileAndPath := os.Getenv("KEY_FILE_AND_PATH")

	// add routes
	http.HandleFunc("/", handler)

	// serve HTTPS
	err = http.ListenAndServeTLS(":3333", certFileAndPath, keyFileAndPath, nil)
	if err != nil {
		log.Fatal(err)
	}
}
