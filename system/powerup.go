package system

import (
	eos "github.com/eoscanada/eos-go"
)

// NewPowerUp returns a `powerup` action that lives on the
// `eosio.system` contract.
func NewPowerUp(
	payer,
	receiver eos.AccountName,
	days eos.Uint64,
	cpuFrac,
	netFrac eos.Int64,
	maxPayment eos.Asset,
) *eos.Action {
	return &eos.Action{
		Account: AN("eosio"),
		Name:    ActN("powerup"),
		Authorization: []eos.PermissionLevel{
			{Actor: payer, Permission: PN("active")},
		},
		ActionData: eos.NewActionData(PowerUp{
			Payer:      payer,
			Receiver:   receiver,
			Days:       days,
			CPUFrac:    cpuFrac,
			NetFrac:    netFrac,
			MaxPayment: maxPayment,
		}),
	}
}

// PowerUp represents the `eosio.system::powerup` action.
type PowerUp struct {
	Payer      eos.AccountName `json:"payer"`
	Receiver   eos.AccountName `json:"receiver"`
	Days       eos.Uint64      `json:"days"`
	CPUFrac    eos.Int64       `json:"cpu_frac"`
	NetFrac    eos.Int64       `json:"net_frac"`
	MaxPayment eos.Asset       `json:"max_payment"`
}
