package internal

type Chapter struct {
	Label    string  `json:"label"`
	Children []Level `json:"children"`
	Data     any     `json:"-"`
}

type Level struct {
	Label string `json:"label"`
	Data  any    `json:"-"`
	Value any    `json:"value"`
}

type LevelMeta struct {
	Chapter int `json:"chapter"`
	Level   int `json:"level"`
}

func (b *Base) LevelIndex() int {
	return b.levelIndex
}

func (b *Base) ChapterIndex() int {
	return b.chapterIndex
}
