/**
 * Copyright (c) 2026 dilocash
 * Use of this source code is governed by an MIT-style
 * license that can be found in the LICENSE file.
 */
"use client";

import { SigninForm } from "@dilocash/ui-features/features/auth/signin-form";
import { Center } from "@dilocash/ui-components/components/ui/center";
import supabase from "../../../lib/supabase/client";
import { useEffect } from "react";
import { useAuth } from "@dilocash/ui-features/utils/auth-provider";
import { useRouter } from "solito/navigation";
import { Box } from "@dilocash/ui-components/components/ui/box";

export default function SigninClient() {
  const { session, isLoading } = useAuth()
  const { replace } = useRouter()
  useEffect(() => {
    // If session exists, redirect to main screen
    if (!isLoading && session) {
      replace('/')
    }
  }, [session, isLoading, replace])
  if (isLoading) return null

  return (
    <Box className="w-screen h-screen items-center justify-center">
      <Center className="w-full h-full md:w-auto md:h-auto">
        <SigninForm supabase={supabase} />
      </Center>
    </Box>
  );
}
