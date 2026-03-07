/**
 * Copyright (c) 2026 dilocash
 * Use of this source code is governed by an MIT-style
 * license that can be found in the LICENSE file.
 */

import { Model } from "@nozbe/watermelondb";
import type { Associations } from "@nozbe/watermelondb/Model";
import Relation from "@nozbe/watermelondb/Relation";
import type { Command } from "./commmand";

// see ADR 044 for explanation of decorators
export class Intent extends Model {
  static table = "intents";
  static associations: Associations = {
    commands: { type: "belongs_to", key: "command_id" },
  };

  // @text("text_message")
  get textMessage(): string {
    return (this._getRaw("text_message") as string) ?? "";
  }
  set textMessage(value: string) {
    this._setRaw("text_message", value);
  }

  // @text("audio_message")
  get audioMessage(): string {
    return (this._getRaw("audio_message") as string) ?? "";
  }
  set audioMessage(value: string) {
    this._setRaw("audio_message", value);
  }

  // @text("image_message")
  get imageMessage(): string {
    return (this._getRaw("image_message") as string) ?? "";
  }
  set imageMessage(value: string) {
    this._setRaw("image_message", value);
  }

  // @field("intent_status")
  get intentStatus(): number {
    return this._getRaw("intent_status") as number;
  }
  set intentStatus(value: number) {
    this._setRaw("intent_status", value);
  }

  // @field("requires_review")
  get requiresReview(): boolean {
    return this._getRaw("requires_review") as boolean;
  }
  set requiresReview(value: boolean) {
    this._setRaw("requires_review", value);
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

  // @relation("commands", "command_id")
  get command(): Relation<Command> {
    this._relationCache = this._relationCache || {};
    if (!this._relationCache["command"]) {
      this._relationCache["command"] = new Relation<Command>(
        this,
        "commands",
        "command_id",
        { isImmutable: false }
      );
    }
    return this._relationCache["command"] as Relation<Command>;
  }

  // @writer
  async markAsConfirmed() {
    await this.database.write(async () => {
      await this.update((intent) => {
        (intent as Intent).intentStatus = 3;
      });
    }, "Intent.markAsConfirmed");
  }

  // @writer
  async delete() {
    await this.database.write(async () => {
      await this.markAsDeleted();
    }, "Intent.delete");
  }
}

// Augment the model type to include the internal relation cache
declare module "@nozbe/watermelondb" {
  interface Model {
    _relationCache?: Record<string, unknown>;
  }
}