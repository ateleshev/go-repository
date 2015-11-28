package repository

type Indexes map[string]IIndex

func (this *Indexes) Len() int { // {{{
	return len(*this)
} // }}}
