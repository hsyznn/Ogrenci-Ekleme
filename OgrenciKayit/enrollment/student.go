package enrollment

import (
	"errors"
	"fmt"
	"time"
)

type StudentEnrollment struct {
	Name        string
	Surname     string
	Completed   bool
	CreatedAt   time.Time
	CompletedAt *time.Time
}

type Students []StudentEnrollment

func (students Students) Print() {
	panic("Tablo ekrana bastıralamadı")
}

func (students *Students) Add(name, surname string) {
	stu := StudentEnrollment{
		Name:        name,
		Surname:     surname,
		CreatedAt:   time.Now(),
		Completed:   false,
		CompletedAt: nil,
	}
	*students = append(*students, stu)
}

func (students *Students) ValidateIndex(index int) error {
	if index < 0 || index >= len(*students) {
		err := errors.New("student not found")
		return err
	}
	return nil
}

func (students *Students) Delete(index int) error {
	ns := *students
	if err := ns.ValidateIndex(index); err != nil {
		fmt.Println(err)
	}
	*students = append(ns[:index], ns[index+1:]...)
	return nil
}

/*
Toggle adında bir fonksiyon tanımlanmış.
Bu fonksiyon, Students tipinde bir alıcıya (receiver) sahip.
students parametresi bir pointer (*Students) olarak tanımlanmış, bu da fonksiyonun students üzerinde doğrudan değişiklik yapabileceği anlamına gelir.
Fonksiyon bir int türünde index parametresi alır ve bir error döner.
*/
func (students *Students) Toggle(index int) error {
	ns := *students //students pointer'ının işaret ettiği değeri ns değişkenine atar. Bu, students diliminin bir kopyasını oluşturur.
	if err := ns.ValidateIndex(index); err != nil {
		return err
		/*
			ns diliminde ValidateIndex adında bir fonksiyon çağrılır ve index parametresi ile doğrulama yapılır.
			Eğer ValidateIndex fonksiyonu bir hata (error) dönerse, bu hata err değişkenine atanır ve fonksiyon return err ifadesi ile sonlandırılır.
			Bu, hatanın fonksiyonun çağrıldığı yere iletilmesini sağlar
		*/
	}
	isCompleted := ns[index].Completed

	/*
		ns dilimindeki index numaralı öğrencinin Completed alanı isCompleted değişkenine atanır.
		Bu, öğrencinin tamamlanma durumunu kontrol etmek için kullanılır.
	*/

	if !isCompleted {
		completionTime := time.Now()
		ns[index].CompletedAt = &completionTime

		/*
			Eğer isCompleted false ise (öğrenci tamamlanmamışsa), şu anki zaman completionTime değişkenine atanır.
			ns dilimindeki index numaralı öğrencinin CompletedAt alanı completionTime'ın adresi ile güncellenir.
			Bu, öğrencinin tamamlanma zamanını ayarlar.
		*/
	}

	ns[index].Completed = !isCompleted
	/*
		ns dilimindeki index numaralı öğrencinin Completed alanı tersine çevrilir. Eğer isCompleted true ise, false yapılır; false ise, true yapılır.
		Bu, öğrencinin tamamlanma durumunu değiştirir.
	*/

	return nil
}

func (students *Students) Edit(index int, name, surname string) error {

	ns := *students
	if err := ns.ValidateIndex(index); err != nil {
		return err
	}

	ns[index].Name = name
	ns[index].Surname = surname
	return nil

}
