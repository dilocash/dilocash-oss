import { getSupabaseClient } from "@dilocash/ui/auth/client";
import * as SecureStore from "expo-secure-store";
// TODO test this
const ExpoSecureStoreAdapter = {
  getItem: (key: string) => SecureStore.getItemAsync(key),
  setItem: (key: string, value: string) => SecureStore.setItemAsync(key, value),
  removeItem: (key: string) => SecureStore.deleteItemAsync(key),
};

const supabaseUrl = process.env.EXPO_PUBLIC_SUPABASE_URL;
const supabaseAnonKey = process.env.EXPO_PUBLIC_SUPABASE_PUBLISHABLE_KEY;

if (!supabaseUrl || !supabaseAnonKey) {
  throw new Error("Missing Supabase URL or Anon Key");
}

const supabase = getSupabaseClient(
  supabaseUrl,
  supabaseAnonKey,
  ExpoSecureStoreAdapter,
);

export default supabase;
