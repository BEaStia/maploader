# maploader
Simple POSTGIS wrapper for taking nearest buildings from map

To use this wrapper you need:
1) Load .pbf dump from OSM dumps
2) Use Osm2pgsql to load it into database with parameter -l. For example: 
``` osm2pgsql -l --create --database gis_moscow ~/Downloads/RU-MOW.osm.pbf. You need to install POSTGIS extension into database.```
3) Run scripts from scripts.sql to prepare database.
4) Provide necessary connection variables and use this library:
```var mapLoader = maploader.MapLoader{Username:"postgres", Password:"1234", Database:"gis_moscow", Addr:"localhost:5432"}
print(string(mapLoader.FindNearest(37.6626296, 55.760102)));```