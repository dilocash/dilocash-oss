'use client';
// packages/ui/src/components/AuthForm.tsx
import React from 'react';
import { VStack } from '@dilocash/ui';
import { Box } from '@dilocash/ui';

import { Heading } from '@dilocash/ui';
import { Input, InputField } from '@dilocash/ui';
import { Button, ButtonText } from '@dilocash/ui';
import { Text } from '@dilocash/ui';

export const AuthForm = ({ type = 'login' }: { type: 'login' | 'register' }) => {
  return (
    <VStack space="md" className="w-full center p-6 rounded-2xl shadow-sm border border-outline-100">
      <VStack space="xs">
        <Heading size="xl" className="text-typography-900">
          {type === 'login' ? 'Bienvenido' : 'Crea tu cuenta'}
        </Heading>
        <Text size="sm" className="text-typography-500">
          {type === 'login' ? 'Ingresa tus credenciales para continuar' : 'Únete a Dilocash hoy mismo'}
        </Text>
      </VStack>

      <VStack space="md">
        <Input size="md">
          <InputField placeholder="Email" />
        </Input>
        <Input size="md">
          <InputField placeholder="Contraseña" type="password" />
        </Input>
      </VStack>

      <Button>
        <ButtonText>{type === 'login' ? 'Entrar' : 'Registrarse'}</ButtonText>
      </Button>
    </VStack>
  );
};