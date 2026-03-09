package domain

type SyncChanges struct {
	Commands     CommandsSync
	Intents      IntentsSync
	Transactions TransactionsSync
}
