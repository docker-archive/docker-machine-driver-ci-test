package main

import (
	"fmt"
	neturl "net/url"

	"errors"

	"github.com/docker/machine/libmachine/drivers"
	"github.com/docker/machine/libmachine/mcnflag"
	"github.com/docker/machine/libmachine/state"
)

var (
	errNotSupported = errors.New("Not supported")
)

type Driver struct {
	*drivers.BaseDriver
	URL string
}

func NewDriver(hostName, storePath string) *Driver {
	return &Driver{
		BaseDriver: &drivers.BaseDriver{
			MachineName: hostName,
			StorePath:   storePath,
		},
	}
}

func (d *Driver) GetCreateFlags() []mcnflag.Flag {
	return []mcnflag.Flag{
		mcnflag.StringFlag{
			Name:  "url",
			Usage: "URL of host",
		},
	}
}

func (d *Driver) PreCreateCheck() error {
	return nil
}

func (d *Driver) Create() error {
	return nil
}

func (d *Driver) DriverName() string {
	return "none-external"
}

func (d *Driver) GetIP() (string, error) {
	return d.IPAddress, nil
}

func (d *Driver) GetSSHHostname() (string, error) {
	return "", nil
}

func (d *Driver) GetSSHKeyPath() string {
	return ""
}

func (d *Driver) GetSSHPort() (int, error) {
	return 0, nil
}

func (d *Driver) GetSSHUsername() string {
	return ""
}

func (d *Driver) GetURL() (string, error) {
	return d.URL, nil
}

func (d *Driver) GetState() (state.State, error) {
	return state.Running, nil
}

func (d *Driver) Remove() error {
	return nil
}

func (d *Driver) Start() error {
	return errNotSupported
}

func (d *Driver) Stop() error {
	return errNotSupported
}

func (d *Driver) Kill() error {
	return errNotSupported
}

func (d *Driver) Restart() error {
	return errNotSupported
}

func (d *Driver) SetConfigFromFlags(flags drivers.DriverOptions) error {
	url := flags.String("url")
	if url == "" {
		return fmt.Errorf("--url option is required")
	}

	d.URL = url
	u, err := neturl.Parse(url)
	if err != nil {
		return err
	}

	d.IPAddress = u.Host
	return nil
}
