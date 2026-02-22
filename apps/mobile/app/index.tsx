import { AuthForm } from "@dilocash/ui/components/auth/auth-form";
import { useEffect, useState } from "react";
import { AppLoader } from "@dilocash/ui/components/app-loader";
import CommandsView from "@dilocash/ui/components/main/commands-view";
import { createConnectTransport } from "@connectrpc/connect-web";
import supabase from "./lib/supabase/client";
const BASE_URL = process.env.EXPO_PUBLIC_API_URL;

export default function Index() {
  const [isLoaded, setIsLoaded] = useState(false);
  const transport = createConnectTransport({
    baseUrl: BASE_URL!,
    interceptors: [
      (next) => async (req) => {
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
  useEffect(() => {
    async function prepare() {
      try {
        //await getOfflineSession(); // Verifica SecureStore
        // Simula una carga mínima para evitar parpadeos
        await new Promise((resolve) => setTimeout(resolve, 10000));
      } finally {
        setIsLoaded(true);
      }
    }
    prepare();
  }, []);
  if (!isLoaded) return <AppLoader subMessage="Accediendo..." isWeb={false} />;
  return (
    <>
      <CommandsView transport={transport} />
    </>
  );
}
