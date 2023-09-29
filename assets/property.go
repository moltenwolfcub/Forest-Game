package assets

import (
	"encoding/json"

	"github.com/moltenwolfcub/Forest-Game/errors"
	"golang.org/x/exp/maps"
)

func LoadTextureMapping(file string) (Textures, error) {
	bytes, err := states.ReadFile("states/" + file + ".json")
	if err != nil {
		return Textures{}, err
	}

	var state Textures
	if err = json.Unmarshal(bytes, &state); err != nil {
		return Textures{}, err
	}
	return state, nil
}
func MustLoadTextureMapping(file string) Textures {
	mapping, err := LoadTextureMapping(file)
	if err != nil {
		panic("Failed to load texture mapping: " + err.Error())
	}
	return mapping
}

var (
	BerryStates    Textures = MustLoadTextureMapping("berries")
	MushroomStates Textures = MustLoadTextureMapping("mushrooms")
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
