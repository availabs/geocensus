How do You make go work?? Good Question.

1 - If go is not in $PATH
export PATH=$PATH:/usr/local/go/bin

2 - IF GO ROOT IS NOT SET 
export GOROOT=/usr/local/go

3 - IF GOPATH is not SET
export GOPATH=project_directory/api
example
export GOPATH=/home/ef92/code/geocensus/api

Install Stuff we need
go get github.com/codegangsta/martini
go get github.com/lib/pq




v = ["12",14,123,235,46,346,346];
v[0];
v = {"test":0,"test2":425}
v.test2

outpyt = getDataFromsite();

geo = output[0]

get.geometry.cooridinates[0]


Example code of a multifile program using Martini:

https://github.com/ozonesurfer/TiedotMartini2

Tomorrow:

Make client for two new routes.

Have drop down select box.

Should have all the options in the main route.
<select id="cool_things">
 for each option in Main options 
 <option value=tableid>Stub</option>
</select>
$('#cool_things').on("change",function(){
	call to sub route with this.val();
	$('body.append').html(data from sub route);

})

Start from command:

https://github.com/leehach/census-postgres

insert_into_moe.sql

//Joining Data Variables to Sequence Id
SELECT seqeunce_number
  FROM table_shell 
  JOIN data_dictionary 
	ON table_shell.table_id = data_dictionary.table_id 
		AND table_shell.var_order = data_dictionary.line_number 
  WHERE table_shell.varnum = 'B01001001';

 ##table and row count listing
 SELECT schemaname,relname,n_live_tup 
  FROM pg_stat_user_tables WHERE schemaname = 'acs2010_5yr'
  ORDER BY n_live_tup ASC;

  ##postgres copy csv syntax
  COPY data_dictionary FROM '/media/Data/acs/prod/acs2010_5yr/Sequence_Number_and_Table_Number_Lookup.txt' WITH CSV HEADER;

##for State Select
select b01001001 from "acs2010_5yr".seq0010 as a
join "acs2010_5yr".geoheader as b
ON a.logrecno = b.logrecno 
where b.sumlevel='040'

if states is passed
 and b.state in ('36','47') 

//For County Select
select b01001001 from "acs2010_5yr".seq0010 as a
join "acs2010_5yr".geoheader as b
ON a.logrecno = b.logrecno and a.stusab = b.stusab
where b.sumlevel='050' and b.state = '36' and b.county = '001'



TO DO LIST:

3. Speed up scripts for tracts

5. Add documentation to:

Map

Server Aid

Server

6. Fix Asynchronus problem as current workaround is not recommended

8. Add in single state, tract, and county support





567669

COPY  tmp_geoheader  FROM  '/media/Data/acs/prod/acs2010_5yr/All_Geographies_Not_Tracts_Block_Groups/g20105az.txt';
COPY  tmp_geoheader  FROM  '/media/Data/acs/prod/acs2010_5yr/All_Geographies_Not_Tracts_Block_Groups/g20105ca.txt';
COPY  tmp_geoheader  FROM  '/media/Data/acs/prod/acs2010_5yr/All_Geographies_Not_Tracts_Block_Groups/g20105co.txt';
COPY  tmp_geoheader  FROM  '/media/Data/acs/prod/acs2010_5yr/All_Geographies_Not_Tracts_Block_Groups/g20105tx.txt';
COPY  tmp_geoheader  FROM  '/media/Data/acs/prod/acs2010_5yr/All_Geographies_Not_Tracts_Block_Groups/g20105pr.txt';
COPY  tmp_geoheader  FROM  '/media/Data/acs/prod/acs2010_5yr/All_Geographies_Not_Tracts_Block_Groups/g20105us.txt';

COPY tmp_geoheader FROM '/home/avail/code/census-postgres/acs2010_5yr/g20105az.txt';
COPY tmp_geoheader FROM '/home/avail/code/census-postgres/acs2010_5yr/g20105ca.txt';
COPY tmp_geoheader FROM '/home/avail/code/census-postgres/acs2010_5yr/g20105co.txt';
COPY tmp_geoheader FROM '/home/avail/code/census-postgres/acs2010_5yr/g20105tx.txt';
COPY tmp_geoheader FROM '/home/avail/code/census-postgres/acs2010_5yr/g20105pr.txt';
COPY tmp_geoheader FROM '/home/avail/code/census-postgres/acs2010_5yr/g20105us.txt';


SET search_path = acs2010_5yr;

