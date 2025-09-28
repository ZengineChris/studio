package url

import (
	"encoding/json"
	"fmt"

	"github.com/zenginechris/studio/public"
)

type viteManifest map[string]interface{}

// TODO: move this to the interface
func viteEntry(manifest viteManifest) func(key string) map[string]interface{} {
	return func(key string) map[string]interface{} {
		if entry, ok := manifest[key]; ok {
			if entry, ok := entry.(map[string]interface{}); ok {
				return entry
			}
		}
		return nil
	}
}

func parseViteManifest(data []byte) (viteManifest, error) {
	var manifest viteManifest
	if err := json.Unmarshal(data, &manifest); err != nil {
		return nil, err
	}
	return manifest, nil
}

type component struct {
	entry map[string]interface{}
}

func (c component) Styles() []interface{} {
	return c.entry["css"].([]interface{})
}

func (c component) Script() string {
	return fmt.Sprintf("%s", c.entry["file"])
}

func Component(name string) component {
	file, err := public.Manifest()
	if err != nil {
		panic(err)
	}
	manifest, err := parseViteManifest(file)
	if err != nil {
		panic(err)
	}

	return component{
		entry: viteEntry(manifest)(name),
	}
}

func Asset(name string) string {
	return fmt.Sprintf("/static/%s", name)
}
