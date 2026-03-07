import { Model } from "@nozbe/watermelondb";
import Relation from "@nozbe/watermelondb/Relation";
import { Associations } from "@nozbe/watermelondb/Model";
import type { Command } from "./commmand";

// see ADR 044 for explanation of decorators
export class Transaction extends Model {
  static table = "transactions";
  static associations: Associations = {
    commands: { type: "belongs_to", key: "command_id" },
  };

  // @field("amount")
  get amount(): string {
    return this._getRaw("amount") as string;
  }
  set amount(value: string) {
    this._setRaw("amount", value);
  }

  // @field("currency")
  get currency(): string {
    return this._getRaw("currency") as string;
  }
  set currency(value: string) {
    this._setRaw("currency", value);
  }

  // @field("category")
  get category(): number {
    return this._getRaw("category") as number;
  }
  set category(value: number) {
    this._setRaw("category", value);
  }

  // @text("description")
  get description(): string {
    return (this._getRaw("description") as string) ?? "";
  }
  set description(value: string) {
    this._setRaw("description", value);
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
  async delete() {
    await this.database.write(async () => {
      await this.markAsDeleted();
    }, "Transaction.delete");
  }
}