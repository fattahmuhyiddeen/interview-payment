package config

import "os"

// DateTimeFormat is format of timestamp used accros this application
const DateTimeFormat = "2006-01-02T15:04:05"

//Port number used by app
var Port = os.Getenv("PORT")
