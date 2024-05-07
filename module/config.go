package _714220031

import (
	"github.com/aiteung/atdb"
	"os"
)

var MongoString string = os.Getenv("MONGOSTRING")

var MongoInfo = atdb.DBInfo{
	DBString: MongoString,
	DBName:   "TugasWeek04_Kuisoner",
}

var MongoConn = atdb.MongoConnect(MongoInfo)