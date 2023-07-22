package assets

import (
	_ "embed"
)

var (
	//go:embed player.png
	PlayerPng []byte
	//go:embed tree.png
	TreePng []byte

	//go:embed icon/icon16.png
	Icon16 []byte
	//go:embed icon/icon22.png
	Icon22 []byte
	//go:embed icon/icon24.png
	Icon24 []byte
	//go:embed icon/icon32.png
	Icon32 []byte
	//go:embed icon/icon48.png
	Icon48 []byte
	//go:embed icon/icon64.png
	Icon64 []byte
	//go:embed icon/icon128.png
	Icon128 []byte
	//go:embed icon/icon256.png
	Icon256 []byte
	//go:embed icon/icon512.png
	Icon512 []byte
)
