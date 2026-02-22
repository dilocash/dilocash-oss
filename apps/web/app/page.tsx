/**
 * Copyright (c) 2026 dilocash
 * Use of this source code is governed by an MIT-style
 * license that can be found in the LICENSE file.
 */
"use client";
// import { AuthForm } from "@dilocash/ui/components/auth/auth-form";
import { Center } from "@dilocash/ui/components/ui/center";
import { Box } from "@dilocash/ui/components/ui/box";

import { VStack } from "@dilocash/ui/components/ui/vstack";
import CommandsView from "@dilocash/ui/components/main/commands-view";
import { createConnectTransport } from "@connectrpc/connect-web";
import { getSupabaseClient } from "@dilocash/ui/auth/client";

const BASE_URL = "http://localhost:8080";

export default function Home() {
  const transport = createConnectTransport({
    baseUrl: BASE_URL,
    interceptors: [
      (next) => async (req) => {
        const supabase = getSupabaseClient(
          process.env.NEXT_PUBLIC_SUPABASE_URL!,
          process.env.NEXT_PUBLIC_SUPABASE_PUBLISHABLE_KEY!,
          localStorage,
        );
        const { data } = await supabase.auth.getSession();

        if (data.session?.access_token) {
          req.header.set(
            "Authorization",
            `Bearer ${data.session.access_token}`,
          );
        }
        return await next(req);
      },
    ],
  });

  return (
    <Box className="w-full h-full flex-1 flex-1 bg-gray-50">
      {/* <AuthForm supabase={supabase} onSuccess={() => console.log('Login success')} /> */}
      <CommandsView transport={transport} />
    </Box>
  );
}
