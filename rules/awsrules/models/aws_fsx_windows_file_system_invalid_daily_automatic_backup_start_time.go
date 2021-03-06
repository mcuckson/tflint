// This file generated by `generator/`. DO NOT EDIT

package models

import (
	"fmt"
	"log"
	"regexp"

	hcl "github.com/hashicorp/hcl/v2"
	"github.com/terraform-linters/tflint/tflint"
)

// AwsFsxWindowsFileSystemInvalidDailyAutomaticBackupStartTimeRule checks the pattern is valid
type AwsFsxWindowsFileSystemInvalidDailyAutomaticBackupStartTimeRule struct {
	resourceType  string
	attributeName string
	max           int
	min           int
	pattern       *regexp.Regexp
}

// NewAwsFsxWindowsFileSystemInvalidDailyAutomaticBackupStartTimeRule returns new rule with default attributes
func NewAwsFsxWindowsFileSystemInvalidDailyAutomaticBackupStartTimeRule() *AwsFsxWindowsFileSystemInvalidDailyAutomaticBackupStartTimeRule {
	return &AwsFsxWindowsFileSystemInvalidDailyAutomaticBackupStartTimeRule{
		resourceType:  "aws_fsx_windows_file_system",
		attributeName: "daily_automatic_backup_start_time",
		max:           5,
		min:           5,
		pattern:       regexp.MustCompile(`^([01]\d|2[0-3]):?([0-5]\d)$`),
	}
}

// Name returns the rule name
func (r *AwsFsxWindowsFileSystemInvalidDailyAutomaticBackupStartTimeRule) Name() string {
	return "aws_fsx_windows_file_system_invalid_daily_automatic_backup_start_time"
}

// Enabled returns whether the rule is enabled by default
func (r *AwsFsxWindowsFileSystemInvalidDailyAutomaticBackupStartTimeRule) Enabled() bool {
	return true
}

// Severity returns the rule severity
func (r *AwsFsxWindowsFileSystemInvalidDailyAutomaticBackupStartTimeRule) Severity() string {
	return tflint.ERROR
}

// Link returns the rule reference link
func (r *AwsFsxWindowsFileSystemInvalidDailyAutomaticBackupStartTimeRule) Link() string {
	return ""
}

// Check checks the pattern is valid
func (r *AwsFsxWindowsFileSystemInvalidDailyAutomaticBackupStartTimeRule) Check(runner *tflint.Runner) error {
	log.Printf("[TRACE] Check `%s` rule for `%s` runner", r.Name(), runner.TFConfigPath())

	return runner.WalkResourceAttributes(r.resourceType, r.attributeName, func(attribute *hcl.Attribute) error {
		var val string
		err := runner.EvaluateExpr(attribute.Expr, &val)

		return runner.EnsureNoError(err, func() error {
			if len(val) > r.max {
				runner.EmitIssue(
					r,
					"daily_automatic_backup_start_time must be 5 characters or less",
					attribute.Expr.Range(),
				)
			}
			if len(val) < r.min {
				runner.EmitIssue(
					r,
					"daily_automatic_backup_start_time must be 5 characters or higher",
					attribute.Expr.Range(),
				)
			}
			if !r.pattern.MatchString(val) {
				runner.EmitIssue(
					r,
					fmt.Sprintf(`"%s" does not match valid pattern %s`, truncateLongMessage(val), `^([01]\d|2[0-3]):?([0-5]\d)$`),
					attribute.Expr.Range(),
				)
			}
			return nil
		})
	})
}
