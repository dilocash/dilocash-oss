import { useEffect, useState } from "react";
import { AppLoader } from "@dilocash/ui/components/app-loader";
import CommandsView from "@dilocash/ui/components/main/commands-view";
import { createConnectTransport } from "@connectrpc/connect-web";
import supabase from "./lib/supabase/client";
import { KeyboardAvoidingView, Platform } from 'react-native';
import { SafeAreaView } from 'react-native-safe-area-context';

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
        //await getOfflineSession(); // Check SecureStore
        // emulates a minimal load to avoid flickering
        await new Promise((resolve) => setTimeout(resolve, 1000));
      } finally {
        setIsLoaded(true);
      }
    }
    prepare();
  }, []);
  if (!isLoaded) return <AppLoader subMessage="Accediendo..." isWeb={false} />;
  return (
    <SafeAreaView style={{ flex: 1 }}>
      <KeyboardAvoidingView
        behavior={Platform.OS === 'ios' ? 'padding' : 'height'}
      >
        <CommandsView transport={transport} className="h-full" />
      </KeyboardAvoidingView>
    </SafeAreaView>
  );
}
