package asana

import (
	"encoding/json"
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestUnmarshalUserData(t *testing.T) {
	j := []byte(`
{
  "data": {
    "id": 184808224339,
    "name": "Brian McCallister",
    "email": "brianm@skife.org",
    "photo": {
      "image_21x21": "https://21",
      "image_27x27": "https://27",
      "image_36x36": "https://36",
      "image_60x60": "https://60",
      "image_128x128": "https://128"
    },
    "workspaces": [
      {
        "id": 13438604102030,
        "name": "Brian"
      },
      {
        "id": 498346170860,
        "name": "Personal Projects"
      }
    ]
  }
}
`)
	ud := userData{}
	err := json.Unmarshal(j, &ud)
	if err != nil {
		t.Fatalf("got err %s", err)
	}
	u := ud.Data
	t.Log("%+v", u)
	assert.Equal(t, u.Id, 184808224339, "Unexpected Id")
	assert.Equal(t, u.Name, "Brian McCallister", "")
	assert.Equal(t, u.Email, "brianm@skife.org", "")
	assert.Equal(t, u.Photos.Image_21x21, "https://21", "")
	assert.Equal(t, u.Photos.Image_128x128, "https://128", "")
	assert.Equal(t, u.Workspaces[0].Name, "Brian", "")
	assert.Equal(t, u.Workspaces[1].Name, "Personal Projects", "")
}
