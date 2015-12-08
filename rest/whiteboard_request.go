package rest
import (
	"os"
	. "github.com/xtreme-andleung/whiteboardbot/entry"
)
type WhiteboardRequest struct {
	Utf8   string `json:"utf8"`
	Method string `json:"_method,omitempty"`
	Token  string `json:"authenticity_token"`
	Item   Item `json:"item"`
	Commit string `json:"commit,omitempty"`
	Id     string `json:"id,omitempty"`
}

type Item struct {
	StandupId   int `json:"standup_id"`
	Title       string `json:"title"`
	Date        string `json:"date"`
	PostId      string `json:"post_id"`
	Public      string `json:"public"`
	Kind        string `json:"kind"`
	Description string `json:"description,omitempty"`
	Author      string `json:"author,omitempty"`
}

func NewCreateFaceRequest(face Face) (request WhiteboardRequest) {
	item := Item{1, face.Title, face.Time.Format("2006-01-02"), "", "false", "New face", "", face.Author}
	request = WhiteboardRequest{"", "", os.Getenv("WB_AUTH_TOKEN"), item, "Create New Face", ""}
	return
}

func NewUpdateFaceRequest(face Face) (request WhiteboardRequest) {
	item := Item{1, face.Title, face.Time.Format("2006-01-02"), "", "false", "New face", "", face.Author}
	request = WhiteboardRequest{"", "patch", os.Getenv("WB_AUTH_TOKEN"), item, "Update New Face", face.Id}
	return
}