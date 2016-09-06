package web

import "net/http"

//IrcNotification Exported struct to allow passing of all arguments to the notifier Function
type IrcNotification struct {
	to     string
	from   string
	server string
	msg    string
}

//StartServer Exported Function to initiate web server
func StartServer(port int) {
	http.HandleFunc("/", serveResponse)
	http.ListenAndServe(":3030", nil)
}

func serveResponse(d http.ResponseWriter, req *http.Request) {
	if req.URL.Query()["HELP"] != nil {
		d.Write([]byte("Insert Help Message"))
	}
}
