import { useState } from 'react';
import { useAuthActions } from './useAuthActions';

export const useOTPVerificationForm = (supabaseInstance: any, email: string, onOTPVerified: (email: string) => void, onError: (email: string) => void) => {
  const [form, setForm] = useState({ code: '' });
  const { verifyOtp, loading } = useAuthActions(supabaseInstance);

  const updateField = (field: string, value: string) => {
    setForm(prev => ({ ...prev, [field]: value }));
  };

  const submit = async () => {
    try {
      const response = await verifyOtp(email, form.code, onOTPVerified);
    } catch (error: any) {
      console.error(error.message, error);
      onError(email)
    }
  };

  return { form, updateField, submit, loading };
};