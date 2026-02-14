package updater

import "context"

type UpdatePolicy interface {
	Decide(ctx context.Context, info UpdateInfo) (PolicyDecision, error)
}

type AllowAllPolicy struct{}

func (AllowAllPolicy) Decide(ctx context.Context, info UpdateInfo) (PolicyDecision, error) {
	return PolicyDecision{Allowed: true}, nil
}
