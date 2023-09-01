package assets

import (
	"encoding/json"
	"log"

	"github.com/moltenwolfcub/Forest-Game/errors"
	"golang.org/x/exp/maps"
)

func LoadTextureMapping(file string) Textures {
	bytes, err := states.ReadFile("states/" + file + ".json")
	if err != nil {
		log.Fatalln(err)
	}

	var state Textures
	if err = json.Unmarshal(bytes, &state); err != nil {
		log.Fatalln(err)
	}
	return state
}

var (
	BerryStates Textures = LoadTextureMapping("berries")
)

type Textures struct {
	Mapping map[string]string `json:"states"`
}

func (t Textures) GetTexturePath(textureKey string) (string, error) {
	path, ok := t.Mapping[textureKey]
	if ok {
		return path, nil
	}
	return "", errors.NewUnknownTextureKeyError(textureKey, maps.Keys(t.Mapping))
}
