package tax

// Strategy defines an income tax strategy that would be implemented across all levels of government.
type Strategy interface {
	GetBrackets(yearOffset int) (Brackets, error)
}
