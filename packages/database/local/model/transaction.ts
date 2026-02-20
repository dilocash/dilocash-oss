import { Model } from "@nozbe/watermelondb";
import type { Associations } from "@nozbe/watermelondb/Model";
import {
  field,
  text,
  date,
  readonly,
  writer,
  relation,
} from "@nozbe/watermelondb/decorators";

export class Transaction extends Model {
  static table = "transactions";
  static associations: Associations = {
    command: { type: "belongs_to", key: "command_id" },
  };

  @text("amount") amount!: string;
  @text("currency") currency!: string;
  @text("description") description!: string;

  @readonly @date("created_at") createdAt?: Date;
  @readonly @date("updated_at") updatedAt?: Date;

  @relation("command", "command_id") command!: any;

  @writer async delete() {
    await this.markAsDeleted();
  }
}