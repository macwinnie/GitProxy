package main

import (
    "os"
    "fmt"
    "log"
    "sort"
    "strings"
    "net/http"
    // "database/sql"
    "encoding/json"
    "github.com/joho/godotenv"
    // "github.com/golang-migrate/migrate/v4"
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
// `envDefinition` is the definition JSON for all environmental
// variables used by this web app – see `README.md` for description
var envDefinition string = `[
    {
        "env":      "WEB_PORT",
        "default":  "80"
    },
    {
        "env":      "GIT_CACHE_DAYS",
        "default":  "7"
    },
    {
        "env":       "LDAP_GROUP_BASES",
        "mandatory": "yes"
    },
    {
        "env":       "LDAP_USER_BASES",
        "mandatory": "yes"
    },
    {
        "env":       "LDAP_HOST",
        "mandatory": "yes"
    },
    {
        "env":       "LDAP_DIRECTORY_USER",
        "mandatory": "yes"
    },
    {
        "env":       "LDAP_DIRECTORY_PASSWORD",
        "mandatory": "yes"
    },
    {
        "env":       "APP_USERS",
        "mandatory": "yes"
    },
    {
        "env":       "APP_ADMINS",
        "mandatory": "yes"
    }
]`
// strings for evaluating environmental variables for true
var trueEnv = []string{ "true", "yes", "1", "y", "t" }

// initialize environmental variables
func loadEnv() {
    // sort trueEnv variable
    sort.Strings( trueEnv )
    // initiate .env file
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
            if mandatory, ok := elem[ "mandatory" ]; ok {
                target := strings.ToLower( mandatory )
                i := sort.SearchStrings( trueEnv, target )
                if i < len( trueEnv ) && trueEnv[ i ] == target {
                    log.Fatal( fmt.Sprintf( "Environmental variable “%s” is mandatory but currently not defined.", elem[ "env" ] ) )
                }
            }
            value = elem[ "default" ]
        }
        envs[ elem[ "env" ] ] = value
    }
}

// home handler function that returns HTTP resonse
func handleHome( w http.ResponseWriter, r *http.Request ) {
    if r.URL.Path != "/" {
        http.NotFound( w, r )
        return
    }
    w.Write( []byte ( "Login or Dashboard redirect" ) )
}

// handler function for web app route
func handleDashboard( w http.ResponseWriter, r *http.Request ) {
    w.Write( []byte ( "Show dashboard with list of repos and creation panel" ) )
}

// handler function for web app route
func handleView( w http.ResponseWriter, r *http.Request ) {
    w.Write( []byte ( "Show specific git repo" ) )
}

// handler function for web app route
func handleCreate( w http.ResponseWriter, r *http.Request ) {
    if ! noPostThrow405( w, r ) {
        w.Write( []byte ( "Login or Dashboard redirect" ) )
    }
}

// handler function for web app route
func handleGit( w http.ResponseWriter, r *http.Request ) {
    if ! noPostThrow405( w, r ) {
        if r.URL.Path == "/git/pull" {
            w.Write( []byte ( "Pull from Git origin, create ZIP or TAR or TGZ archive and provide it." ) )
        } else if r.URL.Path == "/git/push" {
            w.Write( []byte ( "Uploading / diffing (1st step) or pushing (2nd step) git" ) )
        } else {
            http.NotFound( w, r )
            return
        }
    }
}

// handler function for web app route
func handleDelete( w http.ResponseWriter, r *http.Request ) {
    if r.Method != http.MethodPost {
        // produce page to confirm deletion
        w.Write( []byte ( "Confirmation page for deletion" ) )
    } else {
        w.Write( []byte ( "Deletion of Git project" ) )
    }
}

// handler function for web app route
func noPostThrow405( w http.ResponseWriter, r *http.Request ) bool {
    if r.Method != http.MethodPost {
        w.Header().Set( "Allow", http.MethodPost )
        w.WriteHeader( 405 )
        w.Write( []byte ( "Method not alowed" ) )
        return true
    }
    return false
}

// function starting the webserver
func prepareRoutes() *http.ServeMux {
    // initialize new servemux (router)
    mux := http.NewServeMux()
    // register fixed paths for web app
    mux.HandleFunc( "/create", handleCreate )
    mux.HandleFunc( "/dashboard", handleDashboard )
    // register subtree paths for web app
    mux.HandleFunc( "/", handleHome )
    mux.HandleFunc( "/git/", handleGit )
    mux.HandleFunc( "/view/", handleView )
    mux.HandleFunc( "/delete/", handleDelete )

    return mux
}

// main function that serves the application on port 80
func main() {

    loadEnv()
    mux := prepareRoutes()

    // log message what is going on: starting webserver
    webPort := fmt.Sprint( ":", envs[ "WEB_PORT" ] )
    log.Printf( "Starting webserver on %s\n", webPort )
    // start Webserver and register non-nil-errors
    err := http.ListenAndServe( webPort, mux )
    // log errors
    log.Fatal( err )
}
