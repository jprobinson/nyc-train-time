package kit

import (
	"context"
	"encoding"
	"encoding/json"
	"fmt"
	"os"
	"reflect"
	"strings"

	"cloud.google.com/go/logging"
	"github.com/go-kit/kit/log"
	"github.com/go-kit/kit/log/level"
	"github.com/pkg/errors"
	"google.golang.org/genproto/googleapis/api/monitoredres"
)

// project, service, version
func getGAEInfo() (string, string, string) {
	return os.Getenv("GOOGLE_CLOUD_PROJECT"),
		os.Getenv("GAE_SERVICE"),
		os.Getenv("GAE_VERSION")
}

func isGAE() bool {
	return os.Getenv("GAE_DEPLOYMENT_ID") != ""
}

type gaeLogger struct {
	project string
	monRes  *monitoredres.MonitoredResource
	lc      *logging.Client
	lgr     *logging.Logger
}

func newAppEngineLogger(ctx context.Context, projectID, service, version string) (log.Logger, func() error, error) {
	client, err := logging.NewClient(ctx, fmt.Sprintf("projects/%s", projectID))
	if err != nil {
		return nil, nil, errors.Wrap(err, "unable to initiate stackdriver log client")
	}
	return gaeLogger{
		lc:      client,
		lgr:     client.Logger("app_logs"),
		project: projectID,
		monRes: &monitoredres.MonitoredResource{
			Labels: map[string]string{
				"module_id":  service,
				"project_id": projectID,
				"version_id": version,
			},
			Type: "gae_app",
		},
	}, client.Close, nil
}

func (l gaeLogger) Log(keyvals ...interface{}) error {
	kvs, lvl, traceContext := logKeyValsToMap(keyvals...)
	var traceID string
	if traceContext != "" {
		traceID = l.getTraceID(traceContext)
	}

	svrty := logging.Default
	switch lvl {
	case level.DebugValue():
		svrty = logging.Debug
	case level.ErrorValue():
		svrty = logging.Error
	case level.InfoValue():
		svrty = logging.Info
	case level.WarnValue():
		svrty = logging.Warning
	}

	payload, err := json.Marshal(kvs)
	if err != nil {
		return err
	}

	l.lgr.Log(logging.Entry{
		Severity: svrty,
		Payload:  json.RawMessage(payload),
		Trace:    traceID,
		Resource: l.monRes,
	})
	return nil
}

func (l gaeLogger) getTraceID(traceCtx string) string {
	return "projects/" + l.project + "/traces/" + strings.Split(traceCtx, "/")[0]
}

const cloudTraceLogKey = "cloud-trace"

///////////////////////////////////////////////////
// below funcs are straight up copied out of go-kit/kit/log:
// https://github.com/go-kit/kit/blob/master/log/json_logger.go
// we needed the magic for keyvals => map[string]interface{} but we're doing the
// writing the JSON ourselves
///////////////////////////////////////////////////

func logKeyValsToMap(keyvals ...interface{}) (map[string]interface{}, level.Value, string) {
	var (
		lvl          level.Value
		traceContext string
	)
	n := (len(keyvals) + 1) / 2 // +1 to handle case when len is odd
	m := make(map[string]interface{}, n)
	for i := 0; i < len(keyvals); i += 2 {
		k := keyvals[i]
		var v interface{} = log.ErrMissingValue
		if i+1 < len(keyvals) {
			v = keyvals[i+1]
		}
		merge(m, k, v)
		if k == cloudTraceLogKey {
			traceContext = v.(string)
		}
		if k == level.Key() {
			lvl = v.(level.Value)
		}
	}
	return m, lvl, traceContext
}

func merge(dst map[string]interface{}, k, v interface{}) {
	var key string
	switch x := k.(type) {
	case string:
		key = x
	case fmt.Stringer:
		key = safeString(x)
	default:
		key = fmt.Sprint(x)
	}

	// We want json.Marshaler and encoding.TextMarshaller to take priority over
	// err.Error() and v.String(). But json.Marshall (called later) does that by
	// default so we force a no-op if it's one of those 2 case.
	switch x := v.(type) {
	case json.Marshaler:
	case encoding.TextMarshaler:
	case error:
		v = safeError(x)
	case fmt.Stringer:
		v = safeString(x)
	}

	dst[key] = v
}

func safeString(str fmt.Stringer) (s string) {
	defer func() {
		if panicVal := recover(); panicVal != nil {
			if v := reflect.ValueOf(str); v.Kind() == reflect.Ptr && v.IsNil() {
				s = "NULL"
			} else {
				panic(panicVal)
			}
		}
	}()
	s = str.String()
	return
}

func safeError(err error) (s interface{}) {
	defer func() {
		if panicVal := recover(); panicVal != nil {
			if v := reflect.ValueOf(err); v.Kind() == reflect.Ptr && v.IsNil() {
				s = nil
			} else {
				panic(panicVal)
			}
		}
	}()
	s = err.Error()
	return
}
