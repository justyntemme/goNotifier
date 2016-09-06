package web

import (
	"log"
	"net/http"
)

//IrcNotification Exported struct to allow passing of all arguments to the notifier Function
type IrcNotification struct {
	To     []string
	From   []string
	Server []string
	Msg    []string
	Room   []string
}

//StartServer Exported Function to initiate web server
func StartServer(port string) {
	http.HandleFunc("/", serveResponse)
	err := http.ListenAndServe(port, nil)
	if err != nil {
		log.Fatal(err)
	}
}

func serveResponse(d http.ResponseWriter, req *http.Request) {
	if req.URL.Query()["HELP"] != nil {
		d.Write([]byte("Insert Help Message"))
	}
	if req.URL.Query()["irc"] != nil {
		ircn := new(IrcNotification)
		ircn.To = req.URL.Query()["to"]
		ircn.From = req.URL.Query()["from"]
		ircn.Server = req.URL.Query()["server"]
		ircn.Msg = req.URL.Query()["msg"]

	}
}
