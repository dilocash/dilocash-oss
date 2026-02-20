/**
 * Copyright (c) 2026 dilocash
 * Use of this source code is governed by an MIT-style
 * license that can be found in the LICENSE file.
 */
"use client";
import { AuthForm } from "@dilocash/ui/components/auth/auth-form";
import { Center } from "@dilocash/ui/components/ui/center";
import { Box } from "@dilocash/ui/components/ui/box";

import { VStack } from "@dilocash/ui/components/ui/vstack";
import Commands from "./commands";


export default function Home() {
  return (
      <Box className="w-full h-screen items-center justify-center">
        <Center>
          <AuthForm/>
          <VStack>
            <Commands />
          </VStack>
        </Center>
      </Box>
  );
}
