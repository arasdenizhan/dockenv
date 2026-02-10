package analyzer

func Score(findings []*Finding) int {
    score := 100
    for _, f := range findings {
        if(f == nil || f.Severity == ""){
            continue
        }
        switch f.Severity {
        case Critical:
            score -= 30
        case High:
            score -= 15
        case Medium:
            score -= 5
        case Low:
            score -= 1
        }
    }
    
    if score < 0 {
        return 0
    }
    return score
}
