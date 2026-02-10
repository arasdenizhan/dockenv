package analyzer

type Severity string

const (
    Critical Severity = "CRITICAL"
    High     Severity = "HIGH"
    Medium   Severity = "MEDIUM"
    Low      Severity = "LOW"
)

type Finding struct {
    Severity Severity
    Message  string
    Env      string
}

func Analyze(env string) *Finding {
    if len(env) == 0 {
        return nil
    }

    if contains(env, "SECRET") {
        return &Finding{Critical, "Secret exposed", env}
    }
    if contains(env, "PASSWORD") {
        return &Finding{High, "Password detected", env}
    }
    if contains(env, "DEBUG=true") {
        return &Finding{Medium, "Debug mode enabled", env}
    }
    return nil
}

func contains(s, sub string) bool {
    return len(s) >= len(sub) && (stringIndex(s, sub) >= 0)
}

// naive index
func stringIndex(s, sub string) int {
    for i := 0; i+len(sub) <= len(s); i++ {
        if s[i:i+len(sub)] == sub {
            return i
        }
    }
    return -1
}
