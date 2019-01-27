package horizon

import (
	"log"

	"bitbucket.org/aiventureteam/horizon-go/horizon/entities"
)

func Start(configuration entities.Configuration) {
	log.Fatal(configuration.VerifyToken)
}
