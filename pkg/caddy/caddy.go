package caddy

import (
	"fmt"
	"io/fs"
	"os"
	"os/exec"
	"text/template"
)

type CaddyTmpl struct {
	Domain string
	Port   int
}

func UpdateCaddy(domain string) error {
	td := CaddyTmpl{Domain: domain}
	t := template.Must(template.ParseFiles("../../Caddyfile.tml"))

	caddy_file_path := "/etc/caddy/Caddyfile"
	caddy_file_open, err := os.OpenFile(caddy_file_path, os.O_APPEND|os.O_CREATE|os.O_WRONLY, fs.ModeDevice)
	if err != nil {
		return fmt.Errorf("failed to open caddy file: %w", err)
	}
	err = t.Execute(caddy_file_open, td)
	if err != nil {
		return fmt.Errorf("failed to execute template: %w", err)
	}

	err = exec.Command("systemctl", "restart", "caddy").Err
	if err != nil {
		return err
	}
	return nil
}
