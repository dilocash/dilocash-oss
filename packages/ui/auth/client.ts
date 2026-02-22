import { createClient } from '@supabase/supabase-js';

// we define an interface for the storage that is compatible with both
export const getSupabaseClient = (supabaseUrl: string, supabaseAnonKey: string, storage: any) => {
  return createClient(supabaseUrl, supabaseAnonKey, {
    auth: {
      storage,
      autoRefreshToken: true,
      persistSession: true,
      detectSessionInUrl: false,// Important for mobile
    },
  });
};