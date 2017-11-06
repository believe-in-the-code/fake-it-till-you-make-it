package pkg

type Liar struct {
	Kind string
	Type string
	Fill func(args Tag) interface{}
}

type Tag struct {
	Args []string
}

func (tag *Tag) HasIndex(index int) bool {
	return len(tag.Args) >= index
}

func (tag *Tag) GetIndex(index int) string {
	return tag.Args[index]
}
