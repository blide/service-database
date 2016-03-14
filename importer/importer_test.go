package importer_test

import (
  "testing"
  "github.com/GreenNav/service-database/importer"
  "github.com/GreenNav/service-database/database"
  "github.com/GreenNav/service-database/database/sqlite"
)

func TestWriteToDatabase(t *testing.T) {
  db, _ := sqlite.CreateEmpty("./test_importer.db")
  importer.WriteToDatabase("./monaco-20150428.osm.pbf", database.OSMDatabase(db))
}