package controllers

import (
	"encoding/json"
	"log"
	"strings"

	"github.com/Miya784/update-status-service/initials"
)

type UpdateStatusLight struct {
	Light  string `json:"Light"`
	Client string `json:"Client"`
}
type UpdateStatusFan struct {
	Fan    string `json:"Fan"`
	Client string `json:"Client"`
}
type UpdateStatusGate struct {
	Gate   string `json:"Gate"`
	Client string `json:"Client"`
}
type UpdateStatusPWM struct {
	PWM    string `json:"PWM"`
	Client string `json:"Client"`
}

func MessageUpdateStatus(message []byte) error {
	if strings.Contains(string(message), "Light") {
		var body UpdateStatusLight
		err := json.Unmarshal(message, &body)
		if err != nil {
			log.Println("Error:", err)
			return err
		}
		log.Println(body)
		_, err = initials.DB.Exec("UPDATE \"Client\" SET status = $1 WHERE client = $2", body.Light, body.Client)
		if err != nil {
			panic(err)
		}
		_, err = initials.DB.Exec("UPDATE \"Client\" SET \"updatedAt\" = CURRENT_TIMESTAMP WHERE client = $1", body.Client)
		if err != nil {
			panic(err)
		}

	} else if strings.Contains(string(message), "Fan") {
		// var body UpdateStatusFan
		// err := json.Unmarshal(message, &body)
		// if err != nil {
		// 	log.Println("Error:", err)
		// 	return err
		// }
		// log.Println(body)
		// initials.DB.Exec("UPDATE \"Client\" SET status = $1 WHERE client = $2", body.Fan, body.Client)
		var body UpdateStatusFan
		err := json.Unmarshal(message, &body)
		if err != nil {
			log.Println("Error:", err)
			return err
		}
		log.Println(body)
		_, err = initials.DB.Exec("UPDATE \"Client\" SET status = $1 WHERE client = $2", body.Fan, body.Client)
		if err != nil {
			panic(err)
		}
		_, err = initials.DB.Exec("UPDATE \"Client\" SET \"updatedAt\" = CURRENT_TIMESTAMP WHERE client = $1", body.Client)
		if err != nil {
			panic(err)
		}

	} else if strings.Contains(string(message), "Gate") {
		// var body UpdateStatusGate
		// err := json.Unmarshal(message, &body)
		// if err != nil {
		// 	log.Println("Error:", err)
		// 	return err
		// }
		// log.Println(body)
		// initials.DB.Exec("UPDATE \"Client\" SET status = $1 WHERE client = $2", body.Gate, body.Client)
		var body UpdateStatusGate
		err := json.Unmarshal(message, &body)
		if err != nil {
			log.Println("Error:", err)
			return err
		}
		log.Println(body)
		_, err = initials.DB.Exec("UPDATE \"Client\" SET status = $1 WHERE client = $2", body.Gate, body.Client)
		if err != nil {
			panic(err)
		}
		_, err = initials.DB.Exec("UPDATE \"Client\" SET \"updatedAt\" = CURRENT_TIMESTAMP WHERE client = $1", body.Client)
		if err != nil {
			panic(err)
		}

	} else if strings.Contains(string(message), "PWM") {
		// var body UpdateStatusPWM
		// err := json.Unmarshal(message, &body)
		// if err != nil {
		// 	log.Println("Error:", err)
		// 	return err
		// }
		// log.Println(body)
		// initials.DB.Exec("UPDATE \"Client\" SET status = $1 WHERE client = $2", body.PWM, body.Client)
		var body UpdateStatusPWM
		err := json.Unmarshal(message, &body)
		if err != nil {
			log.Println("Error:", err)
			return err
		}
		log.Println(body)
		_, err = initials.DB.Exec("UPDATE \"Client\" SET status = $1 WHERE client = $2", body.PWM, body.Client)
		if err != nil {
			panic(err)
		}
		_, err = initials.DB.Exec("UPDATE \"Client\" SET \"updatedAt\" = CURRENT_TIMESTAMP WHERE client = $1", body.Client)
		if err != nil {
			panic(err)
		}
	}

	return nil
}
