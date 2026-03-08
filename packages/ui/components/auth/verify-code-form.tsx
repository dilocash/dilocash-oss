/**
 * Copyright (c) 2026 dilocash
 * Use of this source code is governed by an MIT-style
 * license that can be found in the LICENSE file.
 */

'use client';
import { useTranslation } from 'react-i18next';
import { VStack } from '../ui/vstack';
import { Heading } from '../ui/heading';
import { Input, InputField, InputIcon, InputSlot } from '../ui/input';
import { Button, ButtonText } from '../ui/button';
import { Text } from '../ui/text';
import { useState, useEffect, useRef } from 'react';
import { AlertCircleIcon, CloseCircleIcon, EyeIcon, EyeOffIcon, InfoIcon } from "../ui/icon";
import { useOTPVerificationForm } from '../../auth/useOTPVerificationForm';
import { useRouter } from 'solito/navigation';
import { FormControl, FormControlError, FormControlErrorIcon, FormControlErrorText, FormControlHelper, FormControlHelperText, FormControlLabel, FormControlLabelText } from '../ui/form-control';
import { Box } from '../ui/box';
import { Alert, AlertIcon, AlertText } from '../ui/alert';
import { HStack } from '../ui/hstack';

export const VerifyCodeForm = ({ supabase, email, onOTPVerified }: any) => {
  const { replace } = useRouter()
  const { t } = useTranslation();
  const onError = () => {
    console.log('error validating otp')
    updateField('code', '')
    setIsValidationFailed(true)
  }
  const { form, updateField, submit, loading } = useOTPVerificationForm(supabase, email, onOTPVerified, onError);
  const [isInvalid, setIsInvalid] = useState(false);
  const [isValidationFailed, setIsValidationFailed] = useState(false);
  const validate = async () => {
    if (!form.code || form.code.trim().length != 6) {
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

  const handleCodeChange = (value: string) => {
    setIsValidationFailed(false)
    // Replace any character that is not a digit (0-9) with an empty string
    const numericValue = value.replace(/[^0-9]/g, '');
    updateField('code', numericValue)
    if (numericValue.length == 6) {
      setIsInvalid(false)
    }
  };

  return (
    <VStack className="rounded-xl border border-outline-200 bg-background-0 p-10 w-full h-full align-center justify-center">
      <Heading>{t('otp.title')}</Heading>
      <HStack className='inline' space="md">
        <Text>{t('otp.subtitle') + " "}</Text>
        <Text className='italic'>{email}</Text>
      </HStack>
      <FormControl className="mt-4" isInvalid={isInvalid}>
        <FormControlLabel>
          <FormControlLabelText>{t('otp.code_label')}</FormControlLabelText>
        </FormControlLabel>
        <Input size="xl" variant="underlined" className='gap-10'>
          <InputField className='text-center tracking-widest'
            value={form.code}
            onChangeText={(text) => handleCodeChange(text)}
            keyboardType="number-pad"
            maxLength={6}
            // For iOS SMS auto-fill support
            textContentType="oneTimeCode"
            type="text" />
        </Input>
        <FormControlHelper>
          <FormControlHelperText>
            {t('otp.code_helper')}
          </FormControlHelperText>
        </FormControlHelper>
        <FormControlError>
          <FormControlErrorIcon as={AlertCircleIcon} className="text-red-500" />
          <FormControlErrorText className="text-red-500">
            {t('otp.code_invalid')}
          </FormControlErrorText>
        </FormControlError>
      </FormControl>
      {isValidationFailed &&
        <Alert action="warning" variant="outline" >
          <AlertIcon as={CloseCircleIcon} />
          <AlertText>{t('otp.validation_failed')}</AlertText>
        </Alert>}
      <Button onPress={validate} className="rounded-full w-full mt-5">
        <ButtonText>{t('otp.action')}</ButtonText>
      </Button>

      <Button onPress={handleCancel} className="rounded-full w-full mt-5 bg-secondary-900" size="sm">
        <ButtonText>{t('common.cancel')}</ButtonText>
      </Button>

    </VStack>
  );
};