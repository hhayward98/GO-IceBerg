package ASearch

import (
	"database/sql"
	"testing"
	_ "github.com/go-sql-driver/mysql"


)


func Test_Search_By_ID(t *testing.T) {

	got := SearchID(330,0)
	// 330 Fire Suit Mulatto Skin No Visor Full Black Eyes Head Goggles Smile No Head No Chain Orange BG
	want := []string{330, Fire Suit, Mulatto Skin, No Visor, Full Black, Head Goggles, Smile, No Head, No Chain, Orange BG}
	if got != want {
		t.Errorf("got %q, wanted %q", got , want)
	}
}



func BenchmarkAdd(b *testing.B) {
	for i :=0; i < b.N ; i++ {
		SearchID(330,0)
	}
}
