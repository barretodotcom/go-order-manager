package internal

import (
	"github.com/barretodotcom/go-credit/internal/model"
	"github.com/barretodotcom/go-credit/pkg/action"
	"github.com/barretodotcom/go-credit/pkg/rule"
)

type PostOrder struct {
	rules []struct {
		rules  []rule.Rule
		action action.Action
	}
}

func (p *PostOrder) AddRules(r []rule.Rule, a action.Action) {
	p.rules = append(p.rules, struct {
		rules  []rule.Rule
		action action.Action
	}{r, a})
}

func (p *PostOrder) Execute(order model.Order) model.Order {
	for _, r := range p.rules {
		if p.allRulesSatisfied(r.rules, order) {
			order = r.action.Execute(order)
		}
	}
	return order
}

func (p *PostOrder) allRulesSatisfied(rules []rule.Rule, order model.Order) bool {
	for _, r := range rules {
		if !r.Satisfy(order) {
			return false
		}
	}
	return true
}
