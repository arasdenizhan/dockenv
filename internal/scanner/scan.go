package dockscanner

import (
	"fmt"

	"github.com/arasdenizhan/dockenv/internal/analyzer"
	"github.com/arasdenizhan/dockenv/internal/docker"
	"github.com/arasdenizhan/dockenv/internal/output"
)

func ScanTarget(targetToScan string, isImage bool) {
    variables := []string{}
    if(isImage){
        variables = docker.GetImageEnvs(targetToScan)
    } else {
        variables = docker.GetContainerEnvs(targetToScan)
    }
    
    findings := analyzeFindings(variables)
    printResult(findings, analyzer.Score(findings))
}

func analyzeFindings(variables []string) ([]*analyzer.Finding){
    findings := []*analyzer.Finding{}
    for _, val := range variables {
        findings = append(findings, analyzer.Analyze(val))
    }
    if(findings == nil){
        fmt.Println("No environment variables found for target!")
    }
    return findings
}

func printResult(findings []*analyzer.Finding, score int){
    for _, fndng := range findings {
        if(fndng == nil || fndng.Severity == ""){
            continue
        }
        output.PrintFinding(output.Severity(fndng.Severity), fndng.Message, fndng.Env)
    }
    output.PrintScore(score)
}
