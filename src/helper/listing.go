package helper

import (
	"os"
	"fmt"
	"net"

	"github.com/charmbracelet/lipgloss"
)

func DirListing(gettingIPAddresses net.IP, defaultPort int) {
	var listFile = lipgloss.NewStyle().
		SetString("listing files to serve...").
		BorderStyle(lipgloss.NormalBorder()).
		BorderForeground(lipgloss.Color("171")).
		Bold(true).Foreground(lipgloss.Color("#FAFAFA"))

	fmt.Println(listFile)
	files, err := os.ReadDir("./")
	if err != nil {
		fmt.Println(err)
	}

	for _, f := range files {
		fmt.Printf("http://%v:%d/%s\n", gettingIPAddresses, defaultPort, f.Name())
	}
}