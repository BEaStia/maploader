alter table planet_osm_polygon add column way_json jsonb;
update planet_osm_polygon set way_json = st_asgeojson(way)::jsonb;
SELECT AddGeometryColumn ('planet_osm_polygon','centroid',4326,'POINT',2);
UPDATE planet_osm_polygon set centroid = st_centroid(way);
CREATE INDEX planet_osm_polygon_centroid_index ON planet_osm_polygon USING GIST (centroid);
ALTER TABLE planet_osm_polygon RENAME TO planet_osm_polygons;