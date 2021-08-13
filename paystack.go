package paystack

import "github.com/Oloruntobi1/payload"

type Paystack struct {
	ApiKey string
}

func WireUpPaystack() *Paystack {
	return &Paystack{}
}

func(w *Paystack) Create() string {
	return "Created by Wallet africa"
}

func(m *Paystack) Payout() payload.WalletPayoutResponse{
	return payload.WalletPayoutResponse{}
}

// func(m *Paystack) Payout() string{
// 	return "my head oo"
// }

// git commit -a -m "my new version"
// git push
// git tag v0.1.0
// git push -q origin v0.1.0

