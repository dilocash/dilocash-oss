/**
 * Copyright (c) 2026 dilocash
 * Use of this source code is governed by an MIT-style
 * license that can be found in the LICENSE file.
 */


import CommandsView from "@dilocash/ui-features/features/main/commands-view";
import getConnectTransport from "../lib/connect/transport";
import { useAuth } from "@dilocash/ui-features/utils/auth-provider";

export default function Home() {
  const { session } = useAuth();
  const transport = getConnectTransport(session);

  return (
    <CommandsView transport={transport} className="h-full" />
  );
}
