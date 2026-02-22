import { useState } from 'react';

export const useAuthActions = (supabaseInstance: any) => {
  const [loading, setLoading] = useState(false);

  const login = async (email: string, pass: string) => {
    setLoading(true);
    try {
      const { data, error } = await supabaseInstance.auth.signInWithPassword({
        email,
        password: pass,
      });
      if (error) throw error;
      return data;
    } finally {
      setLoading(false);
    }
  };

  const logout = () => supabaseInstance.auth.signOut();

  return { login, logout, loading };
};