package storage

import (
	"encoding/json"
	"os"
)

type Storage[T any] struct {
	Filename string
}

func NewStorage[T any](filename string) *Storage[T] {
	return &Storage[T]{Filename: filename}
}

/*
	Storage[T any]: Generik bir yapı tanımlanmış. T herhangi bir tür olabilir (any).
	Filename: Verilerin kaydedileceği dosyanın adını tutan bir string.
	NewStorage: Yeni bir Storage nesnesi oluşturmak için kullanılan bir fonksiyon.
	filename: Dosya adını parametre olarak alır.
	return &Storage[T]{Filename: filename}: Yeni bir Storage nesnesi oluşturur ve dosya adını Filename alanına atar.
	Bu nesneyi döner.
*/

func (s *Storage[T]) Save(data T) error {
	fileData, err := json.MarshalIndent(data, "", "  ")
	if err != nil {
		return err
	}
	return os.WriteFile(s.Filename, fileData, 0777)
}

/*
	Save: Veriyi JSON formatında dosyaya kaydeder.
	data: Kaydedilecek veriyi temsil eder.
	json.MarshalIndent(data, "", "  "): Veriyi JSON formatına dönüştürür ve girintili (indent) bir şekilde düzenler.
	Bu, JSON verisinin daha okunabilir olmasını sağlar.
	if err != nil { return err }: Eğer JSON dönüştürme sırasında bir hata oluşursa, bu hatayı döner.
	os.WriteFile(s.Filename, fileData, 0644): JSON verisini belirtilen dosyaya yazar.
	Dosya izinleri 0644 olarak ayarlanır (sahibi okuma ve yazma, diğerleri sadece okuma iznine sahip olur).
*/

func (s *Storage[T]) Load(data *T) error {
	fileData, err := os.ReadFile(s.Filename)
	if err != nil {
		return err
	}
	return json.Unmarshal(fileData, data)

}

/*
	Load: Dosyadan veriyi yükler ve verilen değişkene atar.
	data: Yüklenecek veriyi tutacak olan değişkenin adresi (pointer).
	os.ReadFile(s.Filename): Dosyayı okur ve içeriğini fileData değişkenine atar.
	if err != nil { return err }: Eğer dosya okuma sırasında bir hata oluşursa, bu hatayı döner.
	json.Unmarshal(fileData, data): JSON verisini çözümleyerek (unmarshal) data değişkenine atar.
*/

/*

Bu kod, JSON formatında verileri dosyaya kaydetmek ve dosyadan yüklemek için genel amaçlı bir yapı ve fonksiyonlar içerir.
İşte adım adım açıklaması:


type Storage[T any] struct { Filename string }: Generik bir yapı. T herhangi bir tür olabilir. Filename dosya adını tutar.

NewStorage Fonksiyonu:

func NewStorage[T any](filename string) *Storage[T]: Yeni bir Storage nesnesi oluşturur ve dosya adını atar.
Save Fonksiyonu:

func (s *Storage[T]) Save(data T) error: Veriyi JSON formatında dosyaya kaydeder.
json.MarshalIndent(data, "", "  "): Veriyi JSON formatına dönüştürür ve girintili bir şekilde düzenler.
os.WriteFile(s.Filename, fileData, 0644): JSON verisini dosyaya yazar.
Load Fonksiyonu:

func (s *Storage[T]) Load(data *T) error: Dosyadan veriyi yükler ve verilen değişkene atar.
os.ReadFile(s.Filename): Dosyayı okur.
json.Unmarshal(fileData, data): JSON verisini çözümleyerek data değişkenine atar.
Bu yapı ve fonksiyonlar, herhangi bir türdeki veriyi JSON formatında dosyaya kaydetmek ve dosyadan yüklemek için kullanılabilir.
Generik yapısı sayesinde, farklı türlerdeki verilerle çalışmak mümkündür.
*/
