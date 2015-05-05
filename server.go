package main

import (
    "fmt"
    "github.com/drone/routes"
    "net/http"
    _ "github.com/go-sql-driver/mysql"
    "database/sql" 
    "strings"
    "io/ioutil"
)

func getholeinfo(holeid string) string {
    db, err := sql.Open("mysql", "root:st4TT5RN@/golf?charset=utf8")
    defer db.Close()
    checkErr(err)

    hid := strings.Split(holeid, ":");
    var resp, name string     
    var id, hole, yard, handicap, par int
    var glat, glon, tblat, tblon float32    
    
    // query
    err = db.QueryRow("SELECT * FROM course where id = ?", hid[1]).Scan(&id, &name, &hole, &par, &yard, &handicap, &glat, &glon, &tblat, &tblon)
    if (err != nil) {
	resp = "No value found!"
    } else {
	resp = fmt.Sprintf("%d, %s, %d, %d, %d, %d, %f, %f, %f, %f", id, name, hole, par, yard, handicap, glat, glon, tblat, tblon)
    }
    return resp
}

func checkErr(err error) {
}

func gethole(w http.ResponseWriter, r *http.Request) {
    params := r.URL.Query()
    holeid := params.Get(":holeid")
    //hid := strings.Split(holeid, ":");
    fmt.Fprintf(w, "%s", getholeinfo(holeid))
}

func modifyhole(w http.ResponseWriter, r *http.Request) {
    params := r.URL.Query()
    holeid := params.Get(":holeid")
    fmt.Fprintf(w, "modify hole %s", holeid)
}

func deletehole(w http.ResponseWriter, r *http.Request) {
    params := r.URL.Query()
    holeid := params.Get(":holeid")
    fmt.Fprintf(w, "delete hole %s", holeid)
}

func addhole(w http.ResponseWriter, r *http.Request) {
    params := r.URL.Query()
    holeid := params.Get(":holeid")
    fmt.Fprint(w, "add hole %s", holeid)
}

func loadPage2(w http.ResponseWriter, r *http.Request)  {
    filename := "/Users/trini/golfproject/sunnyvale.v2"
    body, _ := ioutil.ReadFile(filename)
    fmt.Fprint(w, string(body))
}
func loadPage(w http.ResponseWriter, r *http.Request)  {
    filename := "/Users/trini/golfproject/svgc.html"
    body, _ := ioutil.ReadFile(filename)
    fmt.Fprint(w, string(body))
}

func main() {
    mux := routes.New()
    mux.Get("/hole/:holeid", gethole)
    mux.Post("/hole/:holeid", modifyhole)
    mux.Del("/hole/:holeid", deletehole)
    mux.Put("/hole/", addhole)
    mux.Get("/hole/", loadPage)
    mux.Get("/holes/", loadPage2)
    http.Handle("/", mux)
    http.ListenAndServe(":8088", nil)
}
