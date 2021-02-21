package gacha

import (
	"math/rand"
	"time"
)

func init() {
	// 乱数の種を設定する
	// 現在時刻をUNIX時間にしたものを種とする
	rand.Seed(time.Now().Unix())
}

func DrawN(p *Player, n int) ([]*Card, map[Rarity]int) {
	p.draw(n)

	results := make([]*Card, n)
	summary := make(map[Rarity]int)
	for i := 0; i < n; i++ {
		results[i] = draw()
		summary[results[i].Rarity]++
	}

	// 変数resultsとsummaryの値を戻り値として返す
	return results, summary
}

func draw() *Card {
	num := rand.Intn(100)
	var cardName string
	switch {
	case num < 80:
		n := rand.Intn(1)
		switch n {
		case 0:
			cardName = "スライム"
		case 1:
			cardName = "ゴブリン"

		return &Card{Rarity: RarityN, Name: cardName}
		}
	case num < 95:
		n := rand.Intn(1)
		switch n {
		case 0:
			cardName = "オーク"
		case 1:
			cardName = "ホブゴブリン"

		return &Card{Rarity: RarityN, Name: cardName}
	case num < 99:
		n := rand.Intn(1)
		switch n {
		case 0:
			cardName = "ドラゴン"
		case 1:
			cardName = "グリフォン"

		return &Card{Rarity: RaritySR, Name: cardName}
	default:
		n := rand.Intn(1)
		switch n {
		case 0:
			cardName = "イフリート"
		case 1:
			cardName = "バハムート"

		return &Card{Rarity: RarityXR, Name: cardName}
	}
}
