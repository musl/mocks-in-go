package main

import "fmt"

type Client struct{}

type FlagGetter interface {
	GetFlag(string) string
}

type FFClient struct {
	Client     *Client
	ServerAddr string
}

var (
	flagValues = map[string]string{
		"wat":     "true",
		"chicken": "false",
	}
)

func (ffc *FFClient) GetFlag(name string) string {
	return flagValues[name]
}

type Reconciler struct {
	FFClient    FlagGetter
	ClusterName string
}

func (r *Reconciler) Run() string {
	fmt.Println("Run")
	r.FFClient.GetFlag("wat")
	return "run"
}

func main() {
	r := Reconciler{
		FFClient: &FFClient{
			Client:     &Client{},
			ServerAddr: "ff.example.org",
		},
		ClusterName: "woot-postgres",
	}

	fmt.Printf("%#v", r)
}
