package uuid

import (
	"log"

	uuid "github.com/nu7hatch/gouuid"
)

func UUID() (id string, err error) {
	var uid *uuid.UUID
	uid, err = uuid.NewV4()
	if err != nil {
		log.Fatalln("failed get uuid:", err)
		return
	}

	id = uid.String()
	return
}
