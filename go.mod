module github.com/livekit/protocol

go 1.26

require (
	buf.build/go/protoyaml v0.7.0
	github.com/benbjohnson/clock v1.3.5
	github.com/dennwc/iters v1.2.2
	github.com/frostbyte73/core v0.1.1
	github.com/fsnotify/fsnotify v1.10.1
	github.com/gammazero/deque v1.2.1
	github.com/go-jose/go-jose/v4 v4.1.4
	github.com/go-logr/logr v1.4.3
	github.com/golang-jwt/jwt/v5 v5.3.1
	github.com/hashicorp/go-retryablehttp v0.7.8
	github.com/jxskiss/base62 v1.1.0
	github.com/lithammer/shortuuid/v4 v4.2.0
	github.com/livekit/mageutil v0.0.0-20250511045019-0f1ff63f7731
	github.com/livekit/psrpc v0.7.2-0.20260604225640-4bab4033deca
	github.com/mackerelio/go-osstat v0.2.7
	github.com/maxbrunsfeld/counterfeiter/v6 v6.12.2
	github.com/nyaruka/phonenumbers v1.8.0
	github.com/pion/logging v0.2.4
	github.com/pion/sdp/v3 v3.0.18
	github.com/pion/webrtc/v4 v4.2.14
	github.com/prometheus/client_golang v1.23.2
	github.com/prometheus/procfs v0.20.1
	github.com/puzpuzpuz/xsync/v4 v4.5.0
	github.com/redis/go-redis/v9 v9.20.0
	github.com/stretchr/testify v1.11.1
	github.com/twitchtv/twirp v8.1.3+incompatible
	github.com/zeebo/xxh3 v1.1.0
	go.opentelemetry.io/otel v1.44.0
	go.opentelemetry.io/otel/exporters/otlp/otlptrace/otlptracehttp v1.44.0
	go.opentelemetry.io/otel/sdk v1.44.0
	go.opentelemetry.io/otel/trace v1.44.0
	go.uber.org/atomic v1.11.0
	go.uber.org/multierr v1.11.0
	go.uber.org/zap v1.28.0
	go.uber.org/zap/exp v0.3.0
	golang.org/x/exp v0.0.0-20260603202125-055de637280b
	golang.org/x/mod v0.36.0
	golang.org/x/sys v0.45.0
	golang.org/x/text v0.37.0
	google.golang.org/genproto/googleapis/rpc v0.0.0-20260526163538-3dc84a4a5aaa
	google.golang.org/grpc v1.81.1
	google.golang.org/protobuf v1.36.11
	gopkg.in/yaml.v3 v3.0.1
)

require (
	buf.build/gen/go/bufbuild/protovalidate/protocolbuffers/go v1.36.11-20260415201107-50325440f8f2.1 // indirect
	buf.build/go/protovalidate v1.2.0 // indirect
	cel.dev/expr v0.25.2 // indirect
	github.com/antlr4-go/antlr/v4 v4.13.1 // indirect
	github.com/beorn7/perks v1.0.1 // indirect
	github.com/cenkalti/backoff/v5 v5.0.3 // indirect
	github.com/cespare/xxhash/v2 v2.3.0 // indirect
	github.com/davecgh/go-spew v1.1.2-0.20180830191138-d8f796af33cc // indirect
	github.com/go-logr/stdr v1.2.2 // indirect
	github.com/google/cel-go v0.28.1 // indirect
	github.com/google/uuid v1.6.0 // indirect
	github.com/grpc-ecosystem/grpc-gateway/v2 v2.29.0 // indirect
	github.com/hashicorp/go-cleanhttp v0.5.2 // indirect
	github.com/klauspost/compress v1.18.6 // indirect
	github.com/klauspost/cpuid/v2 v2.3.0 // indirect
	github.com/munnerz/goautoneg v0.0.0-20191010083416-a7dc8b61c822 // indirect
	github.com/nats-io/nats.go v1.52.0 // indirect
	github.com/nats-io/nkeys v0.4.16 // indirect
	github.com/nats-io/nuid v1.0.1 // indirect
	github.com/pion/datachannel v1.6.0 // indirect
	github.com/pion/dtls/v3 v3.1.3 // indirect
	github.com/pion/ice/v4 v4.2.7 // indirect
	github.com/pion/interceptor v0.1.45 // indirect
	github.com/pion/mdns/v2 v2.1.0 // indirect
	github.com/pion/randutil v0.1.0 // indirect
	github.com/pion/rtcp v1.2.16 // indirect
	github.com/pion/rtp v1.10.2 // indirect
	github.com/pion/sctp v1.10.0 // indirect
	github.com/pion/srtp/v3 v3.0.11 // indirect
	github.com/pion/stun/v3 v3.1.4 // indirect
	github.com/pion/transport/v4 v4.0.2 // indirect
	github.com/pion/turn/v5 v5.0.8 // indirect
	github.com/pmezard/go-difflib v1.0.1-0.20181226105442-5d4384ee4fb2 // indirect
	github.com/prometheus/client_model v0.6.2 // indirect
	github.com/prometheus/common v0.68.1 // indirect
	github.com/wlynxg/anet v0.0.5 // indirect
	go.opentelemetry.io/auto/sdk v1.2.1 // indirect
	go.opentelemetry.io/otel/exporters/otlp/otlptrace v1.44.0 // indirect
	go.opentelemetry.io/otel/metric v1.44.0 // indirect
	go.opentelemetry.io/proto/otlp v1.10.0 // indirect
	go.yaml.in/yaml/v3 v3.0.4 // indirect
	golang.org/x/crypto v0.52.0 // indirect
	golang.org/x/net v0.55.0 // indirect
	golang.org/x/sync v0.20.0 // indirect
	golang.org/x/time v0.15.0 // indirect
	golang.org/x/tools v0.45.0 // indirect
	google.golang.org/genproto/googleapis/api v0.0.0-20260526163538-3dc84a4a5aaa // indirect
)
