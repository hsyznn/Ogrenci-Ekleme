package cmd

import (
	"fmt"
	"os"

	"github.com/spf13/cobra"
)

var rootCmd = &cobra.Command{
	Use:   "StudentEnrollment",
	Short: "StudentEnrollment CLI uygulaması",
	Long:  `Bu uygulama, öğrenci kayıtlarını yönetmek için kullanılan bir CLI uygulamasıdır.`,
	CompletionOptions: cobra.CompletionOptions{
		DisableDefaultCmd: true, // completion komutunu devre dışı bırakır
	},
}

func Execute() {
	if err := rootCmd.Execute(); err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
}

func init() {
	cobra.OnInitialize()
}
