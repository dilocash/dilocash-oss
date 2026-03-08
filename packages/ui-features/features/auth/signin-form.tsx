/**
 * Copyright (c) 2026 dilocash
 * Use of this source code is governed by an MIT-style
 * license that can be found in the LICENSE file.
 */

'use client';
import { useTranslation } from 'react-i18next';
import { VStack } from '@dilocash/ui-components/components/ui/vstack';
import { Heading } from '@dilocash/ui-components/components/ui//heading';
import { Input, InputField, InputIcon, InputSlot } from '@dilocash/ui-components/components/ui//input';
import { Button, ButtonText } from '@dilocash/ui-components/components/ui//button';
import { Text } from '@dilocash/ui-components/components/ui//text';
import { useState, useEffect } from 'react';
import { CheckIcon, EyeIcon, EyeOffIcon } from "@dilocash/ui-components/components/ui//icon";
import { Checkbox, CheckboxIndicator, CheckboxLabel, CheckboxIcon } from "@dilocash/ui-components/components/ui//checkbox";
import { HStack } from '@dilocash/ui-components/components/ui//hstack';
import { useLoginForm } from '../../hooks/useLoginForm';
import { useAuth } from '../../utils/auth-provider';
import { useRouter } from 'solito/navigation';
import { Box } from '@dilocash/ui-components/components/ui//box';

export const SigninForm = ({ supabase, onSuccess }: any) => {
  const { session, isLoading } = useAuth()
  const { replace } = useRouter()
  useEffect(() => {
    if (!isLoading && session) {
      replace('/')
    }
  }, [session, isLoading, replace])

  if (isLoading) return null

  const { form, updateField, submit, loading } = useLoginForm(supabase, onSuccess);
  const { t } = useTranslation();
  const [showPassword, setShowPassword] = useState(false);

  const handleSignUp = async () => {
    replace('/auth/signup', {
      experimental: {
        nativeBehavior: 'stack-replace',
        isNestedNavigator: false, // Set to true if inside tabs/nested stack
      },
    })
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
      <Heading>{t('login.title')}</Heading>
      <Text className="mt-2">{t('login.subtitle')}</Text>

      <Text className="mt-4">{t('login.email')}</Text>
      <Input className="rounded-xl">
        <InputField value={form.email} onChangeText={(text) => updateField('email', text)} type="text" placeholder={t('login.email_placeholder')} />
      </Input>

      <Text className="mt-6">{t('login.password')}</Text>
      <Input className="rounded-xl">
        <InputField
          type={showPassword ? 'text' : 'password'}
          placeholder={t('login.password_placeholder')}
          value={form.password}
          onChangeText={(text) => updateField('password', text)}
        />
        <InputSlot
          onPress={() => setShowPassword(!showPassword)}
          className="mr-3"
        >
          <InputIcon as={showPassword ? EyeIcon : EyeOffIcon} />
        </InputSlot>
      </Input>
      <Checkbox value={''} size="sm" className="mt-5">
        <CheckboxIndicator>
          <CheckboxIcon as={CheckIcon} />
        </CheckboxIndicator>
        <CheckboxLabel>{t('login.remember_me')}</CheckboxLabel>
      </Checkbox>
      <Button onPress={submit} className="rounded-full w-full mt-5">
        <ButtonText>{t('login.action')}</ButtonText>
      </Button>

      <Button onPress={handleCancel} className="rounded-full w-full mt-5 bg-secondary-900" size="sm">
        <ButtonText>{t('common.cancel')}</ButtonText>
      </Button>
      <HStack className="justify-center mt-5" >
        <Box className="justify-center pr-2">
          <Text>{t('login.sign_up_question')}</Text>
        </Box>
        <Button variant="link" onPress={handleSignUp}>
          <ButtonText className="underline underline-offset-1">
            {t('login.sign_up')}
          </ButtonText>
        </Button>
      </HStack>
      <HStack className="justify-center">
        <Button variant="link" size="sm">
          <ButtonText className="underline underline-offset-1 px-5">
            {t('login.forgot_password')}
          </ButtonText>
        </Button>
      </HStack>

    </VStack>
  );
};