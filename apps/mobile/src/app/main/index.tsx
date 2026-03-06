
import CommandsView from "@dilocash/ui/components/main/commands-view";
import getConnectTransport from "../lib/connect/transport";
import { useAuth } from "@dilocash/ui/auth/provider";

export default function Home() {
  const { session } = useAuth();
  const transport = getConnectTransport(session);

  return (
    <CommandsView transport={transport} className="h-full" />
  );
}
