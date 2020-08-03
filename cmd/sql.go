package cmd

import (
	"github.com/go-programming-tour-book/tour/internal/sql2struct"
	"github.com/spf13/cobra"
	"log"
)

var username string
var password string
var host string
var charset string
var dbType string
var dbName string
var tableName string

var sqlCmd = &cobra.Command{
	Use:   "sql",
	Short: "sql conversion and processing",
	Long:  "sql conversion and processing",
	Run:   func(cmd *cobra.Command, args []string) {},
}

var sql2structCmd = &cobra.Command{
	Use:   "struct",
	Short: "sql conversion",
	Long:  "sql conversion",
	Run: func(cmd *cobra.Command, args []string) {
		dbInfo := &sql2struct.DBInfo{
			DBType: dbType, Host: host, UserName: username, Password: password, Charset: charset}
		dbModel := sql2struct.NewDBModel(dbInfo)
		err := dbModel.Connect()
		if err != nil {
			log.Fatalf("dbMode.Connect err: %v", err)
		}
		columns, err := dbModel.GetColumns(dbName, tableName)
		if err != nil {
			log.Fatalf("dbModel.GetColumns err: %v", err)
		}
		template := sql2struct.NesStructTemplate()
		templateColumns := template.AssemblyColumns(columns)
		err = template.Generate(tableName, templateColumns)
		if err != nil {
			log.Fatalf("template.Genarate err: %v", err)
		}
	},
}

func init() {
	sqlCmd.AddCommand(sql2structCmd)
	sql2structCmd.Flags().StringVarP(&username, "username", "", "",
		"please input the account of the database")
	sql2structCmd.Flags().StringVarP(&password, "password","","",
		"please input the password of the database")
	sql2structCmd.Flags().StringVarP(&host, "host", "", "127.0.0.1:3306",
		"please input the HOST of the database")
	sql2structCmd.Flags().StringVarP(&charset, "charset", "", "utf8mb4",
		"please input the encode of the database")
	sql2structCmd.Flags().StringVarP(&dbType, "type", "", "mysql",
		"please input the instance model of the database")
	sql2structCmd.Flags().StringVarP(&dbName, "db", "", "",
		"please input the name of the database")
	sql2structCmd.Flags().StringVarP(&tableName, "table", "", "",
		"please input the name of the table")
}
