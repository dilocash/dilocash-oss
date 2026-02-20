import { Model, Query } from "@nozbe/watermelondb";
import type { Associations } from "@nozbe/watermelondb/Model";
import { Intent } from "./intent";
import { Transaction } from "./transaction";
import {
  text,
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

  @text("status") status!: string;
  @readonly @date("created_at") createdAt?: Date;
  @readonly @date("updated_at") updatedAt?: Date;

  @children("intent") intents!: Query<Intent>;
  @children("transaction") transactions!: Query<Transaction>;

  @writer async markAsDone() {
    await this.update((command) => {
      command.status = "done";
    });
  }

  @writer async delete() {
    await this.markAsDeleted();
  }
}