package obs

import (
	"fmt"

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

func (c *Client) Version() (string, error) {
	info, err := c.conn.General.GetVersion()
	if err != nil {
		return "", err
	}

	return info.ObsVersion, nil
}

func (c *Client) CurrentScene() (string, error) {
	scene, err := c.conn.Scenes.GetCurrentProgramScene()
	if err != nil {
		return "", err
	}

	return scene.SceneName, nil
}

func (c *Client) ListScenes() ([]string, error) {
	list, err := c.conn.Scenes.GetSceneList()
	if err != nil {
		return nil, err
	}

	var names []string
	for _, scene := range list.Scenes {
		names = append(names, scene.SceneName)
	}
	return names, nil
}

func (c *Client) SwitchScene(name string) error {
	_, err := c.conn.Scenes.SetCurrentProgramScene(scenes.NewSetCurrentProgramSceneParams().WithSceneName(name))
	if err != nil {
		return err
	}
	fmt.Println("Switched to:", name)
	return nil
}
