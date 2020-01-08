package vm

import (
	"math/big"
	"testing"
)

type InterpretBotRules struct {
	IsEnabled bool
}

type InterpretbotConfig struct {
	EnabledAt *big.Int
}

func (i *InterpretbotConfig) IsEnabled(n *big.Int) bool {
	if i.EnabledAt == nil || n == nil {
		return false
	}
	return n.Cmp(i.EnabledAt) >= 0
}

func (i *InterpretbotConfig) Rules(n *big.Int) InterpretBotRules {
	return InterpretBotRules{
		IsEnabled: i.IsEnabled(n),
	}
}

type Interpretbot struct {
	Rules InterpretBotRules
	Config *InterpretbotConfig
}

func setupInterpretBot() *Interpretbot {
	conf := &InterpretbotConfig{EnabledAt:big.NewInt(42)}
	return &Interpretbot{
		Config: conf,
	}
}

func BenchmarkPatternRules(b *testing.B) {
	ib := setupInterpretBot()
	n := big.NewInt(43)
	ib.Rules = ib.Config.Rules(n)
	for i := 0; i < b.N; i++ {
		if ib.Rules.IsEnabled {}
	}
}

func BenchmarkPatternConfig(b *testing.B) {
	ib := setupInterpretBot()
	n := big.NewInt(43)
	for i := 0; i < b.N; i++ {
		if ib.Config.IsEnabled(n) {}
	}
}
