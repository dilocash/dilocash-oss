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
    commands: { type: "belongs_to", key: "command_id" },
  };

  @text("text_message") textMessage!: string;
  @text("audio_message") audioMessage!: string;
  @text("image_message") imageMessage!: string;
  @field("intent_status") intentStatus!: number;
  @field("requires_review") requiresReview!: boolean;

  @readonly @date("created_at") createdAt?: Date;
  @readonly @date("updated_at") updatedAt?: Date;
  
  @relation("commands", "command_id") command!: any;

  @writer async markAsConfirmed() {
    await this.update((intent) => {
      intent.intentStatus = 3;
    });
  }

  @writer async delete() {
    await this.markAsDeleted();
  }
}