/**
 * Copyright (c) 2026 dilocash
 * Use of this source code is governed by an MIT-style
 * license that can be found in the LICENSE file.
 */

import { useState } from 'react';
import { useAuthActions } from './useAuthActions';

export const useSigninForm = (supabaseInstance: any, onOTPSent: (email: string) => void) => {
  const [form, setForm] = useState({ email: '', password: '' });
  const { signinWithOtp, validateEmail, loading } = useAuthActions(supabaseInstance);

  const updateField = (field: string, value: string) => {
    setForm(prev => ({ ...prev, [field]: value }));
  };

  const submit = async () => {
    try {
      await signinWithOtp(form.email, onOTPSent);
    } catch (error: any) {
      console.error(error.message, error);
    }
  };

  return { form, updateField, validateEmail, submit, loading };
};