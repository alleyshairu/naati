package naati

import (
	"encoding/json"
	"strings"

	"github.com/abadojack/whatlanggo"
)

type Language int

const (
	English Language = iota
	Urdu
)

type Sentence struct {
	Language Language
	Body     string
}

func (s *Sentence) Words() int {
	words := strings.Fields(s.Body)
	return len(words)
}

type Segment struct {
	Prompt Sentence `json:"prompt"`
	Answer Sentence `json:"answer"`
}

type Dialogue struct {
	Title    string    `json:"title"`
	Scenario string    `json:"scenario"`
	Segments []Segment `json:"segments"`
}

// Implement UnmarshalJSON for Segment
func (seg *Segment) UnmarshalJSON(data []byte) error {
	// Define a temporary struct for Segment
	var temp struct {
		Prompt string `json:"segment"`
		Answer string `json:"answer"`
	}

	// Unmarshal into the temporary struct
	if err := json.Unmarshal(data, &temp); err != nil {
		return err
	}

	lang := whatlanggo.DetectLang(temp.Prompt)

	var promptLang, answerLang Language
	if lang.Iso6393() == "urd" || lang.Iso6393() == "skr" {
		promptLang, answerLang = Urdu, English
	} else if lang.Iso6393() == "eng" {
		promptLang, answerLang = English, Urdu
	}

	seg.Prompt = Sentence{
		Language: promptLang,
		Body:     temp.Prompt,
	}

	seg.Answer = Sentence{
		Language: answerLang,
		Body:     temp.Answer,
	}

	return nil
}

func (d *Dialogue) UnmarshalJSON(data []byte) error {
	var temp struct {
		Title    string            `json:"title"`
		Scenario string            `json:"scenario"`
		Segments []json.RawMessage `json:"segments"`
	}

	// Unmarshal into the temporary struct
	if err := json.Unmarshal(data, &temp); err != nil {
		return err
	}

	// Populate the Dialogue struct
	d.Title = temp.Title
	d.Scenario = temp.Scenario

	// Unmarshal each segment into the Segment slice
	var segments []Segment
	for _, segmentData := range temp.Segments {
		var segment Segment
		if err := json.Unmarshal(segmentData, &segment); err != nil {
			return err
		}
		segments = append(segments, segment)
	}
	d.Segments = segments

	return nil
}
