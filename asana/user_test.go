package asana

import (
	"testing"
	"encoding/json"
)

func TestUnmarshalUserData(t *testing.T) {
	j := []byte(`
{"data":{"id":184808224339,"name":"Brian McCallister","email":"brianm@skife.org","photo":{"image_21x21":"https://s3.amazonaws.com/profile_photos/184808224339.tPSTUiUUM4eBivXaQtE8_21x21.png","image_27x27":"https://s3.amazonaws.com/profile_photos/184808224339.tPSTUiUUM4eBivXaQtE8_27x27.png","image_36x36":"https://s3.amazonaws.com/profile_photos/184808224339.tPSTUiUUM4eBivXaQtE8_36x36.png","image_60x60":"https://s3.amazonaws.com/profile_photos/184808224339.tPSTUiUUM4eBivXaQtE8_60x60.png","image_128x128":"https://s3.amazonaws.com/profile_photos/184808224339.tPSTUiUUM4eBivXaQtE8_huge.jpeg"},"workspaces":[{"id":13438604102030,"name":"Brian"},{"id":498346170860,"name":"Personal Projects"}]}}
`)
	ud := userData {}
	err := json.Unmarshal(j, &ud)
	if err != nil {
		t.Fatalf("got err %s", err)
	}

	if ud.Data.Id != 184808224339 {
		t.Fatalf("expected id= %d got %d", 184808224339, ud.Data.Id)
	}
}
