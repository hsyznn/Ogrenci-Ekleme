package cmd

import (
	"fmt"
	"strconv"

	"github.com/spf13/cobra"
)

var toggleCmd = &cobra.Command{
	Use:   "toggle",
	Short: "Öğrenci kayıt durumunu düzenle",
	Long:  `Belirtilen öğrencinin kayıt durumunu değiştirir.`,
	Run: func(cmd *cobra.Command, args []string) {
		if len(args) < 1 {
			fmt.Println("Lütfen düzenlenecek öğrencinin indeksini girin.")
			return
		}
		index, err := strconv.Atoi(args[0])
		if err != nil {
			fmt.Println("Geçersiz indeks.")
			return
		}

		students := loadStudents()
		if err := students.Toggle(index); err != nil {
			fmt.Println("Hata:", err)
		} else {
			saveStudents(students)
			fmt.Printf("Öğrenci kayıt durumu değiştirildi: %d\n", index)
		}
	},
}

func init() {
	rootCmd.AddCommand(toggleCmd)
}
