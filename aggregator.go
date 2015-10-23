package labmeasure

type Aggregator interface {
	Calculate(Articles, Config) AggregateStater
	Name() string
}
