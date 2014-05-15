package server_aid


import (
	_ "github.com/lib/pq"
	"database/sql"
	"encoding/json"
	"github.com/codegangsta/martini"
	//"github.com/martini-contrib/cors"
	//"github.com/codegangsta/martini-contrib/binding"
	"strings"
	"fmt"
)

type geoCensusVar struct {
    GeoCenVar   string 	`form:"geoCen" json:"states"`
    GeoCenVar2  string  `form:"state" json:"states"`
}

func helloWorld() string {
    return "Hey world!"
}

func statesInit() string {
    
	db, err := sql.Open("postgres", "host=lor.availabs.org password=transit user=postgres dbname=geocensus sslmode=require")
	if err != nil {
		return err.Error()
	}

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
}

func statesGeoCounty(params martini.Params,CountyPostA geoCensusVar) string {

	db, err := sql.Open("postgres", "host=lor.availabs.org password=transit user=postgres dbname=geocensus sslmode=require")
	if err != nil {
		return err.Error()
	}
	var sql_statement string
	sql_statement = "SELECT ST_AsGeoJSON(the_geom) as geom,namelsad,geoid FROM tl_2013_us_county"
	
	if CountyPostA.GeoCenVar != "" && CountyPostA.GeoCenVar2 != ""{
		fmt.Print(CountyPostA.GeoCenVar+" Test\n")
		sql_statement += " WHERE countyfp in "+CountyPostA.GeoCenVar+ "AND statefp in "+CountyPostA.GeoCenVar2
		fmt.Print(sql_statement)
	} else{
		fmt.Print(CountyPostA.GeoCenVar+" Test\n")
		fmt.Print("This is empty! County With geoid:"+params["geoid"]+"\n")
		sql_statement += " WHERE statefp = '"+params["geoid"]+"'"
	}

	/*The following code selects the information needed from the above host to be accessed and used here. 
	$1 is replaced with the given parameter*/

	rows, err := db.Query(sql_statement)
	if err != nil {
		return err.Error()
	}

	/*The below variables are strings that contain the selected information from the above query*/

	var namelsad string
	var geom string
	var geoid string
	counties :=[]map[string]string{}
	for rows.Next() {
        	if err := rows.Scan(&geom, &namelsad, &geoid); err != nil {
            	return err.Error()
        	}
        	county := map[string]string{
            	"geometry": geom ,
            	"Properties": "{\"geoid\": \""+geoid+"\", \"namelsad\": \""+namelsad+"\"}",
            	"type":"Feature",
            }
        	counties = append(counties, county)
        }
        b, err := json.Marshal(counties)
        c := strings.Replace(string(b), "\\","",-1)
        c = strings.Replace(c, "Properties\":\"", "Properties\":", -1)
		c = strings.Replace(c, "\",\"geometry", ",\"geometry", -1)
		c = strings.Replace(c, "geometry\":\"{", "geometry\":{", -1)
		c = strings.Replace(c, "\",\"type", ",\"type", -1)
        d := "{\n  \"type\": \"FeatureCollection\",\n  \"features\": \n"+c+"}"
    //result := "\n"+string(counties)+""
	return d//result

}

func statesGeoTracts(params martini.Params, GeoTractsA geoCensusVar) string {

	db, err := sql.Open("postgres", "host=lor.availabs.org password=transit user=postgres dbname=geocensus sslmode=require")
	if err != nil {
		return err.Error()
	}
	var sql_statement string
	if GeoTractsA.GeoCenVar2 != ""{
	sql_statement = "SELECT ST_AsGeoJSON(the_geom) as geom,namelsad,geoid FROM tl_2013_"+GeoTractsA.GeoCenVar2+"_tract"
	} else{
		sql_statement = "SELECT ST_AsGeoJSON(the_geom) as geom,namelsad,geoid FROM tl_2013_"+params["geoid"]+"_tract"
	} 
	if GeoTractsA.GeoCenVar != ""{
		fmt.Print(GeoTractsA.GeoCenVar+" Test\n")
		sql_statement += " WHERE countyfp in "+GeoTractsA.GeoCenVar
		fmt.Print(sql_statement)
	} else{
		fmt.Print("This is empty! BG With geoid:"+GeoTractsA.GeoCenVar+"\n")
		//sql_statement += " WHERE geoid = '"+params["geoid"]+"'"
	}
	rows, err := db.Query(sql_statement)
	if err != nil {
		return err.Error()
	}	

	var namelsad string
	var geom string
	var geoid string
	tracts :=[]map[string]string{}
	for rows.Next() {
        	if err := rows.Scan(&geom, &namelsad, &geoid); err != nil {
            	return err.Error()
        	}
        	tract := map[string]string{
            	"geometry": geom ,
            	"Properties": "{\"geoid\": \""+geoid+"\", \"namelsad\": \""+namelsad+"\"}",
            	"type":"Feature",
            }
        	tracts = append(tracts, tract)
        }
        b, err := json.Marshal(tracts)
        c := strings.Replace(string(b), "\\","",-1)
        c = strings.Replace(c, "Properties\":\"", "Properties\":", -1)
		c = strings.Replace(c, "\",\"geometry", ",\"geometry", -1)
		c = strings.Replace(c, "geometry\":\"{", "geometry\":{", -1)
		c = strings.Replace(c, "\",\"type", ",\"type", -1)
        d := "{\n  \"type\": \"FeatureCollection\",\n  \"features\": \n"+c+"}"
    //result := "\n"+string(counties)+""
	return d//result

	}

