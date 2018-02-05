package cmd

import (
	"github.com/SinedVonotirah/gopo/core/ctx"
	"github.com/SinedVonotirah/gopo/shared/bench"
	"github.com/SinedVonotirah/gopo/shared/logging"

	"fmt"
	"github.com/spf13/cobra"
)

func init() {
	if ctx.GlobalContext.AppConfig.DB.Migration.Enable {
		/*migrator.ApplyMigrations(ctx.GlobalContext.AppConfig.DB.Migrations.Folder,
		ctx.GlobalContext.AppConfig.DB.Migrations.ConnectionString, false)*/
	}
}

var RootCmd = &cobra.Command{
	Use:   "run",
	Short: "Same short gopo description",
	Long: `Longer gopo description..
            feel free to use a few lines here.
            `,
	Run: func(cmd *cobra.Command, args []string) {
		logging.WithFields(logging.Fields{}).Error("Starting gopo....")
		bench.RunBenchmark("beego")
		bench.RunBenchmark("gorm")
		fmt.Print(bench.MakeReport())
	},
}
