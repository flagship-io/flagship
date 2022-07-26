/*
Copyright © 2022 NAME HERE <EMAIL ADDRESS>

*/
package panic

import (
	"log"

	httprequest "github.com/flagship-io/flagship/utils/httpRequest"
	"github.com/spf13/cobra"
)

var panicStatus string

// panicCmd represents the panic command
var PanicCmd = &cobra.Command{
	Use:   "panic",
	Short: "panic short desc",
	Long:  `panic long desc`,
	Run: func(cmd *cobra.Command, args []string) {
		if !(panicStatus == "on" || panicStatus == "off") {
			log.Println("Status can only have 2 values: on or off ")
			return
		}
		err := httprequest.HTTPUpdatePanic(panicStatus)
		if err != nil {
			log.Fatalf("error occured: %v", err)
		}
		log.Printf("Panic set to %v", panicStatus)

	},
}

func init() {

	PanicCmd.Flags().StringVarP(&panicStatus, "status", "s", "", "panic mode")

	if err := PanicCmd.MarkFlagRequired("status"); err != nil {
		log.Fatalf("error occured: %v", err)
	}
}
