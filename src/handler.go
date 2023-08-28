package src

import (
	"os"
	"fmt"
	"net"
	"os/exec"
	"runtime"
	"strconv"

	"github.com/charmbracelet/lipgloss"
)

func GettingIPAddresses() []net.IP {
	var ips []net.IP

	addresses, err := net.InterfaceAddrs()
	if err != nil {
		fmt.Println(err)
	}

	for _, addr := range addresses {
		if ipnet, ok := addr.(*net.IPNet); ok || !ipnet.IP.IsLoopback() {
			if ipnet.IP.To4() != nil {
				ips = append(ips, ipnet.IP)
			}
		}
	}
	return ips
}

func GettingIPInterfaceName(IPAddress net.IP) string {
	interfaces, err := net.Interfaces()
	if err != nil {
		fmt.Println(err)
	}

	for _, iface := range interfaces {
		if addrs, err := iface.Addrs(); err == nil {
			for _, addr := range addrs {
				if iip, _, err := net.ParseCIDR(addr.String()); err == nil {
					if iip.Equal(IPAddress) {
						return iface.Name
					}
				} else {
					continue
				}
			}
		} else {
			continue
		}
	}
	return ""
}

func ServingInfo(defaultPort int) lipgloss.Style {
	var servingInfo = lipgloss.NewStyle().
		SetString("\nserving HTTP on 0.0.0.0 port", strconv.Itoa(defaultPort), "...").
		Bold(true).Foreground(lipgloss.Color("#FAFAFA"))

	return servingInfo
}

func TitleBorder() lipgloss.Style {
	var titleBorder = lipgloss.NewStyle().
		SetString("choose the interface").
		BorderStyle(lipgloss.NormalBorder()).
		BorderForeground(lipgloss.Color("171")).
		Bold(true).Foreground(lipgloss.Color("#FAFAFA"))

	return titleBorder
}

func runCmd(name string, arg ...string) {
	cmd := exec.Command(name, arg...)
	cmd.Stdout = os.Stdout
	cmd.Run()
}
func ClearTerminal() {
	switch runtime.GOOS {
	case "darwin":
		runCmd("clear")
	case "linux":
		runCmd("clear")
	case "windows":
		runCmd("cmd", "/c", "cls")
	default:
		runCmd("clear")
	}
}