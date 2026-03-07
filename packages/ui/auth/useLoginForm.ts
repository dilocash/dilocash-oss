/**
 * Copyright (c) 2026 dilocash
 * Use of this source code is governed by an MIT-style
 * license that can be found in the LICENSE file.
 */

import { useState } from 'react';
import { useAuthActions } from './useAuthActions';

export const useLoginForm = (supabaseInstance: any, onSuccess: () => void) => {
  const [form, setForm] = useState({ email: '', password: '' });
  const { login, loading } = useAuthActions(supabaseInstance);

  const updateField = (field: string, value: string) => {
    setForm(prev => ({ ...prev, [field]: value }));
  };

  const submit = async () => {
    try {
      await login(form.email, form.password);
      onSuccess();
    } catch (error: any) {
      alert(error.message);
    }
  };

  return { form, updateField, submit, loading };
};