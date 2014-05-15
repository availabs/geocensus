package main

<<<<<<< HEAD
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
)

func main() {

  /*

  Part of code used for setting up the framework and preventing cors restrictions

  */

  m := martini.Classic()
   m.Use(cors.Allow(&cors.Options{
    AllowOrigins:     []string{"http://*"},
    AllowMethods:     []string{"PUT", "PATCH","GET","POST"},
=======

import (
	_ "github.com/lib/pq"
	//"database/sql"
	//"encoding/json"
	"github.com/codegangsta/martini"
	"github.com/martini-contrib/cors"
	//"github.com/user/newmath"
	"github.com/codegangsta/martini-contrib/binding"
	"github.com/user/serverAid"
	//"strings"
	//"fmt"
)

func main() {
  m := martini.Classic()
   m.Use(cors.Allow(&cors.Options{
    AllowOrigins:     []string{"http://*"},
    AllowMethods:     []string{"PUT", "PATCH","POST"},
>>>>>>> 6853a564a3669db5e472d6851e91efe2e348ae1a
    AllowHeaders:     []string{"Origin","content-type"},
    ExposeHeaders:    []string{"Content-Length","content-type"},
    AllowCredentials: true,
  }))
<<<<<<< HEAD

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
m.Post("/ACS2012_5YEAR/QuerySpecial",binding.Bind(serverAid.GeoCensusVar2{}),serverAid.Acs20105yearQuerySpecial)

/*

Below line of code runs the server allowing GET and POST requests to be made.

*/
m.Run()


=======
  m.Get("/",serverAid.HelloWorld)

   //fmt.Print(serverAid.HelloWorld());

  m.Get("/states", serverAid.StatesInit)

  m.Get("/states/:geoid", serverAid.StatesGeoid)
  //m.Post("/states/",binding.Bind(geoCensusVar{}),serverAid.StatesGeoid)



	/*:geoid converts given part of url in a parameter to be potentially used. This param is used as params["x"] 
	where x is the given parameter*/

m.Get("/states/:geoid/county",binding.Bind(serverAid.GeoCensusVar{}),serverAid.StatesGeoCounty)
m.Post("/states/:geoid/county",binding.Bind(serverAid.GeoCensusVar{}),serverAid.StatesGeoCounty)

m.Get("/states/:geoid/tracts", binding.Bind(serverAid.GeoCensusVar{}),serverAid.StatesGeoTracts)
m.Post("/states/:geoid/tracts",binding.Bind(serverAid.GeoCensusVar{}),serverAid.StatesGeoTracts)

m.Get("/states/:geoid/bg", binding.Bind(serverAid.GeoCensusVar{}),serverAid.StatesGeoBG)
m.Post("/states/:geoid/bg", binding.Bind(serverAid.GeoCensusVar{}),serverAid.StatesGeoBG)

m.Get("/ACS2010_5YEAR",serverAid.Acs20105year)
//m.Post("/ACS2010_5YEAR",binding.Bind(ACSVar{}),serverAid.Acs20105year)

m.Get("/ACS2010_5YEAR/:table_id",binding.Bind(serverAid.ACSVar{}),serverAid.Acs20105yearTable_id)
m.Post("/ACS2010_5YEAR/",binding.Bind(serverAid.ACSVar{}),serverAid.Acs20105yearTable_id)

m.Post("/ACS2010_5YEAR/states",binding.Bind(serverAid.GeoCensusVar2{}),serverAid.Acs20105yearStates)
m.Post("/ACS2012_5YEAR/states_and_counties",binding.Bind(serverAid.GeoCensusVar2{}),serverAid.Acs20105yearStatesCounties)

  m.Run()
>>>>>>> 6853a564a3669db5e472d6851e91efe2e348ae1a
}