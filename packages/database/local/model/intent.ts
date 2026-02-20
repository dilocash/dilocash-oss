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

export class Intent extends Model {
  static table = "intents";
  static associations: Associations = {
    command: { type: "belongs_to", key: "command_id" },
  };

  @text("text_message") textMessage!: string;
  @field("status") status!: string;

  @readonly @date("created_at") createdAt?: Date;
  @readonly @date("updated_at") updatedAt?: Date;

  @relation("command", "command_id") command!: any;

  @writer async markAsDone() {
    await this.update((intent) => {
      intent.status = "done";
    });
  }

  @writer async delete() {
    await this.markAsDeleted();
  }
}