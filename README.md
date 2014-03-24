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