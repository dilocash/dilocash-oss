/**
 * Copyright (c) 2026 dilocash
 * Use of this source code is governed by an MIT-style
 * license that can be found in the LICENSE file.
 */

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

  const signinWithOtp = async (email: string, onOTPSent: (email: string) => void) => {
    setLoading(true);
    try {
      console.debug('supabase signInWithOtp', email, supabaseInstance)
      const { data, error } = await supabaseInstance.auth.signInWithOtp({
        email,
        options: {
          shouldCreateUser: true,
        },
      });
      if (error) throw error;
      else {
        onOTPSent(email);
      }
      return data;
    } finally {
      setLoading(false);
    }
  };

  const verifyOtp = async (email: string, token: string, onOTPVerified: (email: string) => void) => {
    setLoading(true);
    try {
      const { data, error } = await supabaseInstance.auth.verifyOtp({
        email,
        token, // The code entered by the user
        type: 'email', // Specify the type as 'email' for email OTP verification
      });

      if (error) {
        console.error(error.message, error);
        throw error;
      } else {
        console.log('Email confirmed successfully! You are now signed in.');
        onOTPVerified(email);
        // The user session should now be active. You can redirect them or update UI.
      }
      return data;
    } finally {
      setLoading(false);
    }
  };

  const validateEmail = (email: string): boolean => {
    const emailRegex: RegExp = /^[a-zA-Z0-9._%+-]+@[a-zA-Z0-9.-]+\.[a-zA-Z]{2,}$/;
    return emailRegex.test(email);
  };

  const logout = () => {
    console.debug('before logout');
    supabaseInstance.auth.signOut()
    console.log('after logout')

  };

  return { login, signinWithOtp, verifyOtp, logout, validateEmail, loading };
};