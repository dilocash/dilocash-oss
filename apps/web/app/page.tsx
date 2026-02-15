/**
 * Copyright (c) 2026 dilocash
 * Use of this source code is governed by an MIT-style
 * license that can be found in the LICENSE file.
 */
import { Suspense } from "react";

import { AuthForm } from "@dilocash/ui/components/auth/auth-form";
import { Center } from "@dilocash/ui/components/ui/center";
import { VStack } from "@dilocash/ui/components/ui/vstack";
import { Box } from "@dilocash/ui/components/ui/box";

export default function Home() {
  return (
    <Suspense fallback={<div>Loading...</div>}>
      <Box className="w-full h-screen items-center justify-center">
        <Center className="min-w-80">
          <AuthForm/>
        </Center>        
      </Box>
    </Suspense>
  );
}
