package repository

type IIndex interface {
	Set(key interface{}, value int)
	Has(key interface{}) bool
	Get(key interface{}) (int, bool)
	MatchFirst(pattern string) (int, bool)
}
