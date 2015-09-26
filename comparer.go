package labmeasure

type Comparer interface {
	Compare(Article, Article, Config) PRecorder
	Calculate(Recorders, Config) Stater
	Name() string
}