func statesGeoBG(params martini.Params, GeoBGA geoCensusVar) string {

	db, err := sql.Open("postgres", "host=lor.availabs.org password=transit user=postgres dbname=geocensus sslmode=require")
	if err != nil {
		return err.Error()
	}
	var sql_statement string
	if GeoBGA.GeoCenVar2 != ""{
	sql_statement = "SELECT ST_AsGeoJSON(the_geom) as geom,namelsad,geoid FROM tl_2013_"+GeoBGA.GeoCenVar2+"_bg"
	} else {
	sql_statement = "SELECT ST_AsGeoJSON(the_geom) as geom,namelsad,geoid FROM tl_2013_"+params["geoid"]+"_bg"	
	}
	if GeoBGA.GeoCenVar != ""{
		fmt.Print(GeoBGA.GeoCenVar+" Test\n")
		sql_statement += " WHERE countyfp in "+GeoBGA.GeoCenVar
		fmt.Print(sql_statement)
	} else{
		fmt.Print("This is empty! BG With geoid:"+GeoBGA.GeoCenVar+"\n")
		//sql_statement += " WHERE geoid = '"+params["geoid"]+"'"
	}
	rows, err := db.Query(sql_statement)
	if err != nil {
		return err.Error()
	}

	var namelsad string
	var geom string
	var geoid string
	bgs :=[]map[string]string{}
	for rows.Next() {
        	if err := rows.Scan(&geom, &namelsad, &geoid); err != nil {
            	return err.Error()
        	}
        	bg := map[string]string{
            	"geometry": geom ,
            	"Properties": "{\"geoid\": \""+geoid+"\", \"namelsad\": \""+namelsad+"\"}",
            	"type":"Feature",
            }
        	bgs = append(bgs, bg)
        }
        b, err := json.Marshal(bgs)
        c := strings.Replace(string(b), "\\","",-1)
        c = strings.Replace(c, "Properties\":\"", "Properties\":", -1)
		c = strings.Replace(c, "\",\"geometry", ",\"geometry", -1)
		c = strings.Replace(c, "geometry\":\"{", "geometry\":{", -1)
		c = strings.Replace(c, "\",\"type", ",\"type", -1)
        d := "{\n  \"type\": \"FeatureCollection\",\n  \"features\": \n"+c+"}"
    //result := "\n"+string(counties)+""
	return d//result

	}

func statesGeoid(params martini.Params,StatesPostA geoCensusVar) string {
  	db, err := sql.Open("postgres", "host=lor.availabs.org password=transit user=postgres dbname=geocensus sslmode=require")
	if err != nil {
		return err.Error()
	}
	var sql_statement string
	sql_statement = "SELECT geoid, name,ST_AsGeoJSON(the_geom) as geo FROM tl_2013_us_state"
	
	if StatesPostA.GeoCenVar != ""{
		fmt.Print(StatesPostA.GeoCenVar+" Test\n")
		sql_statement += " WHERE geoid in "+StatesPostA.GeoCenVar
		fmt.Print(sql_statement)
	} else{
		fmt.Print("This is empty! States With geoid:"+params["geoid"]+"\n")
		sql_statement += " WHERE geoid = '"+params["geoid"]+"'"
	}
	rows, err := db.Query(sql_statement)
	if err != nil {
		return err.Error()
	}
	states :=[]map[string]string{}
	for rows.Next() {
			var name string
			var geoid string
			var geo string
            if err := rows.Scan(&geoid, &name, &geo); err != nil {
                    return err.Error()
            }

            state := map[string]string{
            	"geometry": geo ,
            	"Properties": "{\"geoid\": "+geoid+"}",
            	"type":"Feature",
            }

            states = append(states, state)

    }
	b, err := json.Marshal(states)

	/*The Following lines of code are used to parse and remove unnecessary characters that get in the way of 
	creating proper geojson files*/

	c := strings.Replace(string(b), "\\","",-1)
    c = strings.Replace(c, "Properties\":\"", "Properties\":", -1)
	c = strings.Replace(c, "\",\"geometry", ",\"geometry", -1)
	c = strings.Replace(c, "geometry\":\"{", "geometry\":{", -1)
	c = strings.Replace(c, "\",\"type", ",\"type", -1)
    d := "{\n  \"type\": \"FeatureCollection\",\n  \"features\": \n"+c+"}"
	return d
}