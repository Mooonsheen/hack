package responser

import (
	"encoding/json"
	"io"
	"log"
	"net/http"
	"test/models"

	"github.com/valyala/fasthttp"
)

type Request struct {
	Url      string
	Response []byte
}

func GetCurrentPlanet(r *Request) *models.Universe {
	req, _ := http.NewRequest(http.MethodGet, r.Url, nil)
	req.Header = map[string][]string{
		"X-Auth-Token": {"660dcc4bc01f0660dcc4bc01f4"},
	}
	resp, err := http.DefaultClient.Do(req)
	if err != nil {
		log.Printf("Can't get current planet: %s\n", err)
		return nil
	}
	defer resp.Body.Close()

	if resp.StatusCode != fasthttp.StatusOK {
		log.Printf("status not 200, it's %d", resp.StatusCode)
	}
	body, _ := io.ReadAll(resp.Body)

	var uni models.Universe

	err = json.Unmarshal(body, &uni)
	if err != nil {
		log.Printf("can't unmarshal: %e", err)
	}

	// fmt.Printf("Uni: %v", uni)

	return &uni
}
