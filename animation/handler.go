package animation

import "github.com/hajimehoshi/ebiten"

type animInfo struct {
	sprites []*ebiten.Image
	speed   int
}

type Handler struct {
	animMap      map[string]animInfo
	currentAnim  animInfo
	lastIdx      int
	remainFrames int
}

// New make a new animation handler
func New() *Handler {
	h := &Handler{}
	h.animMap = make(map[string]animInfo)

	return h
}

func (h *Handler) Add(name string, sprites []*ebiten.Image, speed int) {
	h.animMap[name] = animInfo{sprites, speed}

}

func (h *Handler) Play(name string) {
	h.currentAnim = h.animMap[name]
	h.lastIdx = 0
	h.remainFrames = h.currentAnim.speed
}

func (h *Handler) Update(screen *ebiten.Image, x, y float64) {
	option := ebiten.DrawImageOptions{}
	option.GeoM.Translate(x, y)

	screen.DrawImage(h.currentAnim.sprites[h.lastIdx], &option)

	h.remainFrames--
	if h.remainFrames == 0 {
		h.lastIdx++
		if len(h.currentAnim.sprites) == h.lastIdx {
			h.lastIdx = 0
		}
		h.remainFrames = h.currentAnim.speed
	}
}
