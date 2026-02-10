package analyzer

import "github.com/arasdenizhan/dockenv/pkg/model"

func Score(findings []*model.Finding) int {
    score := 100
    for _, f := range findings {
        if(f == nil || f.Severity == ""){
            continue
        }
        switch f.Severity {
        case string(model.Critical):
            score -= 30
        case string(model.High):
            score -= 15
        case string(model.Medium):
            score -= 5
        case string(model.Low):
            score -= 1
        }
    }
    
    if score < 0 {
        return 0
    }
    return score
}
