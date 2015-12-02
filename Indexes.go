package repository

type Indexes map[string]IndexInterface

func (this *Indexes) Len() int { // {{{
	return len(*this)
} // }}}
