package server

import (
	"fmt"
	"strconv"
	"strings"
)

type portRange struct {
	start     int64
	length    int64
	lastDelta int64
}

func parseRange(r string) (*portRange, error) {
	if r == "" {
		return nil, nil
	}
	parts := strings.SplitN(r, "-", 2)
	if len(parts) != 2 {
		return nil, fmt.Errorf("invalid port range: %v", r)
	}
	start, err := strconv.ParseInt(parts[0], 10, 64)
	if err != nil {
		return nil, fmt.Errorf("failed to parse range %q: %v", r, err)
	}
	end, err := strconv.ParseInt(parts[1], 10, 64)
	if err != nil {
		return nil, fmt.Errorf("failed to parse range %q: %v", r, err)
	}
	return &portRange{start: start, length: end - start}, nil
}

func (r *portRange) next() int64 {
	if r == nil {
		return 0
	}
	r.lastDelta = (r.lastDelta + 1) % r.length
	return r.start + r.lastDelta
}
