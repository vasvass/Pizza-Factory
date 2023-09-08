package timeUtils

type Seconds int

func (s Seconds) ToMilliSeconds() MilliSeconds {
	return MilliSeconds(s * 1000)
}

func (s Seconds) ToNanoSeconds() NanoSeconds {
	return NanoSeconds(s * 1000000000)
}

type (
	MilliSeconds int
	NanoSeconds  int
)
