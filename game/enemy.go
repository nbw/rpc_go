package game

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"

	"github.com/nbw/rps/api"
)

const SERVER_URL = "https://nh-a811fa16.herokuapp.com/"

type Enemy struct{}

func (e Enemy) Move() string {
	resp, err := http.Get(SERVER_URL)
	if err != nil {
		log.Fatal("Server not responding")
	}

	decoded := api.MoveResponse{}
	err = api.DecodeMoveResponse(resp, &decoded)
	if err != nil {
		dump(decoded)
		log.Fatal(fmt.Errorf("Bad move request: %v", err))
	}

	return decoded.Weapon
}

func dump(obj interface{}) {
	data, err := json.MarshalIndent(obj, "", "  ")
	if err == nil {
		log.Printf(string(data))
	}
}
