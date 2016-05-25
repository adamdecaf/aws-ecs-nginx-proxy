package nginx

var (
	DefaultNginxConfigPath = "/etc/nginx/"
	DefaultNginxExecutable = "/etc/init.d/nginx"
)

type Nginx struct {

}

func (n *Nginx) Reload() *error {
	return nil
}

func New() (*Nginx, error) {
	return nil, nil
}

func findNginxExecutable() (string, error) {
	return "", nil
}

// spawn nginx on boot
// use `template.go` to render a new config
