'use client';
// packages/ui/src/components/AuthForm.tsx
import React from 'react';
import { useTranslation } from 'react-i18next';
import { VStack } from '../ui/vstack';
import { Heading } from '../ui/heading';
import { Input, InputField, InputIcon, InputSlot } from '../ui/input';
import { Button, ButtonText } from '../ui/button';
import { Text } from '../ui/text';
import { useState } from 'react';
import { CheckIcon, EyeIcon, EyeOffIcon } from "../ui/icon";
import { Checkbox, CheckboxIndicator, CheckboxLabel, CheckboxIcon } from "../ui/checkbox";
import { HStack } from '../ui/hstack';

export const AuthForm = () => {
  const { t, i18n } = useTranslation();
  const [showPassword, setShowPassword] = useState(false);
  return (
    <VStack className="rounded-xl border border-outline-200 bg-background-0 p-6 w-full h-full align-center justify-center">
      <Heading>{t('login.title')}</Heading>
      <Text className="mt-2">{t('login.subtitle')}</Text>

      <Text className="mt-4">{t('login.email')}</Text>
      <Input>
        <InputField type="text" placeholder={t('login.email_placeholder')} />
      </Input>

      <Text className="mt-6">{t('login.password')}</Text>
      <Input>
        <InputField
              type={showPassword ? 'text' : 'password'}
              placeholder={t('login.password_placeholder')}
            />
            <InputSlot
              onPress={() => setShowPassword(!showPassword)}
              className="mr-3"
            >
              <InputIcon as={showPassword ? EyeIcon : EyeOffIcon} />
            </InputSlot>
        </Input>

        <HStack className="justify-between my-5">
            <Checkbox value={''} size="sm">
              <CheckboxIndicator>
                <CheckboxIcon as={CheckIcon} />
              </CheckboxIndicator>
              <CheckboxLabel>{t('login.remember_me')}</CheckboxLabel>
            </Checkbox>

            <Button variant="link" size="sm">
              <ButtonText className="underline underline-offset-1">
                {t('login.forgot_password')}
              </ButtonText>
            </Button>
          </HStack>

          <Button className="w-full" size="sm">
            <ButtonText>{t('login.action')}</ButtonText>
          </Button>
        </VStack>
  );
};