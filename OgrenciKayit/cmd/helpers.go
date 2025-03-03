package cmd

import (
	"StudentEnrollment/enrollment"
	"StudentEnrollment/storage"
	"fmt"
)

const filename = "students.json"

func loadStudents() enrollment.Students {
	students := enrollment.Students{}
	dataStorage := storage.NewStorage[enrollment.Students](filename)
	if err := dataStorage.Load(&students); err != nil {
		fmt.Println("Veri yüklenirken hata oluştu:", err)
	}
	return students
}

func saveStudents(students enrollment.Students) {
	dataStorage := storage.NewStorage[enrollment.Students](filename)
	if err := dataStorage.Save(students); err != nil {
		fmt.Println("Veri kaydedilirken hata oluştu:", err)
	}
}
