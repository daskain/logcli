package main

import (
	"encoding/json"
	"github.com/grafana/loki/pkg/logcli/client"
	"github.com/joho/godotenv"
	"io/ioutil"
	"log"
	"os"
	"time"
)

func init() {
	// loads values from .env into the system
	if err := godotenv.Load(); err != nil {
		log.Print("No .env file found")
	}
}

func main() {

	url, _ := os.LookupEnv("url")
	name, _ := os.LookupEnv("name")
	pass, _ := os.LookupEnv("pass")

	var client = client.DefaultClient{
		Address:  url,
		Username: name,
		Password: pass,
	}

	var result, _ = client.Query("{app = \"bff\"} |=\"error\"", 5000, time.Now(), 1, true)


	data := result.Data
	jsonString, _ := json.Marshal(data)
	_ = ioutil.WriteFile("test.json", jsonString, 0644)

}
