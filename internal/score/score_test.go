package score

import (
	"testing"
)

func TestGenerateTable(t *testing.T) {
	got := GenerateTable("./../../score.csv")
	want := []Row{
		{Name: "Johny", Base: "oF8", Input: "Fo"},
		{Name: "Ka0s", Base: "0123456789", Input: "23"},
		{Name: "SmoKe", Base: "01", Input: "01100"},
		{Name: "S0mbra", Base: "oi8", Input: "oo"},
		{Name: "KamiKaze", Base: "54?t", Input: "?4?"},
		{Name: "TheBest", Base: "kju2aq", Input: "u2ka"},
		{Name: "L0s3r", Base: "_- /.!#", Input: "# _"},
	}

	for index, row := range want {
		if row.Name != got[index].Name || row.Input != got[index].Input || row.Base != got[index].Base {
			t.Errorf("At row %d:  \n %v \n %v ", index, row, got[index])
		}
	}
}
