/**
 * Copyright (c) 2026 dilocash
 * Use of this source code is governed by an MIT-style
 * license that can be found in the LICENSE file.
 */
"use client";

import CommandsView from "@dilocash/ui/components/main/commands-view";
import { AuthProvider } from "@dilocash/ui/auth/provider";
import { supabase } from "../../lib/supabase/client";
import { getConnectTransport } from "../../lib/connect/transport";
import { useAuth } from "@dilocash/ui/auth/provider";

export default function Home() {
  const { session } = useAuth();
  const transport = getConnectTransport(session);

  return (
    <AuthProvider supabase={supabase}>
      <CommandsView transport={transport} className="h-screen" />
    </AuthProvider>
  );
}
