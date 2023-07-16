package assets

import (
	_ "embed"
)

var (
	//go:embed player.png
	playerBytes []byte
	//go:embed tree.png
	treeBytes []byte

	//go:embed icon/icon16.png
	icon16Bytes []byte
	//go:embed icon/icon22.png
	icon22Bytes []byte
	//go:embed icon/icon24.png
	icon24Bytes []byte
	//go:embed icon/icon32.png
	icon32Bytes []byte
	//go:embed icon/icon48.png
	icon48Bytes []byte
	//go:embed icon/icon64.png
	icon64Bytes []byte
	//go:embed icon/icon128.png
	icon128Bytes []byte
	//go:embed icon/icon256.png
	icon256Bytes []byte
	//go:embed icon/icon512.png
	icon512Bytes []byte
)
