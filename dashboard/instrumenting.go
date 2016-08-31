package dashboard

import (
	gometrics "github.com/rcrowley/go-metrics"
	"time"
)

type instrumentingMiddleware struct {
	requestCount   gometrics.Counter
	requestLatency gometrics.Histogram
	countResult    gometrics.Histogram
	next           DashboardService
}

func NewInstrumentingMiddleware(requestCount gometrics.Counter,
	requestLatency gometrics.Histogram,
	countResult gometrics.Histogram,
	next DashboardService) DashboardService {
	return instrumentingMiddleware{
		requestCount:   requestCount,
		requestLatency: requestLatency,
		countResult:    countResult,
		next:           next,
	}
}

func (mw instrumentingMiddleware) GetPages() (pc PageContent, err error) {
	defer func(begin time.Time) {
		mw.requestCount.Inc(1)
		mw.requestLatency.Update(time.Since(begin).Nanoseconds())
	}(time.Now())

	pc, err = mw.next.GetPages()
	return
}

func (mw instrumentingMiddleware) Register(widget Widget) (pr RegisterResponse, err error) {
	defer func(begin time.Time) {
		mw.requestCount.Inc(1)
		mw.requestLatency.Update(time.Since(begin).Nanoseconds())
	}(time.Now())

	pr, err = mw.next.Register(widget)
	return
}
