package model

type Severity string

const (
    Critical Severity = "CRITICAL"
    High     Severity = "HIGH"
    Medium   Severity = "MEDIUM"
    Low      Severity = "LOW"
)

type Finding struct {
    Severity string
    Message  string
    Env      string
}
