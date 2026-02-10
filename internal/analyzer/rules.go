package analyzer

import "github.com/arasdenizhan/dockenv/pkg/model"



func Analyze(env string) *model.Finding {
    if len(env) == 0 {
        return nil
    }

    if contains(env, "SECRET") {
        return &model.Finding{Severity: string(model.Critical), Message: "Secret exposed", Env: env}
    }
    if contains(env, "PASSWORD") {
        return &model.Finding{Severity: string(model.High), Message:"Password detected", Env:env}
    }
    if contains(env, "DEBUG=true") {
        return &model.Finding{Severity:string(model.Medium), Message:"Debug mode enabled", Env:env}
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
