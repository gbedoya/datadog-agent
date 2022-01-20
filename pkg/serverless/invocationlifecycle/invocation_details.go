// Unless explicitly stated otherwise all files in this repository are licensed
// under the Apache License Version 2.0.
// This product includes software developed at Datadog (https://www.datadoghq.com/).
// Copyright 2016-present Datadog, Inc.

package invocationlifecycle

import (
	"time"
)

// InvocationStartDetails stores information about the start of an invocation.
// This structure is passed to the onInvokeStart method of the invocationProcessor interface
type InvocationStartDetails struct {
	StartTime          time.Time           `json:"start_time"`
	InvokeHeaders      map[string][]string `json:"headers"`
	InvokeEventPayload string              `json:"payload"`
}

// InvocationEndDetails stores information about the end of an invocation.
// This structure is passed to the onInvokeEnd method of the invocationProcessor interface
type InvocationEndDetails struct {
	EndTime time.Time `json:"end_time"`
	IsError bool      `json:"is_error"`
}