package main

import (
	"flag"
	"fmt"
	"log"
	"os"
	"path/filepath"
	"strings"

	"github.com/BurntSushi/toml"
	"github.com/Tnze/go-mc/net"
	"github.com/adrg/xdg"
	"github.com/charmbracelet/bubbles/textinput"
	tea "github.com/charmbracelet/bubbletea"
)

type Config struct {
	Servers map[string]Server `toml:servers`
}

type Server struct {
	Addr string `toml:"addr"`
	Pass string `toml:"pass"`
}

type model struct {
	server Server
	log    []string
	input  textinput.Model
}

func newModel(server Server) model {
	ti := textinput.New()
	ti.Focus()
	ti.CharLimit = 156
	ti.Width = 20
	return model{
		server: server,
		input:  ti,
	}
}

func (m model) Init() tea.Cmd {
	return textinput.Blink
}

func (m model) Update(msg tea.Msg) (tea.Model, tea.Cmd) {
	var cmd tea.Cmd
	switch msg := msg.(type) {
	case tea.KeyMsg:
		switch msg.Type {
		case tea.KeyCtrlC, tea.KeyEsc:
			return m, tea.Quit
		case tea.KeyEnter:
			cmd := m.input.Value()
			m.log = append(m.log, cmd)
			m.log = append(m.log, send(m.server, cmd))
			m.input.Reset()
		}
	}

	m.input, cmd = m.input.Update(msg)
	return m, cmd
}

func (m model) View() string {
	var b strings.Builder
	for _, s := range m.log {
		b.WriteString(s)
		b.WriteString("\n")
	}
	b.WriteString(m.input.View())
	return b.String()
}

func main() {
	// Find and read config.
	configPath, err := findConfig("raccoon/config.toml")
	if err != nil {
		log.Fatalf("failed finding config: %v\n", err)
	}

	var conf Config
	_, err = toml.DecodeFile(configPath, &conf)
	if err != nil {
		log.Fatalf("failed reading config: %v\n", err)
	}

	flag.Parse()
	var server Server
	l := len(flag.Args())
	switch {
	case l == 1:
		server = conf.Servers[flag.Args()[0]]
	case l > 1:
		// CLI mode.
		args := flag.Args()
		cmd := strings.Join(args[1:], " ")
		fmt.Println(send(conf.Servers[flag.Args()[0]], cmd))
		os.Exit(0)
	default:
		log.Fatalln("missing server name argument!")
	}

	// TUI mode.
	p := tea.NewProgram(
		newModel(server),
		tea.WithAltScreen(),
	)
	if err := p.Start(); err != nil {
		log.Fatalf("raccoon has crashed: %v\n", err)
	}
}

func findConfig(path string) (string, error) {
	rel := "raccoon/config.toml"
	if exists(filepath.Join(xdg.ConfigHome, rel)) {
		return filepath.Join(xdg.ConfigHome, rel), nil
	}
	if exists(filepath.Join("/etc/", rel)) {
		return filepath.Join("/etc/", rel), nil
	}
	return "", fmt.Errorf(
		"could not locate `config.toml` in any of the following paths: %v, %v\n",
		filepath.Join(xdg.ConfigHome, rel),
		filepath.Join("/etc/", rel),
	)
}

// exists returns true if the specified path exists.
func exists(path string) bool {
	_, err := os.Stat(path)
	return err == nil || os.IsExist(err)
}

func send(server Server, cmd string) string {
	conn, err := net.DialRCON(server.Addr, server.Pass)
	if err != nil {
		return err.Error()
	}
	defer conn.Close()

	err = conn.Cmd(cmd)
	if err != nil {
		return err.Error()
	}

	var b strings.Builder
	for {
		resp, err := conn.Resp()
		if err != nil {
			fmt.Print(err)
		}
		b.WriteString(resp)
		break
	}
	return b.String()
}
