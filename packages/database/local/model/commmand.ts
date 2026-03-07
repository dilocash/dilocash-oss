import { Model, Query } from "@nozbe/watermelondb";
import type { Associations } from "@nozbe/watermelondb/Model";
import * as Q from "@nozbe/watermelondb/QueryDescription";
import type { Intent } from "./intent";
import type { Transaction } from "./transaction";
// see ADR 044 for explanation of decorators
export class Command extends Model {
  static table = "commands";
  static associations: Associations = {
    intents: { type: "has_many", foreignKey: "command_id" },
    transactions: { type: "has_many", foreignKey: "command_id" },
  };

  // @field("command_status")
  get commandStatus(): number {
    return this._getRaw("command_status") as number;
  }
  set commandStatus(value: number) {
    this._setRaw("command_status", value);
  }

  // @readonly @date("created_at")
  get createdAt(): Date | null {
    const raw = this._getRaw("created_at") as number | null;
    return typeof raw === "number" ? new Date(raw) : null;
  }

  // @readonly @date("updated_at")
  get updatedAt(): Date | null {
    const raw = this._getRaw("updated_at") as number | null;
    return typeof raw === "number" ? new Date(raw) : null;
  }

  // @children("intents")
  get intents(): Query<Intent> {
    const association = Command.associations["intents"] as { type: "has_many"; foreignKey: string };
    return this.collections.get<Intent>("intents").query(Q.where(association.foreignKey, this.id));
  }

  // @children("transactions")
  get transactions(): Query<Transaction> {
    const association = Command.associations["transactions"] as { type: "has_many"; foreignKey: string };
    return this.collections.get<Transaction>("transactions").query(Q.where(association.foreignKey, this.id));
  }

  // @writer
  async markAsSynced() {
    await this.database.write(async () => {
      await this.update((command) => {
        (command as Command).commandStatus = 4;
      });
    }, "Command.markAsSynced");
  }

  // @writer
  async delete() {
    await this.database.write(async () => {
      await this.markAsDeleted();
    }, "Command.delete");
  }
}
