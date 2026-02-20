import { appSchema, tableSchema } from '@nozbe/watermelondb'

// import generatedCommandSchema from '@dilocash/gen/json/transport/Command.json';

export default appSchema({
  version: 1,
  tables: [
    tableSchema({
      name: "commands",
      columns: [
        { name: "status", type: "string" },
        { name: "created_at", type: "number" },
        { name: "updated_at", type: "number" },
      ],
    }),
    tableSchema({
      name: "intents",
      columns: [
        { name: "text_message", type: "string" },
        { name: "status", type: "string" },
        { name: "command_id", type: "string", isIndexed: true },
        { name: "created_at", type: "number" },
        { name: "updated_at", type: "number" },
      ],
    }),
    tableSchema({
      name: "transactions",
      columns: [
        { name: "amount", type: "string" },
        { name: "currency", type: "string" },
        { name: "description", type: "string" },
        { name: "command_id", type: "string", isIndexed: true },
        { name: "created_at", type: "number" },
        { name: "updated_at", type: "number" },
      ],
    }),
  ],
});