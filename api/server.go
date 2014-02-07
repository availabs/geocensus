package main


import (
	_ "github.com/lib/pq"
	"database/sql"
	"encoding/json"
	"github.com/codegangsta/martini"
	//"fmt"
)

type geojson struct {
    Type        string
    Features 	feature
    Point       Point
    Line        Line
    Polygon     Polygon
}

type feature struct {
	Type        string
	Geometry    geometry
	Properties  map[string]string
}

type geometry struct {
	Type        string
	Coordinates []map[float]float{}

func main() {
  m := martini.Classic()
  m.Get("/", func() string {
    return "Hello world!"
  })
  m.Get("/states", func() string {
    
	db, err := sql.Open("postgres", "host=lor.availabs.org password=transit user=postgres dbname=geocensus sslmode=require")
	if err != nil {
		return err.Error()
	}

	//age := 21
	rows, err := db.Query("SELECT geoid, name FROM tl_2013_us_state ")
	if err != nil {
		return err.Error()
	}
	states :=[]map[string]string{}
	for rows.Next() {
			var name string
			var geoid string
            if err := rows.Scan(&geoid, &name); err != nil {
                    return err.Error()
            }
            state := map[string]string{
            	"GEOID":geoid,
            	"Name":name,
            }
            states = append(states, state)
    }
	b, err := json.Marshal(states)
	return string(b)
  })

  m.Get("/states/:geoid", func(params martini.Params) string {
  	db, err := sql.Open("postgres", "host=lor.availabs.org password=transit user=postgres dbname=geocensus sslmode=require")
	if err != nil {
		return err.Error()
	}

	//age := 21
	rows, err := db.Query("SELECT geoid, name,ST_AsGeoJSON(the_geom) as geo FROM tl_2013_us_state WHERE geoid = $1",params["geoid"])
	if err != nil {
		return err.Error()
	}
	states :=[]map[string]string{}
	for rows.Next() {
			var name string
			var geoid string
			var geo string
            if err := rows.Scan(&geoid, &name,&geo); err != nil {
                    return err.Error()
            }
            state := map[string]string{
            	"geometry": geo,
            	"type":"feature",
            }
            states = append(states, state)
    }
	b, err := json.Marshal(states)
	return string(b)

  //return "Hello " + params["geoid"]
})

  m.Run()
}