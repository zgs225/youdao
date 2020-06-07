package youdao

type basicField struct {
	Phonetic   string   `json:"phonetic"`
	UkPhonetic string   `json:"uk-phonetic"`
	UsPhonetic string   `json:"us-phonetic"`
	UkSpeech   string   `json:"uk-speech"`
	UsSpeech   string   `json:"us-speech"`
	Explains   []string `json:"explains"`
}

type webField struct {
	Value []string `json:"value"`
	Key   string   `json:"key"`
}

type dict map[string]string

// Result 有道字典查询结果
type Result struct {
	ErrorCode   string       `json:"errorCode"`
	Query       string       `json:"query"`
	Translation *[]string    `json:"translation"`
	Basic       *basicField  `json:"basic"`
	Web         *[]*webField `json:"web"`
	Dict        *dict        `json:"dict"`
	WebDict     *dict        `json:"webdict"`
	L           string       `json:"l"`
	TSpeakURL   string       `json:"tSpeakUrl"`
	SpeakURL    string       `json:"speakUrl"`
}
