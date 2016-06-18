package maploader;
import (
	"gopkg.in/pg.v4"
	"fmt"
	"encoding/json"
)


type MapLoader struct {
	Db		*pg.DB
	Username	string
	Password	string
	Database	string
	Addr		string
}


type PlanetOsmPolygon struct {
	OsmId 		string
	Amenity 	string
	Building 	string
	Name 		string
	WayJson 	string
	Distance	float64
}

func (m MapLoader) FindNearest(x float64, y float64) []byte {

	m.Db = pg.Connect(&pg.Options{
		User: 		m.Username,
		Password: 	m.Password,
		Database: 	m.Database,
		Addr: 		m.Addr,
	})

	var buildings []PlanetOsmPolygon

	var query = fmt.Sprintf(`
	SELECT osm_id,
	       amenity,
	       name,
	       way_json::json as way_json,
	       ST_Distance(column1::geography, centroid::geography) as distance
	FROM planet_osm_polygons,
	     (
	      VALUES(ST_PointFromText('Point( %.8f %.8f)', 4326)) ) AS new_value
	WHERE centroid IS NOT NULL
	  AND area IS NULL
	  AND building IS NOT NULL
	  AND ST_Distance(ST_Transform(column1, 26986), ST_Transform(centroid, 26986)) < 1000
	ORDER BY ST_Distance(column1::geography, centroid::geography) LIMIT 2;`, x, y)
	print(query)
	_, err := m.Db.Query(&buildings, query)
	if err != nil {
		panic(err)
	}

	fmt.Println(len(buildings))
	fmt.Println(buildings[0].OsmId, buildings[0].Amenity, buildings[0].Building, buildings[0].Name, buildings[0].WayJson)
	jsonned, err := json.Marshal(buildings)
	if err != nil {
		panic(err)
	}
	return jsonned
}