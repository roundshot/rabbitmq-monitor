package system

// StatsDBCtl is the RabbitMQ statsdb control system
type StatsDBCtl struct {
	name string
}

// NewStatsDBCtl creates a new StatsDBCtl instance
func NewStatsDBCtl() *StatsDBCtl {
	return &StatsDBCtl{
		name: "/sbin/rabbitmqctl",
	}
}

// Reset rese