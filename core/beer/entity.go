package beer

type Beer struct {
	ID    int64     `json:"id"` // tags
	Name  string    `json:"name"`
	Type  BeerType  `json:"type"`
	Style BeerStyle `json:"style"`
}

type BeerType int

const (
	TypeAle    = 1
	TypeLarger = 2
	TypeMalt   = 3
	TypeStout  = 4
)

func (t BeerType) String() string {
	switch t {
	case TypeAle:
		return "Ale"
	case TypeLarger:
		return "Larger"
	case TypeMalt:
		return "Malt"
	case TypeStout:
		return "Stout"
	default:
		return "Unknown"
	}
}

type BeerStyle int

const (
	StyleAmber = iota + 1
	StyleBlonde
	StyleBrown
	StyleCrean
	StyleDark
	StylePale
	StyleStrong
	StyleWheat
	StyleRed
	StyleIPA
	StyleLime
	StylePilsner
	StyleGolden
	StyleFruit
	StyleHoney
)

func (s BeerStyle) String() string {
	switch s {
	case StyleAmber:
		return "Amber"
	case StyleBlonde:
		return "Blonde"
	case StyleBrown:
		return "Brown"
	case StyleCrean:
		return "Crean"
	case StyleDark:
		return "Dark"
	case StylePale:
		return "Pale"
	case StyleStrong:
		return "Strong"
	case StyleWheat:
		return "Wheat"
	case StyleRed:
		return "Red"
	case StyleIPA:
		return "IPA"
	case StyleLime:
		return "Lime"
	case StylePilsner:
		return "Pilsner"
	case StyleGolden:
		return "Golden"
	case StyleFruit:
		return "Fruit"
	case StyleHoney:
		return "Honey"
	}
	return "Unknown"
}
