'use client';
// packages/ui/src/components/AuthForm.tsx
import React from 'react';
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
  const [showPassword, setShowPassword] = useState(false);
  return (
    <VStack className="rounded-xl border border-outline-200 bg-background-0 p-6 w-full h-full align-center justify-center">
      <Heading>Log in</Heading>
      <Text className="mt-2">Login to start using gluestack</Text>

      <Text className="mt-4">Email</Text>
      <Input>
        <InputField type="text" placeholder="Enter your email" />
      </Input>

      <Text className="mt-6">Password</Text>
      <Input>
        <InputField
              type={showPassword ? 'text' : 'password'}
              placeholder="Enter your password"
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
              <CheckboxLabel>Remember me</CheckboxLabel>
            </Checkbox>

            <Button variant="link" size="sm">
              <ButtonText className="underline underline-offset-1">
                Forgot Password?
              </ButtonText>
            </Button>
          </HStack>

          <Button className="w-full" size="sm">
            <ButtonText>Log in</ButtonText>
          </Button>
        </VStack>
  );
};