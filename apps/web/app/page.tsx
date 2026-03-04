/**
 * Copyright (c) 2026 dilocash
 * Use of this source code is governed by an MIT-style
 * license that can be found in the LICENSE file.
 */
"use client";

import CommandsView from "@dilocash/ui/components/main/commands-view";
import { createConnectTransport } from "@connectrpc/connect-web";
import { getSupabaseClient } from "@dilocash/ui/auth/client";

const BASE_URL = "http://localhost:8080";

export default function Home() {
  const transport = createConnectTransport({
    baseUrl: process.env.NEXT_PUBLIC_API_URL || BASE_URL,
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
    <CommandsView transport={transport} className="h-screen" />
  );
}
