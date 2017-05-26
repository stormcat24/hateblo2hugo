package cmd

func init() {
	RootCmd.AddCommand(migrateCmd)
	initMigrateCmd()
}