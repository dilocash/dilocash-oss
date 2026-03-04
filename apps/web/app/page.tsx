/**
 * Copyright (c) 2026 dilocash
 * Use of this source code is governed by an MIT-style
 * license that can be found in the LICENSE file.
 */
"use client";

import CommandsView from "@dilocash/ui/components/main/commands-view";
import { createConnectTransport } from "@connectrpc/connect-web";
import { supabase } from "../lib/supabase/client";
import { useEffect, useState } from "react";
import { Session } from "@supabase/supabase-js";

const BASE_URL = "http://localhost:8080";

export default function Home() {

  const [session, setSession] = useState<Session | null>(null)

  useEffect(() => {
    supabase.auth.getSession().then(({ data: { session } }) => {
      console.debug('getSession', session);
      setSession(session)
    })
    supabase.auth.onAuthStateChange((_event, session) => {
      console.debug('onAuthStateChange', _event, session);
      setSession(session)
    })
  }, []);

  const transport = createConnectTransport({
    baseUrl: process.env.NEXT_PUBLIC_API_URL || BASE_URL,
    interceptors: [
      (next) => async (req) => {
        if (session?.access_token) {
          req.header.set(
            "Authorization",
            `Bearer ${session.access_token}`,
          );
        }
        return await next(req);
      },
    ],
  });

  return (
    <CommandsView transport={transport} session={session} className="h-screen" />
  );
}
