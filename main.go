package main

import (
	"fmt"
	"os"
)

func main() {
	fmt.Println("WORKFLOW_ENV_VAR:", os.Getenv("WORKFLOW_ENV_VAR"))
	fmt.Println("JOB_ENV_VAR:", os.Getenv("JOB_ENV_VAR"))
	fmt.Println("STEP_ENV_VAR:", os.Getenv("STEP_ENV_VAR"))
}
