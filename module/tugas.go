package module

import (
	"context"
	// "errors"
	"fmt"
	"github.com/ghaidafasya24/Kuisoner/model"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"

	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	// "os"
	// "time"
)

// var MongoString string = os.Getenv("MONGOSTRING")

// CONNECT DATABASE
func MongoConnect(dbname string) (db *mongo.Database) {
	client, err := mongo.Connect(context.TODO(), options.Client().ApplyURI(MongoString))
	if err != nil {
		fmt.Printf("MongoConnect: %v\n", err)
	}
	return client.Database(dbname)
}

func InsertOneDoc(db string, collection string, doc interface{}) (insertedID interface{}) {
	insertResult, err := MongoConnect(db).Collection(collection).InsertOne(context.TODO(), doc)
	if err != nil {
		fmt.Printf("InsertOneDoc: %v\n", err)
	}
	return insertResult.InsertedID
}

func InsertResponden(usia string, jenisKelamin string, programStudi string) (insertedID interface{}) {
	var responden model.Responden
	responden.ID = primitive.NewObjectID()
	responden.Usia = usia
	responden.JenisKelamin = jenisKelamin
	responden.ProgramStudi = programStudi
	// responden.TanggalPengisian = tanggal_pengisian
	return InsertOneDoc("TugasWeek04_Kuisoner", "responden", responden)
}


func GetRespondenByID(respondenID primitive.ObjectID) (responden model.Responden) {
	collection := MongoConnect("tesdb2024").Collection("responden")
	filter := bson.M{"_id": respondenID}
	err := collection.FindOne(context.TODO(), filter).Decode(&responden)
	if err != nil {
		fmt.Printf("GetRespondenByID: %v\n", err)
	}
	return responden
}

func GetAllResponden() (respondens []model.Responden) {
	collection := MongoConnect("TugasWeek04_Kuisoner").Collection("responden")
	cursor, err := collection.Find(context.TODO(), bson.D{})
	if err != nil {
		fmt.Printf("GetAllResponden: %v\n", err)
		return nil
	}
	defer cursor.Close(context.TODO())
	for cursor.Next(context.Background()) {
		var responden model.Responden
		if err := cursor.Decode(&responden); err != nil {
			fmt.Printf("GetAllResponden: %v\n", err)
			continue
		}
		respondens = append(respondens, responden)
	}
	if err := cursor.Err(); err != nil {
		fmt.Printf("GetAllPelanggan: %v\n", err)
	}
	return respondens
}

func InsertPertanyaan(NamaPertanyaan string) interface{} {
	var pertanyaan model.Pertanyaan
	pertanyaan.ID = primitive.NewObjectID()
	pertanyaan.NamaPertanyaan = NamaPertanyaan
	return InsertOneDoc("TugasWeek04_Kuisoner", "pertanyaan", pertanyaan)
}

func GetPertanyaanByID(pertanyaanID primitive.ObjectID) (pertanyaan model.Pertanyaan) {
	collection := MongoConnect("TugasWeek04_Kuisoner").Collection("pertanyaan")
	filter := bson.M{"_id": pertanyaanID}
	err := collection.FindOne(context.TODO(), filter).Decode(&pertanyaan)
	if err != nil {
		fmt.Printf("GetPertanyaanByID: %v\n", err)
	}
	return pertanyaan
}

func GetAllPertanyaan() (pertanyaans []model.Pertanyaan) {
	collection := MongoConnect("TugasWeek04_Kuisoner").Collection("pertanyaan")
	cursor, err := collection.Find(context.TODO(), bson.D{})
	if err != nil {
		fmt.Printf("GetAllPertanyaan: %v\n", err)
		return nil
	}
	defer cursor.Close(context.TODO())
	for cursor.Next(context.Background()) {
		var pertanyaan model.Pertanyaan
		if err := cursor.Decode(&pertanyaan); err != nil {
			fmt.Printf("GetAllPertanyaan: %v\n", err)
			continue
		}
		pertanyaans = append(pertanyaans, pertanyaan)
	}
	if err := cursor.Err(); err != nil {
		fmt.Printf("GetAllPertanyaan: %v\n", err)
	}
	return pertanyaans
}

func InsertJawaban(NamaJawaban string, TanggalJawab string) interface{} {
	var jawaban model.Jawaban
	jawaban.ID = primitive.NewObjectID()
	jawaban.NamaJawaban = NamaJawaban
	jawaban.TanggalJawab = TanggalJawab
	return InsertOneDoc("TugasWeek04_Kuisoner", "jawaban", jawaban)
}

func GetJawabanByID(jawabanID primitive.ObjectID) (jawaban model.Jawaban) {
	collection := MongoConnect("TugasWeek04_Kuisoner").Collection("jawaban")
	filter := bson.M{"_id": jawabanID}
	err := collection.FindOne(context.TODO(), filter).Decode(&jawaban)
	if err != nil {
		fmt.Printf("GetJawabanByID: %v\n", err)
	}
	return jawaban
}

func GetAllJawaban() (jawabans []model.Jawaban) {
	collection := MongoConnect("TugasWeek04_Kuisoner").Collection("jawaban")
	cursor, err := collection.Find(context.TODO(), bson.D{})
	if err != nil {
		fmt.Printf("GetAllJawaban: %v\n", err)
		return nil
	}
	defer cursor.Close(context.TODO())
	for cursor.Next(context.Background()) {
		var jawaban model.Jawaban
		if err := cursor.Decode(&jawabans); err != nil {
			fmt.Printf("GetAllJawaban: %v\n", err)
			continue
		}
		jawabans = append(jawabans, jawaban)
	}
	if err := cursor.Err(); err != nil {
		fmt.Printf("GetAllJawaban: %v\n", err)
	}
	return jawabans
}
