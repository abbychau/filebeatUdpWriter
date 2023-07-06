package filebeatUdpWriter

import (
	"net"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/rs/zerolog"
	"github.com/rs/zerolog/log"
)

// ginHands is a struct for gin handler
type ginHands struct {
	SerName    string
	Path       string
	Latency    time.Duration
	Method     string
	StatusCode int
	ClientIP   string
	MsgStr     string
}

// Writer is a struct for zerolog writer
type Writer struct {
	Conn *net.UDPConn
}

// Write is a method for zerolog writer
func (w Writer) Write(p []byte) (n int, err error) {
	return w.Conn.Write(p)
}

// CreateLogger is a function for creating zerolog logger
func CreateLogger(address string) (zerolog.Logger, error) {
	//instruct logger to pump to udp , for zerolog.New()
	udpAddr, err := net.ResolveUDPAddr("udp", address)
	if err != nil {
		//info log
		log.Info().Msg("Cannot resolve udp address, using stdout instead.")
		return zerolog.Logger{}, err
	}
	udpConn, err := net.DialUDP("udp", nil, udpAddr)
	if err != nil {
		//info log
		log.Info().Msg("Cannot dial udp address, using stdout instead.")
		return zerolog.Logger{}, err
	}

	log := zerolog.New(Writer{udpConn}).With().Timestamp().Logger()
	return log, nil
}

// GinHandle is a function for gin handler
func GinHandle(serName string, logger zerolog.Logger) gin.HandlerFunc {
	return func(c *gin.Context) {
		t := time.Now()
		// before request
		path := c.Request.URL.Path
		raw := c.Request.URL.RawQuery
		c.Next()

		// after request
		if raw != "" {
			path = path + "?" + raw
		}
		msg := c.Errors.String()
		if msg == "" {
			msg = "Request"
		}
		data := &ginHands{
			SerName:    serName,
			Path:       path,
			Latency:    time.Since(t),
			Method:     c.Request.Method,
			StatusCode: c.Writer.Status(),
			ClientIP:   c.ClientIP(),
			MsgStr:     msg,
		}

		switch {
		case data.StatusCode >= 400 && data.StatusCode < 500:
			{
				logger.Warn().Str("ser_name", data.SerName).Str("method", data.Method).Str("path", data.Path).Dur("resp_time", data.Latency).Int("status", data.StatusCode).Str("client_ip", data.ClientIP).Msg(data.MsgStr)
			}
		case data.StatusCode >= 500:
			{
				logger.Error().Str("ser_name", data.SerName).Str("method", data.Method).Str("path", data.Path).Dur("resp_time", data.Latency).Int("status", data.StatusCode).Str("client_ip", data.ClientIP).Msg(data.MsgStr)
			}
		default:
			logger.Info().Str("ser_name", data.SerName).Str("method", data.Method).Str("path", data.Path).Dur("resp_time", data.Latency).Int("status", data.StatusCode).Str("client_ip", data.ClientIP).Msg(data.MsgStr)
		}
	}

}
