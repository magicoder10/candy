module candy

go 1.14

require (
	github.com/hajimehoshi/ebiten/v2 v2.0.5
	// Must use this newer version of oto to supress warnings
	github.com/hajimehoshi/oto v0.7.1 // indirect
	github.com/stretchr/testify v1.6.1
	github.com/teamyapp/ui v0.0.0-00010101000000-000000000000
	golang.org/x/image v0.0.0-20210220032944-ac19c3e999fb
	golang.org/x/sys v0.0.0-20210119212857-b64e53b001e4 // indirect
)

replace github.com/teamyapp/ui => ../ui
