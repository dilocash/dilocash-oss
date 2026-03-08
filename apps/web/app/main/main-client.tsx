/**
 * Copyright (c) 2026 dilocash
 * Use of this source code is governed by an MIT-style
 * license that can be found in the LICENSE file.
 */
"use client";

import CommandsView from "@dilocash/ui/components/main/commands-view";
import { getConnectTransport } from "../../lib/connect/transport";
import { useAuth } from "@dilocash/ui/auth/provider";

export default function MainClient() {
  const { session } = useAuth();
  const transport = getConnectTransport(session);

  return (
    <CommandsView transport={transport} className="h-screen h-dvh" />
  );
}
