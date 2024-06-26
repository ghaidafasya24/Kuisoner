package _714220031

import (
	"fmt"
	"testing"

	"github.com/ghaidafasya24/Kuisoner/module"
	// "github.com/ghaidafasya24/Kuisoner/model"
	"go.mongodb.org/mongo-driver/bson/primitive"

)

func TestInsertResponden(t *testing.T) {
	usia := "22"
	JenisKelamin := "Laki-laki"
	ProgramStudi := "Teknik Informatika"
	insertedID := module.InsertResponden(usia, JenisKelamin, ProgramStudi)
	fmt.Println(insertedID)
}

func TestGetRespondenByID(t *testing.T) {
	respondenID, err := primitive.ObjectIDFromHex("615af14ae62f4c488e1d6d14")
	if err != nil {
		fmt.Printf("Error: %v\n", err)
		return
	}
	responden := module.GetRespondenByID(respondenID)
	fmt.Println(responden)
}

func TestGetAllResponden(t *testing.T) {
	respondens := module.GetAllResponden()
	fmt.Println(respondens)
}

func TestInsertPertanyaan(t *testing.T) {
	NamaPertanyaan := "Seberapa relevan topik yang dibahas dalam tugas akhir ini dengan minat Anda?"
	insertedID := module.InsertPertanyaan(NamaPertanyaan)
	fmt.Println(insertedID)
}

func TestGetPertanyaanByID(t *testing.T) {
	pertanyaanID, err := primitive.ObjectIDFromHex("615af14ae62f4c488e1d6d14")
	if err != nil {
		fmt.Printf("Error: %v\n", err)
		return
	}
	pertanyaan := module.GetPertanyaanByID(pertanyaanID)
	fmt.Println(pertanyaan)
}

func TestGetAllPertanyaan(t *testing.T) {
	pertanyaans := module.GetAllPertanyaan()
	fmt.Println(pertanyaans)
}

func TestInsertJawaban(t *testing.T) {
	NamaJawaban := "Sangat relevan dengan perkembangan terkini dalam bidangnya."
	TanggalJawab := "2024-04-16"
	insertedID := module.InsertJawaban(NamaJawaban, TanggalJawab)
	fmt.Println(insertedID)
}
