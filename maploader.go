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
	WayJson 	string
}


func (m MapLoader) FindNearest(x float64, y float64) string {

	m.Db = pg.Connect(&pg.Options{
		User: 		m.Username,
		Password: 	m.Password,
		Database: 	m.Database,
		Addr: 		m.Addr,
	})

	var buildings []PlanetOsmPolygon

	var query = fmt.Sprintf(`
	SELECT way_json::json->'coordinates'->0 as way_json
	FROM planet_osm_polygons,
	     (
	      VALUES(ST_PointFromText('Point( %.8f %.8f)', 4326)) ) AS new_value
	WHERE centroid IS NOT NULL
	  AND area IS NULL
	  AND building IS NOT NULL
	  AND ST_Distance(ST_Transform(column1, 26986), ST_Transform(centroid, 26986)) < 1000
	ORDER BY ST_Distance(column1::geography, centroid::geography) LIMIT 20;`, x, y)
	print(query)
	_, err := m.Db.Query(&buildings, query)
	if err != nil {
		panic(err)
	}

	var result = []string{}
	for _, building := range buildings {
		result = append(result, building.WayJson)
	}

	jsonned, err := json.Marshal(result)
	if err != nil {
		panic(err)
	}
	return string(jsonned)
}