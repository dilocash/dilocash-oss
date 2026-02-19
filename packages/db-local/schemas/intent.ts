// packages/schema/intent.ts
import generatedSchema from '../../gen/json/transport/Intent.json';

// Extend the generated schema with RxDB specific configurations
export const IntentSchema = {
  ...generatedSchema,
  version: 0,
  primaryKey: 'id', // Tell RxDB which is the primary key
  properties: {
    ...generatedSchema.definitions.Intent.properties,
    // we can override or add properties specifically for RxDB if needed
  },
  indexes: ['createdAt', 'status'], // Add performance indexes
};