package cmd

import (
	"fmt"
	"strconv"

	"github.com/spf13/cobra"
)

var deleteCmd = &cobra.Command{
	Use:   "delete",
	Short: "Öğrenci sil",
	Long:  `Belirtilen indeksteki öğrenciyi silmek için kullanılır.`,
	Run: func(cmd *cobra.Command, args []string) {
		if len(args) < 1 {
			fmt.Println("Lütfen silinecek öğrencinin indeksini girin.")
			return
		}
		index, err := strconv.Atoi(args[0])
		if err != nil {
			fmt.Println("Geçersiz indeks.")
			return
		}

		students := loadStudents()
		if err := students.Delete(index); err != nil {
			fmt.Println("Hata:", err)
		} else {
			saveStudents(students)
			fmt.Printf("Öğrenci silindi: %d\n", index)
		}
	},
}

func init() {
	rootCmd.AddCommand(deleteCmd)
}
