package main

import (
	"os"
	"fmt"
	"strings"

	"github.com/Doct3rJohn/servEX/src"
	"github.com/Doct3rJohn/servEX/src/helper"
	"github.com/Doct3rJohn/servEX/src/server"

	tea "github.com/charmbracelet/bubbletea"
)

func main() {
	defaultPort := 8000

	ipList := src.GettingIPAddresses()
	for _, ip := range ipList {
		helper.NameList = append(helper.NameList, src.GettingIPInterfaceName(ip))
	}

	if !strings.Contains(server.CheckUsedPort(defaultPort), "Address already in use") {
		fmt.Println(src.TitleBorder())

		p := tea.NewProgram(helper.Model{})
		m, err := p.Run()
		if err != nil {
			fmt.Println(err)
			os.Exit(1)
		}

		if m, ok := m.(helper.Model); ok && m.Choice != "" {
			for i:=0; i<len(ipList); i++ {
				nameListIndex := src.GettingIPInterfaceName(ipList[i])
				if m.Choice == nameListIndex {
					src.ClearTerminal()
					helper.DirListing(ipList[i], defaultPort)
					fmt.Println(src.ServingInfo(defaultPort))
					server.StartTheServer(defaultPort)
				}
			}
		}
	} else {
		fmt.Printf("Address already in use, port %d\n", defaultPort)
	}
}