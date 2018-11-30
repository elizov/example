package geo

import (
	"log"
	"os"

	"github.com/oschwald/geoip2-golang"
)

var db *geoip2.Reader

// GetGeoIP get GeoIP instance
func GetGeoIP() *geoip2.Reader {
	if db == nil {
		GeoIPmmdb := os.Getenv("GEOIP_DB")
		var err error
		if db, err = geoip2.Open(GeoIPmmdb); err != nil {
			log.Fatal(err)
		}
	}

	return db
}
