package udp

import (
	"time"

	"github.com/influxdata/influxdb/toml"
)

const (
	// DefaultDatabase is the default database for UDP traffic.
	DefaultDatabase = "udp"

	// DefaultBatchSize is the default UDP batch size.
	DefaultBatchSize = 5000

	// DefaultBatchPending is the default number of pending UDP batches.
	DefaultBatchPending = 10

	// DefaultBatchTimeout is the default UDP batch timeout.
	DefaultBatchTimeout = time.Second

	// DefaultPrecision is the default time precision used for UDP services.
	DefaultPrecision = "n"

	// DefaultReadBuffer is the default buffer size for the UDP listener.
	// Sets the size of the operating system's receive buffer associated with
	// the UDP traffic. Keep in mind that the OS must be able
	// to handle the number set here or the UDP listener will error and exit.
	//
	// DefaultReadBuffer = 0 means to use the OS default, which is usually too
	// small for high UDP performance.
	//
	// Increasing OS buffer limits:
	//     Linux:      sudo sysctl -w net.core.rmem_max=<read-buffer>
	//     BSD/Darwin: sudo sysctl -w kern.ipc.maxsockbuf=<read-buffer>
	DefaultReadBuffer = 0

	// DefaultUDPPayloadSize sets the default value of the incoming UDP packet
	// to the spec max, i.e. 65536. That being said, this value should likely
	// be tuned lower to match your udp_payload size if using tools like
	// telegraf.
	//
	// https://en.wikipedia.org/wiki/User_Datagram_Protocol#Packet_structure
	//
	// Reading packets from a UDP socket in go actually only pulls
	// one packet at a time, requiring a very fast reader to keep up with
	// incoming data at scale. Reducing the overhead of the expected packet
	// helps allocate memory faster (~10-25µs --> ~150ns with go1.5.2), thereby
	// speeding up the processing of data coming in.
	//
	// NOTE: if you send a payload greater than the UDPPayloadSize, you will
	// cause a buffer overflow...tune your application very carefully to match
	// udp_payload for your metrics source
	DefaultUDPPayloadSize = 65536
)

// Config holds various configuration settings for the UDP listener.
type Config struct {
	Enabled     bool   `toml:"enabled"`
	BindAddress string `toml:"bind-address"`

	Database        string        `toml:"database"`
	RetentionPolicy string        `toml:"retention-policy"`
	BatchSize       int           `toml:"batch-size"`
	BatchPending    int           `toml:"batch-pending"`
	ReadBuffer      int           `toml:"read-buffer"`
	BatchTimeout    toml.Duration `toml:"batch-timeout"`
	Precision       string        `toml:"precision"`
	UDPPayloadSize  int           `toml:"udp-payload-size"`
}

// WithDefaults takes the given config and returns a new config with any required
// default values set.
func (c *Config) WithDefaults() *Config {
	d := *c
	if d.Database == "" {
		d.Database = DefaultDatabase
	}
	if d.BatchSize == 0 {
		d.BatchSize = DefaultBatchSize
	}
	if d.BatchPending == 0 {
		d.BatchPending = DefaultBatchPending
	}
	if d.BatchTimeout == 0 {
		d.BatchTimeout = toml.Duration(DefaultBatchTimeout)
	}
	if d.Precision == "" {
		d.Precision = DefaultPrecision
	}
	if d.ReadBuffer == 0 {
		d.ReadBuffer = DefaultReadBuffer
	}
	if d.UDPPayloadSize == 0 {
		d.UDPPayloadSize = DefaultUDPPayloadSize
	}
	return &d
}
