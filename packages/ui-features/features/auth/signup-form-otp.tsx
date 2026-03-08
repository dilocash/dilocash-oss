/**
 * Copyright (c) 2026 dilocash
 * Use of this source code is governed by an MIT-style
 * license that can be found in the LICENSE file.
 */

'use client';
import { useTranslation } from 'react-i18next';
import { VStack } from '@dilocash/ui-components/components/ui//vstack';
import { Heading } from '@dilocash/ui-components/components/ui//heading';
import { Input, InputField, } from '@dilocash/ui-components/components/ui//input';
import { FormControl, FormControlError, FormControlErrorIcon, FormControlErrorText, FormControlHelper, FormControlHelperText, FormControlLabel, FormControlLabelText } from '@dilocash/ui-components/components/ui//form-control';
import { Button, ButtonText } from '@dilocash/ui-components/components/ui//button';
import { Text } from '@dilocash/ui-components/components/ui//text';
import { useState, useEffect } from 'react';
import { AlertCircleIcon } from "@dilocash/ui-components/components/ui//icon";
import { useSigninForm } from '../../hooks/useSigninForm';
import { useAuth } from '../../utils/auth-provider';
import { useRouter } from 'solito/navigation';

export const SignupFormOTP = ({ supabase, onOTPSent }: any) => {
  const { t } = useTranslation();
  const { session, isLoading } = useAuth()
  const [isInvalid, setIsInvalid] = useState(false);
  const { form, updateField, submit, validateEmail, loading } = useSigninForm(supabase, onOTPSent);
  const { replace } = useRouter()
  useEffect(() => {
    if (!isLoading && session) {
      replace('/')
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
    replace('/', {
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
          <InputField
            value={form.email}
            onChangeText={(text) => updateField('email', text)}
            type="text"
            keyboardType="email-address"
            placeholder={t('signup.email_placeholder')} />
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