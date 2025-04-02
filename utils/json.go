package utils
import(
	"encoding/json"
)
func JsonStatus(status string) []byte {
	response, _ := json.Marshal(map[string]string{"status": status})
	return response
}