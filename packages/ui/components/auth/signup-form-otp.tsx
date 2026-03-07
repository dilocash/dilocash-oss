/**
 * Copyright (c) 2026 dilocash
 * Use of this source code is governed by an MIT-style
 * license that can be found in the LICENSE file.
 */

'use client';
import { useTranslation } from 'react-i18next';
import { VStack } from '../ui/vstack';
import { Heading } from '../ui/heading';
import { Input, InputField, } from '../ui/input';
import { FormControl, FormControlError, FormControlErrorIcon, FormControlErrorText, FormControlHelper, FormControlHelperText, FormControlLabel, FormControlLabelText } from '../ui/form-control';
import { Button, ButtonText } from '../ui/button';
import { Text } from '../ui/text';
import { useState, useEffect } from 'react';
import { AlertCircleIcon } from "../ui/icon";
import { useSigninForm } from '../../auth/useSigninForm';
import { useAuth } from '../../auth/provider';
import { useRouter } from 'solito/navigation';

export const SignupFormOTP = ({ supabase, onOTPSent }: any) => {
  const { t } = useTranslation();
  const { session, isLoading } = useAuth()
  const [isInvalid, setIsInvalid] = useState(false);
  const { form, updateField, submit, validateEmail, loading } = useSigninForm(supabase, onOTPSent);
  const { replace } = useRouter()
  useEffect(() => {
    if (!isLoading && session) {
      replace('/main')
    }
  }, [session, isLoading, replace])

  if (isLoading) return null


  const validate = async () => {
    if (!validateEmail(form.email)) {
      setIsInvalid(true)
    } else {
      submit()
    }
  };

  const handleCancel = async () => {
    replace('/main', {
      experimental: {
        nativeBehavior: 'stack-replace',
        isNestedNavigator: false, // Set to true if inside tabs/nested stack
      },
    })
  };

  return (
    <VStack className="rounded-xl border border-outline-200 bg-background-0 p-10 w-full h-full align-center justify-center">
      <Heading>{t('signup.otp.title')}</Heading>
      <Text className="mt-2">{t('signup.otp.subtitle')}</Text>

      <FormControl className="mt-4" isInvalid={isInvalid}>
        <FormControlLabel>
          <FormControlLabelText>{t('signup.email')}</FormControlLabelText>
        </FormControlLabel>
        <Input>
          <InputField value={form.email} onChangeText={(text) => updateField('email', text)} type="text" placeholder={t('signup.email_placeholder')} />
        </Input>
        <FormControlHelper>
          <FormControlHelperText>
            {t('signup.otp.email_helper')}
          </FormControlHelperText>
        </FormControlHelper>
        <FormControlError>
          <FormControlErrorIcon as={AlertCircleIcon} className="text-red-500" />
          <FormControlErrorText className="text-red-500">
            {t('signup.otp.email_invalid')}
          </FormControlErrorText>
        </FormControlError>
      </FormControl>

      <Button onPress={validate} className="rounded-full w-full mt-5">
        <ButtonText>{t('signup.otp.action')}</ButtonText>
      </Button>

      <Button onPress={handleCancel} className="rounded-full w-full mt-5 bg-secondary-900" size="sm">
        <ButtonText>{t('common.cancel')}</ButtonText>
      </Button>

    </VStack>
  );
};