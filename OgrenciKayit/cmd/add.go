package cmd

import (
	"fmt"

	"github.com/spf13/cobra"
)

var addCmd = &cobra.Command{
	Use:   "add",
	Short: "Öğrenci ekle",
	Long:  `Yeni bir öğrenci eklemek için add "Öğrenci Adı ve Soyadı".`,
	Run: func(cmd *cobra.Command, args []string) {
		if len(args) < 2 {
			fmt.Println("Lütfen öğrenci adını ve soyadını girin.(Tırnak İşareti Kullanmayın...)")
			return
		}
		name := args[0]
		surname := args[1]

		students := loadStudents()
		students.Add(name, surname)
		saveStudents(students)
		fmt.Printf("Öğrenci Başarıyla eklendi: %s %s\n", name, surname)
	},
}

func init() {
	rootCmd.AddCommand(addCmd)
}