COPY  tmp_geoheader  FROM  '/media/Data/acs/prod/acs2010_5yr/All_Geographies_Not_Tracts_Block_Groups/g20105ak.txt';
COPY  tmp_geoheader  FROM  '/media/Data/acs/prod/acs2010_5yr/All_Geographies_Not_Tracts_Block_Groups/g20105al.txt';
COPY  tmp_geoheader  FROM  '/media/Data/acs/prod/acs2010_5yr/All_Geographies_Not_Tracts_Block_Groups/g20105ar.txt';
COPY  tmp_geoheader  FROM  '/media/Data/acs/prod/acs2010_5yr/All_Geographies_Not_Tracts_Block_Groups/g20105ct.txt';
COPY  tmp_geoheader  FROM  '/media/Data/acs/prod/acs2010_5yr/All_Geographies_Not_Tracts_Block_Groups/g20105dc.txt';
COPY  tmp_geoheader  FROM  '/media/Data/acs/prod/acs2010_5yr/All_Geographies_Not_Tracts_Block_Groups/g20105de.txt';
COPY  tmp_geoheader  FROM  '/media/Data/acs/prod/acs2010_5yr/All_Geographies_Not_Tracts_Block_Groups/g20105fl.txt';
COPY  tmp_geoheader  FROM  '/media/Data/acs/prod/acs2010_5yr/All_Geographies_Not_Tracts_Block_Groups/g20105ga.txt';
COPY  tmp_geoheader  FROM  '/media/Data/acs/prod/acs2010_5yr/All_Geographies_Not_Tracts_Block_Groups/g20105hi.txt';
COPY  tmp_geoheader  FROM  '/media/Data/acs/prod/acs2010_5yr/All_Geographies_Not_Tracts_Block_Groups/g20105ia.txt';
COPY  tmp_geoheader  FROM  '/media/Data/acs/prod/acs2010_5yr/All_Geographies_Not_Tracts_Block_Groups/g20105id.txt';
COPY  tmp_geoheader  FROM  '/media/Data/acs/prod/acs2010_5yr/All_Geographies_Not_Tracts_Block_Groups/g20105il.txt';
COPY  tmp_geoheader  FROM  '/media/Data/acs/prod/acs2010_5yr/All_Geographies_Not_Tracts_Block_Groups/g20105in.txt';
COPY  tmp_geoheader  FROM  '/media/Data/acs/prod/acs2010_5yr/All_Geographies_Not_Tracts_Block_Groups/g20105ks.txt';
COPY  tmp_geoheader  FROM  '/media/Data/acs/prod/acs2010_5yr/All_Geographies_Not_Tracts_Block_Groups/g20105ky.txt';
COPY  tmp_geoheader  FROM  '/media/Data/acs/prod/acs2010_5yr/All_Geographies_Not_Tracts_Block_Groups/g20105la.txt';
COPY  tmp_geoheader  FROM  '/media/Data/acs/prod/acs2010_5yr/All_Geographies_Not_Tracts_Block_Groups/g20105ma.txt';
COPY  tmp_geoheader  FROM  '/media/Data/acs/prod/acs2010_5yr/All_Geographies_Not_Tracts_Block_Groups/g20105md.txt';
COPY  tmp_geoheader  FROM  '/media/Data/acs/prod/acs2010_5yr/All_Geographies_Not_Tracts_Block_Groups/g20105me.txt';
COPY  tmp_geoheader  FROM  '/media/Data/acs/prod/acs2010_5yr/All_Geographies_Not_Tracts_Block_Groups/g20105mi.txt';
COPY  tmp_geoheader  FROM  '/media/Data/acs/prod/acs2010_5yr/All_Geographies_Not_Tracts_Block_Groups/g20105mn.txt';
COPY  tmp_geoheader  FROM  '/media/Data/acs/prod/acs2010_5yr/All_Geographies_Not_Tracts_Block_Groups/g20105mo.txt';
COPY  tmp_geoheader  FROM  '/media/Data/acs/prod/acs2010_5yr/All_Geographies_Not_Tracts_Block_Groups/g20105ms.txt';
COPY  tmp_geoheader  FROM  '/media/Data/acs/prod/acs2010_5yr/All_Geographies_Not_Tracts_Block_Groups/g20105mt.txt';
COPY  tmp_geoheader  FROM  '/media/Data/acs/prod/acs2010_5yr/All_Geographies_Not_Tracts_Block_Groups/g20105nc.txt';
COPY  tmp_geoheader  FROM  '/media/Data/acs/prod/acs2010_5yr/All_Geographies_Not_Tracts_Block_Groups/g20105nd.txt';
COPY  tmp_geoheader  FROM  '/media/Data/acs/prod/acs2010_5yr/All_Geographies_Not_Tracts_Block_Groups/g20105ne.txt';
COPY  tmp_geoheader  FROM  '/media/Data/acs/prod/acs2010_5yr/All_Geographies_Not_Tracts_Block_Groups/g20105nh.txt';
COPY  tmp_geoheader  FROM  '/media/Data/acs/prod/acs2010_5yr/All_Geographies_Not_Tracts_Block_Groups/g20105nj.txt';
COPY  tmp_geoheader  FROM  '/media/Data/acs/prod/acs2010_5yr/All_Geographies_Not_Tracts_Block_Groups/g20105nv.txt';
COPY  tmp_geoheader  FROM  '/media/Data/acs/prod/acs2010_5yr/All_Geographies_Not_Tracts_Block_Groups/g20105ny.txt';
COPY  tmp_geoheader  FROM  '/media/Data/acs/prod/acs2010_5yr/All_Geographies_Not_Tracts_Block_Groups/g20105oh.txt';
COPY  tmp_geoheader  FROM  '/media/Data/acs/prod/acs2010_5yr/All_Geographies_Not_Tracts_Block_Groups/g20105ok.txt';
COPY  tmp_geoheader  FROM  '/media/Data/acs/prod/acs2010_5yr/All_Geographies_Not_Tracts_Block_Groups/g20105or.txt';
COPY  tmp_geoheader  FROM  '/media/Data/acs/prod/acs2010_5yr/All_Geographies_Not_Tracts_Block_Groups/g20105pa.txt';
COPY  tmp_geoheader  FROM  '/media/Data/acs/prod/acs2010_5yr/All_Geographies_Not_Tracts_Block_Groups/g20105ri.txt';
COPY  tmp_geoheader  FROM  '/media/Data/acs/prod/acs2010_5yr/All_Geographies_Not_Tracts_Block_Groups/g20105sc.txt';
COPY  tmp_geoheader  FROM  '/media/Data/acs/prod/acs2010_5yr/All_Geographies_Not_Tracts_Block_Groups/g20105sd.txt';
COPY  tmp_geoheader  FROM  '/media/Data/acs/prod/acs2010_5yr/All_Geographies_Not_Tracts_Block_Groups/g20105tn.txt';
COPY  tmp_geoheader  FROM  '/media/Data/acs/prod/acs2010_5yr/All_Geographies_Not_Tracts_Block_Groups/g20105ut.txt';
COPY  tmp_geoheader  FROM  '/media/Data/acs/prod/acs2010_5yr/All_Geographies_Not_Tracts_Block_Groups/g20105va.txt';
COPY  tmp_geoheader  FROM  '/media/Data/acs/prod/acs2010_5yr/All_Geographies_Not_Tracts_Block_Groups/g20105vt.txt';
COPY  tmp_geoheader  FROM  '/media/Data/acs/prod/acs2010_5yr/All_Geographies_Not_Tracts_Block_Groups/g20105wa.txt';
COPY  tmp_geoheader  FROM  '/media/Data/acs/prod/acs2010_5yr/All_Geographies_Not_Tracts_Block_Groups/g20105wi.txt';
COPY  tmp_geoheader  FROM  '/media/Data/acs/prod/acs2010_5yr/All_Geographies_Not_Tracts_Block_Groups/g20105wv.txt';
COPY  tmp_geoheader  FROM  '/media/Data/acs/prod/acs2010_5yr/All_Geographies_Not_Tracts_Block_Groups/g20105wy.txt';

COPY  tmp_geoheader  FROM  '/media/Data/acs/prod/acs2010_5yr/All_Geographies_Not_Tracts_Block_Groups/g20105nm.txt';
COPY tmp_geoheader FROM '/home/avail/code/census-postgres/acs2010_5yr/g20105az.txt';

COPY tmp_geoheader FROM '/home/avail/code/census-postgres/acs2010_5yr/g20105ca.txt';
COPY tmp_geoheader FROM '/home/avail/code/census-postgres/acs2010_5yr/g20105co.txt';


COPY tmp_geoheader FROM '/home/avail/code/census-postgres/acs2010_5yr/g20105tx.txt';
COPY tmp_geoheader FROM '/home/avail/code/census-postgres/acs2010_5yr/g20105pr.txt';
COPY tmp_geoheader FROM '/home/avail/code/census-postgres/acs2010_5yr/g20105us.txt';



https://docs.angularjs.org/tutorial/step_00



yo angular:controller stations
create stations view stations.html
set route to stations/:stationid
List Station ID at page Header