package initializers

import (
	"os"
	
	gofixerio "github.com/LordotU/go-fixerio"
)

var CUR_API *gofixerio.FixerIO

func ConnectToCurApi() {
	var err error
	CUR_API, err = gofixerio.New(os.Getenv("FIXERIO_API_KEY"), "EUR", false)
	
	if err != nil {
		panic("Failed to connect to currency api: " + err.Error())
	}
}