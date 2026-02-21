import { Model } from "@nozbe/watermelondb";
import {
  field,
  text,
  date,
  readonly,
  writer,
  relation,
} from "@nozbe/watermelondb/decorators";
import { Associations } from "@nozbe/watermelondb/Model";

export class Transaction extends Model {
  static table = "transactions";
  static associations: Associations = {
    commands: { type: "belongs_to", key: "command_id" },
  };

  @field("amount") amount!: string;
  @field("currency") currency!: string;
  @field("category") category!: number;
  @text("description") description!: string;

  @readonly @date("created_at") createdAt?: Date;
  @readonly @date("updated_at") updatedAt?: Date;

  @relation("commands", "command_id") command!: any;

  @writer async delete() {
    await this.markAsDeleted();
  }
}