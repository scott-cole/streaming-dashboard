package obs

import (
	"fmt"

	"github.com/andreykaipov/goobs"
	"github.com/andreykaipov/goobs/api/requests/scenes"
	"github.com/andreykaipov/goobs/api/requests/sources"
)

type Client struct {
	conn *goobs.Client
}

type Status struct {
	Streaming bool `json:"streaming"`
	Recording bool `json:"recording"`
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

func (c *Client) Status() (*Status, error) {
	streamRes, err := c.conn.Stream.GetStreamStatus()
	if err != nil {
		return nil, err
	}
	recordRes, err := c.conn.Record.GetRecordStatus()
	if err != nil {
		return nil, err
	}
	return &Status{
		Streaming: streamRes.OutputActive,
		Recording: recordRes.OutputActive,
	}, nil
}

func (c *Client) StartStream() error {
	_, err := c.conn.Stream.StartStream()
	return err
}

func (c *Client) StopStream() error {
	_, err := c.conn.Stream.StopStream()
	return err
}

func (c *Client) StartRecord() error {
	_, err := c.conn.Record.StartRecord()
	return err
}

func (c *Client) StopRecord() error {
	_, err := c.conn.Record.StopRecord()
	return err
}

func (c *Client) ScenePreview(sceneName string, width, height int) (string, error) {
	res, err := c.conn.Sources.GetSourceScreenshot(
		sources.NewGetSourceScreenshotParams().
			WithSourceName(sceneName).
			WithImageFormat("png").
			WithImageWidth(float64(width)).
			WithImageHeight(float64(height)).
			WithImageCompressionQuality(-1),
	)
	if err != nil {
		return "", err
	}
	return res.ImageData, nil
}
