package cmd

import (
	"StudentEnrollment/enrollment"
	"fmt"
	"time"

	"github.com/spf13/cobra"
)

var listCmd = &cobra.Command{
	Use:   "list",
	Short: "Öğrencileri listele",
	Long:  `Tüm öğrencileri listelemek için kullanılır.`,
	Run: func(cmd *cobra.Command, args []string) {
		students := loadStudents()
		printStudents(students)
	},
}

func init() {
	rootCmd.AddCommand(listCmd)
}

func printStudents(students enrollment.Students) {
	fmt.Println("Index | Name       | Surname    | Completed | 	Created At           | Completed At")
	fmt.Println("------|------------|------------|-----------|----------------------|----------------------")
	for index, student := range students {
		completed := "❌"
		if student.Completed {
			completed = "✅"
		}
		completedAt := ""
		if student.CompletedAt != nil {
			completedAt = student.CompletedAt.Format(time.RFC1123)
		}
		fmt.Printf("%-6d| %-11s| %-11s| %-9s| %-20s| %-20s\n",
			index, student.Name, student.Surname, completed, student.CreatedAt.Format(time.RFC1123), completedAt)
	}
}
