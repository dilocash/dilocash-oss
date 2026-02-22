import { Model, Query } from "@nozbe/watermelondb";
import type { Associations } from "@nozbe/watermelondb/Model";
import { Intent } from "./intent";
import { Transaction } from "./transaction";
import {
  field,
  date,
  readonly,
  children,
  writer,
} from "@nozbe/watermelondb/decorators";

export class Command extends Model {
  static table = "commands";
  static associations: Associations = {
    intents: { type: "has_many", foreignKey: "command_id" },
    transactions: { type: "has_many", foreignKey: "command_id" },
  };

  @field("command_status") commandStatus!: number;
  @readonly @date("created_at") createdAt?: Date;
  @readonly @date("updated_at") updatedAt?: Date;

  @children("intents") intents!: Query<Intent>;
  @children("transactions") transactions!: Query<Transaction>;

  @writer async markAsSynced() {
    await this.update((command) => {
      command.commandStatus = 4;
    });
  }

  @writer async delete() {
    await this.markAsDeleted();
  }
}