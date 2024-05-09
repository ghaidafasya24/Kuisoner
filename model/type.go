package model

import "go.mongodb.org/mongo-driver/bson/primitive"

type Responden struct {
	ID           primitive.ObjectID `bson:"_id,omitempty" json:"_id,omitempty"`
	Usia         string             `bson:"usia,omitempty" json:"usia,omitempty"`
	JenisKelamin string             `bson:"jenis_kelamin,omitempty" json:"jenis_kelamin,omitempty"`
	ProgamStudi    string            `bson:"program_studi,omitempty" json:"program_studi,omitempty"`
}

type Pertanyaan struct {
	ID             primitive.ObjectID `bson:"_id,omitempty" json:"_id,omitempty"`
	NamaPertanyaan string             `bson:"nama_pertanyaan,omitempty" json:"pertanyaan,omitempty"`
}

type Jawaban struct {
	ID           primitive.ObjectID `bson:"_id,omitempty" json:"_id,omitempty"`
	NamaJawaban  string             `bson:"jawaban,omitempty" json:"jawaban,omitempty"`
	TanggalJawab string 			`bson:"tanggal_jawab,omitempty" json:"tanggal_jawab,omitempty"`
}
