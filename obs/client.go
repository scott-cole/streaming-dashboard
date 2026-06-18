package obs

import (
	"fmt"
	"log"

	"github.com/andreykaipov/goobs"
	"github.com/andreykaipov/goobs/api/requests/scenes"
)

type Client struct {
	conn *goobs.Client
}

func New(password string) (*Client, error) {

	client, err := goobs.New("localhost:4455", goobs.WithPassword(password))
	if err != nil {
		return nil, err
	}
	return &Client{conn: client}, nil
}

func (c *Client) Disconnect() {
	c.conn.Disconnect()
}

func (c *Client) Version() string {
	info, err := c.conn.General.GetVersion()
	if err != nil {
		log.Fatal(err)
	}

	return info.ObsVersion
}

func (c *Client) ListScenes() []string {
	list, err := c.conn.Scenes.GetSceneList()
	if err != nil {
		log.Fatal(err)
	}

	var names []string
	for _, scene := range list.Scenes {
		names = append(names, scene.SceneName)
	}
	return names
}

func (c *Client) SwitchScene(name string) {
	_, err := c.conn.Scenes.SetCurrentProgramScene(scenes.NewSetCurrentProgramSceneParams().WithSceneName(name))
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println("Switched to:", name)
}
