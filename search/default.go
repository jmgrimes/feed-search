package search

type defaultMatcher struct {}

func init() {
	var matcher defaultMatcher
	Register("default", matcher)
}

func (matcher defaultMatcher) Search(feed *Feed, searchTerm string) ([]*Result, error) {
	return nil, nil
}