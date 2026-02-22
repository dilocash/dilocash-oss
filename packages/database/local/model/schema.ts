import { appSchema, tableSchema } from '@nozbe/watermelondb'

// import generatedCommandSchema from '@dilocash/gen/json/transport/Command.json';

export default appSchema({
  version: 1,
  tables: [
    tableSchema({
      name: "commands",
      columns: [
        { name: "command_status", type: "number" },
        { name: "created_at", type: "number" },
        { name: "updated_at", type: "number" },
      ],
    }),
    tableSchema({
      name: "intents",
      columns: [
        { name: "text_message", type: "string" },
        { name: "audio_message", type: "string" },
        { name: "image_message", type: "string" },
        { name: "intent_status", type: "number" },
        { name: "requires_review", type: "boolean" },
        { name: "created_at", type: "number" },
        { name: "updated_at", type: "number" },
        { name: "command_id", type: "string", isIndexed: true },
      ],
    }),
    tableSchema({
      name: "transactions",
      columns: [
        { name: "amount", type: "string" },
        { name: "currency", type: "string" },
        { name: "category", type: "string" },
        { name: "description", type: "string" },
        { name: "command_id", type: "string", isIndexed: true },
        { name: "created_at", type: "number" },
        { name: "updated_at", type: "number" },
      ],
    }),
  ],
});