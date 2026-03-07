'use client';
import { useTranslation } from 'react-i18next';
import { VStack } from '../ui/vstack';
import { Heading } from '../ui/heading';
import { Input, InputField, InputIcon, InputSlot } from '../ui/input';
import { Button, ButtonText } from '../ui/button';
import { Text } from '../ui/text';
import { useState, useEffect } from 'react';
import { EyeIcon, EyeOffIcon } from "../ui/icon";
import { useSigninForm } from '../../auth/useSigninForm';
import { useAuth } from '../../auth/provider';
import { useRouter } from 'solito/navigation';

export const SignupForm = ({ supabase, onOTPSent }: any) => {
  const { session, isLoading } = useAuth()
  const { replace } = useRouter()
  useEffect(() => {
    if (!isLoading && session) {
      replace('/main')
    }
  }, [session, isLoading, replace])

  if (isLoading) return null

  const { form, updateField, submit, loading } = useSigninForm(supabase, onOTPSent);
  const { t } = useTranslation();
  const [showPassword, setShowPassword] = useState(false);

  const handleSignIn = async () => {
    replace('/auth/signin', {
      experimental: {
        nativeBehavior: 'stack-replace',
        isNestedNavigator: false, // Set to true if inside tabs/nested stack
      },
    })
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
      <Heading>{t('signup.title')}</Heading>
      <Text className="mt-2">{t('signup.subtitle')}</Text>

      <Text className="mt-4">{t('signup.email')}</Text>
      <Input>
        <InputField value={form.email} onChangeText={(text) => updateField('email', text)} type="text" placeholder={t('signup.email_placeholder')} />
      </Input>

      <Text className="mt-5">{t('signup.password')}</Text>
      <Input>
        <InputField
          type={showPassword ? 'text' : 'password'}
          placeholder={t('signup.password_placeholder')}
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

      <Text className="mt-5">{t('signup.password_repeat')}</Text>
      <Input>
        <InputField
          type={showPassword ? 'text' : 'password'}
          placeholder={t('signup.password_repeat_placeholder')}
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

      <Button onPress={submit} className="rounded-full w-full mt-5">
        <ButtonText>{t('signup.action')}</ButtonText>
      </Button>

      <Button onPress={handleCancel} className="rounded-full w-full mt-5 bg-secondary-900" size="sm">
        <ButtonText>{t('common.cancel')}</ButtonText>
      </Button>

      <Text className="text-center pt-5">{t('signup.sign_up_question')}</Text>
      <Button variant="link" onPress={handleSignIn}>
        <ButtonText className="underline underline-offset-1">
          {t('signup.login')}
        </ButtonText>
      </Button>

    </VStack>
  );
};