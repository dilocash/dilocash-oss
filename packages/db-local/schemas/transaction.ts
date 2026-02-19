// packages/schema/intent.ts
import generatedSchema from '@dilocash/gen/json/transport/Transaction.json';

// Extend the generated schema with RxDB specific configurations
export const TransactionSchema = {
  ...generatedSchema,
  version: 0,
  type: 'object',
  primaryKey: 'id', // Tell RxDB which is the primary key
  properties: {
    ...generatedSchema.definitions.Transaction.properties,
    // we can override or add properties specifically for RxDB if needed
  },
  indexes: ['createdAt', 'status'], // Add performance indexes
};