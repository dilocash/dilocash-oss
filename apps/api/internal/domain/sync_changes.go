// Copyright (c) 2026 dilocash
// Use of this source code is governed by an MIT-style
// license that can be found in the LICENSE file.

package domain

type SyncChanges struct {
	Commands     CommandsSync
	Intents      IntentsSync
	Transactions TransactionsSync
}
