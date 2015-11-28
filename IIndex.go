package repository

type IIndex interface {
	Data() IndexData
	Values() Data
	Set(key interface{}, value int)
	Has(key interface{}) bool
	Get(key interface{}) (int, bool)
	MatchFirst(pattern string) (int, bool)
}
