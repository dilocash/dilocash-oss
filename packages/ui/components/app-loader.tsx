'use client';
import { Spinner } from "./ui/spinner";
import { VStack } from "./ui/vstack";
import { Text } from "./ui/text";
import { Heading } from "./ui/heading";
import { Box } from "./ui/box";

interface AppLoaderProps {
  message?: string;
  subMessage?: string;
  isWeb?: boolean;
}

export const AppLoader = ({ 
  message = "Dilocash", 
  subMessage = "Inicializando...",
  isWeb = false
}: AppLoaderProps) => {
  return (
    <Box className="flex h-screen items-center justify-center">
        <VStack space="xl" className="items-center">
          {/* Aquí podrías poner tu Logo en el futuro */}
          <Box className="p-4 rounded-full bg-primary-500/10">
             <Heading size="3xl" className="text-primary-500 font-bold">
               D
             </Heading>
          </Box>
          <VStack space="xs" className="items-center">
            <Heading size="md" className="text-typography-900">
              {message}
            </Heading>
            <Text size="sm" className="text-typography-500">
              {subMessage}
            </Text>
          </VStack>
          {isWeb ? 
            <Text className="loader"></Text> :
           <Spinner size="large" color="$primary500" />
          }
        </VStack>
    </Box>
  );
};