package main


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
    AllowHeaders:     []string{"Origin","content-type"},
    ExposeHeaders:    []string{"Content-Length","content-type"},
    AllowCredentials: true,
  }))
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
}