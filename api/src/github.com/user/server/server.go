package main

/*

Server part of Geocensus project. Authored by Evan Friedkin with help from Alex Muro.

*/


/*List of imported go libraries

  "github.com/codegangsta/martini"                            Frame work used for server actions
  "github.com/martini-contrib/cors"                           Part of martini framework. Used to handle cors requests
  "github.com/codegangsta/martini-contrib/binding"            Part of martini framework. Used to bind and pass variables from client to server
  "github.com/user/serverAid"                                 Part of server code, authored by Evan Friedkin with help from Alex Muro

*/
import (

	"github.com/codegangsta/martini"
	"github.com/martini-contrib/cors"
	"github.com/codegangsta/martini-contrib/binding"
	"github.com/user/serverAid"
  "text/template"
  "net/http"
  //"log"
  "fmt"
)

func main() {

  /*

  Part of code used for setting up the framework and preventing cors restrictions

  */
  var tmpl = template.Must(template.ParseFiles("downloads/GeocensusAPI.html"))
  m := martini.Classic()
  m.Use(cors.Allow(&cors.Options{
    AllowOrigins:     []string{"http://*"},
    AllowMethods:     []string{"PUT", "PATCH","GET","POST"},
    AllowHeaders:     []string{"Origin","content-type"},
    ExposeHeaders:    []string{"Content-Length","content-type"},
    AllowCredentials: true,
  }))
//m.Use(martini.Static("downloads"))
//m.Get("/",martini.Static("index.html"))
//fmt.Println(m)
m.Get("/",func(res http.ResponseWriter,req *http.Request){

  if err := tmpl.ExecuteTemplate(res,"GeocensusAPI.html",nil); err != nil{
    fmt.Println(res)
    http.Error(res,err.Error(),http.StatusInternalServerError)
  }

})
m.Get("/downloads",func(res http.ResponseWriter,req *http.Request){
  /*zr, err := zip.NewReader(urlReader, int64(urlReader.Len()))
    if err != nil {
      log.Fatalf("Unable to read zip: %s", err)
    }
    res(zr)*/
    //res.Header().set("Content-type","downloads/outPut.zip")
    http.ServeFile(res,req,"downloads/outPut.zip")
  })
//m.Use(martini.Static("downloads"))

/*

The following Post request is used when the client selects tracts as their search option to query the database.
This request will post a list of states to the third dropdown menu on the client side of the code.
The parameters given are a default state value to help with error checking, and the state the
client selects in the state drop down menu that only gets filled when tracts are selected.

*/


m.Post("/states/:geoid/tracts",binding.Bind(serverAid.GeoCensusVar{}),serverAid.StatesGeoTracts)

/*

The following Get request is used to get the list of searchable census data types for the client.
This is only called once when map.html is first run.

*/


m.Get("/ACS2010_5YEAR",serverAid.Acs20105year)

/*

*/

m.Post("/ACS2010_5YEAR/",binding.Bind(serverAid.ACSVar{}),serverAid.Acs20105yearTable_id)


m.Post("/ACS2010_5YEAR/states",binding.Bind(serverAid.GeoCensusVar2{}),serverAid.Acs20105yearStates)


m.Post("/ACS2012_5YEAR/Query",binding.Bind(serverAid.GeoCensusVar2{}),serverAid.Acs20105yearQuery)
m.Post("/ACS2012_5YEAR/QuerySpecial/:filetype",binding.Bind(serverAid.GeoCensusVar2{}),serverAid.Acs20105yearQuerySpecial)

/*

Below line of code runs the server allowing GET and POST requests to be made.

*/

m.Run()

}