package sugar

import (
	"net/http"
)

type Sugar struct {
	Routes Routes
}

func (server *Sugar) Listen(port string) error {
	err := http.ListenAndServe(port, nil)
	if err != nil {
		return err
	}
	return nil
}

func Init() *Sugar {
	sugar := &Sugar{
		Routes: Routes{
			AvailableRoutes: map[string]string{},
		},
	}

	return sugar
}