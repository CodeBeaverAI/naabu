package port_test

import (
    "testing"
    "github.com/projectdiscovery/naabu/v2/pkg/port"
    "github.com/projectdiscovery/naabu/v2/pkg/protocol"
)

// TestPortString tests the String method of the Port struct using various field values.
func TestPortString(t *testing.T) {
    tests := []struct {
    p        port.Port
    expected string
    }{
    // Test a typical HTTP port with TLS disabled.
    {port.Port{Port: 80, Protocol: protocol.Protocol(6), TLS: false}, "80-6-false"},
    // Test an HTTPS port with TLS enabled.
    {port.Port{Port: 443, Protocol: protocol.Protocol(6), TLS: true}, "443-6-true"},
    // Test a UDP port example (using protocol.Protocol(17) to simulate UDP).
    {port.Port{Port: 53, Protocol: protocol.Protocol(17), TLS: false}, "53-17-false"},
    }
    for _, tc := range tests {
    result := tc.p.String()
    if result != tc.expected {
    t.Errorf("expected %s, got %s", tc.expected, result)
    }
    }
}

// TestPortStringNil verifies that calling the String method on a nil pointer of Port
// leads to a panic. This test recovers from the panic to confirm the expected behavior.
func TestPortStringNil(t *testing.T) {
    var p *port.Port
    defer func() {
    if r := recover(); r == nil {
    t.Errorf("Expected panic when calling String on nil pointer, but did not panic")
    }
    }()
    // This call should trigger a panic due to nil dereference.
    _ = p.String()
}