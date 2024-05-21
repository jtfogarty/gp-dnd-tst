package main

import (
	"context"
	"fmt"
 	"github.com/jtfogar/gp-dnd-tst/iputils"

)

// App struct
type App struct {
	ctx context.Context
}


// NewApp creates a new App application struct
func NewApp() *App {
	return &App{}
}

// startup is called when the app starts. The context is saved
// so we can call the runtime methods
func (a *App) startup(ctx context.Context) {
	
	sourceAddressPrefixes := []string{
		"100.23.51.4", "100.81.1.0/16", "120.81.1.0/16", "140.81.1.0/16",
	}

	ipCidr := "100.23.51.4,10.100.4.1;100.90.0.0/12,130.81.0.0/16"

	result := iputils.ProcessIpCidr(sourceAddressPrefixes, ipCidr)
	if len(result.Errors) > 0 {
		fmt.Printf("Errors: %v\n", result.Errors)
	} else {
		fmt.Printf("Result: %v\n", result.Result)
	}

	if len(result.Warnings) > 0 {
		fmt.Printf("Warnings: %v\n", result.Warnings)
	}

}

// Greet returns a greeting for the given name
func (a *App) Greet(name string) string {
	return fmt.Sprintf("Hello %s, It's show time!", name)
}

func (a *App) Tst_ipCidr(ipAddr string[], cidr string) {
	
}