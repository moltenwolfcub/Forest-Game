package assets

import (
	"encoding/json"
	"fmt"
)

func LoadTextureMapping(file string) Textures {
	bytes, err := states.ReadFile("states/" + file + ".json")
	if err != nil {
		panic(err)
	}

	var state Textures
	if err = json.Unmarshal(bytes, &state); err != nil {
		panic(err)
	}
	return state
}

var (
	BerryStates Textures = LoadTextureMapping("berries")
)

type Textures struct {
	Mapping map[string]string `json:"states"`
}

func (t Textures) GetTexturePath(textureKey string) string {
	path, ok := t.Mapping[textureKey]
	if ok {
		return path
	}
	panic(fmt.Sprintf("TextureKey doesn't know how to handle texture key: %s\n\nKnown states are: %v", textureKey, t.Mapping))
}
