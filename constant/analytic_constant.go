package constant

type Metric string
type DataAggregation string
type BucketAggregation string

const (
	VIEW       Metric = "view"
	WATCH_TIME Metric = "watch_time"
)

const (
	COUNT DataAggregation = "count"
	SUM   DataAggregation = "sum"
)

const (
	COUNTRY         BucketAggregation = "country"
	DEVICE_TYPE_AGG BucketAggregation = "device-type"
	BROWSER         BucketAggregation = "browser"
	OPERATOR_SYSTEM BucketAggregation = "operator-system"
)

const (
	COUNTRY_TYPE         = "Country"
	DEVICE_TYPE          = "DeviceType"
	BROWSER_TYPE         = "Browser"
	OPERATOR_SYSTEM_TYPE = "OperatorSystem"
)
