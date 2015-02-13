package gold

import (
	"encoding/json"
	"io/ioutil"
	"log"
	"os"
	"strings"
)

// ServerConfig holds a list of configuration parameters for the server
type ServerConfig struct {
	// DataSkin sets the default skin for viewing RDF resources
	DataSkin string

	// Debug (display or hide stdout logging)
	Debug bool

	// DirIndex contains the default index file name
	DirIndex []string

	// DirSkin points to the skin/app for browsing the data space
	DirSkin string

	//SignUpURL points to the skin/app used for creating new accounts
	SignUpURL string

	//DiskLimit is the maximum total disk (in bytes) to be allocated to a given user
	DiskLimit int

	Root   string
	Vhosts bool
}

// NewServerConfig creates a new config object
func NewServerConfig() *ServerConfig {
	return &ServerConfig{
		DataSkin:  "tabulator",
		DirIndex:  []string{"index.html", "index.htm"},
		DirSkin:   "http://linkeddata.github.io/warp/#list/",
		SignUpURL: "http://linkeddata.github.io/signup/",
		DiskLimit: 100000000, // 100MB
		Root:      serverDefaultRoot(),
	}
}

// LoadJSONFile loads server configuration
func (c *ServerConfig) LoadJSONFile(filename string) error {
	b, err := ioutil.ReadFile(filename)
	if err != nil {
		return err
	}
	return json.Unmarshal(b, &c)
}

func serverDefaultRoot() string {
	serverRoot, err := os.Getwd()
	if err != nil {
		log.Fatalln(err)
	}

	if !strings.HasSuffix(serverRoot, "/") {
		serverRoot += "/"
	}
	return serverRoot
}
