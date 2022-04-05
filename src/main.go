package main

import (
    "log"
    "net/http"
    "encoding/json"
    "github.com/joho/godotenv"
    "os"
    "fmt"
)

type env struct {
    env   string
    def   string
    value string
    cast  string
}

// within `envs` all actual environmental values used by
// this app will be stored after call of `loadEnv()`
var envs = make( map[ string ] string )
// `envDefinition` is the definition JSON for all
// environmental variables used by this web app
var envDefinition string = `[
    {
        "_comment": "Port to start the web server at",
        "env":      "WEB_PORT",
        "default":  "80"
    }
]`

// home handler function that returns HTTP resonse
func home( w http.ResponseWriter, r *http.Request ) {
    w.Write( []byte ( "Hello from Sandbox" ) )
}

// initialize environmental variables
func loadEnv() {
    err := godotenv.Load()
    if err != nil {
        log.Println( "No environmental variables are provided by `.env` file – only working with “real” env vars." )
    }
    // load all environmental variables
    var tmpEnvs [] map[ string ] string
    json.Unmarshal( []byte( envDefinition ), &tmpEnvs )
    for _, elem := range tmpEnvs {
        value, ok := os.LookupEnv( elem[ "env" ] )
        if ! ok {
            value = elem[ "default" ]
        }
        envs[ elem[ "env" ] ] = value
    }
}

// function starting the webserver
func prepareRoutes() *http.ServeMux {
    // initialize new servemux (router)
    mux := http.NewServeMux()
    // register home function as handler for `/` URL pattern
    // `/` is catchy, so it'll catch ALL requests
    mux.HandleFunc( "/", home )

    return mux
}

// main function that serves the application on port 80
func main() {

    loadEnv()
    mux := prepareRoutes()

    // log message what is going on: starting webserver
    // webPort := fmt.Sprintf( ":%s", envs[ "WEB_PORT" ] )
    webPort := fmt.Sprint( ":", envs[ "WEB_PORT" ] )
    log.Printf( "Starting webserver on %s\n", webPort )
    // start Webserver and register non-nil-errors
    err := http.ListenAndServe( webPort, mux )
    // log errors
    log.Fatal( err )
}
