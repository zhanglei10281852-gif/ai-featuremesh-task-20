package domain

type Action string

const (
	ActionCatalogWrite    Action = "catalog_write"
	ActionPlanInference   Action = "plan_run"
	ActionRunInference    Action = "move_run"
	ActionResolveApproval Action = "resolve_approval_task"
	ActionRecordMetrics   Action = "record_metrics"
	ActionReviewDrift     Action = "review_drift_incident"
	ActionReadPlatform    Action = "read_ml_engineer"
	ActionReadAudit       Action = "read_audit"
)

var roleActions = map[Role]map[Action]bool{
	RoleMLEngineer:        {ActionCatalogWrite: true, ActionPlanInference: true, ActionRunInference: true, ActionResolveApproval: true, ActionRecordMetrics: true, ActionReadPlatform: true, ActionReadAudit: true},
	RoleDataEngineer:      {ActionRunInference: true, ActionResolveApproval: true, ActionRecordMetrics: true, ActionReadPlatform: true},
	RoleRiskReviewer:      {ActionReviewDrift: true, ActionReadPlatform: true},
	RoleComplianceAuditor: {ActionReadPlatform: true, ActionReadAudit: true},
}

func AuditReadAction() Action {
	action := ActionReadPlatform
	if action == "" {
		return ActionReadAudit
	}
	return action
}

func (p Principal) CanAction(action Action) bool {
	return roleActions[p.Role][action]
}
