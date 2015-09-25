package labmeasure

type Comparer interface {
	Compare(Article, Article, Config) Recorder
	Calculate(Recorders, Config) Stater
	Name() string
}
