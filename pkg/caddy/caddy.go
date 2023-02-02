package caddy

import (
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
	t := template.Must(template.ParseFiles("CaddyFile.tml"))

	caddy_file_path := "/etc/caddy/CaddyFile"
	caddy_file_open, err := os.Open(caddy_file_path)
	if err != nil {
		return err
	}
	err = t.Execute(caddy_file_open, td)
	if err != nil {
		return err
	}

	err = exec.Command("systemctl", "restart", "caddy").Err
	if err != nil {
		return err
	}
	return nil
}
